package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	hasql "golang.yandex/hasql/sqlx"
)

type User struct {
	Id        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, usr *User) (*User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, updateFn func(*User) (bool, error)) (*User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CheckUserExist(ctx context.Context, email string) (bool, error)
}

type PostgresRepo struct {
	cluster *hasql.Cluster
}

func newUserRepository(cluster *hasql.Cluster) UserRepository {
	return &PostgresRepo{
		cluster: cluster,
	}
}

func (r *PostgresRepo) CreateUser(
	ctx context.Context,
	user *User,
) (*User, error) {
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

func (r *PostgresRepo) UpdateUser(
	ctx context.Context,
	id uuid.UUID,
	updateFn func(*User) (bool, error),
) (*User, error) {
	db := r.cluster.Primary().DBx()
	user, err := r.GetUserById(ctx, id)
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

func (r *PostgresRepo) SaveUser(ctx context.Context, user User) error {
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

func (r *PostgresRepo) GetUserById(ctx context.Context, id uuid.UUID) (*User, error) {
	db := r.cluster.StandbyPreferred().DBx()
	var user User
	query := `SELECT id, first_name, email FROM "user" WHERE id = $1`
	if err := db.GetContext(ctx, &user, query, id); err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}
	return &user, nil
}

func (r *PostgresRepo) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	db := r.cluster.StandbyPreferred().DBx()
	var user User
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
