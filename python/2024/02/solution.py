import argparse
from typing import List

INPUT_PATH = "input-files/2024/02/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/02/input-sample.txt"

is_sample_mode = False


def read_input() -> List[str]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> list[list[int]]:
    split_lines = [list(map(int, line.split())) for line in input]

    return split_lines


def test_for_safety(report: List[int]) -> bool:
    increasing = None  # None means we haven't determined the trend yet

    for i in range(1, len(report)):
        diff = report[i] - report[i - 1]

        if abs(diff) > 3 or abs(diff) < 1:
            return False

        # Determine the trend if not already determined
        if increasing is None:
            if diff > 0:
                increasing = True
            elif diff < 0:
                increasing = False

        # Check if the trend is violated
        if (increasing and diff < 0) or (not increasing and diff > 0):
            return False

    return True


def part1():
    input = read_input()
    reports = parse_input(input)

    safe_count = sum(1 for report in reports if test_for_safety(report))

    print(f"Safe count (part 1): {safe_count}")


def part2():
    input = read_input()
    reports = parse_input(input)

    # Count the reports that are safe, or have a safe report when a single
    # number is removed
    safe_count = sum(
        1 for report in reports
        if test_for_safety(report) or
        any(
            test_for_safety(report[:i] + report[i+1:])
            for i in range(len(report))
        )
    )

    print(f"Safe count (part 2): {safe_count}")


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
