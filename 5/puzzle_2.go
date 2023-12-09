package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

type AlmanacSeed struct {
	SeedStart int
	SeedRange int
}

type AlmanacMap struct {
	DestinationRangeStart int
	SourceRangeStart int
	RangeLength int
}

type Almanac struct {
	Seeds []AlmanacSeed
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
			for i := 0; i < len(seed_nums_str); i += 2 {
				seed_num_start, _ := strconv.Atoi(seed_nums_str[i])
				seed_num_range, _ := strconv.Atoi(seed_nums_str[i + 1])
				almanac.Seeds = append(almanac.Seeds, AlmanacSeed{
					SeedStart: seed_num_start,
					SeedRange: seed_num_range,
				})
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

	mapped_seed_ranges := almanac.Seeds
	counter := 0
	for _, map_type := range maps {
		curr_map := almanac.Maps[map_type]
		new_mapped_seed_ranges := []AlmanacSeed{}
		processed_mapped_seed_ranges := []AlmanacSeed{}

		// fmt.Println("mapped: ", mapped_seed_ranges)

		for _, map_item := range curr_map {
			for _, seed_range := range mapped_seed_ranges {
				if seed_range.SeedStart < map_item.SourceRangeStart || seed_range.SeedStart >= map_item.SourceRangeStart + map_item.RangeLength {
					continue
				}

				range_length := min(seed_range.SeedRange, map_item.SourceRangeStart + map_item.RangeLength - seed_range.SeedStart)
				offset := map_item.DestinationRangeStart - map_item.SourceRangeStart
				new_mapped_seed_ranges = append(new_mapped_seed_ranges, AlmanacSeed{
					SeedStart: seed_range.SeedStart + offset,
					SeedRange: range_length,
				})
				processed_mapped_seed_ranges = append(processed_mapped_seed_ranges, AlmanacSeed{
					SeedStart: seed_range.SeedStart,
					SeedRange: range_length,
				})

				// fmt.Println(map_item, seed_range, AlmanacSeed{
				// 	SeedStart: seed_range.SeedStart + offset,
				// 	SeedRange: range_length,
				// })
			}
		}

		for _, mapped_seed_range := range mapped_seed_ranges {
			missing_seed_ranges := []AlmanacSeed{mapped_seed_range}
			for {
				// fmt.Println("here 1", processed_mapped_seed_ranges, missing_seed_ranges)

				if len(missing_seed_ranges) == 0 {
					break
				}

				if len(processed_mapped_seed_ranges) == 0 {
					new_mapped_seed_ranges = append(new_mapped_seed_ranges, missing_seed_ranges...)
					break
				}

				missing_seed_range := missing_seed_ranges[0]
				found := false
				for _, processed_mapped_seed_range := range processed_mapped_seed_ranges {
					if processed_mapped_seed_range.SeedStart >= missing_seed_range.SeedStart + missing_seed_range.SeedRange || missing_seed_range.SeedStart >= processed_mapped_seed_range.SeedStart + processed_mapped_seed_range.SeedRange {
						// fmt.Println(processed_mapped_seed_range, missing_seed_range, "continue")
						continue
					}

					new_missing_seed_ranges := []AlmanacSeed{}

					// Missing range before
					new_missing_seed_range_before := min(missing_seed_range.SeedStart + missing_seed_range.SeedRange, processed_mapped_seed_range.SeedStart) - max(missing_seed_range.SeedStart, processed_mapped_seed_range.SeedStart)
					// fmt.Println("new_missing_seed_range_before", new_missing_seed_range_before)
					if new_missing_seed_range_before > 0 {
						new_missing_seed_ranges = append(new_missing_seed_ranges, AlmanacSeed{
							SeedStart: min(missing_seed_range.SeedStart, processed_mapped_seed_range.SeedStart),
							SeedRange: new_missing_seed_range_before,
						})
					}

					// Missing range after
					new_missing_seed_range_after := max(missing_seed_range.SeedStart + missing_seed_range.SeedRange, processed_mapped_seed_range.SeedStart + processed_mapped_seed_range.SeedRange) - min(missing_seed_range.SeedStart + missing_seed_range.SeedRange, processed_mapped_seed_range.SeedStart + processed_mapped_seed_range.SeedRange)
					// fmt.Println("new_missing_seed_range_after", new_missing_seed_range_after)
					if new_missing_seed_range_after > 0 {
						new_missing_seed_ranges = append(new_missing_seed_ranges, AlmanacSeed{
							SeedStart: missing_seed_range.SeedStart + processed_mapped_seed_range.SeedRange,
							SeedRange: new_missing_seed_range_after,
						})
					}

					missing_seed_ranges = missing_seed_ranges[1:]
					missing_seed_ranges = append(missing_seed_ranges, new_missing_seed_ranges...)

					found = true
					break
				}

				if !found {
					found := false
					for _, map_item := range curr_map {
						if missing_seed_range.SeedStart < map_item.SourceRangeStart || missing_seed_range.SeedStart >= map_item.SourceRangeStart + map_item.RangeLength {
							continue
						}

						range_length := min(missing_seed_range.SeedRange, map_item.SourceRangeStart + map_item.RangeLength - missing_seed_range.SeedStart)
						offset := map_item.DestinationRangeStart - map_item.SourceRangeStart
						new_mapped_seed_ranges = append(new_mapped_seed_ranges, AlmanacSeed{
							SeedStart: missing_seed_range.SeedStart + offset,
							SeedRange: range_length,
						})
						processed_mapped_seed_ranges = append(processed_mapped_seed_ranges, AlmanacSeed{
							SeedStart: missing_seed_range.SeedStart,
							SeedRange: range_length,
						})

						found = true
						break
					}

					if !found {
						new_mapped_seed_ranges = append(new_mapped_seed_ranges, missing_seed_range)
					}
					missing_seed_ranges = missing_seed_ranges[1:]
				}
			}

			// fmt.Println("here 2", processed_mapped_seed_ranges, missing_seed_ranges, new_mapped_seed_ranges)
		}

		mapped_seed_ranges = new_mapped_seed_ranges

		if counter++; counter == 5 {
			// break
		}
	}

	lowest_destination := 1000000000000000000
	for _, destination := range mapped_seed_ranges {
		if destination.SeedStart < lowest_destination {
			lowest_destination = destination.SeedStart
		}
	}

	fmt.Println("Day 5 Puzzle 2:", lowest_destination)
}
