package day21

import (
	"Aoc2015/lib/utility"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

//func generateDrugsArray(drugs []Drug, in []byte, index int) {
//	contents := bytes.SplitN(in, []byte(" => "), 2)
//
//	drugs[index] = Drug{
//		initial:     contents[0],
//		replacement: contents[1],
//	}
//}

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
	health: 109,
	damage: 8,
	armor:  2,
}

func playGame(p_damage, p_armor int) (player_win bool) {
	player_turn := true
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

	for boss_health >= 0 && player_health >= 0 {
		if player_turn {
			boss_health -= p_damage
			player_turn = false
			continue
		}
		player_health -= boss_damage
		player_turn = true
	}

	return player_health > 0
}

func part1() int {
	// choose a weapon
	lowest_cost := math.MaxInt
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

	index := 0
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			break
		}
		//generateDrugsArray(drugs[:], line, index)
		index += 1
	}
	result := part1()

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day21 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			break
		}
		//generateDrugsArray(drugs[:], line, index)
		index += 1
	}
	//result := 0
	result := part2()

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day21 part 2 is:", result)
}
