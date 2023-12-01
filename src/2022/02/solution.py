WIN_POINTS = 6
LOSE_POINTS = 0
TIE_POINTS = 3

shape_points_map = {
    "X": 1,
    "Y": 2,
    "Z": 3
}

outcome_points_map = {
    ("X", "A"): TIE_POINTS,
    ("Y", "B"): TIE_POINTS,
    ("Z", "C"): TIE_POINTS,

    ("X", "C"): WIN_POINTS,
    ("X", "B"): LOSE_POINTS,

    ("Y", "A"): WIN_POINTS,
    ("Y", "C"): LOSE_POINTS,

    ("Z", "B"): WIN_POINTS,
    ("Z", "A"): LOSE_POINTS,
}

throw_win_map = {
    "C": "X",
    "A": "Y",
    "B": "Z"
}

throw_tie_map = {
    "A": "X",
    "B": "Y",
    "C": "Z",
}

throw_lose_map = {
    "B": "X",
    "C": "Y",
    "A": "Z"
}


def round(theirs, ours):
    points = shape_points_map[ours]
    points += outcome_points_map[ours, theirs]

    return points


def throw_round(theirs, ours):
    new_ours = ours

    if ours == "X":
        # Lose
        new_ours = throw_lose_map[theirs]
    elif ours == "Y":
        # Tie
        new_ours = throw_tie_map[theirs]
    elif ours == "Z":
        # Win
        new_ours = throw_win_map[theirs]

    points = shape_points_map[new_ours]
    points += outcome_points_map[new_ours, theirs]

    return points


input_path = "src/2022/02/input.txt"

total_points = 0
throw_total_points = 0

with open(input_path, "r") as input_file:
    for line in input_file:
        theirs, ours = line.replace("\n", "").split(" ")
        total_points += round(theirs, ours)
        throw_total_points += throw_round(theirs, ours)

print(f"Total points: {total_points}")
print(f"Throw total points: {throw_total_points}")
