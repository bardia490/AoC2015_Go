package day22

// I should try Dijkstra's algorithm

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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
	Shield spellType = iota
	Poison
	Recharge
	Missile
	Drain
)

type spell struct {
	cost       int
	duration   int
	damage     int
	heal       int
	armor      int
	manaRefill int
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
		cost:       229,
		duration:   5,
		manaRefill: 101,
	},
}

type playerStats struct {
	mana   int
	health int
	armor  int
}

// returns true if the spell is active
func isSpellActive(s spellType, durations []int) bool {
	// since there is only three spells active at each round we can just use a simple for loop for searching
	// 0 => Shield, 1 => Poison, 2 => Recharge
	switch s {
	case Shield:
		return durations[0] > 0
	case Poison:
		return durations[1] > 0
	case Recharge:
		return durations[2] > 0
	}
	return false // this should not happen
}

// min_mana will updated at the end of every turn and is the answer to the problem
// duration: shows the time left for every spell, at the begining everything is zero (inactive)
// 0 => Shield, 1 => Poison, 2 => Recharge
func playGame(p playerStats, b Boss, durations []int, mana_cost int, min_mana *int) {
	//Stop if we are already worse than the best
	if mana_cost >= *min_mana {
		return
	}
	// check active effects and apply any that are active
	if durations[0] != 0 {
		p.armor = spells[Shield].armor
		durations[0] -= 1
		if durations[0] == 0 {
			p.armor = 0
		}
	}
	if durations[1] != 0 {
		b.health -= spells[Poison].damage
		durations[1] -= 1
	}
	if durations[2] != 0 {
		p.mana += spells[Recharge].manaRefill
		durations[2] -= 1
	}
	// if boss is dead update the result variable
	if b.health < 1 {
		if mana_cost < *min_mana {
			*min_mana = mana_cost
		}
		return
	}
	// choose a new effect to start the round
	for spell_type, spell := range spells {
		durations_copy := make([]int, len(durations))
		copy(durations_copy, durations)
		// if it was already active use continue to loop back and choose a new effect
		if isSpellActive(spell_type, durations_copy) {
			continue
		}
		// add the mana cost
		if p.mana < spell.cost {
			continue
		}
		next_mana_cost := mana_cost + spell.cost
		next_p := p
		next_p.mana -= spell.cost
		next_b := b
		// add effect to active effects (if necessary)
		switch spell_type {
		case Missile:
			next_b.health -= 4
		case Drain:
			next_b.health -= spell.damage
			next_p.health += spell.heal
		case Shield:
			durations_copy[0] = spell.duration
		case Poison:
			durations_copy[1] = spell.duration
		case Recharge:
			durations_copy[2] = spell.duration
		}
		// if boss is dead update the result variable
		if next_b.health < 1 {
			if next_mana_cost < *min_mana {
				*min_mana = next_mana_cost
			}
			return
		}
		// the bosses turn
		effective_armor := 0
		if durations_copy[0] > 0 { // Shield
			effective_armor = spells[Shield].armor
			durations_copy[0]--
		}

		if durations_copy[1] > 0 { // Poison
			next_b.health -= spells[Poison].damage
			durations_copy[1]--
		}

		if durations_copy[2] > 0 { // Recharge
			next_p.mana += spells[Recharge].manaRefill
			durations_copy[2]--
		}

		if next_b.health < 1 {
			if next_mana_cost < *min_mana {
				*min_mana = next_mana_cost
			}
			continue
		}
		next_p.health -= max(b.damage-effective_armor, 1)
		// check to see if the player has died
		if next_p.health < 1 { // we lost
			return
		}
		// call this function again to proceed to the next state
		playGame(next_p, next_b, durations_copy, next_mana_cost, min_mana)
	}
}

