package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Number struct {
	X      int
	Y      int
	Value  int
	Length int
	Marker string
	MX     int
	MY     int
	IsPart bool
}

func NewNumber(s string, x int, y int) Number {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	length := len(s)
	return Number{X: x - length, Y: y, Value: value, Length: length, IsPart: false}

}

type Schematic struct {
	Rows   []string
	Width  int
	Height int
}

func NewSchematic(lines []string) Schematic {
	return Schematic{Rows: lines, Width: len(lines[0]), Height: len(lines)}
}

func (s *Schematic) ParseNumbers() []Number {
	numbers := []Number{}

	curNumber := ""
	for y, row := range s.Rows {
		x := 0
		for x < s.Width {
			if unicode.IsDigit(rune(row[x])) {
				curNumber += string(row[x])
			} else if len(curNumber) > 0 {
				numbers = append(numbers, s.CheckPart(NewNumber(curNumber, x, y)))
				curNumber = ""
			}
			x++
		}
		if len(curNumber) > 0 {
			numbers = append(numbers, s.CheckPart(NewNumber(curNumber, x, y)))
			curNumber = ""
		}
	}
	return numbers
}

func (s *Schematic) CheckPart(num Number) Number {
	newNum := num
	startY := num.Y - 1
	endY := num.Y + 1
	if startY < 0 {
		startY = 0
	}
	if endY > s.Height-1 {
		endY = s.Height - 1
	}

	startX := num.X - 1
	endX := num.X + num.Length
	if startX < 0 {
		startX = 0
	}
	if endX > s.Width-1 {
		endX = s.Width - 1
	}

	y := startY
	for y <= endY {
		x := startX
		for x <= endX {
			if !unicode.IsDigit(rune(s.Rows[y][x])) && s.Rows[y][x] != '.' {

				newNum.Marker = string(s.Rows[y][x])
				newNum.MX = x
				newNum.MY = y
				newNum.IsPart = true
			}
			x++
		}
		y++
	}

	// return false
	return newNum
}

func main() {
	data := ReadData("data.txt")

	smc := NewSchematic(data)

	Part1(smc)
	Part2(smc)

}

func Part1(smc Schematic) {
	start_time := time.Now()

	ans := 0
	numbers := smc.ParseNumbers()
	for _, n := range numbers {
		if n.IsPart {
			ans += n.Value
		}
	}

	elapsed := time.Since(start_time)
	fmt.Printf("Part 1 answer: %v\n", ans)
	fmt.Printf("Part 1 Time: %s\n\n", elapsed)
}

func Part2(smc Schematic) {
	start_time := time.Now()

	gears := []Number{}
	numbers := smc.ParseNumbers()
	for _, n := range numbers {
		nn := smc.CheckPart(n)
		if nn.Marker == "*" {
			gears = append(gears, smc.CheckPart(n))
		}
	}

	ans := 0
	for i, n := range gears {
		j := i + 1
		for j < len(gears) {
			if n.MX == gears[j].MX && n.MY == gears[j].MY {
				ans += n.Value * gears[j].Value
			}
			j++
		}
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
