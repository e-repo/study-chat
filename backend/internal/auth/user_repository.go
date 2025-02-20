package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	hasql "golang.yandex/hasql/sqlx"
)

type user struct {
	Id        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
}

type userRepository interface {
	createUser(ctx context.Context, user *user) (*user, error)
	updateUser(ctx context.Context, id uuid.UUID, updateFn func(*user) (bool, error)) (*user, error)
	getUserById(ctx context.Context, id uuid.UUID) (*user, error)
	getUserByEmail(ctx context.Context, email string) (*user, error)
	checkUserExist(ctx context.Context, email string) (bool, error)
}

type PostgresRepo struct {
	cluster *hasql.Cluster
}

func newUserRepository(cluster *hasql.Cluster) userRepository {
	return &PostgresRepo{
		cluster: cluster,
	}
}

func (r *PostgresRepo) createUser(
	ctx context.Context,
	user *user,
) (*user, error) {
	db := r.cluster.Primary().DBx()

	query := `
		INSERT INTO "user" (id, first_name, email, password)
		VALUES (:id, :first_name, :email, :password)
	`
	_, err := db.NamedExecContext(ctx, query, user)
	if err != nil {
		return nil, fmt.Errorf("ошибка сохранения пользователя: %w", err)
	}

	return user, nil
}

func (r *PostgresRepo) updateUser(
	ctx context.Context,
	id uuid.UUID,
	updateFn func(*user) (bool, error),
) (*user, error) {
	db := r.cluster.Primary().DBx()
	user, err := r.getUserById(ctx, id)
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
	query := `
		UPDATE "user"
		SET first_name = :first_name, email = :email, password = :password
		WHERE id = :id
	`
	_, err = db.NamedExecContext(ctx, query, user)
	if err != nil {
		return nil, fmt.Errorf("ошибка обновления пользователя: %w", err)
	}
	return user, nil
}

func (r *PostgresRepo) SaveUser(ctx context.Context, user user) error {
	db := r.cluster.Primary().DBx()
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

func (r *PostgresRepo) getUserById(ctx context.Context, id uuid.UUID) (*user, error) {
	db := r.cluster.StandbyPreferred().DBx()
	var user user
	query := `SELECT id, first_name, email FROM "user" WHERE id = $1`
	if err := db.GetContext(ctx, &user, query, id); err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}
	return &user, nil
}

func (r *PostgresRepo) getUserByEmail(ctx context.Context, email string) (*user, error) {
	db := r.cluster.StandbyPreferred().DBx()
	var user user
	query := `SELECT id, first_name, email, password FROM "user" WHERE email = $1`
	if err := db.GetContext(ctx, &user, query, email); err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}
	return &user, nil
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

func (r *PostgresRepo) checkUserExist(ctx context.Context, email string) (bool, error) {
	db := r.cluster.Primary().DBx()

	query := `SELECT EXISTS (SELECT 1 FROM "user" WHERE email = $1)`
	var exist bool
	err := db.GetContext(ctx, &exist, query, email)
	if err != nil {
		return false, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return exist, nil
}
