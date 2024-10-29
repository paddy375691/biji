import requests
import json

def dingtalk_msg(url, content):
    headers = {'Content-Type': 'application/json;charset=utf-8'}
    data = {
            "msgtype": "text",
            "text": {
                "content": content
            }
        }
    res = requests.post(url=url, headers=headers, data=json.dumps(data))
    return res.text
if __name__ == "__main__":
    content = "发布失败啦！！！"
    url = "https://oapi.dingtalk.com/robot/send?access_token=6b55c27305708f34c37165f22814965e48fc7cd92efca115ee7a27a807dc6345"
    print(dingtalk_msg(url, content))