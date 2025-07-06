package database

import (
	"fmt"
	"log"

	"github.com/elghazx/anime-crud-api/internal/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(conf config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.Pass,
		conf.User,
		conf.Name,
		conf.Tz,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get *sql.DB object:", err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	return db
}
