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
  total = 0
  for value in parse_input():
    hashed_value = hash_algorithm(value)
    total += hashed_value

  print(f"Total (part 1): {total}")


def part2():
  boxes = [[] for _ in range(256)]
  focal_lengths = {}

  for instruction in parse_input():
    if "-" in instruction:
      label = instruction[:-1]
      index = hash_algorithm(label)
      if label in boxes[index]:
        boxes[index].remove(label)
    else:
      label, length = instruction.split("=")
      length = int(length)

      index = hash_algorithm(label)
      if label not in boxes[index]:
        boxes[index].append(label)

      focal_lengths[label] = length

  total = 0

  for box_number, box in enumerate(boxes, 1):
    for lens_slot, label in enumerate(box, 1):
      total += box_number * lens_slot * focal_lengths[label]

  print(f"Total part 2: {total}")


def main():
  part1()
  part2()


if __name__ == "__main__":
  main()
