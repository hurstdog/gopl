// Package tempconf performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celcius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheiht) Celcius { return Celcius((f - 32) * 5 / 9) }
