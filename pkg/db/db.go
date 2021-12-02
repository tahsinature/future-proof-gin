package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.DBname,
		config.DB.Port)

	var err error
	db, err = ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func SyncForce() {
	db.AutoMigrate(&models.User{}, &models.Article{})
	fmt.Println("SyncForce Done...")
}

func ConnectDB(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("DB Connected...")
	return db, nil
}

func GetDB() (*gorp.DbMap, *gorm.DB) {
	return nil, db
}

var RedisClient *_redis.Client

func InitRedis() {
	host := config.Redis.Host
	password := config.Redis.Password
	db, _ := strconv.ParseInt(config.Redis.DB, 10, 8)

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     host,
		Password: password,
		DB:       int(db),
	})

	response := RedisClient.Ping()
	if response.Err() != nil {
		log.Panicf("Redis Connect Error: %s", response.Err())
	}

	fmt.Println("Redis Connected...")
}

func GetRedis() *_redis.Client {
	return RedisClient
}
