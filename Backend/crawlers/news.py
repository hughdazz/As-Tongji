from urllib import request
from bs4 import BeautifulSoup


class News:
    def get():
        tail = ['.htm', '/3150.htm']
        result = []

        for i in tail:

            response = request.urlopen('https://news.tongji.edu.cn/tjkx1'+i)
            html = response.read().decode('utf-8')
            soup = BeautifulSoup(html)

            def has_id(tag):
                return tag.name == 'li' and tag.has_attr('id')

            for event in soup.find_all(has_id):
                info = {}
                if "http" in event.a["href"]:
                    info["href"] = event.a["href"]
                else:
                    info["href"] = "https://news.tongji.edu.cn/"+event.a["href"]
                info["title"] = event.a["title"]
                info["summary"] = event.find_all('span')[1].get_text()
                result.append(info)
        return result
