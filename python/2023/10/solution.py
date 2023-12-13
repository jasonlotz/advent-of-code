INPUT_PATH = "input-files/2023/10/input.txt"


def parse_input():
  with open(INPUT_PATH, "r") as input_file:
    input = input_file.read().strip().splitlines()

  return input


class PipeLoop:
  def __init__(self, input):
    self.map = input
    self.directions = {
        "left": (0, -1), "right": (0, 1), "up": (-1, 0), "down": (1, 0)}
    self.pipe_directions = {
        ("L", "left"): "up",
        ("F", "up"): "right",
        ("7", "right"): "down",
        ("J", "down"): "left",
        ("F", "left"): "down",
        ("7", "up"): "left",
        ("J", "right"): "up",
        ("L", "down"): "right",
        ("-", "left"): "left",
        ("-", "right"): "right",
        ("|", "up"): "up",
        ("|", "down"): "down",
    }
    self.start = self.find_start_point()
    self.loop = self.follow_pipe()
    self.enclosed = self.find_enclosed()

  def find_start_point(self):
    for i in range(len(self.map)):
      for j in range(len(self.map[0])):
        if self.map[i][j] == "S":
          return i, j

  def traverse(self, current_pipe, current_direction):
    (r, c) = current_pipe

    next_direction = self.pipe_directions[(
        self.map[r][c], current_direction)]

    next_pipe = tuple(
        map(sum, zip(current_pipe, self.directions[next_direction])))

    return next_pipe, next_direction

  def from_start_point(self):
    for direction in self.directions:
      (dr, dc) = tuple(
          map(sum, zip(self.start, self.directions[direction])))

      if (self.map[dr][dc], direction) in self.pipe_directions:
        return (dr, dc), direction

  def follow_pipe(self):
    current, direction = self.from_start_point()
    (r, c) = current
    loop = {current: (self.map[r][c], direction)}

    while current != self.start:
      current, direction = self.traverse(current, direction)
      (r, c) = current
      loop[current] = self.map[r][c], direction
    # Replace "S"
    else:
      for pipe in ("L", "F", "7", "J", "|", "-"):
        if (pipe, direction) in self.pipe_directions:
          next_direction = self.pipe_directions[(pipe, direction)]
          next_pipe = tuple(
              map(sum, zip(current, self.directions[next_direction])))
          if next_pipe == list(loop)[0]:
            loop[current] = (pipe, direction)
            return loop

  def find_enclosed(self):
    enclosed = set()
    inside_loop = {}
    loop_direction = 0

    # Get loop"s direction
    for r in range(len(self.map)):
      for c in range(len(self.map[0])):
        if (r, c) in self.loop:
          if self.loop[(r, c)][1] == "up" or self.pipe_directions[self.loop[(r, c)]] == "up":
            loop_direction = "up"
          elif self.loop[(r, c)][1] == "down" or self.pipe_directions[self.loop[(r, c)]] == "down":
            loop_direction = "down"
          # Get most outside loop's direction to check
          if inside_loop == {}:
            inside_loop = {loop_direction: True}
        elif loop_direction in inside_loop:
          enclosed.add((r, c))
    return enclosed


def main():
  input = parse_input()
  pipes = PipeLoop(input)
  print("Part 1:", len(pipes.loop) // 2)
  print("Part 2:", len(pipes.enclosed))


if __name__ == "__main__":
  main()
