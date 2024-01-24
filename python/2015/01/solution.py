INPUT_PATH = "input-files/2015/01/input.txt"

with open(INPUT_PATH, "r") as input_file:
    for line in input_file:
        left = len(line.split("("))
        right = len(line.split(")"))

        print("Floor #: " + str(left - right))
