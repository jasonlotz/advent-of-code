import * as fs from "fs";

type KeyPosition = [number, number];
type Direction = "U" | "D" | "L" | "R";
type KeyPad = string[][];

const input = fs.readFileSync("input-files/2016/02/input.txt", "utf-8");

function move(
  currentKey: KeyPosition,
  direction: Direction,
  keypad: KeyPad,
): KeyPosition {
  let newKey: KeyPosition | undefined = undefined;

  if (direction === "U") {
    newKey = [currentKey[0] - 1, currentKey[1]];
  } else if (direction === "D") {
    newKey = [currentKey[0] + 1, currentKey[1]];
  } else if (direction === "L") {
    newKey = [currentKey[0], currentKey[1] - 1];
  } else if (direction === "R") {
    newKey = [currentKey[0], currentKey[1] + 1];
  }

  if (newKey && checkKey(newKey, keypad)) {
    return newKey;
  } else {
    return currentKey;
  }
}

function checkKey(key: KeyPosition, keypad: KeyPad): boolean {
  return keypad[key[0]][key[1]] !== "";
}

function processInstructions(keypad: KeyPad) {
  const instructions = input.split("\n");
  let currentKeyPosition: KeyPosition = [2, 2];
  let code = "";

  instructions.forEach((instruction) => {
    for (let i = 0; i < instruction.length; i++) {
      const direction: Direction = instruction[i] as Direction;
      const newKey = move(currentKeyPosition, direction, keypad);
      currentKeyPosition = newKey;
    }

    code += keypad[currentKeyPosition[0]][currentKeyPosition[1]];
  });

  console.log(`Part 2: ${code}`);
}

function part1() {
  const keypad = [
    ["", "", "", "", ""],
    ["", "1", "2", "3", ""],
    ["", "4", "5", "6", ""],
    ["", "7", "8", "9", ""],
    ["", "", "", "", ""],
  ];

  processInstructions(keypad);
}

function part2() {
  const keypad = [
    ["", "", "", "", "", "", ""],
    ["", "", "", "1", "", "", ""],
    ["", "", "2", "3", "4", "", ""],
    ["", "5", "6", "7", "8", "9", ""],
    ["", "", "A", "B", "C", "", ""],
    ["", "", "", "D", "", "", ""],
    ["", "", "", "", "", "", ""],
  ];

  processInstructions(keypad);
}

function main() {
  part1();
  part2();
}

main();
