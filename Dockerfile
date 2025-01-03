# Go lang image
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/exe app/main/main.go

EXPOSE 8080

CMD ["./bin/exe"]
