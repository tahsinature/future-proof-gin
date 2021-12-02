package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq" // import postgres
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*sql.DB
}

var (
	db1 *gorp.DbMap
	db2 *gorm.DB
)

// Init ...
func Init() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.User, config.DB.Password, config.DB.DBname)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", config.DB.Host, config.DB.User, config.DB.Password, config.DB.DBname)

	var err error
	db1, db2, err = ConnectDB(dbinfo, dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func SyncForce() {
	fmt.Println("SyncForce will be executed in 2 sec...")
	time.Sleep(time.Second * 2)
	db2.AutoMigrate(&models.User{}, &models.Article{})
}

func ConnectDB(dataSourceName string, dsn string) (db1 *gorp.DbMap, db2 *gorm.DB, err error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	// dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests

	db2, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return dbmap, db2, nil
}

func GetDB() (*gorp.DbMap, *gorm.DB) {
	return db1, db2
}

var RedisClient *_redis.Client

func InitRedis(selectDB ...int) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})
}

// GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}
