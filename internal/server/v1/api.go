package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler{
	r:=chi.NewRouter()
	bk := &BookingRouter{
		Repository: &data.BookingRepository{
			Data: data.New()
		}
	}
	r.Mount("/bookings", bk.Routes())
	return r
}