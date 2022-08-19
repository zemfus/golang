package repo

import (
	"context"

	"boobot/kernel/domain/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type session struct {
	connPool *pgxpool.Pool
}

func NewSession(connPool *pgxpool.Pool) *session {
	return &session{
		connPool: connPool,
	}
}

func (s session) GetByID(ctx context.Context, id int) (*models.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s session) GetByUserID(ctx context.Context, userID int) (*models.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s session) Create(ctx context.Context, session *models.Session) error {
	_, err := s.connPool.Exec(ctx,
		"INSERT INTO sessions(user_id, code) VALUES($1, $2)",
		session.UserID, session.Code)

	return err
}

func (s session) Delete(ctx context.Context, ID int) error {
	//TODO implement me
	panic("implement me")
}
