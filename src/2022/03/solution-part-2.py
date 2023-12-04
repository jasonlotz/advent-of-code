from string import ascii_lowercase
from string import ascii_uppercase

INPUT_PATH = 'src/2022/03/input.txt'


def find_matching_character(strs: list):
    sets = []

    for str in strs:
        sets.append(set(str))

    match_set = set.intersection(*sets)

    if match_set:
        return match_set.pop()
    else:
        return None


char_score_map = {}
for index, letter in enumerate(ascii_lowercase, 1):
    char_score_map[letter] = index
for index, letter in enumerate(ascii_uppercase, 1):
    char_score_map[letter] = index + 26


def main():
    total_score = 0

    with open(INPUT_PATH, 'r') as input_file:
        group = []

        for index, line in enumerate(input_file, 1):
            group.append(line.strip())

            if index % 3 == 0:
                match = find_matching_character(group)
                if match:
                    score = char_score_map[match]
                total_score += score
                group = []

    print(f'Total Score: {total_score}')


if __name__ == "__main__":
    main()
