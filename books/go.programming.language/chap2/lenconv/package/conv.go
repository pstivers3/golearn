package lenconv

// CmToIn converts centimeters to inches
func CmToIn(cm Cm) In {
	return In(cm/2.54)
}

// InToCm converts inches to centimeters 
func InToCm(in In) Cm {
	return Cm(in*2.54)
}
