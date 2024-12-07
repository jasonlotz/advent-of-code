import argparse
import itertools
from typing import List, Tuple

INPUT_PATH = "input-files/2024/07/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/07/input-sample.txt"

is_sample_mode = False


def read_input() -> List[str]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> List[Tuple[int, List[int]]]:
    parsed_input: List[Tuple[int, List[int]]] = []

    for line in input:
        split_line = line.split(":")
        test_value = int(split_line[0])
        parsed_numbers = [int(num) for num in split_line[1].split()]
        parsed_input.append((test_value, parsed_numbers))

    return parsed_input


def part1():
    input = parse_input(read_input())

    total_calibration_result = 0

    for test_value, numbers in input:
        operations = ["*", "+"]

        op_combos = itertools.product(operations, repeat=len(numbers) - 1)

        for op_combo in op_combos:
            result = numbers[0]
            for i, op in enumerate(op_combo):
                if op == "*":
                    result *= numbers[i + 1]
                elif op == "+":
                    result += numbers[i + 1]

            if result == test_value:
                total_calibration_result += test_value
                break

    print(f"Total calibration result (part 1): {total_calibration_result}")


def part2():
    input = parse_input(read_input())

    total_calibration_result = 0

    for test_value, numbers in input:
        operations = ["*", "+", "||"]

        op_combos = itertools.product(operations, repeat=len(numbers) - 1)

        for op_combo in op_combos:
            result = numbers[0]
            for i, op in enumerate(op_combo):
                if op == "*":
                    result *= numbers[i + 1]
                elif op == "+":
                    result += numbers[i + 1]
                elif op == "||":
                    # Concatenate the result and the next number
                    result = int(str(result) + str(numbers[i + 1]))

            if result == test_value:
                total_calibration_result += test_value
                break

    print(f"Total calibration result (part 2): {total_calibration_result}")


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
