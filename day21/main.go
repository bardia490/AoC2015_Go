package day21

import (
	"Aoc2015/lib/utility"
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Item struct {
	cost   int
	damage int
	armor  int
}

var weapons = [...]Item{
	{8, 4, 0},  //"Dagger"
	{10, 5, 0}, //"Shortsword"
	{25, 6, 0}, //"Warhammer"
	{40, 7, 0}, //"Longsword"
	{74, 8, 0}, //"Greataxe"
}

var armors = [...]Item{
	{13, 0, 1},  //Leather
	{31, 0, 2},  //Chainmail
	{53, 0, 3},  //Splintmail
	{75, 0, 4},  //Bandedmail
	{102, 0, 5}, //Platemail
	{0, 0, 0},   //No armor
}
var rings = [...]Item{
	{25, 1, 0},  // Damage +1
	{50, 2, 0},  // Damage +2
	{100, 3, 0}, // Damage +3
	{20, 0, 1},  // Defense +1
	{40, 0, 2},  // Defense +2
	{80, 0, 3},  // Defense +3
	{0, 0, 0},   // Empty 1
	{0, 0, 0},   // Empty 2
}

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
	} else if bytes.Equal(contents[0], []byte("Armor")) {
		boss.armor = val
	}
	return nil
}

// return true if the player wins based on the players stats
func playGame(p_damage, p_armor int) (player_win bool) {
	player_health := 100
	boss_damage := boss.damage
	boss_health := boss.health

	if p_damage <= boss.armor {
		p_damage = 1
	} else {
		p_damage -= boss.armor
	}

	if boss_damage <= p_armor {
		boss_damage = 1 // right
	} else {
		boss_damage -= p_armor
	}

	// how many rounds should be played before the player dealts the finishing blow
	rounds := boss_health / p_damage
	// update the healths after playing ${rounds} rounds of the game
	boss_health -= rounds * p_damage
	player_health -= rounds * boss_damage
	// check to see if anyone has lost already
	if player_health <= 0 || boss_health <= 0 {
		return player_health > 0
	}
	// if both have some health left, player will win because it will go first
	// and it will definitely knock out the boss since there is no more full rounds left
	return true
}

func part1() int {
	lowest_cost := math.MaxInt
	// choose a weapon
	for _, weapon := range weapons {
		// choose an armor
		for _, armor := range armors {
			// generate subsets with two items for the rings
			for subset := range utility.GenerateSubsets(rings[:]) {
				if len(subset) != 2 {
					continue
				}
				total_cost := weapon.cost + armor.cost + subset[0].cost + subset[1].cost
				total_damage := weapon.damage + subset[0].damage + subset[1].damage // armors don't add any damage
				total_armor := armor.armor + subset[0].armor + subset[1].armor      // weapon doesn't have any armor
				winner := playGame(total_damage, total_armor)
				if winner && total_cost < lowest_cost {
					lowest_cost = total_cost
				}
			}
		}
	}
	return lowest_cost
}

// the same code as above but instead we look for the loosing matches
func part2() int {
	// choose a weapon
	biggest_cost := math.MinInt
	for _, weapon := range weapons {
		// choose an armor
		for _, armor := range armors {
			// generate subsets with two items for the rings
			for subset := range utility.GenerateSubsets(rings[:]) {
				if len(subset) != 2 {
					continue
				}
				total_cost := weapon.cost + armor.cost + subset[0].cost + subset[1].cost
				total_damage := weapon.damage + subset[0].damage + subset[1].damage // armors don't add any damage
				total_armor := armor.armor + subset[0].armor + subset[1].armor      // weapon doesn't have any armor

				if !playGame(total_damage, total_armor) && total_cost > biggest_cost {
					biggest_cost = total_cost
				}
			}
		}
	}
	return biggest_cost
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
