package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func ConnectRedis() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	address := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	database, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	ctx := context.Background()

	// Konfigurasi koneksi Redis
	RDB := redis.NewClient(&redis.Options{
		Addr:     address,  // Ganti dengan alamat Redis
		Password: password, // Kosong jika tanpa password
		DB:       database, // Gunakan DB 0
	})

	// Coba ping Redis untuk memastikan koneksi berhasil
	pong, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
	fmt.Println("Redis connection successfully:", pong)
}
