FROM golang:1.22.12
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build ./...
RUN go test -run '^$' ./...
CMD ["go", "run", "./cmd/server", "-addr", ":8080"]
