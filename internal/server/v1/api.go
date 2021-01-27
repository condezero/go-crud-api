package v1

import (
	"net/http"

	"github.com/condezero/go-crud-api/internal/database"
	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()
	bk := &BookingRouter{
		Repository: &database.BookingRepository{
			Data: database.New(),
		},
	}
	r.Mount("/bookings", bk.Routes())
	return r
}
