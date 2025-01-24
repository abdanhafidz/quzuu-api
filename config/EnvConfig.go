package config

import "os"

var TCP_ADDRESS string

func init() {
	HOST_ADDRESS := os.Getenv("HOST_ADDRESS")
	HOST_PORT := os.Getenv("HOST_PORT")
	TCP_ADDRESS = HOST_ADDRESS + ":" + HOST_PORT
	// Menampilkan nilai variabel lingkungan
}
