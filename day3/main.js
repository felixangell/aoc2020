const fs = require('fs')
const data = fs.readFileSync("./input").toString()

const lines = data.split("\n")

// the lines should be same width
const width = (i) => lines[i].length
const height = lines.length

const traverseBy = (xSlope, ySlope) => {
    let treeCount = 0
    const traverse = (xa, ya) => {
        console.log(xa, xa % width(ya), width(ya))
        const i = lines[ya][xa % width(ya)]
        treeCount += i === '#' ? 1 : 0
    
        if (ya + ySlope >= height) {
            console.log(ya, height)
            return
        }
        traverse(xa + xSlope, ya + ySlope)
    }
    traverse(xSlope, ySlope)
    return treeCount
}

const slopes = [
    [1,1],
    [3,1],
    [5,1],
    [7,1],
    [1,2]
]

// part 2
console.log(slopes.map(([x, y]) => traverseBy(x, y)).reduce((a, b) => a * b))