FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
RUN go get github.com/lib/pq
CMD ["/app/main"]