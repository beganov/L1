package main

import "fmt"

var array []float64 = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.1, -0.1}

type Temperatures []float64

func main() {
	tempMap := make(map[int]Temperatures, 10)
	for _, i := range array {
		tempMap[(int(i) - int(i)%10)] = append(tempMap[(int(i)-int(i)%10)], i)
	}
	fmt.Println(tempMap)
}

func (t Temperatures) String() string {
	return fmt.Sprintf("%.1f", t)
}
