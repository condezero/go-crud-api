package v1

import (
	"github.com/go-chi/chi"
	"encoding/json"
	"net/http"
	"github.com/condezero/go-crud-api/pkg/booking"
)

type BookingRouter struct {
	Repository booking.Repository
}

func (br *BookingRouter) CreateHandler(w http.ResponseWriter, r *http.Request){
	var b booking.Booking
	
	err := json.NewDecoder(r.Body).Decode(&b)
	if err !=nil{
		respose.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx :=r.Context()
	err = br.Repository.Create(ctx, &b)
	if err != nil{
		respose.HTTPError(w,r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf(fmt.Sprintf("%s%d", r.URL.String(), b.BookingId))
    response.JSON(w,r,http.StatusCreated, response.Map{"booking": b})
}

func (bk *BookingRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	 ctx := r.Context()
	
	 bookings, err := bk.Repository.GetAll(ctx)
	 if err != nil {
		 response.HTTPError(w, r, http.StatusNotFound, err.Error())
	 }
	 response.JSON(w,r, http.StatusOK, response.Map("bookings": bookings))
}

func (bk *BookingRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	
	idStr := chi.URLParam(r, "bookingid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx :=r.Context()

	b, err:= bk.Repository.GetOne(ctx, uint(id))

	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
	}
	response.JSON(w,r, http.StatusOK, response.Map("booking": b))
}
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

func (bk *BookingRouter) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", ur.GetAllHandler)

    r.Post("/", ur.CreateHandler)

    r.Get("/{bookingid}", ur.GetOneHandler)

    r.Put("/{bookingid}", ur.UpdateHandler)

    r.Delete("/{bookingid}", ur.DeleteHandler)
	return r
}