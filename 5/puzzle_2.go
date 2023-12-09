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
	for _, map_type := range maps {
		curr_map := almanac.Maps[map_type]
		new_mapped_seed_ranges := []AlmanacSeed{}
		seed_ranges_to_be_mapped := []AlmanacSeed{}

		for _, seed_range := range mapped_seed_ranges {
			seed_range_splits := []AlmanacSeed{}

			for _, map_item := range curr_map {
				overlapping_start := max(seed_range.SeedStart, map_item.SourceRangeStart)
				overlapping_range := min(seed_range.SeedStart + seed_range.SeedRange, map_item.SourceRangeStart + map_item.RangeLength) - max(seed_range.SeedStart, map_item.SourceRangeStart)

				if overlapping_range <= 0 {
					continue
				}

				seed_range_splits = append(seed_range_splits, AlmanacSeed{
					SeedStart: overlapping_start,
					SeedRange: overlapping_range,
				})

				if overlapping_start == seed_range.SeedStart && overlapping_range == seed_range.SeedRange {
					break
				}
			}

			// The original seed range might not be mapped completely, let's add
			// the missing splits to the list of seed ranges to be mapped.
			if len(seed_range_splits) == 0 {
				seed_range_splits = append(seed_range_splits, seed_range)
			} else {
				// Sort the seed range splits.
				for i := 0; i < len(seed_range_splits); i++ {
					for j := i + 1; j < len(seed_range_splits); j++ {
						if seed_range_splits[i].SeedStart > seed_range_splits[j].SeedStart {
							seed_range_splits[i], seed_range_splits[j] = seed_range_splits[j], seed_range_splits[i]
						}
					}
				}


				// Add the missing splits to seed_range_splits.
				// start
				missing_seed_range_splits := []AlmanacSeed{}
				if seed_range_splits[0].SeedStart > seed_range.SeedStart {
					missing_seed_range_splits = append(missing_seed_range_splits, AlmanacSeed{
						SeedStart: seed_range.SeedStart,
						SeedRange: seed_range_splits[0].SeedStart - seed_range.SeedStart,
					})
				}

				// end
				if seed_range_splits[len(seed_range_splits) - 1].SeedStart + seed_range_splits[len(seed_range_splits) - 1].SeedRange < seed_range.SeedStart + seed_range.SeedRange {
					missing_seed_range_splits = append(missing_seed_range_splits, AlmanacSeed{
						SeedStart: seed_range_splits[len(seed_range_splits) - 1].SeedStart + seed_range_splits[len(seed_range_splits) - 1].SeedRange,
						SeedRange: seed_range.SeedStart + seed_range.SeedRange - (seed_range_splits[len(seed_range_splits) - 1].SeedStart + seed_range_splits[len(seed_range_splits) - 1].SeedRange),
					})
				}

				// Add the missing splits in between the already existing splits.
				for i := 0; i < len(seed_range_splits) - 1; i++ {
					if seed_range_splits[i].SeedStart + seed_range_splits[i].SeedRange < seed_range_splits[i + 1].SeedStart {
						missing_seed_range_splits = append(missing_seed_range_splits, AlmanacSeed{
							SeedStart: seed_range_splits[i].SeedStart + seed_range_splits[i].SeedRange,
							SeedRange: seed_range_splits[i + 1].SeedStart - (seed_range_splits[i].SeedStart + seed_range_splits[i].SeedRange),
						})
					}
				}

				seed_range_splits = append(seed_range_splits, missing_seed_range_splits...)
			}

			seed_ranges_to_be_mapped = append(seed_ranges_to_be_mapped, seed_range_splits...)
		}

		// Map the seed ranges
		for _, seed_range := range seed_ranges_to_be_mapped {
			mapped := false
			for _, map_item := range curr_map {
				overlapping_start := max(seed_range.SeedStart, map_item.SourceRangeStart)
				overlapping_range := min(seed_range.SeedStart + seed_range.SeedRange, map_item.SourceRangeStart + map_item.RangeLength) - max(seed_range.SeedStart, map_item.SourceRangeStart)

				if overlapping_range <= 0 {
					continue
				}

				new_mapped_seed_ranges = append(new_mapped_seed_ranges, AlmanacSeed{
					SeedStart: map_item.DestinationRangeStart + (overlapping_start - map_item.SourceRangeStart),
					SeedRange: overlapping_range,
				})

				mapped = true
				break
			}

			if !mapped {
				new_mapped_seed_ranges = append(new_mapped_seed_ranges, seed_range)
			}
		}

		mapped_seed_ranges = new_mapped_seed_ranges
	}

	lowest_destination := 1000000000000000000
	for _, destination := range mapped_seed_ranges {
		if destination.SeedStart < lowest_destination {
			lowest_destination = destination.SeedStart
		}
	}

	fmt.Println("Day 5 Puzzle 2:", lowest_destination)
}
