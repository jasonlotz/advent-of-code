import * as fs from "fs";

const input = fs.readFileSync("input-files/2023/01/input.txt", "utf-8");

const numberMap: { [key: string]: number } = {
    zero: 0,
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9,
};

let total = 0;

input.split("\n").forEach((line) => {
    const numbers: [number?] = [];

    for (let i = 0; i < line.length; i++) {
        // 1, 2, 3 ..
        const number = parseInt(line[i]);
        if (!isNaN(number)) {
            numbers.push(number);
            continue;
        }

        // one, two, three ..
        for (const key in numberMap) {
            if (line.substring(i, i + key.length) === key) {
                numbers.push(numberMap[key]);
                continue;
            }
        }
    }

    const lineTotal = numbers[0]! * 10 + numbers[numbers.length - 1]!;
    console.log(line, numbers, lineTotal);
    total += lineTotal;
});

console.log(`Total: ${total}`);
