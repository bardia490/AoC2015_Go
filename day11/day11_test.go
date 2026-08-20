package day11

import "testing"

func TestCorrectPassword(t *testing.T) {
	pass := []byte{'g', 'h', 'j', 'a', 'a', 'b', 'c', 'c'}
	if !checkPassword(pass) {
		t.Fatalf("oh shit")
	}
}

func TestForwardPassword(t *testing.T) {
	answers := []string{"xy", "xz", "ya"}
	pass := []byte{'x', 'x'}

	for _, answer := range answers {
		forwadPassword(pass)
		if answer != string(pass) {
			t.Fatalf("the answer was: %s, and got: %s", answer, pass)
		}
	}
}

func TestCheckPassword(t *testing.T) {
	answers := []string{"abcdffaa", "ghjaabcc"}
	results := []string{generateNewPassword([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}), generateNewPassword([]byte{'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'})}

	for index, answer := range answers {
		if answer != results[index] {
			t.Fatalf("the answer was: %s, and got: %s", answer, results[index])
		}
	}
}
