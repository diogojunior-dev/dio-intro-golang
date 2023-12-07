package main

import "fmt"

const EbulicaoDaAguaEmKelvin = 373

func kelvinToCelsius(tempK float64) float64 {
	return tempK - 273
}

func main() {
	fmt.Printf("O ponto de ebulição da água em Celsius é %v\n", kelvinToCelsius(EbulicaoDaAguaEmKelvin))
}
