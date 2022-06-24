const puppeteer = require('puppeteer');
class Login {
    constructor(stream) {
        puppeteer.launch({ headless: false, defaultViewport: { width: 1280, height: 720 } }).
            then((browser) => {
                this.browser = browser;
                return browser.newPage();
            }).then((page) => {
                this.page = page;
                this.page.on('response',
                    response => {
                        if (response.url().includes('nidp/app/login?sid=0&sid=0/getCaptcha=1')) {
                            // 获取验证码响应后返回到前端
                            response.text().then((body) => {
                                stream(body)
                            })
                        }
                        // session_id
                        if (response.url().includes('api/sessionservice/session/login')) {
                            response.text().then((body) => {
                                stream(body)
                            })
                        }
                        // common_msg
                        if (response.url().includes('api/commonservice/commonMsgPublish/findMyCommonMsgPublish')) {
                            response.text().then((body) => {
                                stream(body)
                            })
                        }
                        // exam
                        if (response.url().includes('api/electionservice/undergraduateExamQuery/getStudentListPage')) {
                            response.text().then((body) => {
                                stream(body)
                            })
                        }
                        // timetab
                        if (response.url().includes('api/electionservice/reportManagement/findStudentTimetab')) {
                            response.text().then((body) => {
                                stream(body)
                            })
                        }
                        // canvas
                        if (response.url().includes('api/v1/planner/items')) {
                            response.text().then((body) => {
                                stream(body)
                            })
                        }
                    }
                )
            })
    }

    start() {
        this.page.goto('https://ids.tongji.edu.cn:8443/nidp/app/login?id=Login&sid=0&option=credential&sid=0&target=https%3A%2F%2Fids.tongji.edu.cn%3A8443%2Fnidp%2Foauth%2Fnam%2Fauthz%3Fscope%3Dprofile%26response_type%3Dcode%26redirect_uri%3Dhttps%3A%2F%2F1.tongji.edu.cn%2Fapi%2Fssoservice%2Fsystem%2FloginIn%26client_id%3D5fcfb123-b94d-4f76-89b8-475f33efa194')
            .then((_) => {

            })
    }
    input(username, password) {
        this.page.$("#username").then((username_ele) => {
            return username_ele.type(username)
        }).then((_) => {
            return this.page.$("#password")
        }).then((password_ele) => {
            return password_ele.type(password)
        }).then((_) => {
            return this.page.$('#reg')
        }).then((submit_ele) => {
            return submit_ele.click()
        }).then((_) => {
            return this.page.$('.verify-img-panel')
        }).then((_) => { })
    }
    click(points) {
        this.page.$('.back-img').then((img_ele) => {
            return img_ele.boundingBox()
        }).then((box) => {
            this.page.mouse.click(box.x + points[0].x, box.y + points[0].y, { delay: 80 }).then((_) => {
                return this.page.mouse.click(box.x + points[1].x, box.y + points[1].y, { delay: 80 })
            }).then((_) => {
                return this.page.mouse.click(box.x + points[2].x, box.y + points[2].y, { delay: 80 })
            }).then((_) => {

            })
        })
    }
    exam() {
        this.page.goto('https://1.tongji.edu.cn/StuExamEnquiries')
            .then((_) => {

            })
    }
    timetab() {
        this.page.goto('https://1.tongji.edu.cn/GraduateStudentTimeTable')
            .then((_) => {

            })
    }
    canvas() {
        this.page.goto('http://canvas.tongji.edu.cn/')
            .then((_) => {

            })
    }
}
module.exports = {
    Login
}