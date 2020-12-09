const fs = require('fs')
const data = fs.readFileSync("./input").toString()

const sleft = (r) => r.splice(0, Math.ceil(r.length / 2))
const sright = (r) => r.splice(Math.ceil(r.length / 2), r.length)

const seq = (len) => Array.from(Array(len).keys())

const calcPos = (line) => {
    let rows = seq(127)
    let cols = seq(8)

    for (const step of [...line]) {
        switch (step) {
            case 'F':
                rows = sleft(rows)
                break
            case 'B':
                rows = sright(rows)
                break
            case 'L':
                cols = sleft(cols)
                break
            case 'R':
                cols = sright(cols)
                break
        }
    }

    const [row] = rows
    const [col] = cols
    return [row, col]
}

const seats = data.split("\n").map((line) => calcPos(line))
const seatPos = seats.map(([row, col]) => row * 8 + col)

console.log('highest seat', seatPos.sort((a, b) => b - a)[0])

// part 2 gross solution

let seatData = Array(127 * 8)
seatPos.forEach((i) => seatData[i] = 'x')

for (let y = 0; y < 8; y++) {
    let line = ''
    for (let x = 0; x < 127; x++) {
        const pos = x + y * 127
        const dp = seatData[pos] || '_'
        if (dp == '_' && seatData[pos+1] == 'x' && seatData[pos-1] == 'x') {
            console.log('missing seat', pos)
        }
        line += dp
    }
    console.log(line)
}