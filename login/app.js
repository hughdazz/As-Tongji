var express = require('express');
var ele = require('./ele');
var app = express();
const puppeteer = require('puppeteer');
const { application } = require('express');
pages = []
app.use('/public', express.static('public'));

app.get('/', function (req, res) {
    res.send('Hello World');
})
app.get('/start', function (req, res) {
    let page = ele.start()
    pages.push(page)
    res.send('ok');
})
app.get('/submit', function (req, res) {
    console.log(req.query.username + req.query.password);
    page = pages.pop()
    page.then((page) => {
        ele.input(page, req.query.username, req.query.password).then(
            (json) => {
                res.send(json)
            }
        )
    })
})
var server = app.listen(8081, function () {
    var host = server.address().address
    var port = server.address().port
    console.log("应用实例，访问地址为 http://127.0.0.1:%s", port)
})

