/**
OBU: Onboard Unit

sits in trucks, and sends out gps per interval
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var sendInterval = time.Second

type OBUData struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

// // returns Lat, Long
// func genLocation() (float64, float64) {
// 	return genCoord(), genCoord()
// }

func genLatLong() (float64, float64) {
	n := float64(rand.Intn(100) + 1) // 1~101
	f := rand.Float64()              // 0.0~1.0
	return n + f
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	obuIDs := generateOBUIDs(20)
	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := genLatLong()
			data := OBUData{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}
			fmt.Printf("%+v\n", data)
		}
		time.Sleep(sendInterval)
	}
}
