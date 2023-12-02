package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Round struct {
	Red   int
	Green int
	Blue  int
}

func NewSet(data string) Round {
	red := 0
	green := 0
	blue := 0

	cubes := strings.Split(data, ",")
	for _, c := range cubes {
		to_add := strings.Split(strings.TrimSpace(c), " ")
		amt, err := strconv.Atoi(to_add[0])
		if err != nil {
			log.Fatal(err)
		}
		color := to_add[1]
		if color == "red" {
			red += amt
		} else if color == "green" {
			green += amt
		} else if color == "blue" {
			blue += amt
		}
	}
	round := Round{Red: red, Green: green, Blue: blue}
	return round
}

type Game struct {
	Id     int
	Rounds []Round
}

func NewGame(data string) Game {

	temp := strings.Split(data, ":")
	idstring := strings.Split(temp[0], " ")[1]
	id, err := strconv.Atoi(idstring)
	if err != nil {
		log.Fatal(err)
	}
	rounds := strings.Split(temp[1], ";")

	newgame := Game{Id: id, Rounds: []Round{}}

	for _, r := range rounds {
		newgame.Rounds = append(newgame.Rounds, NewSet(strings.TrimSpace(r)))
	}

	return newgame
}

func (g *Game) isValid(redLimit int, greenLimit int, blueLimit int) bool {
	for _, round := range g.Rounds {
		if round.Red > redLimit || round.Green > greenLimit || round.Blue > blueLimit {
			return false
		}
	}
	return true
}

func (g *Game) smallestValues() int {
	sRed := 0
	sGreen := 0
	sBlue := 0

	for _, round := range g.Rounds {
		if round.Red > sRed {
			sRed = round.Red
		}
		if round.Green > sGreen {
			sGreen = round.Green
		}
		if round.Blue > sBlue {
			sBlue = round.Blue
		}
	}
	ans := sRed * sGreen * sBlue
	return ans
}

func main() {
	lines := ReadData("data.txt")
	games := []Game{}
	for _, line := range lines {
		games = append(games, NewGame(line))
	}

	Part1(games)
	Part2(games)

}

func Part1(games []Game) {
	start_time := time.Now()
	ans := 0

	for _, game := range games {
		if game.isValid(12, 13, 14) {
			ans += game.Id
		}
	}

	elapsed := time.Since(start_time)
	fmt.Printf("Part 1 answer: %v\n", ans)
	fmt.Printf("Part 1 Time: %s\n\n", elapsed)
}

func Part2(games []Game) {
	start_time := time.Now()
	ans := 0
	for _, game := range games {
		ans += game.smallestValues()
	}

	elapsed := time.Since(start_time)
	fmt.Printf("Part 2 answer: %v\n", ans)
	fmt.Printf("Part 2 Time: %s\n\n", elapsed)
}

func ReadData(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}
