package tempconv

import "fmt"

type Celcius float64
type Kelvin float64

const (
	AbsoluteZeroK Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func (c Celcius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
