INPUT_PATH = "input-files/2023/15/input.txt"


def parse_input():
  return open(INPUT_PATH, "r").read().split(",")


def hash_algorithm(value):
  hash_value = 0

  for ch in value:
    hash_value += ord(ch)

    hash_value = (hash_value * 17) % 256

  return hash_value


def part1():
  input = parse_input()
  total = 0

  for value in input:
    hashed_value = hash_algorithm(value)
    total += hashed_value

  print(f"Total (part 1): {total}")


def part2():
  # TODO: Implement part 2
  # for each value in input:
  # parse the first two letter and get the hash value (which is the map key)
  # parse the remainder of the string
  # if -, then remove the label from the linked list of labels for the map key
  # if =X, then append (or replace if exists) the label to the linked list of labels for the map key
  pass


def main():
  part1()
  part2()


if __name__ == "__main__":
  main()
