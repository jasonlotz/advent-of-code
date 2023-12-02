from string import ascii_lowercase
from string import ascii_uppercase


def find_matching_character_in_halves(str):
    mid_index = len(str) // 2
    left_str = str[:mid_index]
    right_str = str[mid_index:]

    left_set = set(left_str)
    right_set = set(right_str)
    match_set = left_set.intersection(right_set)

    if match_set:
        return match_set.pop()
    else:
        return None


char_score_map = {}
for index, letter in enumerate(ascii_lowercase, 1):
    char_score_map[letter] = index
for index, letter in enumerate(ascii_uppercase, 1):
    char_score_map[letter] = index + 26


input_path = 'src/2022/03/input.txt'
total_score = 0

with open(input_path, 'r') as input_file:
    for line in input_file:
        match = find_matching_character_in_halves(line)
        if match:
            score = char_score_map[match]
        total_score += score

print(f'Total Score: {total_score}')
