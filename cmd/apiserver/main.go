package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/senivaser/BEonGo/internal/app/apiserver"
	"github.com/senivaser/BEonGo/internal/app/model"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path_to_config_file")
}

func main() {
	flag.Parse()
	serverConfig := apiserver.NewConfig()
	_, errServer := toml.DecodeFile(configPath, serverConfig)
	if errServer != nil {
		log.Fatal(errServer)
	}

	storeConfig := model.NewConfig()
	_, errStore := toml.DecodeFile(configPath, storeConfig)
	if errStore != nil {
		log.Fatal(errStore)
	}

	store, _ := createStore(storeConfig)
	server := apiserver.New(serverConfig, store)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

func createStore(config *model.Config) (*model.Store, []error) {
	store, errors := model.NewStore(config)

	if len(errors) > 0 {
		fmt.Printf("Create Store Errors: %v", errors)
		return nil, errors
	}

	return store, errors
}
