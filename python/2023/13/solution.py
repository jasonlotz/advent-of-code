INPUT_PATH = "input-files/2023/13/input.txt"


def parse_input():
    input = open(INPUT_PATH, "r").read().split("\n\n")

    squares = [line.splitlines() for line in input]

    return squares


def find_mirror(square):
    for r in range(1, len(square)):
        above = square[:r][::-1]
        below = square[r:]

        # Trick for making the array length the same equivalent to the shorter array
        above = above[:len(below)]
        below = below[:len(above)]

        if above == below:
            return r

    return 0


def find_smudge(square):
    for r in range(1, len(square)):
        above = square[:r][::-1]
        below = square[r:]

        if sum(sum(0 if a == b else 1 for a, b in zip(x, y)) for x, y in zip(above, below)) == 1:
            return r

    return 0


def part1():
    squares = parse_input()

    total = 0
    for square in squares:
        print(square)
        row = find_mirror(square)
        print(row)
        total += row * 100

        col = find_mirror(list(zip(*square)))
        total += col

    print(f"Total (part 1): {total}")


def part2():
    squares = parse_input()

    total = 0
    for square in squares:
        print(square)
        row = find_smudge(square)
        print(row)
        total += row * 100

        col = find_smudge(list(zip(*square)))
        total += col

    print(f"Total (part 2): {total}")


def main():
    part1()
    part2()


if __name__ == "__main__":
    main()
