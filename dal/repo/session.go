package repo

import (
	"context"

	"boobot/kernel/domain/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type session struct {
	connPool *pgxpool.Pool
}

func NewSession(connPool *pgxpool.Pool) Session {
	return &session{
		connPool: connPool,
	}
}

func (s session) GetByID(ctx context.Context, id int) (*models.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s session) GetByUserID(ctx context.Context, userID int) (*models.Session, error) {
	rows, err := s.connPool.Query(ctx, `SELECT
		code,
			FROM sessions WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var session = models.Session{}
	if rows.Next() {
		err = rows.Scan(&session.Code)
		if err != nil {
			return nil, err
		}
	}

	return &session, nil
}

func (s session) Create(ctx context.Context, session *models.Session) error {
	_, err := s.connPool.Exec(ctx,
		"INSERT INTO sessions(user_id, code) VALUES($1, $2)",
		session.UserID, session.Code)

	return err
}

func (s session) ExistsByCodeAndUserID(ctx context.Context, userID int, code int) (bool, error) {
	rows, err := s.connPool.Exec(ctx,
		`SELECT id FROM sessions WHERE user_id = $1 AND code = $2`,
		userID, code)
	if err != nil {
		return false, err
	}

	if rows.RowsAffected() != 0 {
		return true, nil
	}

	return false, nil
}

func (s session) Delete(ctx context.Context, ID int) error {
	//TODO implement me
	panic("implement me")
}
