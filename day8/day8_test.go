package day8

import "testing"

func TestSolution1(t *testing.T) {
	number_of_characters_answers := []int{2, 5, 10, 6, 32, 31}
	number_of_characters_got_answers := []int{CountNumberOfCharacters(`""`), CountNumberOfCharacters(`"abc"`), CountNumberOfCharacters(`"aaa\"aaa"`), CountNumberOfCharacters(`"\x27"`), CountNumberOfCharacters(`"ap\"smgsuexjrbuqs\"mpbstogj\"x"`), CountNumberOfCharacters(`"ebvptcjqjhc\"n\"p\"dxrphegr\\"`)}

	for index, answer := range number_of_characters_answers {
		if answer != number_of_characters_got_answers[index] {
			t.Fatalf("the answer was: %d, and got: %d", answer, number_of_characters_got_answers[index])
		}
	}

	number_of_characters_in_memory_answers := []int{0, 3, 7, 1, 27, 25}
	number_of_characters_in_memory_got_answers := []int{CountNumberOfInMemoryCharacters(`""`), CountNumberOfInMemoryCharacters(`"abc"`), CountNumberOfInMemoryCharacters(`"aaa\"aaa"`), CountNumberOfInMemoryCharacters(`"\x27"`), CountNumberOfInMemoryCharacters(`"ap\"smgsuexjrbuqs\"mpbstogj\"x"`), CountNumberOfInMemoryCharacters(`"ebvptcjqjhc\"n\"p\"dxrphegr\\"`)}

	for index, answer := range number_of_characters_in_memory_answers {
		if answer != number_of_characters_in_memory_got_answers[index] {
			t.Fatalf("the answer was: %d, and got: %d", answer, number_of_characters_in_memory_got_answers[index])
		}
	}
}

func TestSolution2(t *testing.T) {
	number_of_characters_answers := []int{2, 5, 10, 6, 32, 31}
	number_of_characters_got_answers := []int{CountNumberOfCharacters(`""`), CountNumberOfCharacters(`"abc"`), CountNumberOfCharacters(`"aaa\"aaa"`), CountNumberOfCharacters(`"\x27"`), CountNumberOfCharacters(`"ap\"smgsuexjrbuqs\"mpbstogj\"x"`), CountNumberOfCharacters(`"ebvptcjqjhc\"n\"p\"dxrphegr\\"`)}

	for index, answer := range number_of_characters_answers {
		if answer != number_of_characters_got_answers[index] {
			t.Fatalf("the answer was: %d, and got: %d", answer, number_of_characters_got_answers[index])
		}
	}

	number_of_characters_in_memory_answers := []int{6, 9, 16, 11, 42, 43}
	number_of_characters_in_memory_got_answers := []int{CountNumberOfCharacterEncodings(`""`), CountNumberOfCharacterEncodings(`"abc"`), CountNumberOfCharacterEncodings(`"aaa\"aaa"`), CountNumberOfCharacterEncodings(`"\x27"`), CountNumberOfCharacterEncodings(`"ap\"smgsuexjrbuqs\"mpbstogj\"x"`), CountNumberOfCharacterEncodings(`"ebvptcjqjhc\"n\"p\"dxrphegr\\"`)}

	for index, answer := range number_of_characters_in_memory_answers {
		if answer != number_of_characters_in_memory_got_answers[index] {
			t.Fatalf("the answer was: %d, and got: %d, at index: %d", answer, number_of_characters_in_memory_got_answers[index], index)
		}
	}
}
