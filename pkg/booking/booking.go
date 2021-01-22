package booking

import "fmt"

type Booking struct {
	BookingId uint   `json:"bookingId,omitempty"`
	Price     Number `json:"price,omitempty"`
}
type Number float64

func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", n)), nil
}
