package bbus
import (
	"errors"
	"math"
	"strings"
)


type Geo struct {
	Lng float64
	Lat float64
}

type Point struct {
	Lng float64
	Lat float64
}

type GeoGroup struct {
	Geos []*Geo
	Type int
}

type Polyline struct {
	Bound []*Point
	Points []*Point
}

var offsetArray [6]float64 = [6]float64{1.289059486E7, 8362377.87, 5591021, 3481989.83, 1678043.12, 0}
var offsetNumArray [6][10]float64 = [6][10]float64{
	[10]float64{    1.410526172116255E-8,      8.98305509648872E-6,    -1.9939833816331,   200.9824383106796,  -187.2403703815547,      91.6087516669843, -23.38765649603339, 2.57121317296198, -0.03801003308653, 1.73379812E7},
	[10]float64{   -7.435856389565537E-9,      8.983055097726239E-6,   -0.78625201886289,  96.32687599759846,  -1.85204757529826,      -59.36935905485877, 47.40033549296737, -16.50741931063887, 2.28786674699375, 1.026014486E7},
	[10]float64{	-3.030883460898826E-8,      8.98305509983578E-6,     0.30071316287616,  59.74293618442277,   7.357984074871,        -25.38371002664745, 13.45380521110908, -3.29883767235584, 0.32710905363475, 6856817.37},
	[10]float64{   -1.981981304930552E-8,      8.983055099779535E-6,    0.03278182852591,  40.31678527705744,   0.65659298677277,      -4.44255534477492, 0.85341911805263, 0.12923347998204, -0.04625736007561, 4482777.06},
	[10]float64{    3.09191371068437E-9,       8.983055096812155E-6,    6.995724062E-5,    23.10934304144901,  -2.3663490511E-4,       -0.6321817810242, -0.00663494467273, 0.03430082397953, -0.00466043876332, 2555164.4},
	[10]float64{    2.890871144776878E-9,      8.983055095805407E-6,   -3.068298E-8,       7.47137025468032,   -3.53937994E-6,         -0.02145144861037, -1.234426596E-5, 1.0322952773E-4, -3.23890364E-6, 826088.5},
}

func DecodePoint(code string)(point *Point, err error){
	geo, _, err := CodeToGeo(code)
	if err != nil {
		return
	}
	point = geo[0].GeoToPoint()
	return
}

func DecodePolyline(code string)(polyline *Polyline, err error){
	polylines := strings.Split(code, "|")
	polyline = &Polyline{}
	e, f, g := polylines[0], polylines[1], polylines[2]
	geo, _, err := CodeToGeo(g)
	for _, v := range geo {
		polyline.Points = append(polyline.Points, v.GeoToPoint())
	}
	geo, _, err = CodeToGeo(e)
	polyline.Bound = append(polyline.Bound, geo[0].GeoToPoint())
	geo, _, err = CodeToGeo(f)
	polyline.Bound = append(polyline.Bound, geo[0].GeoToPoint())
	return
}

//A-Z a-z 0-9 + / : 0 - 63
func charToNum(char string)(charNum uint, err error){
	charNum = uint(char[0])
	if "A" <= char && char <= "Z" {
		charNum -= 65
	}else if "a" <= char && char <= "z" {
		charNum -= 71
	}else if "0" <= char && char <= "9" {
		charNum += 4
	}else if char == "+" {
		charNum = 62
	}else if char == "/" {
		charNum = 63
	}else{
		err = errors.New("无法解释字符串")
	}
	return
}

func decodeOffsetGeo(code string)(lng, lat int, err error){
	var i, j uint
	for i, j = 0, 0; i < 4; i++ {
		j, err = charToNum(code[i:i+1])
		if err != nil {
			return
		}
		lng += int(j << (6 * i))
		j, err = charToNum(code[i+4:i+5])
		if err != nil {
			return
		}
		lat += int(j << (6 * i))
	}
	if lng > 8388608{
		lng = 8388608 - lng
	}
	if lat > 8388608{
		lat = 8388608 - lat
	}
	return
}

func CodeToGeo(codeStr string)(geos []*Geo, geoType int, err error){
	switch codeStr[0:1] {
	case ".":
		geoType = 2
	case "-":
		geoType = 1
	case "*":
		geoType = 0
	default:
		geoType = -1
	}
	for codeStr, i, j := codeStr[1:], 0, len(codeStr) - 1; i < j; {
		switch codeStr[i:i+1] {
		case "=":
			if j - i < 13 {
				err = errors.New("无效字符串")
				return
			}
			var lng, lat uint64
			var tmp, k uint
			for codeStr:= codeStr[i: i + 13]; k < 6; k++ {
				tmp, err = charToNum(codeStr[k+1:k+2])
				if err != nil {
					return
				}
				lng += uint64(tmp) << (6 * k)
				tmp, err = charToNum(codeStr[k+7:k+8])
				if err != nil {
					return
				}
				lat += uint64(tmp) << (6 * k)
			}
			geos = append(geos, &Geo{
				Lng: float64(lng),
				Lat: float64(lat),
			})
			i += 13
		case ";":
			i++
		default:
			if j - i < 8 {
				err = errors.New("无效字符串")
				return
			}
			var lng, lat int
			lng, lat, err = decodeOffsetGeo(codeStr[i:i+8])
			if err != nil {
				return
			}
			geos = append(geos, &Geo{
				Lng: geos[len(geos) - 1].Lng + float64(lng),
				Lat: geos[len(geos) - 1].Lat + float64(lat),
			})
			i += 8
		}
	}
	for key, _ := range geos{
		geos[key].Lng /= 100
		geos[key].Lat /= 100
	}
	return
}

func(this *Geo) GeoToPoint()(point *Point){
	lng, lat := math.Abs(this.Lng), math.Abs(this.Lat)
	for key, _ := range offsetArray {
		if lat >= offsetArray[key] {
			c, d := offsetNumArray[key][0] + offsetNumArray[key][1] * lng, lat / offsetNumArray[key][9]
			d = offsetNumArray[key][2] + offsetNumArray[key][3] * d + offsetNumArray[key][4] * d * d + offsetNumArray[key][5] * d * d * d + offsetNumArray[key][6] * d * d * d * d + offsetNumArray[key][7] * d * d * d * d * d + offsetNumArray[key][8] * d * d * d * d * d * d
			if this.Lng < 0 {
				c *= -1
			}
			if this.Lat < 0 {
				d *= -1
			}
			point = &Point{
				Lng: c,
				Lat: d,
			}
			break
		}
	}
	return
}
