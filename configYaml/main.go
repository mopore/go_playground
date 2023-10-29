package main

import (
    "fmt"
    "log"
    "os"

    "gopkg.in/yaml.v3"
    "github.com/spf13/viper"
)


func readWithViper() {
    viper.SetConfigFile("config.yaml")
    viper.ReadInConfig()
    host := viper.GetString("database.host")
    log.Printf("Host from viper: %s\n", host)
}


func readWithStruct() {
    content, err := os.ReadFile("config.yaml")
    if err != nil {
        log.Fatal(err)
    }

    var config struct {
        ApiVersion string `yaml:"apiVersion"`
        Database struct {
            Host string `yaml:"host"` 
            Port int `yaml:"port"`
            User string `yaml:"user"`
            Password string `yaml:"password"`
            Name string `yaml:"name"`
        } `yaml:"database"`
    }

    err = yaml.Unmarshal(content, &config)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Host from scruct: %v\n", config.Database.Host)
}


func main() {
    readWithViper()
    fmt.Println("--------------------------------------------------")
    readWithStruct()
}
