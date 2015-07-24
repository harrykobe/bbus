package bbus

import (
	"github.com/harrykobe/bbus/vm"
	"encoding/json"
)

type JPoint struct {
	Point *JGeo `json:"point"`
	W []*JGeo `json:"W"`
	Type int `json:"type"`
}

type JPolyline struct {
	Ua *JBound `json:"Ua"`
	W []*JGeo `json:"W"`
	Type int `json:"type"`
}

type JBound struct {
	Uk *JGeo `json:"Uk"`
	Kl *JGeo `json:"kl"`
	Pe float64 `json:"pe"`
	Qe float64 `json:"qe"`
	Ue float64 `json:"ue"`
	Ve float64 `json:"ve"`
}

type JGeo struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func JDecodePolyline(geo string)(polyline JPolyline, err error){
	value, err := vm.VM.Run(`JSON.stringify(lb("` + geo + `", true))`)
	err = json.Unmarshal([]byte(value.String()), &polyline)
	return
}

func  JDecodePoint(geo string)(point JPoint, err error){
	value, err := vm.VM.Run(`JSON.stringify(lb("` + geo + `", true))`)
	err = json.Unmarshal([]byte(value.String()), &point)
	return
}