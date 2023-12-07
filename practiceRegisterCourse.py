import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/registerCourse'  # サーバーの実際のURLに置き換えてください
name = input("username>>")
classid = input("classid>>")
# 送信するデータを準備
data = {'person_name': name, 'class_id': classid}  # 送信するデータをここで設定

# HTTP POSTリクエストを送信
response = requests.post(url, data)

# レスポンスをJSONとしてパース
json_response = response.json()

# レスポンスを出力
print(json_response)