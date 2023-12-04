package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	Id       int
	Winners  []int
	Picks    []int
	WinCount int
	Matches  []int
}

func NewCard(line string) Card {
	data := strings.Split(line, ":")
	temp := strings.Fields(data[0])
	id, err := strconv.Atoi(temp[len(temp)-1])
	if err != nil {
		panic(err)
	}

	numbers := strings.Split(data[1], "|")
	winners := []int{}
	for _, winner := range strings.Fields(numbers[0]) {
		w, err := strconv.Atoi(winner)
		if err != nil {
			fmt.Println(winner)
			panic(err)
		}
		winners = append(winners, w)
	}
	picks := []int{}
	for _, pick := range strings.Fields(numbers[1]) {
		p, err := strconv.Atoi(pick)
		if err != nil {
			panic(err)
		}
		picks = append(picks, p)
	}

	matches := []int{}
	for _, pick := range picks {
		for _, winner := range winners {
			if pick == winner {
				matches = append(matches, pick)
			}
		}
	}
	return Card{Id: id, Winners: winners, Picks: picks, WinCount: len(matches), Matches: matches}
}

func (c *Card) GetScore() int {

	if c.WinCount == 0 {
		return 0
	}
	ans := 1
	for i := 1; i < c.WinCount; i++ {
		ans *= 2
	}
	return ans
}

func main() {
	lines := ReadData("data.txt")

	cards := []*Card{}
	for _, line := range lines {
		newCard := NewCard(line)
		cards = append(cards, &newCard)
	}

	Part1(cards)
	Part2(cards)

}

func Part1(cards []*Card) {
	start_time := time.Now()

	ans := 0
	for _, card := range cards {
		if card.WinCount <= 1 {
			ans += card.WinCount
		} else {
			ans += card.GetScore()
		}
	}

	elapsed := time.Since(start_time)
	fmt.Printf("Part 1 answer: %v\n", ans)
	fmt.Printf("Part 1 Time: %s\n\n", elapsed)

}

func Part2(cards []*Card) {
	start_time := time.Now()

	ans := 0
	cardcount := make([]int, len(cards))

	for i, card := range cards {
		cardcount[i]++
		for j := 0; j < cardcount[i]; j++ {
			for k := 0; k < card.WinCount; k++ {
				cardcount[i+k+1]++
			}
		}
		ans += cardcount[i]
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
