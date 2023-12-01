# A, X = Rock, 1 point
# B, Y = Paper, 2 points
# C, Z = Scissors, 3 points

shape_points_map = {
    'X': 1,
    'Y': 2,
    'Z': 3
}

WIN_POINTS = 6
LOSE_POINTS = 0
TIE_POINTS = 3

outcome_points = {
    ('X', 'A'): TIE_POINTS,
    ('Y', 'B'): TIE_POINTS,
    ('Z', 'C'): TIE_POINTS,

    ('X', 'C'): WIN_POINTS,
    ('X', 'B'): LOSE_POINTS,

    ('Y', 'A'): WIN_POINTS,
    ('Y', 'C'): LOSE_POINTS,

    ('Z', 'B'): WIN_POINTS,
    ('Z', 'A'): LOSE_POINTS,
}


def round(theirs, ours):
    points = shape_points_map[ours]
    points += outcome_points[ours, theirs]

    return points


input_string = 'src/2022/02/input.txt'

total_points = 0

with open(input_string, 'r') as input_file:
    for line in input_file:
        theirs, ours = line.replace("\n", "").split(" ")
        total_points += round(theirs, ours)

print(f"Total points: {total_points}")
