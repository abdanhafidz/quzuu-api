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

# Buat file .env dengan variabel environment yang dibutuhkan
RUN echo "DB_HOST=db.chbzgcdzhplyetsizwui.supabase.co" >> .env && \
    echo "DB_USER=postgres" >> .env && \
    echo "DB_PASSWORD=QuzuuAPIDEV2025" >> .env && \
    echo "DB_PORT=5432" >> .env && \
    echo "DB_NAME=postgres" >> .env && \
    echo "SALT=@quzuuAp12025" >> .env

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]
