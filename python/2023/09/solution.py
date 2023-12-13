INPUT_PATH = "input-files/2023/09/input.txt"


def parse_input_file():
  sequences = []

  with open(INPUT_PATH, "r") as input_file:
    for line in input_file:
      sequences.append(list(map(int, line.strip().split(" "))))

  return sequences


def extrapolate_next_value(sequence):
  next_value = 0

  sequence_stack = generate_sequence_stack(sequence)
  sequence_stack.reverse()

  for s in sequence_stack:
    next_value += s[-1]

  return next_value


def extrapolate_first_value(sequence):
  first_value = 0

  sequence_stack = generate_sequence_stack(sequence)
  sequence_stack.reverse()

  for s in sequence_stack:
    first_value = s[0] - first_value

  return first_value


def generate_sequence_stack(sequence):
  sequence_stack = [sequence]

  while True:
    differences = []

    for i in range(len(sequence) - 1):
      differences.append(sequence[i + 1] - sequence[i])

    sequence_stack.append(differences)

    if all(difference == 0 for difference in differences):
      break

    sequence = differences

  return sequence_stack


def part_1():
  sequences = parse_input_file()

  next_values = []
  for sequence in sequences:
    next_values.append(extrapolate_next_value(sequence))

  print(f"Sum of next values: {sum(next_values)}")


def part_2():
  sequences = parse_input_file()

  first_values = []
  for sequence in sequences:
    first_values.append(extrapolate_first_value(sequence))

  print(f"Sum of first values: {sum(first_values)}")


def main():
  part_1()
  part_2()


if __name__ == "__main__":
  main()
