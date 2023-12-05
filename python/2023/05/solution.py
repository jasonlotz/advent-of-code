from collections import namedtuple
import sys

INPUT_PATH = "input-files/2023/05/input-sample.txt"

Mapping = namedtuple(
    "Mapping", ["destination_start", "source_start", "range_len"])

seeds = []

maps = {
    "seed-to-soil": [],
    "soil-to-fertilizer": [],
    "fertilizer-to-water": [],
    "water-to-light": [],
    "light-to-temperature": [],
    "temperature-to-humidity": [],
    "humidity-to-location": [],
}

cached_ranges = {
    "seed-to-soil": set(),
    "soil-to-fertilizer": set(),
    "fertilizer-to-water": set(),
    "water-to-light": set(),
    "light-to-temperature": set(),
    "temperature-to-humidity": set(),
    "humidity-to-location": set(), }


def add_mapping(map_name: str, values: list):
    destination_start = int(values[0])
    source_start = int(values[1])
    range_len = int(values[2])

    mapping = Mapping(destination_start, source_start, range_len)

    maps[map_name].append(mapping)


def find_location(seed: int) -> int:
    location = seed

    location = lookup_mapping("seed-to-soil", location)
    location = lookup_mapping("soil-to-fertilizer", location)
    location = lookup_mapping("fertilizer-to-water", location)
    location = lookup_mapping("water-to-light", location)
    location = lookup_mapping("light-to-temperature", location)
    location = lookup_mapping("temperature-to-humidity", location)
    location = lookup_mapping("humidity-to-location", location)

    return location


def lookup_mapping(map_name: str, location: int) -> int:
    for mapping in maps[map_name]:
        if mapping.source_start <= location <= (mapping.source_start + mapping.range_len - 1):
            location = location + \
                (mapping.destination_start - mapping.source_start)
            break
    return location


def part_1():
    locations = []

    with open(INPUT_PATH, "r") as input_file:
        lines = input_file.readlines()

        seeds = lines[0].split(":")[1].strip().split()

        load_maps(lines)

    for seed in seeds:
        location = find_location(int(seed))
        locations.append(location)
        print(f"Seed {seed} is located at {location}")

    print(f"Minimum location for part 1: {min(locations)}")


def load_maps(lines):
    map_name = ""
    for line in lines[2:]:
        if line == "\n":
            continue

        if line[0].isdigit():
            add_mapping(map_name, line.strip().split())
            pass
        else:
            map_name = line.split()[0].strip()
            print(f"Loading map name: {map_name}")


def part_2():
    min_location = sys.maxsize

    with open(INPUT_PATH, "r") as input_file:
        lines = input_file.readlines()

        load_maps(lines)

        seeds = lines[0].split(":")[1].strip().split()
        seed_tuples = [(int(seeds[i-1]), int(seeds[i]))
                       for i in range(1, len(seeds), 2)]

        for seed_tuple in seed_tuples:
            for i in range(seed_tuple[0], seed_tuple[0] + seed_tuple[1]):
                location = find_location(i)
                min_location = min(min_location, location)
                print(f"Seed {i} is located at {location}")

    print(f"Minimum location for part 2: {min_location}")


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
