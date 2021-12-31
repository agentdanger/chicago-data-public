FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o main .
RUN go get github.com/lib/pq
CMD ["/app/main"]