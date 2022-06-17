const ele = require('./ele')
arr = []
function add(x, y) {
    arr.push({ x: x, y: y })
}
let page = await ele.start()
let body = await ele.input(page, '2053516', 'Hugh2804856132')
await ele.click(page, arr)
