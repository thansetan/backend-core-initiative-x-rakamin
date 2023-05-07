package main

import (
	"self-payroll/config"
	"sync"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Infof(".env is not loaded properly")
	// }
	// log.Infof("read .env from file")

	config := config.NewConfig()
	server := initServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
