input_string = "src/2022/01/input.txt"

with open(input_string, "r") as input_file:
    current_calories = 0
    calories = []

    for line in input_file:
        if line == "\n":
            calories.append(current_calories)
            current_calories = 0
            continue
        else:
            current_calories += int(line)

print("Final max calories: " + str(max(calories)))

sorted_calories = sorted(calories, reverse=True)

top_3_calories = 0

for i in range(0, 3):
    calories = sorted_calories[i]
    top_3_calories += calories

    print(f"Calories #: {calories}")

print(f"Final top 3 calories: {str(top_3_calories)}")
