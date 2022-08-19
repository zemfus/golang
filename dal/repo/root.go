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

func (r root) CreateCategory(ctx context.Context, name string) error {
	exec, err := r.connPool.Exec(ctx, "SELECT * FROM Categoru WHERE name = $1", name)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 0 {
		return nil
	}
	_, err = r.connPool.Exec(ctx,
		"INSERT INTO category(name) VALUES($1)", name)

	return err
}

func (r root) GetAllCategory(ctx context.Context) ([]models.Category, error) {
	rows, err := r.connPool.Query(ctx, `SELECT id, name FROM category `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories = []models.Category{}

	for rows.Next() {
		var category = models.Category{}
		err = rows.Scan(&category.ID,
			&category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r root) DeleteCategory(ctx context.Context, ID int) error {
	_, err := r.connPool.Exec(ctx, "DELETE FROM category WHERE id = $1", ID)
	return err
}

func (r root) GetAllInventory(ctx context.Context) ([]models.Inventory, error) {
	rows, err := r.connPool.Query(ctx, "SELECT id, name FROM Inventory")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var inventorys []models.Inventory
	for rows.Next() {
		var inventory models.Inventory
		err := rows.Scan(&inventory.ID,
			&inventory.Name,
		)
		if err != nil {
			return nil, err
		}
		inventorys = append(inventorys, inventory)
	}
	return inventorys, nil

}

func (r root) UpdateInventory(ctx context.Context, inventory *models.Inventory) error {
	_, err := inventory.connPool.Exec(ctx,
		`UPDATE inventory SET
			name = $2,
			description = $3,
			campus_id = $4,
			category_id = $5,
			updated_at = now()
		WHERE id = $6`,
		inventory.Name,
		inventory.Description,
		inventory.Campus,
		inventory.Category,
		inventory.ID)

	return err
}

func (r root) DeleteInventory(ctx context.Context, ID int) error {
	_, err := r.connPool.Exec(ctx, "DELETE FROM inventory WHERE id = $1", ID)
	return err
}

func (r root) CreateInventory(ctx context.Context, inventory *models.Inventory) error {
	exec, err := r.connPool.Exec(ctx, "SELECT * FROM inventory WHERE id = $1", inventory.ID)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 0 {
		return nil
	}
	_, err = r.connPool.Exec(ctx,
		"INSERT INTO inventory(id, name, description, campus_id, category_id) VALUES($1, $2, $3, $4, $5)",
		inventory.ID, inventory.Description, inventory.Campus, inventory.Category)

	return err
}
func (r root) GetInventoryByID(ctx context.Context, ID int) (*models.Inventory, error) {
	rows, err := r.connPool.Query(ctx, `SELECT
		id,
		name,
		description,
		campus_id,
		category_id,
		handle_step
			FROM users WHERE id = $1`, ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory = models.Inventory{}
	if rows.Next() {
		err = rows.Scan(&inventory.ID,
			&inventory.Name,
			&inventory.Description,
			&inventory.Campus,
			&inventory.Category)
		if err != nil {
			return nil, err
		}
	}

	return &inventory, nil
}
