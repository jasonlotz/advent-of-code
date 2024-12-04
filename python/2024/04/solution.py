import argparse
from typing import List, Tuple

INPUT_PATH = "input-files/2024/04/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/04/input-sample.txt"

is_sample_mode = False


def read_input() -> List[str]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> List[str]:
    split_lines = [line.strip() for line in input]

    return split_lines


def find_xmas(input: List[str], x_location: Tuple[int, int]) -> int:
    count = 0

    directions = [
        (1, 0),   # Right
        (0, 1),   # Down
        (-1, 0),  # Left
        (0, -1),  # Up
        (1, 1),   # Down right
        (-1, 1),  # Down_left
        (1, -1),  # Up right
        (-1, -1)  # Up left
    ]

    # Look for xmas in all directions starting from the x_location
    for direction in directions:
        for i, char in enumerate("MAS", start=1):
            x = x_location[0] + direction[0] * i
            y = x_location[1] + direction[1] * i

            if x < 0 or y < 0 or y >= len(input) or x >= len(input[y]):
                break

            if input[y][x] != char:
                break

            if i == 3:
                count += 1

    return count


def part1():
    input = parse_input(read_input())

    xmas_count = 0
    for y, line in enumerate(input):
        for x, char in enumerate(line):
            if char == "X":
                xmas_count += find_xmas(input, (x, y))

    print(f"XMAS count (part 1): {xmas_count}")


def part2():
    input = parse_input(read_input())

    count = 0
    y_max = len(input)
    x_max = len(input[0])

    valid_patterns = [
        ('M', 'M', 'S', 'S'),
        ('M', 'S', 'M', 'S'),
        ('S', 'M', 'S', 'M'),
        ('S', 'S', 'M', 'M'),
    ]

    # Since we are looking for "A" in the middle, we can skip the outside edges
    for y in range(1, y_max - 1):
        for x in range(1, x_max - 1):
            # Skip if the current character is not an "A"
            if input[y][x] != 'A':
                continue

            # Create a tuple of the diagonal characters
            letters = (input[y-1][x-1], input[y-1][x+1],
                       input[y+1][x-1], input[y+1][x+1])

            if letters in valid_patterns:
                count += 1

    print(f"X-MAS count (part 2): {count}")


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
