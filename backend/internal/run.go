package internal

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net/http"
	_ "net/http/pprof" //nolint:gosec // pprof port is not exposed to the internet
	"os"
	"os/signal"
	"strings"
	"study-chat/internal/api"
	"study-chat/internal/config"

	"study-chat/pkg/logger"
	"study-chat/pkg/sentry"
)

func Run(cfg config.Config) error {
	if err := sentry.Init(cfg.Sentry.DSN, cfg.Sentry.Environment); err != nil {
		return fmt.Errorf("failed to init sentry: %w", err)
	}
	logger.Setup()

	g, ctx := errgroup.WithContext(context.Background())
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	if err := startServers(ctx, g, cfg); err != nil {
		return err
	}
	if cfg.Server.PprofPort != "" {
		startPprofServer(ctx, g, cfg)
	}

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("server exited with error: %w", err)
	}
	return nil
}

func startServers(ctx context.Context, g *errgroup.Group, cfg config.Config) error {
	locator, err := config.InitLocator(cfg)
	if err != nil {
		return err
	}

	echoInstance := api.SetupRESTPServer(locator)
	grpcServer := api.SetupGRPCServer(locator)

	address := "0.0.0.0:" + cfg.Server.Port
	server := &http.Server{
		Addr: address,
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				echoInstance.ServeHTTP(w, r)
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

	return nil
}

func startPprofServer(ctx context.Context, g *errgroup.Group, cfg config.Config) {
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
