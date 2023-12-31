import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/registerFriends'  # サーバーの実際のURLに置き換えてください

# 送信するデータを準備
my_name = input("my_name:")
your_name = input("your_name:")
data = {'my_name': my_name, 'your_name': your_name}  # 送信するデータをここで設定

# HTTP POSTリクエストを送信
response = requests.post(url, data)

# レスポンスをJSONとしてパース
json_response = response.json()

# レスポンスを出力
print(json_response)