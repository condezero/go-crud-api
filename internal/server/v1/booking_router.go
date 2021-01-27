package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/condezero/go-crud-api/pkg/booking"
	"github.com/condezero/go-crud-api/pkg/response"
	"github.com/go-chi/chi"
)

// BookingRouter instances Repository
type BookingRouter struct {
	Repository booking.Repository
}

// CreateHandler uses for create a new booking
func (br *BookingRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var b booking.Booking

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = br.Repository.Create(ctx, &b)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), b.BookingId))
	response.JSON(w, r, http.StatusCreated, response.Map{"booking": b})
}

// GetAllHandler Retrieves all rows
func (bk *BookingRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bookings, err := bk.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
	}
	response.JSON(w, r, http.StatusOK, response.Map{"bookings": bookings})
}

// GetOneHandler Retrieve one booking object
func (bk *BookingRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "bookingid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()

	b, err := bk.Repository.GetOne(ctx, uint(id))

	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
	}
	response.JSON(w, r, http.StatusOK, response.Map{"booking": b})
}

// UpdateHandler  Updates price of a booking
func (bk *BookingRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "bookingid")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var b booking.Booking
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = bk.Repository.Update(ctx, uint(id), b)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

// DEleteHandler Deletes a booking
func (bk *BookingRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "bookingid")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = bk.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

// Routes Configure all routes on chi
func (br *BookingRouter) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", br.GetAllHandler)

	r.Post("/", br.CreateHandler)

	r.Get("/{bookingid}", br.GetOneHandler)

	r.Put("/{bookingid}", br.UpdateHandler)

	r.Delete("/{bookingid}", br.DeleteHandler)
	return r
}
