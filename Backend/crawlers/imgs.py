

from urllib import request
from bs4 import BeautifulSoup
import requests as rq
import base64


def download_img(img_url):
    print(img_url)
    r = rq.get(img_url, stream=True)
    print(r.status_code)  # 返回状态码
    pic_base64 = None
    if r.status_code == 200:
        pic_base64 = base64.b64encode(r.content)
        print("done")
    return pic_base64


class Imgs:
    def get():
        result = []

        response = request.urlopen('https://news.tongji.edu.cn/tjyw1.htm')
        html = response.read().decode('utf-8')
        soup = BeautifulSoup(html)

        def has_id(tag):
            return tag.name == 'li' and tag.has_attr('id')

        for event in soup.find_all(has_id):
            img = {"base64": "","href":""}
            img['base64'] = download_img('https://news.tongji.edu.cn/'+event.a.img["src"])
            if "http" in event.a["href"]:
                img["href"] = event.a["href"]
            else:
                img["href"] = "https://news.tongji.edu.cn/"+event.a["href"]
            result.append(img)
        return result
