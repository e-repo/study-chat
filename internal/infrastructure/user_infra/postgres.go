package user_infra

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	hasql "golang.yandex/hasql/sqlx"

	domain "study-chat/internal/domain/user_dmn"
)

type userDB struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
}

type PostgresRepo struct {
	cluster *hasql.Cluster
}

func NewPostgres(cluster *hasql.Cluster) *PostgresRepo {
	return &PostgresRepo{
		cluster: cluster,
	}
}

func (r *PostgresRepo) CreateUser(
	ctx context.Context,
	user *domain.User,
) (*domain.User, error) {
	db := r.cluster.Primary().DBx()

	userDB := userDB{
		ID:        user.ID(),
		FirstName: user.FirstName(),
		Email:     user.Email(),
		Password:  user.PassHash(),
	}
	query := `
		INSERT INTO "user" (id, first_name, email, password)
		VALUES (:id, :first_name, :email, :password)
	`
	_, err := db.NamedExecContext(ctx, query, userDB)
	if err != nil {
		return nil, fmt.Errorf("ошибка сохранения пользователя: %w", err)
	}

	return user, nil
}

func (r *PostgresRepo) UpdateUser(
	ctx context.Context,
	id uuid.UUID,
	updateFn func(*domain.User) (bool, error),
) (*domain.User, error) {
	db := r.cluster.Primary().DBx()
	user, err := r.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}
	updated, err := updateFn(user)
	if err != nil {
		return nil, fmt.Errorf("ошибка обновления пользователя: %w", err)
	}
	if !updated {
		return user, nil
	}
	userDB := userDB{
		ID:        user.ID(),
		FirstName: user.FirstName(),
		Email:     user.Email(),
		Password:  user.PassHash(),
	}
	query := `
		UPDATE "user"
		SET first_name = :first_name, email = :email, password = :password
		WHERE id = :id
	`
	_, err = db.NamedExecContext(ctx, query, userDB)
	if err != nil {
		return nil, fmt.Errorf("ошибка обновления пользователя: %w", err)
	}
	return user, nil
}

func (r *PostgresRepo) SaveUser(ctx context.Context, entity domain.User) error {
	db := r.cluster.Primary().DBx()
	user := userDB{
		ID:        entity.ID(),
		FirstName: entity.FirstName(),
		Email:     entity.Email(),
		Password:  entity.PassHash(),
	}
	query := `
		INSERT INTO "user" (id, first_name, email, password)
		VALUES (:id, :name, :email)
		ON CONFLICT (id) DO UPDATE SET name = :name, email = :email
	`
	_, err := db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("ошибка перезаписи пользователя: %w", err)
	}
	return nil
}

func (r *PostgresRepo) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	db := r.cluster.StandbyPreferred().DBx()
	var user userDB
	query := `SELECT id, first_name, email FROM user WHERE id = $1`
	if err := db.GetContext(ctx, &user, query, id); err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}
	entity, err := domain.NewUser(
		&domain.UserDto{
			Id:       user.ID,
			FistName: user.FirstName,
			Email:    user.Email,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации пользователя: %w", err)
	}
	return entity, nil
}

func (r *PostgresRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	db := r.cluster.Primary().DBx()
	query := `DELETE FROM "user" WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Ошибка удаления пользователя: %w", err)
	}
	return nil
}

func (r *PostgresRepo) CheckUserExist(ctx context.Context, email string) (bool, error) {
	db := r.cluster.Primary().DBx()

	query := `SELECT EXISTS (SELECT 1 FROM "user" WHERE email = $1)`
	var exist bool
	err := db.GetContext(ctx, &exist, query, email)
	if err != nil {
		return false, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return exist, nil
}
