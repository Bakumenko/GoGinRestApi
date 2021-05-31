package main

import (
	"apiserver/pkg"
	"apiserver/pkg/handler"
	"apiserver/pkg/repository"
	"apiserver/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"os"
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

	if err := gotenv.Load(); err != nil {
		log.Fatalf("error occured while loading end file: %s", err.Error())
		return err
	}

	dbConfig := pkg.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	db, err := pkg.NewPostgresDB(dbConfig)
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
