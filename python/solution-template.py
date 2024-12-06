import argparse
from typing import List

INPUT_PATH = "input-files/yyyy/dd/input.txt"
SAMPLE_INPUT_PATH = "input-files/yyyy/dd/input-sample.txt"

is_sample_mode = False


def read_input() -> List[str]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> List[List[str]]:
    split_lines = [line.split() for line in input]

    return split_lines


def part1():
    input = parse_input(read_input())

    print(input)


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
