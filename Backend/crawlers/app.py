import requests
from news import News
from lecture import Lecture
from contest import Contest
from flask import Flask
app = Flask(__name__)


@app.route('/')
def hello_world():
    return 'Hello, World!'


@app.route('/get')
def get():
    r1 = Contest.get()
    r2 = Lecture.get()
    r3 = News.get()
    for i in r1:
        requests.post(url='http://127.0.0.1:8080/public/',
                    json={"tags": ["contest"], "data": i, "expire": 60*60*24})
    for i in r2:
        requests.post(url='http://127.0.0.1:8080/public/',
                    json={"tags": ["lecture"], "data": i, "expire": 60*60*24})
    for i in r3:
        requests.post(url='http://127.0.0.1:8080/public/',
                    json={"tags": ["news"], "data": i, "expire": 60*60*24})

    return {"msg": "ok"}


if __name__ == "__main__":
    app.run()
