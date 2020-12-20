const fs = require('fs')
const data = fs.readFileSync("./test_input").toString()

const lines = data.split("\n")

const bagIndices = new Map()
const unhash = (i) => Object.keys(bagIndices);

const hash = (bag) => {
    if (bagIndices.has(bag)) {
        return bagIndices.get(bag)
    }
    const i = bagIndices.size
    bagIndices.set(bag, i)
    return i
}

const width = (() => Math.ceil(lines.length * 2))()

const adjMatrix = []
const plot = (x, y, val) => adjMatrix[x + y * width] = val
const get = (x, y) => adjMatrix[x + y * width]
const edge = (bag, child, count) => plot(hash(bag), hash(child), count)

for (const line of lines) {
    const [colour, properties] = line.replace(/bag(s)?/g, '').replace('.', '').split('contain')
    const props = properties.split(',').map((property) => property.replace(' ', ''))

    for (const prop of props) {
        const [count, ...attribs] = prop.replace(/bag(s)?/g, '').trim().split(' ')
        if (count == 'no') continue
        const childColour = attribs.join(' ')
        
        console.log(count, '===', colour, '->', childColour, hash(colour), 'and', hash(childColour))
        edge(colour, childColour, parseInt(count))
    }
}

console.log(unhash(2))

const seen = new Map()
let proc = []
proc.push(unhash(2))

while (proc.length) {
    const i = proc.shift()
    if (seen.has(i)) continue
    let c = 0
    bagIndices.forEach((idx) => {
        if (get(i, idx) > 0) c += get(i, idx) > 0
    })
    console.log(c)
    seen.set(i, true)
}