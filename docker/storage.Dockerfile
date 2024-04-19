FROM golang:alpine as builder 

RUN apk add build-base

WORKDIR /build

ADD backend/go.mod .

COPY backend .

RUN CGO_ENABLED=1 go build -o storage cmd/storage/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/storage /build/storage

# COPY backend/.env /build/.env

CMD ["./storage"]