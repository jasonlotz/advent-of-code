INPUT_PATH = "input-files/2015/05/input.txt"

nice_strings = 0
naughty_strings = 0


with open(INPUT_PATH, "r") as input_file:
    for line in input_file:
        vowel_count = 0
        double_letter = False
        found_naughty_string = False

        for i in range(len(line)):
            if line[i] in "aeiou":
                vowel_count += 1
            if i > 0 and line[i] == line[i - 1] and not double_letter:
                double_letter = True
            if (i > 0 and (line[i] == "b" and line[i - 1] == "a") or
                    (line[i] == "d" and line[i - 1] == "c") or
                    (line[i] == "q" and line[i - 1] == "p") or
                    (line[i] == "y" and line[i - 1] == "x")):
                found_naughty_string = True

        if vowel_count >= 3 and double_letter and not found_naughty_string:
            nice_strings += 1
            continue
        else:
            naughty_strings += 1


print(f"Nice strings: {nice_strings}")
print(f"Naughty strings: {naughty_strings}")
