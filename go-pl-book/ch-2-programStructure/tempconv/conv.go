package tempconv

//This file just contains the conversion operators

func KtoC(k Kelvin) Celcius {
	if k == 0 {
		return AbsoluteZeroK
	}
	return Celcius(k - 273)
}

func CtoK(c Celcius) Kelvin {
	if c == AbsoluteZeroK {
		return Kelvin(0)
	}
	return Kelvin(c + 273)
}
