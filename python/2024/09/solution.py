import argparse
from typing import List

INPUT_PATH = "input-files/2024/09/input.txt"
SAMPLE_INPUT_PATH = "input-files/2024/09/input-sample.txt"

is_sample_mode = False


def parse_input() -> str:
    if is_sample_mode:
        file_path = SAMPLE_INPUT_PATH
    else:
        file_path = INPUT_PATH

    with open(file_path, "r") as input_file:
        return input_file.read().strip()


def create_disk_map(data: List[int]) -> List[int]:
    disk_size = sum(data)

    disk_map = [-1] * disk_size

    pointer = 0

    for i, value in enumerate(data):
        if i % 2 == 0:
            for _ in range(value):
                disk_map[pointer] = i//2
                pointer += 1
        else:
            for _ in range(value):
                disk_map[pointer] = -1
                pointer += 1

    return disk_map


def compact_disk_map_by_block(disk_map: List[int]) -> None:
    blanks = [i for i, x in enumerate(disk_map) if x == -1]

    for i in blanks:
        while disk_map[-1] == -1:
            disk_map.pop()
        if len(disk_map) <= i:
            break
        disk_map[i] = disk_map.pop()


def part1():
    data = list(map(int, parse_input()))

    disk_map = create_disk_map(data)

    compact_disk_map_by_block(disk_map)

    checksum = (sum(i * x for i, x in enumerate(disk_map) if x != -1))

    print(f"Checksum (part 1): {checksum}")


def part2():
    data = list(map(int, parse_input()))

    files = {}
    blanks = []

    file_id = 0
    position = 0

    for i, x in enumerate(data):
        if i % 2 == 0:
            # Process files
            if x == 0:
                raise ValueError("File size cannot be 0")
            files[file_id] = (position, x)
            file_id += 1
        else:
            # Process blanks
            if x != 0:
                blanks.append((position, x))
        position += x

    while file_id > 0:
        file_id -= 1  # Decrement to get the last file
        position, size = files[file_id]  # Get the last file position and size

        # For each blank, check if the file can fit in it
        for i, (start, length) in enumerate(blanks):
            if start >= position:
                # The blank is after the file, so we can clean up all
                # the blanks before it
                blanks = blanks[:i]
                break
            if size <= length:
                # It fits! Update the file
                files[file_id] = (start, size)

                # Update the blank
                if size == length:
                    blanks.pop(i)
                else:
                    blanks[i] = (start + size, length - size)
                break

    print(f"Files @ end: {files}")

    checksum = 0
    for file_id, (position, size) in files.items():
        for x in range(position, position + size):
            checksum += file_id * x

    print(f"Checksum (part 2): {checksum}")


def main():
    parser = argparse.ArgumentParser(description="Sample mode?")
    parser.add_argument("--sample", action="store_true",
                        help="Use the sample input")
    args = parser.parse_args()

    if args.sample:
        global is_sample_mode
        is_sample_mode = True

    part1()
    part2()


if __name__ == "__main__":
    main()
