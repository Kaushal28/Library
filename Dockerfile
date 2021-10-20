# Go base image
FROM golang:1.17.2

WORKDIR /opt

COPY . .

RUN go get ./...
RUN go build

CMD ["./library"]