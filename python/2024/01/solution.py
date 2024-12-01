import argparse
from typing import Dict, List, Tuple

INPUT_PATH = "input-files/2024/01/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/01/input-sample.txt"

isSampleMode = False


def read_input() -> List[str]:
    if isSampleMode:
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

    total_distance = 0
    for i in range(len(list1)):
        total_distance += abs(list1[i] - list2[i])

    print(f"Total distance: {total_distance}")


def calc_similarity_score():
    input = read_input()
    list1, list2 = parse_input(input)

    # Create a dict of the first list elements
    list1_dict: Dict[int, int] = {
        value: 0 for _, value in enumerate(list1)}

    # Loop over the second list and count the occurrences of each element
    for _, value in enumerate(list2):
        if value in list1_dict:
            list1_dict[value] += 1

    # The similarity score is the sum of the product of the value
    # and the count of the value in the first list
    similarity_score = 0
    for _, value in enumerate(list1):
        similarity_score += value * list1_dict[value]

    print(f"Similarity score: {similarity_score}")


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true",
                        help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global isSampleMode
        isSampleMode = True

    sum_total_distance()
    calc_similarity_score()


if __name__ == "__main__":
    main()
