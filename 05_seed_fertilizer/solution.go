package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var SEQUENCE = []string{
	"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity",
}

type Row struct {
	SourceStart int
	DestStart   int
	Length      int
}

type Mapping struct {
	Source string
	Dest   string
	Rows   []Row
}

func (m *Mapping) GetDest(idx int) int {
	ans := idx

	for _, r := range m.Rows {
		if idx >= r.SourceStart && idx < r.SourceStart+r.Length {
			ans = (idx - r.SourceStart) + r.DestStart
			return ans
		}
	}
	return ans
}

func main() {
	seedline, maptext := ReadData("data.txt")

	seeds := ParseSeeds(seedline)

	source := make(map[string]*Mapping)
	for _, m := range maptext {
		mapping := ProcessMap(m)
		source[mapping.Source] = mapping
	}

	Part1(seeds, source)
	Part2(seeds, source)

}

func Part1(seeds []int, source map[string]*Mapping) {
	start_time := time.Now()

	ans := seeds[0]
	for _, seed := range seeds {
		temp := GetLocation(seed, source)
		if temp < ans {
			ans = temp
		}
	}
	elapsed := time.Since(start_time)
	fmt.Printf("Part 1 answer: %v\n", ans)
	fmt.Printf("Part 1 Time: %s\n\n", elapsed)

}

func Part2(seeds []int, source map[string]*Mapping) {
	start_time := time.Now()

	var wg sync.WaitGroup

	res := []int{}
	result := make(chan int)
	for i := 0; i < len(seeds)-1; i += 2 {
		wg.Add(1)
		go SeedIterate(seeds[i], seeds[i+1], source, result, &wg)
	}
	for i := 0; i < len(seeds)-1; i += 2 {
		res = append(res, <-result)
	}
	wg.Wait()

	ans := res[0]
	for _, num := range res {
		if num < ans {
			ans = num
		}
	}

	elapsed := time.Since(start_time)
	fmt.Printf("Part 2 answer: %v\n", ans)
	fmt.Printf("Part 2 Time: %s\n\n", elapsed)
}

func SeedIterate(seed int, length int, source map[string]*Mapping, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ret := seed
	for i := seed; i < seed+length; i++ {
		res := GetLocation(i, source)
		if res < ret {
			ret = res
		}
	}
	result <- ret
}

func GetLocation(seed int, source map[string]*Mapping) int {
	idx := seed
	for _, k := range SEQUENCE {
		idx = source[k].GetDest(idx)
	}
	return idx
}

func ProcessMap(text string) *Mapping {
	mapping := &Mapping{}

	lines := strings.Split(text, "\n")
	mapstring := strings.Split(lines[0], "-to-")
	sourceName, destName := mapstring[0], strings.Split(mapstring[1], " ")[0]

	mapping.Source = sourceName
	mapping.Dest = destName

	for _, line := range lines[1:] {
		values := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(values[0])
		sourceStart, _ := strconv.Atoi(values[1])
		length, _ := strconv.Atoi(values[2])
		row := Row{SourceStart: sourceStart, DestStart: destStart, Length: length}
		mapping.Rows = append(mapping.Rows, row)
	}
	return mapping
}

func ParseSeeds(line string) []int {
	seeds := []int{}
	for _, value := range strings.Fields(strings.Split(line, ":")[1]) {
		seed, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func ReadData(filename string) (string, []string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(file), "\n\n")
	return s[0], s[1:]
}
