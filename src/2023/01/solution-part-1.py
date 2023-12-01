input_path = "src/2023/01/input.txt"

with open(input_path, "r") as input_file:
    total = 0

    for line in input_file:
        first = 0
        last = 0

        # first digit
        for c in line[::1]:
            if c.isnumeric():
                first = c
                break

        # last digit
        for c in line[::-1]:
            if c.isnumeric():
                last = c
                break

        total += (int(first) * 10) + int(last)

print(f"Total: {total}")
