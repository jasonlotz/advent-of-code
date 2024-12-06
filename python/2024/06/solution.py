import argparse
from typing import List, Set, Tuple

INPUT_PATH = "input-files/2024/06/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/06/input-sample.txt"

is_sample_mode = False


def read_input() -> List[str]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> List[str]:
    parsed_lines = [line.strip() for line in input]

    return parsed_lines


def turn_right(current_direction: Tuple[int, int]) -> Tuple[int, int]:
    # Use rotation matrix to turn right
    if current_direction == (0, -1):
        return (1, 0)
    elif current_direction == (1, 0):
        return (0, 1)
    elif current_direction == (0, 1):
        return (-1, 0)
    elif current_direction == (-1, 0):
        return (0, -1)

    raise ValueError("Invalid direction")


def find_starting_point(input: List[str]) -> Tuple[int, int]:
    for y, row in enumerate(input):
        for x, char in enumerate(row):
            if char == "^":
                return (x, y)

    raise ValueError("No starting point found")


def follow_path(input: List[str],
                starting_point: Tuple[int, int]) -> Set[Tuple[int, int]]:
    direction = (0, -1)  # Start facing up

    visited: Set[Tuple[int, int]] = set()

    x, y = starting_point

    # Mark the starting point as visited
    visited.add((x, y))

    # Move in the current direction
    while True:
        new_x = x + direction[0]
        new_y = y + direction[1]

        # If we exit the grid, stop
        if (new_x < 0 or new_y < 0
                or new_y >= len(input) or new_x >= len(input[new_y])):
            break

        # If we hit an obstacle, turn right
        # Otherwise, mark the location as visited and continue
        if input[new_y][new_x] == "#":
            direction = turn_right(direction)
        else:
            visited.add((new_x, new_y))
            x, y = new_x, new_y

    return visited


def part1():
    input = parse_input(read_input())

    starting_point = find_starting_point(input)

    visited = follow_path(input, starting_point)

    print(f"Visited {len(visited)} locations (part 1)")


def part2():
    pass


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true",
                        help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global is_sample_mode
        is_sample_mode = True
    part1()
    part2()


if __name__ == "__main__":
    main()
