package beaconwrapper

import (
	"encoding/json"
	"log"
	"testing"
)

func TestUnmarshallStringIntoPathArray(t *testing.T) {
	tt := []struct {
		body       string
		shouldPass bool
	}{
		{`[{"beaconID":1003,"Name":"B03","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522079,"Longitude":-0.079650,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1003,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"},{"beaconID":1004,"Name":"B04","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522059,"Longitude":-0.079514,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1004,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"Phone Booths North","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"},{"beaconID":1005,"Name":"B05","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522053,"Longitude":-0.079478,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1005,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"},{"beaconID":1006,"Name":"B06","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522030,"Longitude":-0.079347,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1006,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"}]`, true},
		{`[{"beaconID" : 100}]`, true},
		{`[{"beacon" : 100}, {"beacon": 100}]`, true},
	}
	for _, tc := range tt {
		t.Run(tc.body, func(t *testing.T) {
			f := &[]Beacon{}
			err := json.Unmarshal([]byte(string(tc.body)), f)
			if tc.shouldPass == true && err != nil {
				t.Error(err)
			}

			if !tc.shouldPass && err == nil {
				t.Error(err)
			}
			log.Println(f)
		})
	}
}

func TestGeneratePairs(t *testing.T) {
	f := &Frame{}
	str := `[{"beaconID":1003,"Name":"B03","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522079,"Longitude":-0.079650,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1003,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"},{"beaconID":1004,"Name":"B04","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522059,"Longitude":-0.079514,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1004,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"Phone Booths North","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"},{"beaconID":1005,"Name":"B05","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522053,"Longitude":-0.079478,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1005,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"},{"beaconID":1006,"Name":"B06","OperatingMode":"Duel","StationID":"9400ZZLUAMA","Latitude":51.522030,"Longitude":-0.079347,"Level":2.0,"UUID":"3f846818-1f85-4f93-a46d-039c3056a4b2 ","Major":900,"Minor":1006,"Namespace":"TBC ","Instance":"TBC ","WaypointType":"","Manufacturer":"TBC","Model":"TBC","BatteryLevel":1,"TransmissionPower":1,"Accuracy":0.0,"Interval":"300"}]`
	err := json.Unmarshal([]byte(string(str)), f)
	if err != nil {
		log.Println("Error", err)
		t.Fail()
	}

	_ = f.GeneratePairs()
}

func TestLookupPath(t *testing.T) {
	tt := []struct {
		source     Beacon
		target     Beacon
		shouldPass bool
	}{
		{Beacon{BeaconID: 1001}, Beacon{BeaconID: 1002}, true},
		{Beacon{BeaconID: 1002}, Beacon{BeaconID: 1001}, true},
		{Beacon{BeaconID: 1009}, Beacon{BeaconID: 1001}, false},
	}
	for _, tc := range tt {
		_, err := lookupPath(tc.source, tc.target)
		if tc.shouldPass == true && err != nil {
			t.Fail()
		}
		if !tc.shouldPass && err == nil {
			t.Fail()
		}
	}
}
