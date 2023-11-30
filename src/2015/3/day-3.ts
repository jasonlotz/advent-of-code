// Run: ts-node src/2015/3/ts/day-3.ts

import * as fs from "fs";

const input = fs.readFileSync("src/2015/3/input.txt", "utf-8");

const position = [0, 0];

function countVisited(input: string): number {
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
    return visited.size;
}

console.log(countVisited(input));
