package repo

import (
	"context"

	"boobot/kernel/domain/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type booking struct {
	connPool *pgxpool.Pool
}

func NewBooking(connPool *pgxpool.Pool) Booking {
	return &booking{
		connPool: connPool,
	}
}

func (b booking) GetByID(ctx context.Context, ID int) (*models.Booking, error) {
	rows, err := b.connPool.Query(ctx,
		`SELECT
			id,
			user_id,
			type,
			inventory_id,
			places_id,
			confirm,
			status,
			start_at,
			end_at
		FROM booking
			WHERE ID = $1`,
		ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var booking = models.Booking{}
	if rows.Next() {
		err = rows.Scan(
			&booking.ID,
			&booking.UserID,
			&booking.BookType,
			&booking.PlacesID,
			&booking.Confirm,
			&booking.Status,
			&booking.StartAt,
			&booking.EndAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return &booking, nil
}

func (b booking) Create(ctx context.Context, booking *models.Booking) error {
	_, err := b.connPool.Exec(ctx,
		`INSERT INTO bookings(
		user_id,
		type,
		inventory_id,
		places_id,
		confirm,
		status,
		start_at,
		end_at
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		booking.UserID,
		booking.BookType,
		booking.InventoryID,
		booking.PlacesID,
		booking.Confirm,
		booking.Status,
		booking.StartAt,
		booking.EndAt,
	)

	return err
}

func (b booking) Update(ctx context.Context, booking *models.Booking) error {
	_, err := b.connPool.Exec(ctx,
		`UPDATE
		booking SET
		user_id = $1,
		type = $2,
		inventory_id = $3,
		places_id = $4,
		confirm = $5,
		status = $6,
		start_at = $7,
		end_at = $8
			WHERE id = $9`,
		booking.UserID,
		booking.BookType,
		booking.InventoryID,
		booking.PlacesID,
		booking.Confirm,
		booking.Status,
		booking.StartAt,
		booking.EndAt,
	)

	return err
}

func (b booking) Delete(ctx context.Context, ID int) error {
	_, err := b.connPool.Exec(ctx, "DELETE FROM booking WHERE id = $1", ID)
	return err
}

func (b booking) GetActiveBookings(ctx context.Context, booking *models.Booking) ([]models.Booking, error) {
	rows, err := b.connPool.Query(ctx,
		`SELECT
		id,
		user_id,
		type,
		inventory_id,
		places_id,
		confirm,
		status,
		start_at,
		end_at
			FROM bookings WHERE
		start_at > $1 AND end_at > now() AND end_at < $2 AND (inventory_id = $3 OR places_id = $4)`,
		booking.StartAt, booking.StartAt.AddDate(0, 0, 1), booking.InventoryID, booking.PlacesID)
	if err != nil {
		return nil, err
	}

	var bookings []models.Booking
	for rows.Next() {
		var bkg models.Booking
		err := rows.Scan(
			&bkg.ID,
			&bkg.UserID,
			&bkg.BookType,
			&bkg.InventoryID,
			&bkg.PlacesID,
			&bkg.Confirm,
			&bkg.Status,
			&bkg.StartAt,
			&bkg.EndAt,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, bkg)
	}
	return bookings, nil
}
