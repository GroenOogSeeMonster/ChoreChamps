# Dockerfile
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ChoreChamps .

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/ChoreChamps /app/ChoreChamps

EXPOSE 8282

CMD ["/app/ChoreChamps"]