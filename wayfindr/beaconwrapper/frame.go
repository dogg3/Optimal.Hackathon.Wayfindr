package beaconwrapper

import (
	"log"
	"math"
	"time"
)

// Beacon stores data published from a beacon
type Beacon struct {
	BeaconID          int `json: "beaconID"`
	Name              string
	OperatingMode     string
	StationID         string
	Latitude          float64
	Longitude         float64
	Level             float64
	UUID              string
	Major             int
	Minor             int
	Namespace         string
	Instance          string
	WaypointType      string
	Manufacturer      string
	Model             string
	BatteryLevel      int
	TransmissionPower int
	Accuracy          float32
	Interval          string
	Timestamp         string
}

// Path represents a journey from one beacon to another
type Path struct {
	PathID       string
	StationID    string
	Source       string
	Target       string
	TravelTime   int
	Beginning    string
	Middle       string
	Ending       string
	StartingOnly string
}

// Frame contains data from camera services
type Frame struct {
	Beacon []Beacon
}

// GeneratePairs splits a received frame from a Beacon client into recognised
// pairs
// Returns:
// - []Path an array of recognised paths
// - []Beacon an array of beacons that could not be matched to a path
func (f *Frame) GeneratePairs() []Path {
	s := make([]Path, 2)
	for idx := range f.Beacon {
		p, err := lookupPath(f.Beacon[idx], f.Beacon[idx+1])
		if err != nil {
			s = append(s, p)
		}
	}
	return s
}

// CalcDistance consumes the distance covered, within the time in frame
func (f *Path) CalcDistance() float64 {

	sb, err := lookupBeacon(f.Source)
	if err != nil {
		log.Println(err)
		return 0
	}
	tb, err := lookupBeacon(f.Target)
	if err != nil {
		log.Println(err)
		return 0
	}

	return math.Acos(math.Sin(sb.Longitude*math.Pi/180)*math.Sin(tb.Latitude*math.Pi/180) +
		math.Cos(sb.Latitude*math.Pi/180)*math.Cos(tb.Latitude*math.Pi/180)*
			math.Cos(tb.Longitude*math.Pi/180-sb.Longitude*math.Pi/180)*6371000)
}

// CalcSpeed calculates the speed of the target in the frame
func (f *Path) CalcSpeed() float64 {
	sb, err := lookupBeacon(f.Source)
	if err != nil {
		log.Println(err)
		return 0
	}
	tb, err := lookupBeacon(f.Target)
	if err != nil {
		log.Println(err)
		return 0
	}

	layout := time.RFC3339
	ts, err := time.Parse(layout, sb.Timestamp)
	tt, err := time.Parse(layout, tb.Timestamp)

	return f.CalcDistance() / ts.Sub(tt).Seconds()
}

func lookupBeacon(id string) (Beacon, error) {
	return Beacon{}, nil
}

func lookupPath(arg0, arg1 Beacon) (Path, error) {

	return Path{}, nil
}
