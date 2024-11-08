import * as fs from "fs";

const input = fs.readFileSync("input-files/2023/01/input.txt", "utf-8");

function parseLine(line: string): number {
  let first = 0;
  let last = 0;

  // first digit
  for (const c of line.split("")) {
    const num = parseInt(c);
    if (!isNaN(num)) {
      first = num;
      break;
    }
  }

  // last digit
  for (const c of line.split("").reverse()) {
    const num = parseInt(c);
    if (!isNaN(num)) {
      last = num;
      break;
    }
  }

  return first * 10 + last;
}

let total = 0;
input.split("\n").forEach((line) => {
  const lineValue = parseLine(line);
  console.log(`Line value for ${line}: ${lineValue}`);

  total += lineValue;
});

console.log(`Total: ${total}`);
