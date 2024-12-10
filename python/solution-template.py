import argparse
from typing import List

INPUT_PATH = "input-files/yyyy/dd/input.txt"
SAMPLE_INPUT_PATH = "input-files/yyyy/dd/input-sample.txt"

is_sample_mode = False


def parse_input() -> List[List[str]]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    grid = []
    with open(file_path, "r") as input_file:
        grid = [list(line.strip()) for line in input_file.readlines()]

    return grid


def part1():
    data = parse_input()

    print(data)


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
