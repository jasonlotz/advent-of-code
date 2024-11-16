INPUT_PATH = "input-files/2023/05/input.txt"


def part_1():
    seeds, *map_blocks = open(INPUT_PATH).read().split("\n\n")

    seeds = list(map(int, seeds.split(":")[1].split()))

    for block in map_blocks:
        map_ranges = []
        for line in block.splitlines()[1:]:
            map_ranges.append(list(map(int, line.split())))

        mapped_locations = []
        for seed in seeds:
            for destination_start, source_start, range_len in map_ranges:
                if source_start <= seed < source_start + range_len:
                    mapped_locations.append(
                        seed - source_start + destination_start)
                    break
            else:
                mapped_locations.append(seed)

        seeds = mapped_locations

    print(f"Minimum location - part 1: {min(seeds)}")


def part_2():
    inputs, *map_blocks = open(INPUT_PATH).read().split("\n\n")

    inputs = list(map(int, inputs.split(":")[1].split()))

    seed_tuples = []
    for i in range(0, len(inputs), 2):
        seed_tuples.append((inputs[i], inputs[i] + inputs[i + 1]))

    for block in map_blocks:
        map_ranges = []
        for line in block.splitlines()[1:]:
            map_ranges.append(list(map(int, line.split())))

        new_mapped_ranges = []
        while len(seed_tuples) > 0:
            start, end = seed_tuples.pop()
            for destination_start, source_start, range_len in map_ranges:
                overlap_start = max(start, source_start)
                overlap_end = min(end, source_start + range_len)
                if overlap_start < overlap_end:
                    new_mapped_ranges.append(
                        (overlap_start - source_start + destination_start,
                         overlap_end - source_start + destination_start))
                    if overlap_start > start:
                        seed_tuples.append((start, overlap_start))
                    if end > overlap_end:
                        seed_tuples.append((overlap_end, end))
                    break
            else:
                new_mapped_ranges.append((start, end))
        seed_tuples = new_mapped_ranges

    print(f"Minimum location - part 2: {min(seed_tuples)[0]}")


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
