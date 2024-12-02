import argparse
from collections import Counter
from typing import List, Tuple

INPUT_PATH = "input-files/2024/01/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/01/input-sample.txt"

is_sample_mode = False


def read_input() -> List[str]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.readlines()


def parse_input(input: List[str]) -> Tuple[List[int], List[int]]:
    list1, list2 = zip(*(map(int, line.split()) for line in input))
    return list(list1), list(list2)


def sum_total_distance():
    input = read_input()
    list1, list2 = parse_input(input)

    list1.sort()
    list2.sort()

    total_distance = sum(abs(a - b) for a, b in zip(list1, list2))

    print(f"Total distance: {total_distance}")


def calc_similarity_score():
    input = read_input()
    list1, list2 = parse_input(input)

    counter = Counter()
    counter.update(list2)

    # The similarity score is the sum of the product of the value
    # and the count of the value in the first list
    similarity_score = sum(value * counter[value] for value in list1)

    print(f"Similarity score: {similarity_score}")


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true",
                        help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global is_sample_mode
        is_sample_mode = True

    sum_total_distance()
    calc_similarity_score()


if __name__ == "__main__":
    main()
