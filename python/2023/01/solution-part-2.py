INPUT_PATH = "input-files/2023/01/input.txt"

number_map = {
    "zero": 0,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9
}


def extract_numbers(line):
    numbers = []
    i = 0
    while i < len(line):
        if line[i].isnumeric():
            numbers.append(int(line[i]))
            i += 1
        else:
            for key, value in number_map.items():
                if line[i:i + len(key)] == key:
                    numbers.append(value)
                    i += len(key)
                    break
            else:
                i += 1
    return numbers


with open(INPUT_PATH, "r") as input_file:
    total = sum((int(numbers[0]) * 10) + int(numbers[-1])
                for line in input_file if (numbers := extract_numbers(line)))

print(f"Total: {total}")
