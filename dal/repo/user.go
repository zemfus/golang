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

func (r user) GetByID(ctx context.Context, id int) (*models.User, error) {
	rows, err := r.connPool.Query(ctx, `SELECT
		id,
		nickname,
		email,
		campus_id,
		role,
		handle_step,
		last_msg
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
			&user.HandleStep,
			&user.LastMsg)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (r user) GetAllByCampus(ctx context.Context, campus string) ([]models.User, error) {
	rows, err := r.connPool.Query(ctx, `SELECT 
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

func (r user) GetByNickname(ctx context.Context, nickname string) (*models.User, error) {
	rows, err := r.connPool.Query(ctx, `SELECT
		id,
		nickname,
		email,
		campus_id,
		role,
		handle_step,
		last_msg
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
			&user.HandleStep,
			&user.LastMsg)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (r user) ExistsUsersInCampusByID(ctx context.Context, ID int) (bool, error) {
	cmd, err := r.connPool.Exec(ctx, "SELECT id FROM users WHERE campus_id = $1", ID)
	if err != nil {
		return false, err
	}

	if cmd.RowsAffected() != 0 {
		return true, nil
	}
	return false, nil
}

func (r user) Create(ctx context.Context, user *models.User) error {
	exec, err := r.connPool.Exec(ctx, "SELECT * FROM users WHERE id = $1", user.ID)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 0 {
		return nil
	}
	_, err = r.connPool.Exec(ctx,
		"INSERT INTO users(id, nickname, email, campus_id, role, handle_step) VALUES($1, $2, $3, $4, $5, $6)",
		user.ID, user.Nickname, user.Email, user.CampusID, user.Role, user.HandleStep)

	return err
}

func (r user) Update(ctx context.Context, user *models.User) error {
	_, err := r.connPool.Exec(ctx,
		`UPDATE users SET
			nickname = $1,
			email = $2,
			campus_id = $3,
			role = $4,
			handle_step = $5,
			last_msg = $6,
			updated_at = now()
		WHERE id = $7`,
		user.Nickname,
		user.Email,
		user.CampusID,
		user.Role,
		user.HandleStep,
		user.LastMsg,
		user.ID)

	return err
}

func (r user) Delete(ctx context.Context, ID int) error {
	_, err := r.connPool.Exec(ctx, "DELETE FROM users WHERE id = $1", ID)
	return err
}

func (r user) ExistByID(ctx context.Context, id int) (bool, error) {
	exec, err := r.connPool.Exec(ctx, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	if exec.RowsAffected() != 0 {
		return false, err
	} else {
		return true, err
	}

}

func (r user) ExistByNickname(ctx context.Context, nickname string) (bool, error) {
	exec, err := r.connPool.Exec(ctx, "SELECT * FROM users WHERE nickname = $1", nickname)
	if err != nil {
		return false, err
	}

	if exec.RowsAffected() != 0 {
		return false, err
	} else {
		return true, err
	}
}
