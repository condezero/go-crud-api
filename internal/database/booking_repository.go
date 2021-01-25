package database

import (
	"context"

	"github.com/condezero/go-crud-api/pkg/booking"
)

type BookingRepository struct {
	Data *Data
}

func (bk *BookingRepository) GetAll(ctx context.Context) ([]booking.Booking, error) {

	q := `select bookingId, price from bookings;`

	rows, err := bk.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookings []booking.Booking
	for rows.Next() {
		var b booking.Booking
		rows.Scan(&b.BookingId, &b.Price)
		bookings = append(bookings, b)
	}
	return bookings, nil
}
func (bk *BookingRepository) GetOne(ctx context.Context, bookingid uint) (booking.Booking, error) {

	q := `select bookingId, price from bookings where bookingId=$1;`

	row := bk.Data.DB.QueryRowContext(ctx, q, bookingid)

	var b booking.Booking
	err := row.Scan(&b.BookingId, &b.Price)
	if err != nil {
		return booking.Booking{}, err
	}

	return b, nil
}

func (bk *BookingRepository) Create(ctx context.Context, b *booking.Booking) error {

	q := `INSERT INTO bookings (bookingId, price) VALUES ($1, $2) RETURNING id;`

	row := bk.Data.DB.QueryRowContext(ctx, q, b.BookingId, b.Price)

	var id uint

	err := row.Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (bk *BookingRepository) Update(ctx context.Context, bookingid uint, b booking.Booking) error {

	q := `UPDATE bookings set price=$1 WHERE bookingId= $2`

	stmt, err := bk.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, b.Price, b.BookingId)

	if err != nil {
		return err
	}
	return nil

}

func (bk *BookingRepository) Delete(ctx context.Context, bookingid uint) error {

	q := `DELETE FROM bookings WHERE bookingId=$1`

	stmt, err := bk.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, bookingid)
	if err != nil {
		return err
	}
	return nil

}
