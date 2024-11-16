INPUT_PATH = "input-files/2023/01/input.txt"

with open(INPUT_PATH, "r") as input_file:
    total = 0

    for line in input_file:
        digits = [c for c in line if c.isnumeric()]
        if digits:
            first = digits[0]
            last = digits[-1]
            total += (int(first) * 10) + int(last)

print(f"Total: {total}")
