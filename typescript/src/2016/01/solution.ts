import * as fs from "fs";

const input = fs.readFileSync("input-files/2016/01/input.txt", "utf-8");

type LeftOrRight = "L" | "R";
type Position = [number, number];
type Instruction = [LeftOrRight, number];
type Direction = "N" | "E" | "S" | "W";

function parseInstruction(rawInstruction: string): Instruction {
  const leftOrRight = rawInstruction[0] as LeftOrRight;
  const distance = Number(rawInstruction.slice(1));

  return [leftOrRight, distance];
}

function changeDirection(
  currentDirection: Direction,
  leftOrRight: LeftOrRight,
): Direction {
  if (currentDirection === "N") {
    if (leftOrRight === "L") {
      return "W";
    } else {
      return "E";
    }
  } else if (currentDirection === "E") {
    if (leftOrRight === "L") {
      return "N";
    } else {
      return "S";
    }
  } else if (currentDirection === "S") {
    if (leftOrRight === "L") {
      return "E";
    } else {
      return "W";
    }
  } else if (currentDirection === "W") {
    if (leftOrRight === "L") {
      return "S";
    } else {
      return "N";
    }
  }
  throw new Error("Invalid direction");
}

function updatePosition(
  currentDirection: Direction,
  currentPosition: Position,
  instruction: Instruction,
): Position {
  const [leftOrRight, distance] = instruction;

  if (currentDirection === "N") {
    if (leftOrRight === "L") {
      return [currentPosition[0] - distance, currentPosition[1]];
    } else {
      return [currentPosition[0] + distance, currentPosition[1]];
    }
  } else if (currentDirection === "E") {
    if (leftOrRight === "L") {
      return [currentPosition[0], currentPosition[1] + distance];
    } else {
      return [currentPosition[0], currentPosition[1] - distance];
    }
  } else if (currentDirection === "S") {
    if (leftOrRight === "L") {
      return [currentPosition[0] + distance, currentPosition[1]];
    } else {
      return [currentPosition[0] - distance, currentPosition[1]];
    }
  } else if (currentDirection === "W") {
    if (leftOrRight === "L") {
      return [currentPosition[0], currentPosition[1] - distance];
    } else {
      return [currentPosition[0], currentPosition[1] + distance];
    }
  }
  throw new Error("Invalid direction");
}

let currentPosition: Position = [0, 0];
let currentDirection: Direction = "N";

for (const instruction of input.split(", ")) {
  const parsedInstruction = parseInstruction(instruction);
  const newDirection = changeDirection(currentDirection, parsedInstruction[0]);
  const newPosition = updatePosition(
    currentDirection,
    currentPosition,
    parsedInstruction,
  );

  currentDirection = newDirection;
  currentPosition = newPosition;
}

console.log("Part 1:");
console.log(`Current position: ${currentPosition}`);
console.log(`Distance: ${Math.abs(currentPosition[0]) + Math.abs(currentPosition[1])} blocks away`);
