package main

import "fmt"

func main() {
	F := 32.0  //Fahrenheit
	C := 12.0  //Celsius
	K := 273.0 //Kelvin

	//Converter Fahrenheit em Celsius
	var fahrenheitToCelsius = (F - 32) * 5 / 9
	var celsiusTofahrenheit = (C * 9 / 5) + 32
	var kelvintoCelsius = K - 273
	var celsiusToKelvin = C + 273

	fmt.Printf("A conversão de temperatura de Fahrenheit %g para Celsius %g é %g\n", F, C, fahrenheitToCelsius)
	fmt.Printf("A conversão de temperatura de Celsius %g para Fahrenheit %g é %g\n", C, F, celsiusTofahrenheit)
	fmt.Printf("A conversão de temperatura de Kelvin %g para Celsius %g é %g\n", K, C, kelvintoCelsius)
	fmt.Printf("A conversão de temperatura de Celsius %g para Kelvin %g é %g\n", C, K, celsiusToKelvin)
}