// min_mana will updated at the end of every turn and is the answer to the problem
// duration: shows the time left for every spell, at the begining everything is zero (inactive)
// 0 => Shield, 1 => Poison, 2 => Recharge
// the same thing as part1 but the player takes damage at the start of each round
func playGame2(p playerStats, b Boss, durations []int, mana_cost int, min_mana *int) {
	//Stop if we are already worse than the best
	if mana_cost >= *min_mana {
		return
	}
	//player takes damage at the start of each round
	p.health -= 1
	if p.health < 1 {
		return
	}
	// check active effects and apply any that are active
	if durations[0] != 0 {
		p.armor = spells[Shield].armor
		durations[0] -= 1
		if durations[0] == 0 {
			p.armor = 0
		}
	}
	if durations[1] != 0 {
		b.health -= spells[Poison].damage
		durations[1] -= 1
	}
	if durations[2] != 0 {
		p.mana += spells[Recharge].manaRefill
		durations[2] -= 1
	}
	// if boss is dead update the result variable
	if b.health < 1 {
		if mana_cost < *min_mana {
			*min_mana = mana_cost
		}
		return
	}
	// choose a new effect to start the round
	for spell_type, spell := range spells {
		durations_copy := make([]int, len(durations))
		copy(durations_copy, durations)
		// if it was already active use continue to loop back and choose a new effect
		if isSpellActive(spell_type, durations_copy) {
			continue
		}
		// add the mana cost
		if p.mana < spell.cost {
			continue
		}
		next_mana_cost := mana_cost + spell.cost
		next_p := p
		next_p.mana -= spell.cost
		next_b := b
		// add effect to active effects (if necessary)
		switch spell_type {
		case Missile:
			next_b.health -= 4
		case Drain:
			next_b.health -= spell.damage
			next_p.health += spell.heal
		case Shield:
			durations_copy[0] = spell.duration
		case Poison:
			durations_copy[1] = spell.duration
		case Recharge:
			durations_copy[2] = spell.duration
		}
		// if boss is dead update the result variable
		if next_b.health < 1 {
			if next_mana_cost < *min_mana {
				*min_mana = next_mana_cost
			}
			return
		}
		// the bosses turn
		effective_armor := 0
		if durations_copy[0] > 0 { // Shield
			effective_armor = spells[Shield].armor
			durations_copy[0]--
		}

		if durations_copy[1] > 0 { // Poison
			next_b.health -= spells[Poison].damage
			durations_copy[1]--
		}

		if durations_copy[2] > 0 {
			next_p.mana += spells[Recharge].manaRefill
			durations_copy[2]--
		}

		if next_b.health < 1 {
			if next_mana_cost < *min_mana {
				*min_mana = next_mana_cost
			}
			continue
		}
		next_p.health -= max(b.damage-effective_armor, 1)
		// check to see if the player has died
		if next_p.health < 1 { // we lost
			return
		}
		// call this function again to proceed to the next state
		playGame2(next_p, next_b, durations_copy, next_mana_cost, min_mana)
	}
}

func part1() int {
	// initialize player
	player := playerStats{
		mana:   500,
		health: 50,
		armor:  0,
	}
	// initialize the boss
	boss_clone := Boss{
		damage: boss.damage,
		health: boss.health,
	}
	min_mana := math.MaxInt
	// shows the time left for every spell, at the begining everything is zero (inactive)
	// 0 => Shield, 1 => Poison, 2 => Recharge
	durations := make([]int, 3)
	// start game
	playGame(player, boss_clone, durations, 0, &min_mana)

	return min_mana
}

func part2() int {
	// initialize player
	player := playerStats{
		mana:   500,
		health: 50,
		armor:  0,
	}
	// initialize the boss
	boss_clone := Boss{
		damage: boss.damage,
		health: boss.health,
	}
	min_mana := math.MaxInt
	// shows the time left for every spell, at the begining everything is zero (inactive)
	// 0 => Shield, 1 => Poison, 2 => Recharge
	durations := make([]int, 3)
	// start game
	playGame2(player, boss_clone, durations, 0, &min_mana)

	return min_mana
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
