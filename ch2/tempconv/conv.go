package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celcius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// Kelvin -> Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }

// Kelvin -> Fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
