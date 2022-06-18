const puppeteer = require('puppeteer');
data = []
async function start() {
    const browser = await puppeteer.launch({ headless: false, defaultViewport: { width: 1280, height: 720 } });
    const page = await browser.newPage();// 打开一个新页面
    page.on('response',
        function (response) {
            if (response.url().includes('nidp/app/login?sid=0&sid=0/getCaptcha=1')) {
                // 获取验证码响应后返回到前端
                response.text().then((body) => {
                    data.push(JSON.parse(body))
                })
            }
        }
    )
    await page.goto('https://ids.tongji.edu.cn:8443/nidp/app/login?id=Login&sid=0&option=credential&sid=0&target=https%3A%2F%2Fids.tongji.edu.cn%3A8443%2Fnidp%2Foauth%2Fnam%2Fauthz%3Fscope%3Dprofile%26response_type%3Dcode%26redirect_uri%3Dhttps%3A%2F%2F1.tongji.edu.cn%2Fapi%2Fssoservice%2Fsystem%2FloginIn%26client_id%3D5fcfb123-b94d-4f76-89b8-475f33efa194');
    return page
}
async function input(page, username, password) {
    const username_ele = await page.$("#username");
    await username_ele.type(username);
    const password_ele = await page.$("#password");
    await password_ele.type(password);
    const submit_ele = await page.$('#reg');
    await submit_ele.click();
    let panel = await page.$('.verify-img-panel');
    if (panel) {
        return body = data.pop();
    }
}
async function click(page, points) {
    page.on('response',
        function (response) {
            // console.log(response.url())
            if (response.url().includes('api/sessionservice/session/login')) {
                    response.text().then((body) => {
                        console.log(JSON.parse(body))
                    })
            }
        }
    )
    let panel = await page.$('.back-img');
    let box = await panel.boundingBox();
    console.log(box)

    await page.mouse.click(box.x + points[0].x, box.y + points[0].y, { delay: 100 });
    await page.mouse.click(box.x + points[1].x, box.y + points[1].y, { delay: 100 });
    await page.mouse.click(box.x + points[2].x, box.y + points[2].y, { delay: 100 });
}
module.exports = {
    input,
    start,
    click
}