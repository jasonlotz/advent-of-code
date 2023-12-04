import * as fs from "fs";

interface Dimension {
    l: number;
    w: number;
    h: number;
}

function calculateTotals(input: string) {
    let totalWrappingArea = 0;
    let totalRibbonLength = 0;

    input.split("\n").forEach((line) => {
        const dimensions = line.split("x");
        const dimension: Dimension = {
            l: parseInt(dimensions[0]),
            w: parseInt(dimensions[1]),
            h: parseInt(dimensions[2]),
        };

        totalWrappingArea += calculateSurfaceArea(dimension);
        totalWrappingArea += calculateSmallestSideArea(dimension);

        totalRibbonLength += calculateRibbonWrapLength(dimension);
        totalRibbonLength += calculateRibbonBowLength(dimension);
    });

    return { totalWrappingArea, totalRibbonLength };
}

function calculateSurfaceArea(dimension: Dimension): number {
    const { l, w, h } = dimension;
    return 2 * l * w + 2 * w * h + 2 * h * l;
}

function calculateSmallestSideArea(dimension: Dimension): number {
    const { l, w, h } = dimension;
    const areas = [l * w, w * h, h * l];
    return Math.min(...areas);
}

function calculateRibbonWrapLength(dimension: Dimension): number {
    const { l, w, h } = dimension;
    const lengths = [2 * l + 2 * w, 2 * w + 2 * h, 2 * h + 2 * l];
    return Math.min(...lengths);
}

function calculateRibbonBowLength(dimension: Dimension): number {
    const { l, w, h } = dimension;
    const length = l * w * h;
    return length;
}

const input = fs.readFileSync("input-files/2015/02/input.txt", "utf-8");
const totals = calculateTotals(input);

console.log(`Total wrapping area: ${totals.totalWrappingArea}`);
console.log(`Total ribbon length: ${totals.totalRibbonLength}`);
