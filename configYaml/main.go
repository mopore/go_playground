package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MyConfig struct {
    // Remember to have the field name start with a capital letter
    ApiVersion string `yaml:"apiVersion"`
    Database struct {
        Host string `yaml:"host"` 
        Port int `yaml:"port"`
        User string `yaml:"user"`
        Password string `yaml:"password"`
        Name string `yaml:"name"`
    } `yaml:"database"`
}


func main() {
    content, err := os.ReadFile("config.yaml")
    if err != nil {
        log.Fatal(err)
    }

    var config MyConfig

    err = yaml.Unmarshal(content, &config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Config: %v\n", config)
}

