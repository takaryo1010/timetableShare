import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/removeCourse'  # サーバーの実際のURLに置き換えてください

# 送信するデータを準備
person_id = input("person_id>>")
class_id = input("class_id>>")
data = {'class_id': class_id,'person_id': person_id}  # 送信するデータをここで設定

# HTTP POSTリクエストを送信
response = requests.post(url, data)

# レスポンスをJSONとしてパース
json_response = response.json()

# レスポンスを出力
print(json_response)
