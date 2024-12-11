import argparse
from typing import DefaultDict, Dict, List, Set, Tuple

INPUT_PATH = "input-files/2024/08/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/08/input-sample.txt"

is_sample_mode = False


def parse_input() -> List[List[str]]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        grid = [list(line.strip()) for line in input_file.readlines()]

    return grid


def find_antinodes(
    loc1: Tuple[int, int], loc2: Tuple[int, int], rows: int, cols: int, unlimited=False
) -> List[Tuple[int, int]]:
    antinodes: List[Tuple[int, int]] = []
    row_diff = loc1[0] - loc2[0]
    col_diff = loc1[1] - loc2[1]

    antinode1 = (loc2[0] - row_diff, loc2[1] - col_diff)
    antinode2 = (loc1[0] + row_diff, loc1[1] + col_diff)

    while (
        antinode1[0] >= 0
        and antinode1[0] < rows
        and antinode1[1] >= 0
        and antinode1[1] < cols
    ):
        antinodes.append(antinode1)

        if unlimited:
            antinode1 = (antinode1[0] - row_diff, antinode1[1] - col_diff)
        else:
            break

    while (
        antinode2[0] >= 0
        and antinode2[0] < rows
        and antinode2[1] >= 0
        and antinode2[1] < cols
    ):
        antinodes.append(antinode2)

        if unlimited:
            antinode2 = (antinode2[0] + row_diff, antinode2[1] + col_diff)
        else:
            break

    return antinodes


def part1():
    grid = parse_input()
    rows = len(grid)
    cols = len(grid[0])

    attenna_locs: Dict[str, List[Tuple[int, int]]] = DefaultDict(list)

    # For each attenna location, create antinodes for all of the other attennas
    antinodes_part_1: Set[Tuple[int, int]] = set()
    antinodes_part_2: Set[Tuple[int, int]] = set()

    # Find all of the attenna locations and store in a dictionary
    # keyed by the attenna name
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] != ".":
                attenna_locs[grid[r][c]].append((r, c))
                antinodes_part_2.add((r, c))

    for antenna, locs in attenna_locs.items():
        for i in range(len(locs)):
            for j in range(i + 1, len(locs)):
                antinodes_part_1.update(find_antinodes(locs[i], locs[j], rows, cols))
                antinodes_part_2.update(
                    find_antinodes(locs[i], locs[j], rows, cols, True)
                )

    print(f"Antinodes (part 1): {len(antinodes_part_1)}")
    print(f"Antinodes (part 2): {len(antinodes_part_2)}")


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true", help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global is_sample_mode
        is_sample_mode = True

    part1()


if __name__ == "__main__":
    main()
