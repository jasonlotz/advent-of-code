input_path = "src/2023/01/input.txt"

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

with open(input_path, "r") as input_file:
    total = 0

    for line in input_file:
        numbers = []

        for i in range(len(line)):
            if line[i].isnumeric():
                numbers.append(line[i])
                continue

            for key in number_map.keys():
                if line[i:i+len(key)] == key:
                    numbers.append(number_map[key])
                    continue

        line_total = (int(numbers[0]) * 10) + int(numbers[-1])
        print(line, numbers, line_total)
        total += line_total


print(f"Total: {total}")