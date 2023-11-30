import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/showClassInfoAll'  # サーバーの実際のURLに置き換えてください

# HTTP POSTリクエストを送信
response = requests.get(url)

# レスポンスをJSONとしてパース
json_response = response.json()

# レスポンスを出力
print(json_response)
