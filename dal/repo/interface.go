package repo

import (
	"context"

	"boobot/kernel/domain/models"
)

type User interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetAllByCampus(ctx context.Context, campus string) ([]models.User, error)
	GetByNickname(ctx context.Context, nickname string) (*models.User, error)

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
	GetByID(ctx context.Context, id int) (*models.User, error)

	//GetAllByCampus(ctx context.Context, campus string) ([]models.User, error)
	//GetByNickname(ctx context.Context, nickname string) (*models.User, error)

	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, ID int) error

	ExistByID(ctx context.Context, id int) (bool, error)
	ExistByNickname(ctx context.Context, nickname string) (bool, error)
}

type Root interface {
	GetAllCampuses(ctx context.Context) ([]models.Campus, error)
	CreateCategory(ctx context.Context, Category string) error
	GetAllCategory(ctx context.Context) ([]models.Category, error)
	DeleteCategory(ctx context.Context, name string) error
}

//CREATE get update delete
