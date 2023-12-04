import re
from pprint import pprint
from collections import defaultdict, namedtuple

input_path = "src/2023/03/input.txt"

Point = namedtuple("Point", "x y")

PartNumber = namedtuple("PartNumber", "number point")


def parse_line(line, y_coord) -> (list[PartNumber], set[Point], set[Point]):
    part_numbers: list[PartNumber] = []
    symbols: set[Point] = set()
    asterisks: set[Point] = set()

    for x_coord, c in enumerate(line):
        if not c.isdigit() and c != ".":
            symbols.add(Point(int(x_coord), y_coord))
            if c == "*":
                asterisks.add(Point(int(x_coord), y_coord))

    numbers_dict = dict((m.start(), int(m.group()))
                        for m in re.finditer(r'\d+', line))

    for num in numbers_dict.items():
        part_numbers.append(PartNumber(num[1], Point(num[0], y_coord)))

    return part_numbers, symbols, asterisks


def validate_part_number(part_number: PartNumber, symbols: set[Point]) -> bool:
    number_len = len(str(part_number.number))
    x = part_number.point.x
    y = part_number.point.y

    start = x-1
    end = x + number_len + 1

    # row above
    for i in range(start, end):
        if Point(i, y-1) in symbols:
            return True

    # same row
    if Point(start, y) in symbols:
        return True

    if Point(end-1, y) in symbols:
        return True

    # row below
    for i in range(start, end):
        if Point(i, y+1) in symbols:
            return True

    return False


def find_touching_asterisks(part_number: PartNumber, asterisks: set[Point]) -> list[Point]:
    x = part_number.point.x
    y = part_number.point.y
    number_len = len(str(part_number.number))
    start = x-1
    end = x + number_len + 1

    touching_asterisks: list[Point] = []

    # row above
    for i in range(start, end):
        if Point(i, y-1) in asterisks:
            touching_asterisks.append(Point(i, y-1))

    # same row
    if Point(start, y) in asterisks:
        touching_asterisks.append(Point(start, y))

    if Point(end-1, y) in asterisks:
        touching_asterisks.append(Point(end-1, y))

    # row below
    for i in range(start, end):
        if Point(i, y+1) in asterisks:
            touching_asterisks.append(Point(i, y+1))

    return touching_asterisks


def main():
    all_part_numbers = []
    all_symbols = set()
    all_asterisks = set()
    valid_part_numbers = []
    total = 0
    gearTotal = 0

    asterisk_dict = defaultdict(list)

    with open(input_path, "r") as input_file:
        for i, line in enumerate(input_file):
            numbers, symbols, asterisks = parse_line(line.strip(), i)
            all_part_numbers.extend(numbers)
            all_symbols = set.union(all_symbols, symbols)
            all_asterisks = set.union(all_asterisks, asterisks)

        for part_number in all_part_numbers:
            if validate_part_number(part_number, all_symbols):
                valid_part_numbers.append(part_number)
                total += part_number.number

            touching_asterisks = find_touching_asterisks(
                part_number, all_asterisks)

            for touching_asterisk in touching_asterisks:
                asterisk_dict[touching_asterisk].append(
                    part_number.number)

        for point, ids in asterisk_dict.items():
            print(point, ids)
            if len(ids) == 2:
                gearRatio = ids[0] * ids[1]
                print("Gear Ratio:", gearRatio)
                gearTotal += gearRatio

    print("\nAll Part Numbers:")
    pprint(all_part_numbers)
    print("\nAll Symbols:")
    pprint(all_symbols)
    print("\nAll Asterisks:")
    pprint(all_asterisks)
    print("\nValid Part Numbers:")
    pprint(valid_part_numbers)
    print("\nAsterisk Dict:")
    pprint(asterisk_dict)
    print(f"\nTotal: {total}")
    print(f"Gear Total: {gearTotal}")


if __name__ == "__main__":
    main()
