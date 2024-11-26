from typing import List

INPUT_PATH = "input-files/2015/02/input.txt"


class Present:
    def __init__(self, length: int, width: int, height: int):
        self.length = length
        self.width = width
        self.height = height

    def get_surface_area(self):
        return (2 * self.length * self.width
                + 2 * self.width * self.height
                + 2 * self.height * self.length)

    def get_smallest_side_area(self):
        return min([self.length * self.width,
                    self.width * self.height,
                    self.height * self.length])

    def get_volume(self):
        return self.length * self.width * self.height

    def get_smallest_perimeter(self):
        return min([2 * self.length + 2 * self.width,
                    2 * self.width + 2 * self.height,
                    2 * self.height + 2 * self.length])

    def get_ribbon_length(self):
        return self.get_smallest_perimeter() + self.get_volume()


def parse_input() -> List[str]:
    with (open(INPUT_PATH, "r")) as input_file:
        return input_file.readlines()


def part1():
    input = parse_input()
    total_paper_area = 0
    for line in input:
        l, w, h = map(int, line.split("x"))
        present = Present(l, w, h)
        total_paper_area += (present.get_surface_area() +
                             present.get_smallest_side_area())

    print(f"Part 2 (total area): {total_paper_area}")


def part2():
    input = parse_input()
    total_ribbon_length = 0
    for line in input:
        l, w, h = map(int, line.split("x"))
        present = Present(l, w, h)
        total_ribbon_length += present.get_ribbon_length()

    print(f"Part 1 (total ribbon): {total_ribbon_length}")


def main():
    part1()
    part2()


if __name__ == "__main__":
    main()
