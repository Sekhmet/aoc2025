package main

import "testing"

func TestGetJoltage(t *testing.T) {
	joltage := getJoltage("987654321111111", 2)

	if joltage != 98 {
		t.Errorf("getJoltage(987654321111111, 2) = %d, want 98", joltage)
	}

	joltage = getJoltage("811111111111119", 2)

	if joltage != 89 {
		t.Errorf("getJoltage(811111111111119, 2) = %d, want 89", joltage)
	}

	joltage = getJoltage("987654321111111", 12)

	if joltage != 987654321111 {
		t.Errorf("getJoltage(987654321111111, 12) = %d, want 987654321111", joltage)
	}

	joltage = getJoltage("811111111111119", 12)

	if joltage != 811111111119 {
		t.Errorf("getJoltage(811111111111119, 12) = %d, want 811111111119", joltage)
	}

	joltage = getJoltage("234234234234278", 12)

	if joltage != 434234234278 {
		t.Errorf("getJoltage(234234234234278, 12) = %d, want 434234234278", joltage)
	}
}
