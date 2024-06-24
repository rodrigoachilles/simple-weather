FROM golang:1.22.4 as build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s' -o simple-weather ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/simple-weather .
ENTRYPOINT ["./simple-weather"]
