package repo

import (
	"context"

	"boobot/kernel/domain/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type root struct {
	connPool *pgxpool.Pool
}

func NewRoot(connPool *pgxpool.Pool) Root {
	return &root{
		connPool: connPool,
	}
}

func (r root) GetAllCampuses(ctx context.Context) ([]models.Campus, error) {
	rows, err := r.connPool.Query(ctx, "SELECT id, name FROM Campus")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campuses []models.Campus
	for rows.Next() {
		var campus models.Campus
		err := rows.Scan(&campus.ID,
			&campus.Name,
		)
		if err != nil {
			return nil, err
		}
		campuses = append(campuses, campus)
	}

	return campuses, nil
}
