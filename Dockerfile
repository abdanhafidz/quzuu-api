# Gunakan image dasar Golang versi 1.21.6
FROM golang:1.21.6

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh kode
COPY . .

# Copy file .env (akan dibuat di workflow)
COPY .env.example .env

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]
