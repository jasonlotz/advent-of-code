import argparse
from typing import List

INPUT_PATH = "input-files/yyyy/dd/input.txt"
SAMPLE_INPUT_PATH = "input-files/yyyy/dd/input-sample.txt"

isSampleMode = False


def read_input() -> List[str]:
    if isSampleMode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> None:
    for line in input:
        print(line)
    pass


def part1():
    input = read_input()
    parse_input(input)
    pass


def part2():
    pass


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true",
                        help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global isSampleMode
        isSampleMode = True

    part1()
    part2()


if __name__ == "__main__":
    main()
