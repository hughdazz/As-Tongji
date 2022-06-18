const puppeteer = require('puppeteer');

async function run() {

    const browser = await puppeteer.launch({ headless: false, defaultViewport: { width: 1280, height: 720 } });
    const page = await browser.newPage();// 打开一个新页面
    page.on('response',
        function (response) {
            if (response.url().includes('nidp/app/login?sid=0&sid=0/getCaptcha=1')) {
                // 获取验证码响应后返回到前端
                response.text().then((body) => {
                    console.log(JSON.parse(body)['repCode'])
                })
            }
        }
    )
    await page.goto('https://ids.tongji.edu.cn:8443/nidp/app/login?id=Login&sid=0&option=credential&sid=0&target=https%3A%2F%2Fids.tongji.edu.cn%3A8443%2Fnidp%2Foauth%2Fnam%2Fauthz%3Fscope%3Dprofile%26response_type%3Dcode%26redirect_uri%3Dhttps%3A%2F%2F1.tongji.edu.cn%2Fapi%2Fssoservice%2Fsystem%2FloginIn%26client_id%3D5fcfb123-b94d-4f76-89b8-475f33efa194');

    const username_ele = await page.$("#username");
    await username_ele.type("2053516");
    const password_ele = await page.$("#password");
    await password_ele.type("Hugh2804856132");

    const submit_ele = await page.$('#reg');
    await submit_ele.click();

    let panel = await page.$('.verify-img-panel');
    
    // 调用页面内Dom对象的 screenshot 方法进行截图
    // 为什么有偏移呢？小编也很疑惑
    // await panel.screenshot({
    //     path: 'panel.png'
    // });
}
run();