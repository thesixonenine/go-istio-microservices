FROM golang:1.24.2-alpine
WORKDIR /app
COPY . .
COPY ../proto ../proto
RUN go mod tidy && go build -o server .
EXPOSE 50051
CMD ["./server"]
