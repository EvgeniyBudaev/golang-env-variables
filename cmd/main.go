package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type Config struct {
	Port string `envconfig:"PORT"`
}

func main() {
	var cfg Config
	fmt.Println("Before config: ", cfg)
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	port := os.Getenv("PORT")
	//err := env.Parse(&cfg)
	err := envconfig.Process("MYAPP", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("port: ", port)
	fmt.Println("After config: ", cfg)
}
