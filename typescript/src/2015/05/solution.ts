import fs from "fs";

function validate_input(lines: string[]) {
  let nice_strings = 0;
  let naughty_strings = 0;

  for (const line of lines) {
    if (is_nice(line)) {
      nice_strings += 1;
    } else {
      naughty_strings += 1;
    }
  }

  console.log(`Nice strings: ${nice_strings}`);
  console.log(`Naughty strings: ${naughty_strings}`);
}

function is_nice(input_string: string): boolean {
  let vowel_count = 0;
  let double_letter = false;
  let found_naughty_string = false;

  for (let i = 0; i < input_string.length; i++) {
    if ("aeiou".indexOf(input_string[i]) !== -1) {
      vowel_count += 1;
    }

    if (i > 0 && input_string[i] == input_string[i - 1] && !double_letter) {
      double_letter = true;
    }

    if (
      (i > 0 && input_string[i] == "b" && input_string[i - 1] == "a") ||
      (input_string[i] == "d" && input_string[i - 1] == "c") ||
      (input_string[i] == "q" && input_string[i - 1] == "p") ||
      (input_string[i] == "y" && input_string[i - 1] == "x")
    ) {
      found_naughty_string = true;
    }
  }

  return vowel_count >= 3 && double_letter && !found_naughty_string;
}

function main() {
  const input = fs.readFileSync("input-files/2015/05/input.txt", "utf-8");
  const lines = input.split("\n");
  validate_input(lines);
}

function test() {
  console.assert(is_nice("ugknbfddgicrmopn") == true, "ugknbfddgicrmopn is nice");
  console.assert(is_nice("aaa") == true, "aaa is nice");
  console.assert(is_nice("jchzalrnumimnmhp") == false, "jchzalrnumimnmhp is naughty because no double letter");
  console.assert(is_nice("haegwjzuvuyypxyu") == false, "haegwjzuvuyypxyu is naughty because contains xy");
  console.assert(is_nice("haegwjzuvuyypcdu") == false, "haegwjzuvuyypcdu is naughty because contains cd");
  console.assert(is_nice("vukmuyfstgmscuab") == false, "vukmuyfstgmscuab is naughty because contains ab");
  console.assert(is_nice("dvszwmarrgswjxmb") == false, "dvszwmarrgswjxmb is naughty because contains only one vowel");
}

main();
test();
