FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/github.com/condezero/gocrudapi/
WORKDIR /go/src/github.com/condezero/gocrudapi/
RUN go mod download
COPY . /go/src/github.com/condezero/gocrudapi/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/gocrudapi github.com/condezero/gocrudapi


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/condezero/gocrudapi/build/gocrudapi /usr/bin/gocrudapi

EXPOSE 5000 5000

ENTRYPOINT ["/usr/bin/gocrudapi"]
