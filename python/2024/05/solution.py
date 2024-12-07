import argparse
from typing import List, Tuple

INPUT_PATH = "input-files/2024/05/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/05/input-sample.txt"

is_sample_mode = False


class PuzzleInput:
    def __init__(self):
        self.rules: List[Tuple[int, int]] = []
        self.updates: List[List[int]] = []

    def add_rule(self, rule: Tuple[int, int]):
        self.rules.append(rule)

    def add_update(self, update: List[int]):
        self.updates.append(update)

    def __str__(self):
        return f"Rules:\n{self.rules}\nUpdates:\n{self.updates}"


def read_input() -> str:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.read()


def parse_input(input: str) -> PuzzleInput:
    rules, updates = input.split("\n\n")

    puzzle_input = PuzzleInput()

    for rule in rules.split("\n"):
        rule_parts = rule.split("|")
        puzzle_input.add_rule((int(rule_parts[0]), int(rule_parts[1])))

    for update in updates.strip().split("\n"):
        puzzle_input.add_update([int(num) for num in update.split(",")])

    return puzzle_input


def part1():
    puzzle_input = parse_input(read_input())

    middle_page_total = 0

    for update in puzzle_input.updates:
        is_valid = True

        for rule in puzzle_input.rules:
            try:
                position_before = update.index(rule[0])
                position_after = update.index(rule[1])
            except ValueError:
                continue

            if position_before > position_after:
                is_valid = False
                break

        if is_valid:
            middle_page_total += update[len(update) // 2]

    print(f"Middle page total: {middle_page_total}")


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
