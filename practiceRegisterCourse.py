import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/registerCource'  # サーバーの実際のURLに置き換えてください

# 送信するデータを準備
data = {'person_name': 'taichi', 'class_id' : 3}  # 送信するデータをここで設定

# HTTP POSTリクエストを送信
response = requests.post(url, data=data)

# レスポンスをJSONとしてパース
json_response = response.json()

# レスポンスを出力
print(json_response)
