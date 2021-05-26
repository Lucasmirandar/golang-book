//Pacote tempconv realiza conversões de Celsius e Fahrenheit.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g ºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g ºF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }

//CToF converte uma temperatura em Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converte uma temperatuda em Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//CToK
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.0) }

//KToC
func KToC(k Kelvin) Celsius { return Celsius(k + 273) }

//FToK
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9) + 273.372 }

//KtoF
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k-273.15)*9/5 + 32 }
