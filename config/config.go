package config

import (
	"os"
	"self-payroll/config/postgres"
	"self-payroll/model"
	"strconv"

	"gorm.io/gorm"
)

type (
	config struct {
	}

	Config interface {
		ServiceName() string
		ServicePort() int
		ServiceEnvironment() string
		Database() *gorm.DB
		CreateAdminPosition()
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.InitGorm()
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (c *config) ServicePort() int {
	v := os.Getenv("APP_PORT")
	port, _ := strconv.Atoi(v)

	return port
}

func (c *config) ServiceEnvironment() string {
	return os.Getenv("ENV")
}

func (c *config) CreateAdminPosition() {
	c.Database().Create(&model.Position{
		ID:     69420,
		Name:   "Admin",
		Salary: 0,
	})
}
