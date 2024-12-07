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
    return (-current_direction[1], current_direction[0])


def find_starting_point(input: List[str]) -> Tuple[int, int]:
    for y, row in enumerate(input):
        for x, char in enumerate(row):
            if char == "^":
                return (x, y)

    raise ValueError("No starting point found")


def count_visits(grid: List[str],
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
                or new_y >= len(grid) or new_x >= len(grid[new_y])):
            break

        # If we hit an obstacle, turn right
        # Otherwise, mark the location as visited and continue
        if grid[new_y][new_x] == "#":
            direction = turn_right(direction)
        else:
            visited.add((new_x, new_y))
            x, y = new_x, new_y

    return visited


# Use a fast and slow pointer to detect cycles
def detect_cycle(grid: List[str],
                 starting_point: Tuple[int, int]) -> bool:

    slow_direction = (0, -1)
    fast_direction = (0, -1)

    slow_x, slow_y = starting_point
    fast_x, fast_y = starting_point

    while True:
        # Move the slow pointer
        new_slow_x = slow_x + slow_direction[0]
        new_slow_y = slow_y + slow_direction[1]

        # If we exit the grid, stop
        if (new_slow_x < 0 or new_slow_y < 0
                or new_slow_y >= len(grid)
                or new_slow_x >= len(grid[new_slow_y])):
            return False

        # If we hit an obstacle, turn right
        if (grid[new_slow_y][new_slow_x] == "#" or
                grid[new_slow_y][new_slow_x] == "O"):
            slow_direction = turn_right(slow_direction)
        else:
            slow_x, slow_y = new_slow_x, new_slow_y

        # Move the fast pointer twice
        for _ in range(2):
            new_fast_x = fast_x + fast_direction[0]
            new_fast_y = fast_y + fast_direction[1]

            # If we exit the grid, stop
            if (new_fast_x < 0 or new_fast_y < 0
                    or new_fast_y >= len(grid)
                    or new_fast_x >= len(grid[new_fast_y])):
                return False

            # If we hit an obstacle, turn right
            if (grid[new_fast_y][new_fast_x] == "#" or
                    grid[new_fast_y][new_fast_x] == "O"):
                fast_direction = turn_right(fast_direction)
            else:
                fast_x, fast_y = new_fast_x, new_fast_y

        # If the two pointers meet with the same direction, we have a cycle
        if ((slow_x, slow_y) == (fast_x, fast_y) and
                slow_direction == fast_direction):
            return True


def part1():
    input = parse_input(read_input())

    starting_point = find_starting_point(input)

    visited = count_visits(input, starting_point)

    print(f"Visited {len(visited)} locations (part 1)")


def part2():
    input = parse_input(read_input())

    starting_point = find_starting_point(input)

    visited = count_visits(input, starting_point)

    detected_cycles = 0

    for visited_x, visited_y in visited:
        copy = [line[:] for line in input]

        # Add an obstacle to the visited location
        copy[visited_y] = copy[visited_y][:visited_x] + \
            "O" + copy[visited_y][visited_x + 1:]

        if detect_cycle(copy, starting_point):
            detected_cycles += 1

    print(f"Detected {detected_cycles} cycles (part 2)")


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
