package bbus
import (
	"github.com/astaxie/beego/httplib"
	"fmt"
)

type BuslinesResult struct {
	Content []*BuslineResult `json:"content"`
}

type BuslineResult struct {
	Name string `json:"name"`
	Uid string `json:"uid"`
	TicketPrice int `json:"ticketPrice"`
	Geo string `json:"geo"`
	Company string `json:"company"`
	Stations []*StationResult `json:"stations"`
}

type StationResult struct {
	Geo string `json:"geo"`
	Name string `json:"name"`
	Uid string `json:"uid"`
}

type BusstopsResult struct {
	Content []*BusstopResult `json:"content"`
}

type BusstopResult struct {
	Blinfo []*BuslineResult `json:"blinfo"`
}

func(this *BuslineResult) GeoToPolyline()(polyline *Polyline, err error)  {
	polyline, err = DecodePolyline(this.Geo)
	return
}

func(this *StationResult) GeoToPoint()(point *Point, err error) {
	point, err = DecodePoint(this.Geo)
	return
}

func SeachBusLine(lineName string)(buslinesResult *BuslinesResult, err error){
	buslinesResult = &BuslinesResult{}

	var errCount int
	Here:
	req := httplib.Get("http://api.map.baidu.com/?qt=bl&c=257&wd=" + lineName)
	err = req.ToJson(&buslinesResult)
	if err != nil{
		fmt.Println("SeachBusLine:", lineName, err)
		if errCount > 4 {
			return
		}else{
			errCount = errCount + 1
			goto Here
		}
	}
	return
}

func GetBusLine(uid string)(buslinesResult *BuslinesResult, err error){
	buslinesResult = &BuslinesResult{}
	var errCount int
	Here:
	req := httplib.Get("http://api.map.baidu.com/?qt=bsl&c=257&uid=" + uid)
	err = req.ToJson(&buslinesResult)
	if err != nil{
		fmt.Println("GetBusLine:", uid, err)
		if errCount > 4 {
			return
		}else{
			errCount = errCount + 1
			goto Here
		}
	}
	return
}

func BusStopSeachBusLine(name string)(busstopsResult *BusstopsResult, err error){
	busstopsResult = &BusstopsResult{}
	var errCount int
	Here:
	req := httplib.Get("http://api.map.baidu.com/?qt=s&c=257&wd=" + name + "-公交车站")
	err = req.ToJson(&busstopsResult)
	if err != nil{
		fmt.Println("BusStopSeachBusLine:", name, err)
		if errCount > 4 {
			return
		}else{
			errCount = errCount + 1
			goto Here
		}
	}
	return
}