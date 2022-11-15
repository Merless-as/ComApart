package main

import "fmt"

const (
	CostGaz = 7.04
	CostLight = 4.36
	CostWater = 26.90
	Other = domPhone + Remont + JKX
	domPhone = 150
	Remont = 260
	JKX = 612
	OtherForInter = 1267
)

func main() {
	var oldG, oldL, oldW int
	var g, l, w int
	fmt.Println("Введите предыдущие показания в следующем порядке: газ, эл, вода")
	fmt.Scan(&oldG, &oldL, &oldW)
	fmt.Println("Введите показания в следующем порядке: газ, эл, вода")
	fmt.Scan(&g, &l, &w)
	gaz := float64(g-oldG) * CostGaz
	light := float64(l-oldL) * CostLight
	water := float64(w-oldW) * CostWater
	fmt.Println("Коммуналка получается: ", int(gaz+light+water)+Other)
}
