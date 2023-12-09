package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

type AlmanacMap struct {
	DestinationRangeStart int
	SourceRangeStart int
	RangeLength int
}

type Almanac struct {
	Seeds []int
	ProcessingMap string
	Maps map[string][]AlmanacMap
}

func main() {
	dat, err := os.ReadFile("puzzle_input.txt")
    if err != nil {
		panic(err)
	}
	puzzle := string(dat)

	almanac := Almanac{}
	almanac.Maps = make(map[string][]AlmanacMap)
	maps := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	lines := s.Split(puzzle, "\n")
	for _, line := range lines {
		if line == "" {
			almanac.ProcessingMap = ""
			continue
		}

		if s.HasPrefix(line, "seeds") {
			seed_nums_line := s.Split(line, ": ")[1]
			seed_nums_str := s.Split(seed_nums_line, " ")
			for _, seed_num_str := range seed_nums_str {
				seed_num, _ := strconv.Atoi(seed_num_str)
				almanac.Seeds = append(almanac.Seeds, seed_num)
			}
		}

		if almanac.ProcessingMap != "" {
			map_nums_str := s.Split(line, " ")
			destination, _ := strconv.Atoi(map_nums_str[0])
			source, _ := strconv.Atoi(map_nums_str[1])
			length, _ := strconv.Atoi(map_nums_str[2])
			almanac.Maps[almanac.ProcessingMap] = append(almanac.Maps[almanac.ProcessingMap], AlmanacMap{
				DestinationRangeStart: destination,
				SourceRangeStart: source,
				RangeLength: length,
			})
		}

		for _, map_name := range maps {
			if s.HasPrefix(line, map_name) {
				almanac.ProcessingMap = map_name
				break
			}
		}
	}

	seed_destination := make(map[int]int)
	for _, seed := range almanac.Seeds {
		curr_value := seed
		for _, map_type := range maps {
			curr_map := almanac.Maps[map_type]
			for _, map_item := range curr_map {
				if map_item.SourceRangeStart <= curr_value && curr_value <= map_item.SourceRangeStart + map_item.RangeLength {
					curr_value = map_item.DestinationRangeStart + (curr_value - map_item.SourceRangeStart)
					break
				}
			}
		}
		seed_destination[seed] = curr_value
	}

	lowest_destination := 1000000000000000000
	for _, destination := range seed_destination {
		if destination < lowest_destination {
			lowest_destination = destination
		}
	}

	fmt.Println("Day 5 Puzzle 1:", lowest_destination)
}
