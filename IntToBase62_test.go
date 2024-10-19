package main

import "testing"

func TestIntToBase62(t *testing.T) {

	if IntToBase62(0) != '0' {
		t.Errorf("Expected to be equal to '0'; got: %v", IntToBase62(9))
	}

	if IntToBase62(9) != '9' {
		t.Errorf("Expected to be equal to '9'; got: %v", IntToBase62(9))
	}

	if IntToBase62(10) != 'A' {
		t.Errorf("Expected to be equal to 'A'; got: %v", IntToBase62(10))
	}

	if IntToBase62(35) != 'Z' {
		t.Errorf("Expected to be equal to 'Z'; got: %v", IntToBase62(35))
	}

	if IntToBase62(36) != 'a' {
		t.Errorf("Expected to be equal to 'a'; got: %v", IntToBase62(36))
	}

	if IntToBase62(61) != 'z' {
		t.Errorf("Expected to be equal to 'z'; got: %v", IntToBase62(61))
	}
}
