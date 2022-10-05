package repo

import (
	"context"

	"boobot/kernel/domain/models"
)

type User interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetAllByCampus(ctx context.Context, campus string) ([]models.User, error)
	GetByNickname(ctx context.Context, nickname string) (*models.User, error)
	ExistsUsersInCampusByID(ctx context.Context, ID int) (bool, error)

	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, ID int) error

	ExistByID(ctx context.Context, id int) (bool, error)
	ExistByNickname(ctx context.Context, nickname string) (bool, error)
}

type Session interface {
	GetByID(ctx context.Context, id int) (*models.Session, error)
	GetByUserID(ctx context.Context, userID int) (*models.Session, error)
	ExistsByCodeAndUserID(ctx context.Context, userID int, code int) (bool, error)
	Create(ctx context.Context, session *models.Session) error
	Delete(ctx context.Context, ID int) error
}

type Booking interface {
	GetByID(ctx context.Context, id int) (*models.Booking, error)
	Create(ctx context.Context, booking *models.Booking) error
	Update(ctx context.Context, booking *models.Booking) error
	Delete(ctx context.Context, ID int) error

	GetActiveBookings(ctx context.Context, booking *models.Booking) ([]models.Booking, error)
}

type Root interface {
	GetAllCampuses(ctx context.Context) ([]models.Campus, error)
	CreateCampus(ctx context.Context, category *models.Campus) error
	DeleteCampus(ctx context.Context, ID int) error
	ExistsCampusByID(ctx context.Context, ID int) (bool, error)
	//ExistsUsersInCampus(ctx context.Context) (bool,error)
	UpdateCampus(ctx context.Context, category *models.Campus) error

	GetAllInventoryByCampusID(ctx context.Context, campusID int) ([]models.Inventory, error)
	GetAllInventoryByCampusIDByRole(ctx context.Context, campusID int, role models.Role) ([]models.Inventory, error)
	UpdateInventory(ctx context.Context, inventory *models.Inventory) error
	DeleteInventory(ctx context.Context, ID int) error
	CreateInventory(ctx context.Context, inventory *models.Inventory) error
	GetInventoryByID(ctx context.Context, ID int) (*models.Inventory, error)

	GetAllCategoryByBookType(ctx context.Context, bookType models.BookType) ([]models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, ID int) error
	UpdateCategory(ctx context.Context, category *models.Category) error
	GetCategoryByID(ctx context.Context, ID int) (*models.Category, error)

	GetAllPlacesByCampusIDAndCategoryID(ctx context.Context, CampusID int, CategoryID int) ([]models.Places, error)
	GetAllPlacesByCampusIDAndCategoryIDAndRole(ctx context.Context, CampusID int, CategoryID int, role models.Role) ([]models.Places, error)
	GetAllInventoryByCampusIDAndCategoryID(ctx context.Context, CampusID int, CategoryID int) ([]models.Inventory, error)
	GetAllInventoryByCampusIDAndCategoryIDAndRole(ctx context.Context, CampusID int, CategoryID int, role models.Role) ([]models.Inventory, error)
}

//CREATE get update delete
