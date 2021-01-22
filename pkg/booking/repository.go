package booking

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Booking, error)
	GetOne(ctx context.Context, bookingId uint) (Booking, error)
	Create(ctx context.Context, booking *Booking) error
	Update(ctx context.Context, bookingId uint, booking Booking) error
	Delete(ctx context.Context, bookingId uint) error
}
