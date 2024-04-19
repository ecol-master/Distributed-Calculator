FROM golang:alpine as builder 

WORKDIR /build

ADD backend/go.mod .

COPY backend .

RUN go build -o worker cmd/worker/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/worker /build/worker

# COPY backend/.env /build/.env

CMD ["./worker"]