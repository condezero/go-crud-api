CREATE TABLE IF NOT EXISTS bookings (
    id serial NOT NULL,
    bookingId int NOT NULL,
    price numeric(3,2) NOT NULL,
    CONSTRAINT pk_booking PRIMARY KEY (id)
)