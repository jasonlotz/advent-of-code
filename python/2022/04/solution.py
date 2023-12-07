INPUT_PATH = "input-files/2022/04/input.txt"

assignments = open(INPUT_PATH).read().splitlines()
fully_contained = 0
partially_contained = 0

for assignment in assignments:
    pairs = assignment.split(",")
    pair1 = list(map(int, pairs[0].split("-")))
    pair2 = list(map(int, pairs[1].split("-")))

    if ((pair1[0] >= pair2[0] and pair1[1] <= pair2[1]) or
            (pair2[0] >= pair1[0] and pair2[1] <= pair1[1])):
        fully_contained += 1

    if ((pair1[0] <= pair2[0] and pair1[1] >= pair2[0]) or
            (pair1[0] <= pair2[1] and pair1[1] >= pair2[1]) or
            (pair2[0] <= pair1[0] and pair2[1] >= pair1[0]) or
            (pair2[0] <= pair1[0] and pair2[1] >= pair1[1])):
        partially_contained += 1


print(f"Fully contained: {fully_contained}")
print(f"Partially contained: {partially_contained}")
