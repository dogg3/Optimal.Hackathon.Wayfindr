package main

import (
	"encoding/json"
	"github.com/mattnolf/wayfindr/beaconwrapper"
	"github.com/mattnolf/wayfindr/camerawrapper"
	"io/ioutil"
	"log"
	"net/http"
)

// HandleCameraService does stuff
func HandleCameraService(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	b := &[]camerawrapper.Frame{}

	err = json.Unmarshal([]byte(string(body[:])), b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	/*
		for idx := range b {
			aggregate(b[i])
		}
	*/

}

// HandleBeaconService parses strings into beacon objects
func HandleBeaconService(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	b := &[]beaconwrapper.Beacon{}

	err = json.Unmarshal([]byte(string(body[:])), b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

}
