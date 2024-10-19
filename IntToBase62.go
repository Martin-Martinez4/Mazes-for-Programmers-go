package main

// all digits 0-9
// then A-Z
// then a-z

func IntToBase62(number int) rune {
	if number >= 0 && number <= 9 {
		return rune(number + 48)

	} else if number <= 35 {
		// 65 = A
		//number + 55
		return rune(number + 55)

	} else if number < 62 {
		return rune(number + 61)

	} else {
		panic("Illegal conversion; number too big")
	}
}
