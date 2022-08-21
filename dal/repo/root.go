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

func (r root) UpdateCategory(ctx context.Context, inventory *models.Category) error {
	_, err := r.connPool.Exec(ctx,
		`UPDATE
		inventory SET
		name = $1
		updated_at = now()
			WHERE id = $2`,
		inventory.Name,
		inventory.ID,
	)

	return err
}

func (r root) GetCategoryByID(ctx context.Context, ID int) (*models.Category, error) {
	rows, err := r.connPool.Query(ctx,
		`SELECT
			id,
			name,
		FROM category
			WHERE id = $1`,
		ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category = models.Category{}
	if rows.Next() {
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func (r root) GetAllCampuses(ctx context.Context) ([]models.Campus, error) {
	rows, err := r.connPool.Query(ctx, "SELECT id, name FROM campus")
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

func (r root) CreateCategory(ctx context.Context, category *models.Category) error {
	exec, err := r.connPool.Exec(ctx,
		`SELECT id FROM category WHERE name = $1`, category)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 0 {
		return nil
	}

	_, err = r.connPool.Exec(ctx,
		`INSERT INTO category(
		name,
		) VALUES($1)`,
		category)

	return err
}

func (r root) GetAllCategoryByBookType(ctx context.Context, bookType models.BookType) ([]models.Category, error) {
	rows, err := r.connPool.Query(ctx, `SELECT id, name FROM category WHERE type = $1`, bookType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

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

func (r root) GetAllInventoryByCampusID(ctx context.Context, campusID int) ([]models.Inventory, error) {
	rows, err := r.connPool.Query(ctx,
		`SELECT
		id,
		name,
		description,
		campus_id,
		category_id,
		period,
		permission 
	FROM inventory
		WHERE campus_id = $1`, campusID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []models.Inventory
	for rows.Next() {
		var inventory models.Inventory
		err := rows.Scan(
			&inventory.ID,
			&inventory.Name,
			&inventory.Description,
			&inventory.CampusID,
			&inventory.CategoryID,
			&inventory.Period,
			&inventory.Permission,
		)
		if err != nil {
			return nil, err
		}
		inventories = append(inventories, inventory)
	}
	return inventories, nil

}

func (r root) UpdateInventory(ctx context.Context, inventory *models.Inventory) error {
	_, err := r.connPool.Exec(ctx,
		`UPDATE
		inventory SET
		name = $1,
		description = $2,
		campus_id = $3,
		category_id = $4,
		period = $5,
		permission = $6,
		updated_at = now()
			WHERE id = $7`,
		inventory.Name,
		inventory.Description,
		inventory.CampusID,
		inventory.CategoryID,
		inventory.Period,
		inventory.Permission,
		inventory.ID)

	return err
}

func (r root) DeleteInventory(ctx context.Context, ID int) error {
	_, err := r.connPool.Exec(ctx, "DELETE FROM inventory WHERE id = $1", ID)
	return err
}

func (r root) CreateInventory(ctx context.Context, inventory *models.Inventory) error {
	exec, err := r.connPool.Exec(ctx,
		`SELECT id FROM inventory WHERE id = $1`, inventory.ID)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 0 {
		return nil
	}

	_, err = r.connPool.Exec(ctx,
		`INSERT INTO inventory(
		name,
		description,
		campus_id,
		category_id,
		period,
		permission
		) VALUES($1, $2, $3, $4, $5, $6)`,
		inventory.Name,
		inventory.Description,
		inventory.CampusID,
		inventory.CategoryID,
		inventory.Period,
		inventory.Permission,
	)

	return err
}
func (r root) GetInventoryByID(ctx context.Context, ID int) (*models.Inventory, error) {
	rows, err := r.connPool.Query(ctx,
		`SELECT
			id,
			name,
			description,
			campus_id,
			category_id,
			period,
			permission
		FROM inventory
			WHERE id = $1`,
		ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory = models.Inventory{}
	if rows.Next() {
		err = rows.Scan(
			&inventory.ID,
			&inventory.Name,
			&inventory.Description,
			&inventory.CampusID,
			&inventory.CategoryID,
			&inventory.Period,
			&inventory.Permission,
		)
		if err != nil {
			return nil, err
		}
	}

	return &inventory, nil
}

func (r root) GetAllPlacesByCampusIDAndCategoryID(ctx context.Context, CampusID int, CategoryID int) ([]models.Places, error) {
	rows, err := r.connPool.Query(ctx, `SELECT 
      id,
      name, 
      description, 
      campus_id, 
      category_id, 
      floor,
      room,
      period,
      permission
        FROM places
          WHERE campus_id = $1 AND category_id = $2`, CampusID, CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []models.Places

	for rows.Next() {
		var place = models.Places{}
		err = rows.Scan(&place.ID,
			&place.Name,
			&place.Description,
			&place.CampusID,
			&place.CategoryID,
			&place.Floor,
			&place.Room,
			&place.Period,
			&place.Permission,
		)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}

func (r root) GetAllInventoryByCampusIDAndCategoryID(ctx context.Context, CampusID int, CategoryID int) ([]models.Inventory, error) {
	rows, err := r.connPool.Query(ctx, `SELECT 
      id,
      name, 
      description, 
      campus_id,
      category_id,
      period,
      permission
        FROM inventory
          WHERE campus_id = $1 AND category_id = $2`, CampusID, CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []models.Inventory
	for rows.Next() {
		var inventory = models.Inventory{}

		err = rows.Scan(&inventory.ID,
			&inventory.Name,
			&inventory.Description,
			&inventory.CampusID,
			&inventory.CategoryID,
			&inventory.Period,
			&inventory.Permission,
		)
		if err != nil {
			return nil, err
		}

		inventories = append(inventories, inventory)
	}
	return inventories, nil
}
