const fs = require('fs')
const data = fs.readFileSync("./input").toString()

const lines = data.split("\n")

const groups = []
let ansSet = []
for (const line of lines) {
    if (line.length == 0) {
        groups.push(ansSet)
        ansSet = []
        continue
    }
    ansSet.push([...line])
}

// ans set contains array of arrays.
// ans set -> person_ans[]

// part 1. the unique question 'id' for each ans group
const part1 = groups
    .map((answers) => answers.flat())
    .map((answers) => new Set([...answers]))
    .map((answers) => answers.size)
    .reduce((a, b) => a + b)
console.log(part1)

// part 2.
const commonAnswers = (a) => {
    // create sets for each ans group
    const sets = a.map((ans) => new Set([...ans]))

    const has = (a) => new Set([...sets.map((set) => set.has(a))]).size == 1

    const x = a.flat().filter((ans) => has(ans))
    return new Set([...x])
}

const part2 = groups
    .map((answers) => commonAnswers(answers))
    .map((answers) => answers.size)
    .reduce((a, b) => a + b)
console.log(part2)

// console.log(groups.map((g) => new Set(g)).map((g) => g.size).reduce((a, b) => a + b))