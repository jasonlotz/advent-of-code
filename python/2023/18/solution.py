INPUT_PATH = "input-files/2023/18/input-sample.txt"


def parse_input():
  input = open(INPUT_PATH, "r").read().splitlines()

  instructions = []
  for line in input:
    split_line = line.split(" ")
    instructions.append((split_line[0], int(split_line[1])))

  return instructions


def part1():
  parse_input()


def part2():
  pass


def main():
  part1()
  part2()


if __name__ == "__main__":
  main()
