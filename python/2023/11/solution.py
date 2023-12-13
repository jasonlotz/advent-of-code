INPUT_PATH = "input-files/2023/11/input-sample.txt"


def parse_input():
  return open(INPUT_PATH).read().splitlines()


# Find Manhattan distance between every pair of galaxies
def manhattan_distance_total(grid, expand_scale=0):
  empty_rows = [r for r, row in enumerate(
      grid) if all(ch == "." for ch in row)]
  empty_cols = [c for c, col in enumerate(
      zip(*grid)) if all(ch == "." for ch in col)]

  galaxies = [(r, c) for r, row in enumerate(grid)
              for c, ch in enumerate(row) if ch == "#"]

  print(f"Galaxies: {galaxies}")

  total = 0

  for i, (r1, c1) in enumerate(galaxies):
    for (r2, c2) in galaxies[:i]:
      for r in range(min(r1, r2), max(r1, r2)):
        total += expand_scale if r in empty_rows else 1
      for c in range(min(c1, c2), max(c1, c2)):
        total += expand_scale if c in empty_cols else 1

  return total


def part1():
  grid = parse_input()
  total = manhattan_distance_total(grid, 2)
  print(f"Total (part 1): {total}")


def part2():
  grid = parse_input()
  total = manhattan_distance_total(grid, 1000000)
  print(f"Total (part 2): {total}")


def main():
  part1()
  part2()


if __name__ == "__main__":
  main()
