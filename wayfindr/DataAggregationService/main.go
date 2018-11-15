package main

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/mattnolf/wayfindr/camerawrapper"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

const cfgName string = ".cfg"

func main() {
	SetupConfig()
	Watcher()
	//WebServer()
}

// Watcher tracks changes in data
func Watcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return
	}
	defer watcher.Close()

	if err := watcher.Add(viper.GetString("image")); err != nil {
		log.Println(err)
		return
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:

				jsonFile, err := os.Open(viper.GetString("image"))
				// if we os.Open returns an error then handle it
				if err != nil {
					log.Println(err)
				}

				byteValue, err := ioutil.ReadAll(jsonFile)
				if err != nil {
					log.Println("log", err)
				}

				log.Println(string(byteValue[:]))

				keys := make([]camerawrapper.Frame, 1)

				err = json.Unmarshal([]byte(byteValue), &keys)
				if err != nil {
					log.Println("unmarshall", err)
				}
				for _, val := range keys {
					log.Println(val)

				}
				_ = event

			case err := <-watcher.Errors:
				log.Println("watcher", err)
			}
		}
	}()
	<-done
	return
}

// SetupConfig loads configuration file and watched for modifications
func SetupConfig() {
	viper.SetConfigName(cfgName)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to find configuration file")
	}
	log.Println("Successfully found configuration file " + cfgName)
}
