import argparse
from typing import List

INPUT_PATH = "input-files/2024/05/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/05/input-sample.txt"

is_sample_mode = False

rules = {}
updates = []


def read_input() -> List[str]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> None:
    for line in input:
        if "|" in line:
            # Rule
            a, b = line.strip().split("|")
            if a not in rules:
                rules[a] = []
            rules[a].append(b)
        else:
            # Update
            if line.strip():
                updates.append(line.strip().split(","))


def correct_order(update: list, fix_wrong=False) -> int:
    # For each update, check all of the following updates to see
    # if they are in the correct order
    for i in range(len(update)):
        for j in range(i + 1, len(update)):
            # Get the rule for the second value
            rule = rules.get(update[j], [])

            # If the first value is in the rule, then they are in the
            # wrong order
            if update[i] in rule:
                if fix_wrong:
                    # For part 2, if they are not in the correct order,
                    # swap them and check again
                    update[i], update[j] = update[j], update[i]
                    return correct_order(update, True)
                else:
                    # For part 1, if they are not in the correct order,
                    # return 0 as they are not counted
                    return 0

    return int(update[len(update) // 2])


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true",
                        help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global is_sample_mode
        is_sample_mode = True

    part_1_result = 0
    part_2_result = 0

    parse_input(read_input())

    for update in updates:
        part_1_result += correct_order(update)
        part_2_result += correct_order(update, True)

    print("Part 1:", part_1_result)
    print("Part 2:", part_2_result - part_1_result)


if __name__ == "__main__":
    main()
