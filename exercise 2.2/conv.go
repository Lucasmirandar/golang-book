//conv converte seu argumento numérico para Celsius, Fahrenheit, Pés, Metros, Quilos e Libras.
package main

import (
	"fmt"
	"golang-book/exercise2.1/tempconv"
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
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		m := Metros(t)
		p := Pes(t)
		q := Quilos(t)
		l := Libras(t)
		fmt.Printf("%s = %s, %s = %s\n", m, PToM(p), p, MToP(m))
		fmt.Printf("%s = %s, %s = %s\n", q, LToQ(l), l, QToL(q))
	}

}

type Metros float64
type Pes float64
type Quilos float64
type Libras float64

func (m Metros) String() string { return fmt.Sprintf("%g m", m) }
func (p Pes) String() string    { return fmt.Sprintf("%g p", p) }
func (q Quilos) String() string { return fmt.Sprintf("%g q", q) }
func (l Libras) String() string { return fmt.Sprintf("%q l", l) }

//MToP transforma Metros em Pés
func MToP(m Metros) Pes { return Pes(m * 3.281) }

//PToM transforma Pés em Metros
func PToM(p Pes) Metros { return Metros(p / 3.281) }

//QToL transforma Quilos em Libras
func QToL(q Quilos) Libras { return Libras(q * 2.205) }

//LToQ transforma Libras em Quilos
func LToQ(l Libras) Quilos { return Quilos(l / 2.205) }
