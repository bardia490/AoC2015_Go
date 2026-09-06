package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Boss struct {
	health int
	damage int
	armor  int
}

var boss = Boss{
	health: 0,
	damage: 0,
	armor:  0,
}

func setBossStats(in []byte) (parsing_error error) {
	contents := bytes.Split(in, []byte(": "))
	val, err := strconv.Atoi(string(contents[1]))
	if err != nil {
		return err
	}
	if bytes.Equal(contents[0], []byte("Hit Points")) {
		boss.health = val
	} else if bytes.Equal(contents[0], []byte("Damage")) {
		boss.damage = val
	}
	return nil
}

type spellType int

const (
	Missile spellType = iota
	Drain
	Shield
	Poison
	Recharge
)

type spell struct {
	cost      int
	duration  int
	damage    int
	heal      int
	armor     int
	manRefill int
}

var spells = map[spellType]spell{
	Missile: {
		cost:   53,
		damage: 4,
	},
	Drain: {
		cost:   73,
		damage: 2,
		heal:   2,
	},
	Shield: {
		cost:     113,
		duration: 6,
		armor:    7,
	},
	Poison: {
		cost:     173,
		duration: 6,
		damage:   3,
	},
	Recharge: {
		cost:      229,
		duration:  5,
		manRefill: 101,
	},
}

func part1() int {
	// start round
	// check active effects and apply any that are active
	// choose an effect
	// if it was already active use continue to loop back and choose a new effect
	// add effect to active effects (if necessary)
	// check to see if the player has died, if so then end this round and step back to choose a new effect
	// if boss is dead update the result variable (which will probably be a pointer)
	// call this function again to proceed to the next state
	return 0
}

func part2() int {
	return 0
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Bytes()
		if err := setBossStats(line); err != nil {
			fmt.Println("problem parsing the stat:", err.Error())
		}
	}
	result := part1()

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day21 part 1 is:", result)
}

func Solution2(f *os.File) {
	// we don't even need to parse the file again since the boss is a global variable
	result := part2()

	fmt.Println("the solution to day21 part 2 is:", result)
}
