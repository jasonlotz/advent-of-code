import re
from pprint import pprint
from collections import namedtuple

input_path = "src/2023/03/input.txt"

Point = namedtuple("Point", "x y")

PartNumber = namedtuple("PartNumber", "number point")


def parse_line(line, y_coord) -> (list[PartNumber], set[Point]):
    part_numbers: list[PartNumber] = []
    symbols: set[Point] = set()

    for x_coord, c in enumerate(line):
        if not c.isdigit() and c != ".":
            symbols.add(Point(int(x_coord), y_coord))

    numbers_dict = dict((m.start(), int(m.group()))
                        for m in re.finditer(r'\d+', line))

    for num in numbers_dict.items():
        part_numbers.append(PartNumber(num[1], Point(num[0], y_coord)))

    return part_numbers, symbols


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


def main():
    all_part_numbers = []
    all_symbols = set()
    valid_part_numbers = []
    total = 0

    with open(input_path, "r") as input_file:
        for i, line in enumerate(input_file):
            numbers, symbols = parse_line(line.strip(), i)
            all_part_numbers.extend(numbers)
            all_symbols = set.union(all_symbols, symbols)

        for part_number in all_part_numbers:
            if validate_part_number(part_number, all_symbols):
                valid_part_numbers.append(part_number)
                total += part_number.number

    pprint(all_part_numbers)
    pprint(all_symbols)
    pprint(valid_part_numbers)
    print(f"Total: {total}")


if __name__ == "__main__":
    main()
