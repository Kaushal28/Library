# Go base image
FROM golang:1.17.2

WORKDIR /opt

COPY . .

# install dependencies and build!
RUN go get ./...
RUN go build

# start the application
CMD ["./library"]