package main

import (
	"log"
	_ "time"
)

var (
	m map[string]float64
)

const (
	percentileChange = 10
	timeChange       = 10
)

func aggregate(w Wrapper) {
	log.Println("Speed: ", w.CalcSpeed())

	then, err := time.Parse("timeFormat", w.Timestamp())
	if err != nil {
		log.Println("Failed to aggregate data")
		return
	}
	duration := time.Since(then)

	if m[w.ID] == nil {
		m[w.ID] = w.CalcSpeed()
	}
	if w.CalcSpeed()/m[w.ID()]*100 > percentileChange && duration.Seconds() < timeChange {
		m[w.ID] = m[w.ID]
		commitUpdate()
	}

	commitUpdate()

}

func commitUpdate() {

}
