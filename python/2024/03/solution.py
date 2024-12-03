import argparse
import re

INPUT_PATH = "input-files/2024/03/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/03/input-sample-2.txt"

is_sample_mode = False


def read_input() -> str:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    # read the input file as a string
    with open(file_path, "r") as input_file:
        return input_file.read()


def part1():
    input = read_input()

    # Use a regex to find examples like 'mul(2, 3)' and capture the two digits
    pattern = r"mul\((\d+),(\d+)\)"
    matches = re.findall(pattern, input)

    result = 0
    for x, y in matches:
        result += int(x) * int(y)

    print(f"Part 1: {result}")


def part2():
    input = read_input()

    pattern = r"mul\(\d+,\d+\)|do\(\)|don't\(\)"
    matches = re.findall(pattern, input)

    result = 0
    flag = True  # start with mul instructions as enabled
    for match in matches:
        if match == "do()":
            flag = True
        elif match == "don't()":
            flag = False
        else:
            if flag:
                # extract the two digits
                x, y = map(int, match[4:-1].split(","))
                result += x * y

    print(f"Part 2: {result}")


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
