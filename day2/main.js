var fs = require('fs')
var buffer = fs.readFileSync("./input")

const lines = buffer.toString().split("\n")
const data = lines.map((line) => {
    const parts = line.split(" ")
    return {range: parts[0].split("-"), letter: parts[1].charAt(0), input: parts[2]}
})

const isValid = ({range, letter, input}) => {
    const [min, max] = range
    const letterCount = [...input].filter((c) => c === letter).length
    return letterCount >= min && letterCount <= max
}

const isValidBetter = ({range, letter, input}) => {
    const [min, max] = range
    return input[min-1] == letter ^ input[max-1] == letter
}

console.log(data)
const valid = data.filter((part) => isValidBetter(part))
console.log(valid.length, "passwords")