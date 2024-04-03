# Build stage
FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o bot ./cmd/bot/main.go

## Production stage
FROM golang:1.22 AS final

COPY --from=build /app/bot .

ENTRYPOINT [ "./bot" ]

CMD [ "run" ]