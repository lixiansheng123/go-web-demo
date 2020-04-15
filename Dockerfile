FROM golang:latest
COPY . /
WORKDIR /
RUN go build -o app main.go
ENTRYPOINT ["./app"]
