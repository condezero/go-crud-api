FROM golang:1.15.7-alpine3.13 as builder

COPY go.mod go.sum /go/src/github.com/condezero/gocrudapi/
WORKDIR /go/src/github.com/condezero/gocrudapi/
RUN go mod download
COPY . /go/src/github.com/condezero/gocrudapi/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/gocrudapi /go/src/github.com/condezero/gocrudapi


FROM alpine

#RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/condezero/gocrudapi/build/gocrudapi /usr/bin/gocrudapi

EXPOSE 5001 5001

ENTRYPOINT ["/usr/bin/gocrudapi"]
