import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/registerClass'  # サーバーの実際のURLに置き換えてください

# 送信するデータを準備

class_name = input("class_name:")
class_day = input("day:")
class_time = int(input("period:"))
class_unit = int(input("unit:"))
must_flag = int(input("must:"))
teacher_name = input("teacher_name:")
room = input("room:")
term = input("term:")
dmcs = input("department:")
data = {
        'name' : class_name, 
        'day' : class_day, 'period' : class_time, 
        'unit' : class_unit, 'must' : must_flag,
        'teacher' : teacher_name, 'room' : room, 
        'term' : term, 'department' : dmcs
        }  # 送信するデータをここで設定

# HTTP POSTリクエストを送信
response = requests.post(url, data=data)

# レスポンスをJSONとしてパース
json_response = response.json()

# レスポンスを出力
print(json_response)
