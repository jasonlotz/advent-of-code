// Run: ts-node src/2015/day-1/ts/part-1.ts

import * as fs from "fs";

const input = fs.readFileSync("src/2015/day-1/input.txt", "utf-8");

const left = input.split("(").length - 1;
const right = input.split(")").length - 1;

console.log(`Floor #: ${left - right}`);
