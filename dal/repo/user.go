package repo

import (
	"context"

	"boobot/kernel/domain/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type user struct {
	connPool *pgxpool.Pool
}

func NewUser(connPool *pgxpool.Pool) User {
	return &user{
		connPool: connPool,
	}
}

func (u user) GetByID(ctx context.Context, id int) (*models.User, error) {
	rows, err := u.connPool.Query(ctx, `SELECT
		id,
		nickname,
		email,
		campus_id,
		role,
		handle_step
			FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user = models.User{}
	if rows.Next() {
		err = rows.Scan(&user.ID,
			&user.Nickname,
			&user.Email,
			&user.CampusID,
			&user.Role,
			&user.HandleStep)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u user) GetAllByCampus(ctx context.Context, campus string) ([]models.User, error) {
	rows, err := u.connPool.Query(ctx, `SELECT 
    id, 
    nickname, 
    email, 
    campus_id, 
    role, 
    handle_step 
		FROM users WHERE campus_id = (SELECT 
		                                  id
		                              		FROM campus WHERE name=$1)`, campus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = []models.User{}

	for rows.Next() {
		var user = models.User{}
		err = rows.Scan(&user.ID,
			&user.Nickname,
			&user.Email,
			&user.CampusID,
			&user.Role,
			&user.HandleStep)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u user) GetByNickname(ctx context.Context, nickname string) (*models.User, error) {
	rows, err := u.connPool.Query(ctx, `SELECT
		id,
		nickname,
		email,
		campus_id,
		role,
		handle_step
			FROM users WHERE nickname = $1`, nickname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user = models.User{}
	if rows.Next() {
		err = rows.Scan(&user.ID,
			&user.Nickname,
			&user.Email,
			&user.CampusID,
			&user.Role,
			&user.HandleStep)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u user) Create(ctx context.Context, user *models.User) error {
	exec, err := u.connPool.Exec(ctx, "SELECT * FROM users WHERE id = $1", user.ID)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 0 {
		return nil
	}
	_, err = u.connPool.Exec(ctx,
		"INSERT INTO users(id, nickname, email, campus_id, role, handle_step) VALUES($1, $2, $3, $4, $5, $6)",
		user.ID, user.Nickname, user.Email, user.CampusID, user.Role, user.HandleStep)

	return err
}

func (u user) Update(ctx context.Context, user *models.User) error {
	_, err := u.connPool.Exec(ctx,
		`UPDATE users SET
			nickname = $1,
			email = $2,
			campus_id = $3,
			role = $4,
			handle_step = $5,
			updated_at = now()
		WHERE id = $6`,
		user.Nickname,
		user.Email,
		user.CampusID,
		user.Role,
		user.HandleStep,
		user.ID)

	return err
}

func (u user) Delete(ctx context.Context, ID int) error {
	_, err := u.connPool.Exec(ctx, "DELETE FROM users WHERE id = $1", ID)
	return err
}

func (u user) ExistByID(ctx context.Context, id int) (bool, error) {
	exec, err := u.connPool.Exec(ctx, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	if exec.RowsAffected() != 0 {
		return false, err
	} else {
		return true, err
	}

}

func (u user) ExistByNickname(ctx context.Context, nickname string) (bool, error) {
	exec, err := u.connPool.Exec(ctx, "SELECT * FROM users WHERE nickname = $1", nickname)
	if err != nil {
		return false, err
	}

	if exec.RowsAffected() != 0 {
		return false, err
	} else {
		return true, err
	}
}
