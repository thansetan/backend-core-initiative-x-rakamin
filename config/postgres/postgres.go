package postgres

import (
	"fmt"
	"os"
	"self-payroll/model"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_service := os.Getenv("DB_SERVICE")
	db_port := os.Getenv("DB_PORT")
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db_service, db_user, db_pass, db_name, db_port)
	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}
	db.AutoMigrate(&model.Position{}, &model.User{}, &model.Company{}, &model.Transaction{})
	return db
}
