import * as fs from "fs";

const input = fs.readFileSync("input-files/2015/03/input.txt", "utf-8");

function visited(input: string): Set<string> {
  const position = [0, 0];
  const visited = new Set<string>();

  for (const char of input) {
    switch (char) {
      case "^":
        position[1]++;
        break;
      case "v":
        position[1]--;
        break;
      case ">":
        position[0]++;
        break;
      case "<":
        position[0]--;
        break;
    }

    visited.add(position.toString());
  }
  return visited;
}

function visitedWithRobo(input: string): Set<string> {
  const santaInstructions = [];
  const roboInstructions = [];

  for (let i = 0; i < input.length; i += 1) {
    if (i % 2 == 0) {
      santaInstructions.push(input[i]);
    } else {
      roboInstructions.push(input[i]);
    }
  }

  const santaVisits = visited(santaInstructions.join(""));
  const roboVisits = visited(roboInstructions.join(""));

  const combinedVisits = new Set([...Array.from(santaVisits), ...Array.from(roboVisits)]);

  return combinedVisits;
}

console.log(`Visited - Santa: ${visited(input).size}`);
console.log(`Visited - Santa w/ Robo-Santa: ${visitedWithRobo(input).size}`);
