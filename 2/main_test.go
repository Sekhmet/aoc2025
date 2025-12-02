package main

import "testing"

func TestValidateLevel2(t *testing.T) {
	valid := validateLevel2(565656)

	if valid {
		t.Errorf("validatelevel2(565656) = %t, want false", valid)
	}

	valid = validateLevel2(824824824)

	if valid {
		t.Errorf("validatelevel2(824824824) = %t, want false", valid)
	}

	valid = validateLevel2(73773773)

	if !valid {
		t.Errorf("validatelevel2(73773773) = %t, want true", valid)
	}

	valid = validateLevel2(123456)

	if !valid {
		t.Errorf("validatelevel2(123456) = %t, want true", valid)
	}
}
