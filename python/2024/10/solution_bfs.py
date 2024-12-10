import argparse
from collections import deque
from typing import List

INPUT_PATH = "input-files/2024/10/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/10/input-sample.txt"

is_sample_mode = False


def parse_input() -> List[List[int]]:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    grid = []

    with open(file_path, "r") as input_file:
        grid = [[int(char) for char in line.strip()]
                for line in input_file.readlines()]

    print(grid)

    return grid


def bfs(grid, start_col, start_row):
    rows, cols = len(grid), len(grid[0])
    queue = deque([(start_col, start_row, 0)])  # (col, row, target)
    visited = set()
    paths = 0

    while queue:
        col, row, target = queue.popleft()

        if (col, row) in visited:
            continue

        visited.add((col, row))

        if grid[col][row] == 9:
            paths += 1
            continue

        # Add all possible directions to the queue
        # where the next cell is within the grid and
        # the target is the next number
        for d_col, d_row in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            n_col, n_row = col + d_col, row + d_row
            if (0 <= n_col < rows and 0 <= n_row < cols
                    and grid[n_col][n_row] == target + 1):
                queue.append((n_col, n_row, target + 1))

    return paths


# Note: it would be easy to combine this with the previous function, but I
# wanted to keep the two separate for examples of how to use BFS in different
# ways.  The main difference here is that we keep track of the path as we go
# and we do not use a visited set to keep track of visited nodes which allows
# us to find all unique paths.
def bfs_all_unique(grid, start_col, start_row):
    rows, cols = len(grid), len(grid[0])
    # (col, row, target, path)
    queue = deque([(start_col, start_row, 0, [(start_col, start_row)])])
    unique_paths = []

    while queue:
        col, row, target, path = queue.popleft()

        if grid[col][row] == 9:
            unique_paths.append(path)
            continue

        # Add all possible directions to the queue
        # where the next cell is within the grid and
        # the target is the next number
        for d_col, d_row in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            n_col, n_row = col + d_col, row + d_row
            if (0 <= n_col < rows and 0 <= n_row < cols
                    and grid[n_col][n_row] == target + 1):
                queue.append((n_col, n_row, target + 1,
                             path + [(n_col, n_row)]))

    return len(unique_paths)


def count_paths_bfs(grid):
    total_paths = 0
    total_unique_paths = 0
    for col in range(len(grid)):
        for row in range(len(grid[0])):
            if grid[col][row] == 0:
                paths = bfs(grid, col, row)
                unique_paths = bfs_all_unique(grid, col, row)
                total_paths += paths
                total_unique_paths += unique_paths

    print(f"Total paths: {total_paths}")
    print(f"Total unique paths: {total_unique_paths}")


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true",
                        help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global is_sample_mode
        is_sample_mode = True

    data = parse_input()

    count_paths_bfs(data)


if __name__ == "__main__":
    main()
