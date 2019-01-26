FROM golang:1.9

# Exposing default port
EXPOSE 3000

# Create app directory
WORKDIR /go/src/github.com/gottsohn/go-fun

# Copy files
COPY . .
RUN go get -v -t -d ./...
RUN go install -v ./...
RUN go test -v

CMD [ "go", "run", "main.go"]
