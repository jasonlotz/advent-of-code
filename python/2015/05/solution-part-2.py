INPUT_PATH = "input-files/2015/05/input.txt"

nice_strings = 0
naughty_strings = 0


def check_line(line):
  repeating_letter = False
  pair_of_letters = False

  for i in range(len(line)):
    # Repeating letter
    if i > 1 and line[i] == line[i - 2] and not repeating_letter:
      repeating_letter = True

    # Pair of letters
    if not pair_of_letters:
      pair = line[i:i + 2]
      for j in range(i + 2, len(line) - 1):
        if line[j:j + 2] == pair:
          pair_of_letters = True
          break

    if repeating_letter and pair_of_letters:
      return True

  return False


with open(INPUT_PATH, "r") as input_file:
  for line in input_file:
    if check_line(line):
      nice_strings += 1
    else:
      naughty_strings += 1

print(f"Nice strings: {nice_strings}")
print(f"Naughty strings: {naughty_strings}")
