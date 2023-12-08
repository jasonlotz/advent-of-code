from math import gcd

INPUT_PATH = "input-files/2023/08/input.txt"

instructions, _, * \
    map_input = open(INPUT_PATH).read().splitlines()

network = {}

for line in map_input:
    node, targets = line.split(" = ")
    network[node] = targets[1:-1].split(", ")

nodes = [key for key in network if key.endswith("A")]
cycles = []

for current_node in nodes:
    cycle = []

    current_steps = instructions
    steps = 0
    first_z = None

    while True:
        while steps == 0 or not current_node.endswith("Z"):
            steps += 1
            current_node = network[current_node][0 if current_steps[0] == "L" else 1]
            current_steps = current_steps[1:] + current_steps[0]

        cycle.append(steps)

        if first_z is None:
            first_z = current_node
            steps = 0
        elif current_node == first_z:
            break

    cycles.append(cycle)

nums = [cycle[0] for cycle in cycles]

lcm = nums.pop()

for num in nums:
    lcm = lcm * num // gcd(lcm, num)

print(lcm)
