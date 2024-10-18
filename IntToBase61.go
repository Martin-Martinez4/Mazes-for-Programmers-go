package main

// all digits 0-9
// then A-Z
// then a-z

func IntToBase61(number int) rune {
	if number > 0 && number <= 9 {
		return rune(number)

	} else if number < 62 {
		// 65 = A
		//number + 55
		return rune(number + 55)
	} else {
		panic("Illegal conversion; number too big")
	}
}
