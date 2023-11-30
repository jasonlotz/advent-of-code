// Run: ts-node src/2015/2/ts/day-2.ts

import * as fs from "fs";

function calculateTotals(input: string) {
    let totalWrappingArea = 0;
    let totalRibbonLength = 0;

    input.split("\n").forEach((line) => {
        const dimentions = line.split("x");
        const l = parseInt(dimentions[0]);
        const w = parseInt(dimentions[1]);
        const h = parseInt(dimentions[2]);

        totalWrappingArea += calculateSurfaceArea(l, w, h);
        totalWrappingArea += calculateSmallestSideArea(l, w, h);

        totalRibbonLength += calculateRibbonWrapLength(l, w, h);
        totalRibbonLength += calculateRibbonBowLength(l, w, h);
    });

    return { totalWrappingArea, totalRibbonLength };
}

function calculateSurfaceArea(l: number, w: number, h: number) {
    return 2 * l * w + 2 * w * h + 2 * h * l;
}

function calculateSmallestSideArea(l: number, w: number, h: number) {
    const areas = [l * w, w * h, h * l];
    return Math.min(...areas);
}

function calculateRibbonWrapLength(l: number, w: number, h: number) {
    const lengths = [2 * l + 2 * w, 2 * w + 2 * h, 2 * h + 2 * l];
    return Math.min(...lengths);
}

function calculateRibbonBowLength(l: number, w: number, h: number) {
    const length = l * w * h;
    return length;
}

const input = fs.readFileSync("src/2015/2/input.txt", "utf-8");
const totals = calculateTotals(input);

console.log(`Total wrapping area: ${totals.totalWrappingArea}`);
console.log(`Total ribbon length: ${totals.totalRibbonLength}`);
