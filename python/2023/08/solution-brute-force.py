from itertools import cycle

INPUT_PATH = "input-files/2023/08/input.txt"


def part_1():
  parse_input_file()

  intructions_cycle = cycle(instructions)

  next_node = "AAA"

  steps = 0
  while (True):
    steps += 1
    instruction = next(intructions_cycle)

    next_node = mapping_dict[next_node][0 if instruction[0] == "L" else 1]

    print(steps, instruction, next_node)

    if next_node == "ZZZ":
      print("Found ZZZ on step (part 1):", steps)
      break


def part_2():
  parse_input_file()

  intructions_cycle = cycle(instructions)

  next_nodes = [key for key in mapping_dict.keys() if key.endswith("A")]

  steps = 0
  while (True):
    steps += 1
    instruction = next(intructions_cycle)

    for i, next_node in enumerate(next_nodes):
      next_node = mapping_dict[next_node][0 if instruction[0] == "L" else 1]
      next_nodes[i] = next_node

    all_ends_in_z = True
    for next_node in next_nodes:
      if not next_node.endswith("Z"):
        all_ends_in_z = False
        break

    if all_ends_in_z:
      print("They all end in Z on step (part 2):", steps)
      break


def parse_input_file():
  global mapping_dict
  global instructions

  mapping_dict = {}

  instructions_input, _, * \
      map_input = open("input-files/2023/08/input.txt").read().splitlines()

  instructions = instructions_input.strip()

  for mapping in map_input:
    node, elements = mapping.split("=")

    node = node.strip()

    mapping_dict[node] = elements.strip().replace(
        "(", "").replace(
        ")", "").split(", ")


def main():
  part_1()
  part_2()


if __name__ == "__main__":
  main()
