# 事情がない限り、このコードではなくサイトの方からサインアップをしてください
# サーバーが落ちているときには @fukuda_taichi か @takahashi_ryota に連絡をください

import requests

# サーバーのエンドポイントURLを設定
url = 'http://52.69.43.211/registerPerson'  # サーバーの実際のURLに置き換えてください

# 送信するデータを準備
flag = input('WORNING: このコードは、ユーザー情報に影響を与えます．本当に実行しますか?\nYes >> Y\nNo >> n\n')
if flag == "Y":
    name = input("name:")
    data = {'name': name}  # 送信するデータをここで設定

    # HTTP POSTリクエストを送信
    response = requests.post(url, data)

    # レスポンスをJSONとしてパース
    json_response = response.json()

    # レスポンスを出力
    print(json_response)
