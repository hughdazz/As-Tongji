let ws = require("nodejs-websocket");
console.log("开始建立连接...")
let n = require("./new")
let server = ws.createServer(function (conn) {
    function stream(str) {
        conn.sendText(str)
    }
    let L = new n.Login(stream)
    conn.on("text", function (str) {
        let obj = JSON.parse(str)
        if (obj["action"] == "start") {
            L.start()
        }
        if (obj["action"] == "input") {
            L.input(obj["username"], obj["password"])
        }
        if (obj["action"] == "click") {
            let points = [{ x: obj["points"][0][0], y: obj["points"][0][1] }, { x: obj["points"][1][0], y: obj["points"][1][1] }, { x: obj["points"][2][0], y: obj["points"][2][1] }]
            L.click(points)
        }
        if (obj["action"] == "exam") {
            L.exam()
        }
        if (obj["action"] == "timetab") {
            L.timetab()
        }
        if (obj["action"] == "canvas") {
            L.canvas()
        }
        console.log("message:" + str)
        // conn.sendText("My name is Web Xiu!")
    })
    conn.on("close", function (code, reason) {
        console.log("关闭连接")
    })
    conn.on("error", function (code, reason) {
        console.log("异常关闭")
    })
}).listen(8001)
console.log("WebSocket建立完毕")