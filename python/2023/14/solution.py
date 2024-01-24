INPUT_PATH = "input-files/2023/14/input.txt"


def parse_input():
    grid = open(INPUT_PATH, "r").read()
    grid = list(zip(*grid.splitlines()))  # Rotate the grid 90 degrees

    return grid


def tilt_row(row):
    new_row = [ch for ch in row]
    pointer = 0

    for i in range(len(row)):
        if row[i] == "O":
            new_row[pointer], new_row[i] = new_row[i], new_row[pointer]
            pointer += 1

        if new_row[i] == "#":
            pointer = i + 1

    return new_row


def total_load(grid):
    total_load = 0

    for row in grid:
        weight_multiplier = len(row)

        for i in range(len(row)):
            if row[i] == "O":
                total_load += weight_multiplier
            weight_multiplier -= 1

    return total_load


def part1():
    grid = parse_input()
    new_grid = []

    for row in grid:
        titled_row = tilt_row(row)

        new_grid.append(titled_row)

    total = total_load(new_grid)

    print(f"Total (part 1): {total}")


def part2():
    pass


def main():
    part1()
    part2()


if __name__ == "__main__":
    main()
