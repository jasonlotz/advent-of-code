// Run: ts-node src/2015/1/ts/day-1.ts

import * as fs from "fs";

const input = fs.readFileSync("src/2015/1/input.txt", "utf-8");

// Part 1
const left = input.split("(").length - 1;
const right = input.split(")").length - 1;

console.log(`Floor #: ${left - right}`);

// Part 2
const chars = input.split("");
let floor = 0;

for (let i = 0; i < chars.length; i++) {
    if (chars[i] === "(") {
        floor++;
    } else if (chars[i] === ")") {
        floor--;
    }

    if (floor === -1) {
        console.log(`Basement at position #: ${i + 1}`);
        break;
    }
}
