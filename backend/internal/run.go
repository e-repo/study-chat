package internal

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	hasql "golang.yandex/hasql/sqlx"
	"log/slog"
	"net/http"
	_ "net/http/pprof" //nolint:gosec // pprof port is not exposed to the internet
	"os"
	"os/signal"
	"strings"
	"study-chat/internal/infra/service"
	"study-chat/internal/ui/api"

	userinfra "study-chat/internal/infra/user_infra"
	"study-chat/pkg/logger"
	"study-chat/pkg/postgres"
	"study-chat/pkg/sentry"
)

func Run(cfg service.Config) error {
	if err := sentry.Init(cfg.Sentry.DSN, cfg.Sentry.Environment); err != nil {
		return fmt.Errorf("failed to init sentry: %w", err)
	}
	logger.Setup()

	connData, err := postgres.NewConnectionData(
		cfg.Postgres.Hosts,
		cfg.Postgres.Database,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Port,
		cfg.Postgres.SSL,
	)
	if err != nil {
		return fmt.Errorf("failed to init postgres connection data: %w", err)
	}
	cluster, err := postgres.InitCluster(context.Background(), connData)
	if err != nil {
		return fmt.Errorf("failed to init postgres cluster: %w", err)
	}

	g, ctx := errgroup.WithContext(context.Background())
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	startServers(ctx, g, cluster, cfg)
	if cfg.Server.PprofPort != "" {
		startPprofServer(ctx, g, cfg)
	}

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("server exited with error: %w", err)
	}
	return nil
}

func startServers(ctx context.Context, g *errgroup.Group, cluster *hasql.Cluster, cfg service.Config) {
	locator := initLocator(cluster, cfg)

	httpServer := api.SetupHTTPServer(locator)
	//grpcServer := application.SetupGRPCServer(userRepo, orderRepo, productRepo)

	address := "0.0.0.0:" + cfg.Server.Port
	server := &http.Server{
		Addr: address,
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				panic("gRPC server временно отключен!")
				//grpcServer.ServeHTTP(w, r)
			} else {
				httpServer.ServeHTTP(w, r)
			}
		}), &http2.Server{}),
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
	}

	g.Go(func() error {
		slog.Info("Starting server http and grpc server at " + address)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		slog.Info("Http and grpc server shut down gracefully")
		return nil
	})
	g.Go(func() error {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.InterruptTimeout)
		defer cancel()
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		return nil
	})
}

func startPprofServer(ctx context.Context, g *errgroup.Group, cfg service.Config) {
	pprofAddress := "0.0.0.0:" + cfg.Server.PprofPort
	//nolint:gosec // pprofServer is not exposed to the internet
	pprofServer := &http.Server{Addr: pprofAddress, Handler: http.DefaultServeMux}
	g.Go(func() error {
		slog.Info("Starting pprof server at " + pprofAddress)
		if err := pprofServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		slog.Info("Pprof server shut down gracefully")
		return nil
	})
	g.Go(func() error {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.InterruptTimeout)
		defer cancel()
		err := pprofServer.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		return nil
	})
}

func initLocator(cluster *hasql.Cluster, cfg service.Config) service.LocatorInterface {
	locator := service.NewLocator()
	userRepo := userinfra.NewPostgres(cluster)
	validatorComp := service.NewRuValidator()

	locator.Add(service.ValidatorServiceKey, validatorComp)
	locator.Add(service.ConfigServiceKey, cfg)
	locator.Add(service.UserRepositoryServiceKey, userRepo)

	return locator
}
