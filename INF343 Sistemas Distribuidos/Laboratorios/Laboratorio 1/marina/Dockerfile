FROM golang:1.24.2

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o marina .

EXPOSE 50052

CMD ["./marina"]