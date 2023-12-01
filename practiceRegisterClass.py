import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/registerClass'  # サーバーの実際のURLに置き換えてください

# 送信するデータを準備

class_id = 10000
class_name = '線形代数の応用 100'
class_day = 'Monday'
class_time = 5
class_unit = 2
must_flag = 1
teacher_name = 'zempo'
room = 'W103'
term = 'Fall'
dmcs = 'DM'
data = {
        'class_id' : class_id, 'name' : class_name, 
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
