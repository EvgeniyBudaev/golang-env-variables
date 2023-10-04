package main

import (
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", ".env", "path to config file in .env format")
}

type Config struct {
	Port string `env:"PORT"`
}

func main() {
	flag.Parse()
	var cfg Config
	fmt.Println("Before config: ", cfg)
	//err := env.Parse(&cfg)
	err := envconfig.Process("myapp", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("After config: ", cfg)
}