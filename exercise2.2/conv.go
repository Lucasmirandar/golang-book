//conv converte seu argumento numérico para Celsius, Fahrenheit, Kelvin, Pés, Metros, Quilos e Libras.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		k := Kelvin(t)
		fmt.Printf("%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n%s = %s\n", f, FToC(f), c, CToF(c), k, KToC(k), c, CToK(c), k, KToF(k), f, FToK(f))
		m := Metros(t)
		p := Pes(t)
		q := Quilos(t)
		l := Libras(t)
		fmt.Printf("%s = %s\n%s = %s\n", m, MToP(m), p, PToM(p))
		fmt.Printf("%s = %s\n%s = %s\n", q, QToL(q), l, LToQ(l))
	}

}

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Metros float64
type Pes float64
type Quilos float64
type Libras float64

func (m Metros) String() string     { return fmt.Sprintf("%g m", m) }
func (p Pes) String() string        { return fmt.Sprintf("%g Ft", p) }
func (q Quilos) String() string     { return fmt.Sprintf("%g Kg", q) }
func (l Libras) String() string     { return fmt.Sprintf("%g Lb", l) }
func (c Celsius) String() string    { return fmt.Sprintf("%g ºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g ºF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }

//CToF converte uma temperatura em Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converte uma temperatuda em Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//CToK converte uma temperatura em Celsius para Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.0) }

//KToC converte uma temperatuda em Kelvin para Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + 273) }

//FToK converte uma temperatuda em Fahrenheit para Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9) + 273.372 }

//KtoF converte uma temperatuda em Kelvin para Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k-273.15)*9/5 + 32 }

//MToP transforma Metros em Pés
func MToP(m Metros) Pes { return Pes(m * 3.281) }

//PToM transforma Pés em Metros
func PToM(p Pes) Metros { return Metros(p / 3.281) }

//QToL transforma Quilos em Libras
func QToL(q Quilos) Libras { return Libras(q * 2.205) }

//LToQ transforma Libras em Quilos
func LToQ(l Libras) Quilos { return Quilos(l / 2.205) }
