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
	Port string `envconfig:"PORT"`
}

func main() {
	flag.Parse()
	var cfg Config
	fmt.Println("Before config: ", cfg)
	//err := env.Parse(&cfg)
	err := envconfig.Process("MYAPP", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Port:  %s\n"
	_, err = fmt.Printf(format, cfg.Port)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("After config: ", cfg)
}
