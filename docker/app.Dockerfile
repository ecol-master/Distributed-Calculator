FROM golang:alpine as builder 

WORKDIR /build

ADD backend/go.mod .

COPY backend .

RUN go build -o app cmd/app/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/app /build/app

# COPY backend/.env /build/.env

CMD ["./app"]