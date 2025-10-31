FROM golang:1.24.2

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o submundo .

EXPOSE 50053

CMD ["./submundo"]