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
function newPosition(
  direction: Direction,
  currentPosition: Position,
): Position {
  if (direction === "N") {
    return [currentPosition[0], currentPosition[1] + 1];
  } else if (direction === "E") {
    return [currentPosition[0] + 1, currentPosition[1]];
  } else if (direction === "S") {
    return [currentPosition[0], currentPosition[1] - 1];
  } else if (direction === "W") {
    return [currentPosition[0] - 1, currentPosition[1]];
  }
  throw new Error("Invalid direction");
}

function updatePositions(
  currentDirection: Direction,
  currentPosition: Position,
  distance: number,
): Position[] {
  const result: Position[] = [];

  for (let i = 0; i < distance; i++) {
    const newPos = newPosition(currentDirection, currentPosition);
    result.push(newPos);
    currentPosition = newPos;
  }

  return result;
}

function checkPositionsForHq(
  newPositions: Position[],
  visitedPositions: Position[],
): boolean {
  for (const position of newPositions) {
    if (
      visitedPositions.some((p) => p[0] === position[0] && p[1] === position[1])
    ) {
      console.log("Part 2:");
      console.log(`First visited position twice: ${position}`);
      console.log(
        `Distance: ${Math.abs(position[0]) + Math.abs(position[1])} blocks away`,
      );
      return true;
    }
  }
  return false;
}

function main() {
  let currentPosition: Position = [0, 0];
  let currentDirection: Direction = "N";
  const visitedPositions: Position[] = [];
  let foundHq = false;

  for (const instructionStr of input.split(", ")) {
    const instruction = parseInstruction(instructionStr);
    const [leftOrRight, distance] = instruction;

    const newDirection = changeDirection(currentDirection, leftOrRight);
    const newPositions = updatePositions(
      newDirection,
      currentPosition,
      distance,
    );

    currentDirection = newDirection;
    currentPosition = newPositions.at(-1)!;

    if (!foundHq) {
      foundHq = checkPositionsForHq(newPositions, visitedPositions);
    }

    visitedPositions.push(...newPositions);
  }

  console.log("Part 1:");
  console.log(`Current position: ${currentPosition}`);
  console.log(
    `Distance: ${Math.abs(currentPosition[0]) + Math.abs(currentPosition[1])} blocks away`,
  );
}

main();
