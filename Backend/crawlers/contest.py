from urllib import request
from bs4 import BeautifulSoup


class Contest:
    def get():
        page_num = [2, 3]
        result = []

        for i in page_num:
            response = request.urlopen('https://www.saikr.com/vs?page='+str(i))
            # print(response)
            # 提取响应内容
            html = response.read().decode('utf-8')
            # 打印响应内容
            # print(html)
            soup = BeautifulSoup(html)
            # print(soup)
            for event in soup.find_all('div', class_='fl event4-1-detail-box'):
                info = {}
                
                info["title"] = event.a.attrs['title']
                info["href"] = event.a.attrs['href']
                ps = event.find_all('p')
                info["sponsor"] = ps[0].get_text()
                info["regist_time"] = ps[1].get_text()
                info["level"] = ps[2].get_text()
                info["cotest_time"] = ps[3].get_text()

                result.append(info)

        return result
