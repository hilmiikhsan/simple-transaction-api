package configuration

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(config Config) *gorm.DB {
	username := config.Get("DATASOURCE_USERNAME")
	password := config.Get("DATASOURCE_PASSWORD")
	host := config.Get("DATASOURCE_HOST")
	port := config.Get("DATASOURCE_PORT")
	dbName := config.Get("DATASOURCE_DB_NAME")
	maxPoolOpen, err := strconv.Atoi(config.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPoolIdle, err := strconv.Atoi(config.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPollLifeTime, err := strconv.Atoi(config.Get("DATASOURCE_POOL_LIFE_TIME"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(username+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?parseTime=true"), &gorm.Config{
		Logger: loggerDb,
	})
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	err = db.AutoMigrate(&entity.Customer{})
	err = db.AutoMigrate(&entity.CustomerAddress{})
	err = db.AutoMigrate(&entity.Product{})
	err = db.AutoMigrate(&entity.PaymentMethod{})
	err = db.AutoMigrate(&entity.Order{})
	err = db.AutoMigrate(&entity.OrderProduct{})
	err = db.AutoMigrate(&entity.OrderPayment{})
	if err != nil {
		log.Fatal("Error migrating database")
	}

	return db
}
