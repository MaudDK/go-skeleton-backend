FROM golang:1.22.0

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go mod tidy

RUN go build -o ./bin/app ./cmd/main.go

CMD ["./bin/app"]