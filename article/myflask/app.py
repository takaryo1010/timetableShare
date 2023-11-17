from flask import Flask, render_template, request
import requests

app = Flask(__name__)


@app.route('/')
def index():
    # ここで講義のデータを取得する
    lectures = {
        "Monday": ["微積文法の応用", "", "プログラミング (C/C++)", "プログラミング (C/C++)", "人工知能"],
        "Tuesday": ["西洋近現代史", "科学英語", "", "", "最適化"],
        "Wednesday": ["", "時事英語", "", "", "データ構造とアルゴリズム 2"],
        "Thursday": ["", "", "", "", ""],
        "Friday": ["", "", "プロジェクト", "データベース", ""],
        "Friday": ["", "", "プロジェクト", "データベース", ""],
        "Saturday": ["", "", "", "", ""],
        "Ondemand": ["CF 特論", "IS 特論", "", "", ""],
    }
    return render_template('lectures.html', timetable=lectures)


@app.route("/home")
def index_home():
    return render_template('home.html')


@app.route("/login")
def index_login():
    return render_template('login.html')


@app.route("/timetable_registration")
def index_timetable_registration():
    return render_template('timetable_registration.html')


@app.route("/lecture_list")
def index_lecture_list():
    # サーバーのエンドポイントURLを設定
    url = 'http://52.69.43.211/showClassInfoAll'  # サーバーの実際のURLに置き換えてください

    # HTTP POSTリクエストを送信
    response = requests.get(url)

    # レスポンスをJSONとしてパース
    json_response = response.json()

    # レスポンスを出力
    print(json_response)

    return render_template('lecture_list.html', json=json_response)


@app.route("/lecture_creation", methods=['GET', 'POST'])
def index_lecture_creation():
    if request.method == 'POST':
        class_name = request.form.get('class_name')
        class_day = request.form.get('class_day')
        class_time = request.form.get('class_time')
        class_unit = request.form.get('class_unit')
        must_flag = request.form.get('must_flag')
        teacher_name = request.form.get('teacher_name')
        room = request.form.get('room')
        term = request.form.get('term')
        dmcs = request.form.get('dmcs')
        data = {
            'name': class_name,
            'day': class_day, 'piriod': class_time,
            'unit': class_unit, 'must': must_flag,
            'teacher': teacher_name, 'room': room,
            'term': term, 'Department': dmcs
        }
        url = 'http://52.69.43.211/registerClass'  # サーバーの実際のURLに置き換えてください
        response = requests.post(url, data=data)
        print(response)  # test
    return render_template('lecture_creation.html')


@app.route("/timetable_sharing")
def index_timetable_sharing():
    return render_template('timetable_sharing.html')


if __name__ == '__main__':
    app.run(debug=True)
