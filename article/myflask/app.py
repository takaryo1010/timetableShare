from flask import Flask, request, redirect, url_for, render_template, flash
from flask_sqlalchemy import SQLAlchemy
from flask_login import LoginManager, UserMixin, login_user, logout_user, login_required, current_user
from werkzeug.security import generate_password_hash, check_password_hash
import requests

flag = input("test? \nYes > 1\nNo > 0\n>>>")

app = Flask(__name__)
app.config['SECRET_KEY'] = 'ca448a98'
if flag:
    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:////root/instance/user.db'
else:
    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///users.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)

login_manager = LoginManager()
login_manager.init_app(app)
login_manager.login_view = 'login'


class User(UserMixin, db.Model):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(100), unique=True, nullable=False)
    password_hash = db.Column(db.String(200))

    def set_password(self, password):
        self.password_hash = generate_password_hash(password)

    def check_password(self, password):
        return check_password_hash(self.password_hash, password)


@login_manager.user_loader
def load_user(user_id):
    return User.query.get(int(user_id))


@app.route('/')
@login_required
def index():
    url = 'http://52.69.43.211/showMyClassInfo'
    # test

    # サーバーのエンドポイントURLを設定
    url = 'http://52.69.43.211/showMyClassInfo'  # サーバーの実際のURLに置き換えてください

# 送信するデータを準備
    data = {'name': current_user.username}  # 送信するデータをここで設定
    print(data)

# HTTP POSTリクエストを送信
    response = requests.post(url, data=data)


# レスポンスをJSONとしてパース
    json_response = response.json()
    monday = ["", "", "", "", ""]
    tuesday = ["", "", "", "", ""]
    wednesday = ["", "", "", "", ""]
    thursday = ["", "", "", "", ""]
    friday = ["", "", "", "", ""]
    saturday = ["", "", "", "", ""]
    ondemand = ["", "", "", "", ""]
    for class_ in json_response:
        if class_['day'] == 'Monday':
            monday[int(class_['period'])] = class_['name']
        elif class_['day'] == 'Tuesday':
            tuesday[int(class_['period'])] = class_['name']
        elif class_['day'] == 'Wednesday':
            wednesday[int(class_['period'])] = class_['name']
        elif class_['day'] == 'Thursday':
            thursday[int(class_['period'])] = class_['name']
        elif class_['day'] == 'Friday':
            friday[int(class_['period'])] = class_['name']
        elif class_['day'] == 'Saturday':
            saturday[int(class_['period'])] = class_['name']
        elif class_['day'] == 'Ondemand':
            ondemand[int(class_['period'])] = class_['name']

    lectures = {
        "Monday": monday,
        "Tuesday": tuesday,
        "Wednesday": wednesday,
        "Thursday": thursday,
        "Friday": friday,
        "Saturday": saturday,
        "Ondemand": ondemand
    }
    print (lectures)
    
    # end test
    # ここで講義のデータを取得する
    #lectures = {
    #    "Monday": ["微積文法の応用", "", "プログラミング (C/C++)", "プログラミング (C/C++)", "人工知能"],
    #    "Tuesday": ["西洋近現代史", "科学英語", "", "", "最適化"],
    #    "Wednesday": ["", "時事英語", "", "", "データ構造とアルゴリズム 2"],
    #    "Thursday": ["", "", "", "", ""],
    #    "Friday": ["", "", "プロジェクト", "データベース", ""],
    #    "Friday": ["", "", "プロジェクト", "データベース", ""],
    #    "Saturday": ["", "", "", "", ""],
    #    "Ondemand": ["CF 特論", "IS 特論", "", "", ""],
    #}

    return render_template('lectures.html', timetable=lectures)


@app.route('/test')
def test():

    
    # end test
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

    return render_template('testlecture.html', timetable=lectures)

@app.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        user = User.query.filter_by(username=username).first()
        if user and user.check_password(password):
            login_user(user)
            flash(f'こんにちは、{username}さん。')
            return redirect(url_for('index'))
        else:
            flash('Invalid username or password.')
    return render_template('login.html')


@app.route('/signup', methods=['GET', 'POST'])
def signup():
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        existing_user = User.query.filter_by(username=username).first()
        if existing_user is None:
            url = 'http://52.69.43.211/registerPerson'
            data = {'name': username}  # 送信するデータをここで設定

            response = requests.post(url, data=data)
            json_response = response.json()
            print(json_response)

            if response.ok:
                user = User()
                user.username = username
                user.set_password(password)
                db.session.add(user)
                db.session.commit()
                flash('サインアップが正常に完了しました。')
                return redirect(url_for('login'))
            else:
                flash('外部データベースへの登録に失敗しました。')
        else:
            flash('このユーザー名は既に存在します。')
    return render_template('signup.html')


@app.route('/logout')
@login_required
def logout():
    logout_user()
    flash('ログアウトしました。')
    return redirect(url_for('login'))


@app.route("/timetable_registration")
def index_timetable_registration():
    # サーバーのエンドポイントURLを設定
    url = 'http://52.69.43.211/showClassInfoTimeSpecification'  # サーバーの実際のURLに置き換えてください
    
    data = {"day":"Monday"}
    
    # HTTP POSTリクエストを送信
    response = requests.post(url, data=data)

    # レスポンスをJSONとしてパース
    json_response = response.json()

    # レスポンスを出力
    print(json_response)

    return render_template('timetable_registration.html', json=json_response)

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
        print(class_time)
        class_unit = request.form.get('class_unit')
        must_flag = request.form.get('must_flag')
        teacher_name = request.form.get('teacher_name')
        room = request.form.get('room')
        term = request.form.get('term')
        dmcs = request.form.get('dmcs')
        print(dmcs)
        data = {
            'name': class_name,
            'day': class_day, 'period': class_time,
            'unit': class_unit, 'must': must_flag,
            'teacher': teacher_name, 'room': room,
            'term': term, 'department': dmcs
        }
        url = 'http://52.69.43.211/registerClass'  # サーバーの実際のURLに置き換えてください
        response = requests.post(url, data=data)
        print(response)  # test
    return render_template('lecture_creation.html')


@app.route("/timetable_sharing")
def index_timetable_sharing():
    return render_template('timetable_sharing.html')


if __name__ == '__main__':
    # with app.app_context():
    app.run(host='0.0.0.0', port='80', debug=True)
