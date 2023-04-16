FROM golang:1.20.3-alpine3.17
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .
CMD ["/app/main"]