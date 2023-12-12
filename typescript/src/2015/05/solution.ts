import fs from "fs";

function isNicePart1(inputString: string): boolean {
  let vowelCount = 0;
  let hasDoubleLetter = false;
  let hasFoundNaughtyString = false;

  for (let i = 0; i < inputString.length; i++) {
    if ("aeiou".indexOf(inputString[i]) !== -1) {
      vowelCount += 1;
    }

    if (i > 0 && inputString[i] == inputString[i - 1] && !hasDoubleLetter) {
      hasDoubleLetter = true;
    }

    if (
      (i > 0 && inputString[i] == "b" && inputString[i - 1] == "a") ||
      (inputString[i] == "d" && inputString[i - 1] == "c") ||
      (inputString[i] == "q" && inputString[i - 1] == "p") ||
      (inputString[i] == "y" && inputString[i - 1] == "x")
    ) {
      hasFoundNaughtyString = true;
    }
  }

  return vowelCount >= 3 && hasDoubleLetter && !hasFoundNaughtyString;
}

function isNicePart2(inputString: string): boolean {
  let hasRepeatingLetter = false;
  let hasPairOfLetters = false;

  for (let i = 0; i < inputString.length; i++) {
    if (i > 1 && inputString[i] == inputString[i - 2]) {
      hasRepeatingLetter = true;
    }

    if (!hasPairOfLetters) {
      const pair = inputString[i] + inputString[i + 1];
      for (let j = i + 2; j < inputString.length - 1; j++) {
        if (inputString[j] + inputString[j + 1] == pair) {
          hasPairOfLetters = true;
          break;
        }
      }
    }
  }

  if (hasRepeatingLetter && hasPairOfLetters) {
    return true;
  }

  return false;
}

function main() {
  const input = fs.readFileSync("input-files/2015/05/input.txt", "utf-8");
  const lines = input.split("\n");

  let niceStringsPart1 = 0;
  let naughtyStringsPart1 = 0;

  let niceStringsPart2 = 0;
  let naughtyStringsPart2 = 0;

  for (const line of lines) {
    if (isNicePart1(line)) {
      niceStringsPart1 += 1;
    } else {
      naughtyStringsPart1 += 1;
    }

    if (isNicePart2(line)) {
      niceStringsPart2 += 1;
    } else {
      naughtyStringsPart2 += 1;
    }
  }

  console.log(`Nice strings (part 1): ${niceStringsPart1}`);
  console.log(`Naughty strings (part 1): ${naughtyStringsPart1}`);

  console.log(`Nice strings (part 2): ${niceStringsPart2}`);
  console.log(`Naughty strings (part 2): ${naughtyStringsPart2}`);
}

function testPart1() {
  console.assert(isNicePart1("ugknbfddgicrmopn") == true, "ugknbfddgicrmopn is nice");
  console.assert(isNicePart1("aaa") == true, "aaa is nice");
  console.assert(isNicePart1("jchzalrnumimnmhp") == false, "jchzalrnumimnmhp is naughty because no double letter");
  console.assert(isNicePart1("haegwjzuvuyypxyu") == false, "haegwjzuvuyypxyu is naughty because contains xy");
  console.assert(isNicePart1("haegwjzuvuyypcdu") == false, "haegwjzuvuyypcdu is naughty because contains cd");
  console.assert(isNicePart1("vukmuyfstgmscuab") == false, "vukmuyfstgmscuab is naughty because contains ab");
  console.assert(
    isNicePart1("dvszwmarrgswjxmb") == false,
    "dvszwmarrgswjxmb is naughty because contains only one vowel"
  );
}

function testPart2() {
  console.assert(isNicePart2("qjhvhtzxzqqjkmpb") == true, "qjhvhtzxzqqjkmpb is nice");
  console.assert(isNicePart2("xxyxx") == true, "xxyxx is nice");
  console.assert(isNicePart2("uurcxstgmygtbstg") == false, "uurcxstgmygtbstg is naughty because no repeating letter");
  console.assert(isNicePart2("ieodomkazucvgmuy") == false, "ieodomkazucvgmuy is naughty because no pair of letters");
}

main();
testPart1();
testPart2();
