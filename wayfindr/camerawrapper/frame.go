package camerawrapper

import (
	"math"
)

// Frame contains data from camera services
type Frame struct {
	Person          float64 `json: "person"`
	Entry           float64 `json: "entrytimestamp"`
	Exit            float64 `json: "exittime"`
	CameraID        string  `json: "camera"`
	SourceLatitude  float64 `json: "SourceLatitude"`
	SourceLongitude float64 `json: "SourceLongitude"`
	TargetLatitude  float64 `json: "TargetLatitude"`
	TargetLongitude float64 `json: "TargetLongitude"`
}

func (f *Frame) GetExit() float64 {
	return f.Exit
}

// GetCameraID does
func (f *Frame) GetCameraID() string {
	return f.CameraID
}

// GetSourceLatitude does
func (f *Frame) GetSourceLatitude() float64 {
	return f.SourceLatitude
}

// GetSourceLongitude does
func (f *Frame) GetSourceLongitude() float64 {
	return f.SourceLongitude
}

// GetTargetLatitude does
func (f *Frame) GetTargetLatitude() float64 {
	return f.TargetLatitude
}

// GetTargetLongitude does
func (f *Frame) GetTargetLongitude() float64 {
	return f.TargetLongitude
}

// GetID returns identifier
func (f *Frame) GetID() string {
	return f.CameraID
}

func (f *Frame) tif() float64 {
	return 1
}

// CalcDistance consumes the distance covered, within the time in frame
func (f *Frame) CalcDistance() float64 {
	return math.Acos(math.Sin(f.SourceLongitude*math.Pi/180)*math.Sin(f.TargetLatitude*math.Pi/180) +
		math.Cos(f.SourceLatitude*math.Pi/180)*math.Cos(f.TargetLatitude*math.Pi/180)*
			math.Cos(f.TargetLongitude*math.Pi/180-f.SourceLongitude*math.Pi/180)*6371000)
}

// CalcSpeed calculates the speed of the target in the frame
func (f *Frame) CalcSpeed() float64 {
	return f.CalcDistance() / f.tif()
}
