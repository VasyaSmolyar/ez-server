package adaptor

import (
	"context"
	"errors"
	"ex-server/internal/auth/entity"
	"ex-server/internal/auth/exception"
	"ex-server/pkg/bcrypt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	UniqueViolationErrCode = "23505"
)

func Init(db *pgx.Conn) *AuthRepository {
	return &AuthRepository{db: db}
}

type AuthRepository struct {
	db *pgx.Conn
}

func (repo *AuthRepository) Signin(ctx context.Context, userLogin, pass string) (*entity.User, error) {
	var id, login, hashedPass string
	var query = "select * from users where login=$1"
	if err := repo.db.QueryRow(ctx, query, userLogin).Scan(&id, &login, &hashedPass); err != nil {
		if err == pgx.ErrNoRows {
			return nil, exception.ErrWrongCreds
		} else {
			return nil, err
		}
	}

	if bcrypt.Check(pass, hashedPass) {
		return &entity.User{
			Id:    id,
			Login: login,
		}, nil
	}
	return nil, exception.ErrWrongCreds
}

func (repo *AuthRepository) Get(ctx context.Context, userID string) (*entity.User, error) {
	var login string
	var query = "select login from users where id=$1"
	if err := repo.db.QueryRow(ctx, query, userID).Scan(&login); err != nil {
		if err == pgx.ErrNoRows {
			return nil, exception.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &entity.User{
		Id:    userID,
		Login: login,
	}, nil
}

func (repo *AuthRepository) Signup(ctx context.Context, login, pass string) error {
	if hashed, err := bcrypt.Generate(pass); err == nil {
		var query = "insert into users(login, hashed_pass) values($1, $2)"
		var pgErr *pgconn.PgError

		_, err := repo.db.Exec(ctx, query, login, hashed)
		if errors.As(err, &pgErr) {
			if pgErr.Code == UniqueViolationErrCode {
				return exception.ErrAlreadyCreated
			} else if err != nil {
				return err
			}
		}
		return nil
	}
	return exception.ErrHashingFailed
}
