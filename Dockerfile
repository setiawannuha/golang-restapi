FROM golang:1.12.0-alpine3.9

RUN mkdir /go_apps
# Set the Current Working Directory inside the container
WORKDIR /go-apps

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
# Build the Go app
RUN go build -o ./go-apps .


# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the binary program produced by `go install`
CMD ["make build", "make run", "make migrate-mysql", "make migrate-pg", "make rollback-mysql", "make rollback-pg"]