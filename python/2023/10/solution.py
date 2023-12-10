from pprint import pprint

INPUT_PATH = "input-files/2023/10/input-sample.txt"


N_S = "|"  # is a vertical pipe connecting north and south
E_W = "-"  # is a horizontal pipe connecting east and west
N_E = "L"  # is a 90-degree bend connecting north and east
N_W = "J"  # is a 90-degree bend connecting north and west
S_W = "7"  # is a 90-degree bend connecting south and west
S_E = "F"  # is a 90-degree bend connecting south and east
NO_PIPE = "."  # is ground; there is no pipe in this tile
ANIMAL = "S"  # is the starting position of the animal


def find_animal(maze):
    for i in range(len(maze)):
        for j in range(len(maze[i])):
            if maze[i][j] == ANIMAL:
                print(f"Animal: {i}, {j}")
                return (i, j)


def find_start_end(maze, animal):
    start_end = []

    # North
    north_coord = (animal[0]-1, animal[1])
    north_value = maze[north_coord[0]][north_coord[1]]

    print(f"North: {north_coord} -> {north_value}")
    if north_value == N_S:
        start_end.append(north_coord)

    # South
    south_coord = (animal[0]+1, animal[1])
    south_value = maze[south_coord[0]][south_coord[1]]

    print(f"South: {south_coord} -> {south_value}")
    if south_value == N_S:
        start_end.append(south_coord)

    # East
    east_coord = (animal[0], animal[1]+1)
    east_value = maze[east_coord[0]][east_coord[1]]

    print(f"East: {east_coord} -> {east_value}")
    if east_value in (E_W, N_W, S_W):
        start_end.append(east_coord)

    # West
    west_coord = (animal[0], animal[1]-1)
    west_value = maze[west_coord[0]][west_coord[1]]

    print(f"West: {west_coord} -> {west_value}")
    if west_value in (E_W, N_E, S_E):
        start_end.append(west_coord)

    print(f"Start & end: {start_end}")

    if len(start_end) != 2:
        raise Exception("Invalid maze")

    return start_end[0], start_end[1]


def follow_pipe(maze, animal):
    start, end = find_start_end(maze, animal)

    current = start
    move_count = 0

    while current != end:
        current_value = maze[current[0]][current[1]]

        if current_value == N_S:
            current = (current[0]-1, current[1])
        elif current_value == E_W:
            current = (current[0], current[1]+1)
        elif current_value == N_E:
            current = (current[0]-1, current[1])
        elif current_value == N_W:
            current = (current[0]-1, current[1])
        elif current_value == S_W:
            current = (current[0]+1, current[1])
        elif current_value == S_E:
            current = (current[0]+1, current[1])
        else:
            raise Exception("Invalid maze")

        move_count += 1
        print(f"Move: {move_count}")
        print(f"Current: {current}")


def parse_input():
    maze = []
    with open(INPUT_PATH, "r") as input_file:
        lines = input_file.readlines()

        for line in lines:
            maze.append(list(line.strip()))

    return maze


def part_1():
    maze = parse_input()
    animal = find_animal(maze)
    follow_pipe(maze, animal)


def main():
    part_1()


if __name__ == "__main__":
    main()
