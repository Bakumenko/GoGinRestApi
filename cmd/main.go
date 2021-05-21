package main

import (
	"apiserver/pkg"
	"apiserver/pkg/handler"
	"apiserver/pkg/repository"
	"apiserver/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializing configs: %s", err.Error())
		return err
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: "db.host",
		Port: "db.port",
		Username: "db.username",
		DBName: "db.dbname",
		Password: "db.password",
		SSLMode: "db.sslmode",
	})
	if err != nil {
		log.Fatalf("error occured while connecting to db: %s", err.Error())
		return err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("web.port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
		return err
	}
	return nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
