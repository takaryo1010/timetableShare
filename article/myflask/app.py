from flask import Flask, request, redirect, url_for, render_template, flash
from flask_sqlalchemy import SQLAlchemy
from flask_login import LoginManager, UserMixin, login_user, logout_user, login_required, current_user
from werkzeug.security import generate_password_hash, check_password_hash
import requests
import sys

flag = sys.argv[1]

app = Flask(__name__)
app.config['SECRET_KEY'] = 'ca448a98'
if int(flag) == 0 or int(flag) == 2:
    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:////root/instance/user.db'
    db = SQLAlchemy(app)
elif int(flag) == 1:
    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///users.db'
    db = SQLAlchemy(app)
else:
    pass
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

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
    print(json_response)
    monday = [{}, {}, {}, {}, {}]
    tuesday = [{}, {}, {}, {}, {}]
    wednesday = [{}, {}, {}, {}, {}]
    thursday = [{}, {}, {}, {}, {}]
    friday = [{}, {}, {}, {}, {}]
    saturday = [{}, {}, {}, {}, {}]
    ondemand = [{}, {}, {}, {}, {}]
    for class_ in json_response:
        if class_['day'] == 'Monday':
            monday[int(class_['period']) - 1]  = class_
        elif class_['day'] == 'Tuesday':
            tuesday[int(class_['period']) - 1] = class_
        elif class_['day'] == 'Wednesday':
            wednesday[int(class_['period']) - 1] = class_
        elif class_['day'] == 'Thursday':
            thursday[int(class_['period']) - 1] = class_
        elif class_['day'] == 'Friday':
            friday[int(class_['period']) - 1] = class_
        elif class_['day'] == 'Saturday':
            saturday[int(class_['period']) - 1] = class_
        elif class_['day'] == 'Ondemand':
            ondemand[int(class_['period']) - 1] = class_

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
    username = current_user.username
    
    return render_template('lectures.html', timetable=lectures, username=username)

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
    username = current_user.username

    return render_template('testlecture.html', timetable=lectures, username=username)

@app.route('/login', methods=['GET', 'POST'])
def login():
    users = User.query.all()
    for user in users:
        print(user.username)
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

@app.route("/timetable_registration", methods=['GET', 'POST'])
@login_required
def index_timetable_registration():
    if request.method == 'POST':
        if request.form.get('class_id') == "" or request.form.get('class_id') == None:
            data = {}
            if request.form.get('class_name') != "" and request.form.get('class_name') != None:
                print("aaaaaaaaa")
                data['name'] = request.form.get('class_name')
                print(data['name'])
            if request.form.get('class_day') != "" and request.form.get('class_day') != None:
                data['day'] = request.form.get('class_day')
            if request.form.get('class_time') != "" and request.form.get('class_time') != None:
                data['period'] = request.form.get('class_time')
            if request.form.get('class_unit') != "" and request.form.get('class_unit') != None:
                data['unit'] = request.form.get('class_unit')
            if request.form.get('must_flag') != "" and request.form.get('must_flag') != None:
                data['must'] = request.form.get('must_flag')
            if request.form.get('teacher_name') != "" and request.form.get('teacher_name') != None:
                data['teacher'] = request.form.get('teacher_name')
            if request.form.get('room') != "" and request.form.get('room') != None:
                data['room'] = request.form.get('room')
            if request.form.get('term') != "" and request.form.get('term') != None:
                data['term'] = request.form.get('term')
            if request.form.get('department') != "" and request.form.get('department') != None:
                data['department'] = request.form.get('department')
            url = 'http://52.69.43.211/showClassInfoTimeSpecification'
            print(data)
            response = requests.post(url, data=data)
            json_response = response.json()
            print(json_response)
            return render_template('timetable_registration.html', data=data, json=json_response)
        else:
            url = 'http://52.69.43.211/registerCourse'
            name = current_user.username
            classid = request.form.get('class_id')
            data = {'name': name, 'classid': classid}
            response = requests.post(url, data)
            json_response = response.json()
            print(json_response)
    username = current_user.username
    return render_template('timetable_registration.html', username=username)

@app.route("/timetable_registration_designation", methods=['GET', 'POST'])
@login_required
def index_timetable_registration_designation():
    print(request.args.get('day', default=None), request.args.get('period', default=None))
    url = 'http://52.69.43.211/showClassInfoTimeSpecification'
    data = {'day': request.args.get('day', default=None), 'period': request.args.get('period', default=None)}
    response = requests.post(url, data)
    json_response = response.json()
    print(json_response)
    username = current_user.username
    return render_template('lecture_list.html', data=data, json=json_response, username=username)

@app.route("/lecture_list", methods=['GET', 'POST'])
@login_required
def index_lecture_list():
    if request.method == 'POST':
        if request.form.get('class_id') == "" or request.form.get('class_id') == None:
            data = {}
            if request.form.get('class_name') != "" and request.form.get('class_name') != None:
                print("aaaaaaaaa")
                data['name'] = request.form.get('class_name')
                print(data['name'])
            if request.form.get('class_day') != "" and request.form.get('class_day') != None:
                data['day'] = request.form.get('class_day')
            if request.form.get('class_time') != "" and request.form.get('class_time') != None:
                data['period'] = request.form.get('class_time')
            if request.form.get('class_unit') != "" and request.form.get('class_unit') != None:
                data['unit'] = request.form.get('class_unit')
            if request.form.get('must_flag') != "" and request.form.get('must_flag') != None:
                data['must'] = request.form.get('must_flag')
            if request.form.get('teacher_name') != "" and request.form.get('teacher_name') != None:
                data['teacher'] = request.form.get('teacher_name')
            if request.form.get('room') != "" and request.form.get('room') != None:
                data['room'] = request.form.get('room')
            if request.form.get('term') != "" and request.form.get('term') != None:
                data['term'] = request.form.get('term')
            if request.form.get('department') != "" and request.form.get('department') != None:
                data['department'] = request.form.get('department')
            url = 'http://52.69.43.211/showClassInfoTimeSpecification'
            print(data)
            response = requests.post(url, data=data)
            json_response = response.json()
            print(json_response)
        else:
            url = 'http://52.69.43.211/registerCourse'
            name = current_user.username
            classid = request.form.get('class_id')
            data = {'name': name, 'classid': classid}
            response = requests.post(url, data)
            json_response = response.json()
            print(json_response)
            url = 'http://52.69.43.211/showClassInfoAll'
            response = requests.get(url)
            json_response = response.json()
        
        username = current_user.username
        return render_template('lecture_list.html', data=data, json=json_response, username=username)
        
    else:
        url = 'http://52.69.43.211/showClassInfoAll'
        response = requests.get(url)
        json_response = response.json()
    print(json_response)

    username = current_user.username
    return render_template('lecture_list.html', json=json_response, username=username)

@app.route("/lecture_creation", methods=['GET', 'POST'])
@login_required
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
        
        
    # サーバーのエンドポイントURLを設定
    url = 'http://52.69.43.211/showClassInfoAll'  # サーバーの実際のURLに置き換えてください

    # HTTP POSTリクエストを送信
    response = requests.get(url)

    # レスポンスをJSONとしてパース
    json_response = response.json()

    # レスポンスを出力
    print(json_response)
    username = current_user.username
    return render_template('lecture_list.html',json=json_response, username=username)

@app.route('/timetable_sharing')
@login_required
def share_index():
    
    # サーバーのエンドポイントURLを設定
    
    data = {'my_name':current_user.username}
    friends = []
    friendsTimetables = {}
    response = requests.post('http://52.69.43.211/showFriends', data=data)
    json_response = response.json()
    friends.append(current_user.username)
    if(json_response!=None):
        for y in json_response:
            friends.append( y['your_name'])
        
    for y in friends:
    # 送信するデータを準備
        data = {'name': y}  # 送信するデータをここで設定
        print(data)

        
    # HTTP POSTリクエストを送信
       
        response = requests.post('http://52.69.43.211/showMyClassInfo', data=data)

    # レスポンスをJSONとしてパース
        json_response = response.json()
        print(json_response)
        monday = ["", "", "", "", ""]
        tuesday = ["", "", "", "", ""]
        wednesday = ["", "", "", "", ""]
        thursday = ["", "", "", "", ""]
        friday = ["", "", "", "", ""]
        saturday = ["", "", "", "", ""]
        ondemand = ["", "", "", "", ""]
        for class_ in json_response:
            if class_['day'] == 'Monday':
                monday[int(class_['period']) - 1]  = class_['name']
            elif class_['day'] == 'Tuesday':
                tuesday[int(class_['period']) - 1] = class_['name']
            elif class_['day'] == 'Wednesday':
                wednesday[int(class_['period']) - 1] = class_['name']
            elif class_['day'] == 'Thursday':
                thursday[int(class_['period']) - 1] = class_['name']
            elif class_['day'] == 'Friday':
                friday[int(class_['period']) - 1] = class_['name']
            elif class_['day'] == 'Saturday':
                saturday[int(class_['period']) - 1] = class_['name']
            elif class_['day'] == 'Ondemand':
                ondemand[int(class_['period']) - 1] = class_['name']

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
        friendsTimetables[y]=lectures
    
    return render_template('timetable_sharing.html', friendstimetable=friendsTimetables, myusername=current_user.username)

@app.route('/add_friends', methods=['GET', 'POST'])
@login_required
def index_friend_add():
    if request.method == 'POST':
        add_friend = request.form.get('add_friend')
        remove_friend = request.form.get('remove_friend')
        if add_friend != "" and add_friend != None:
            url = 'http://52.69.43.211/registerFriends'
            data = {'my_name': current_user.username, 'your_name': add_friend}
            response = requests.post(url, data)
            json_response_register = response.json()
            print(json_response_register)
        elif remove_friend != "" and remove_friend != None:
            url = 'http://52.69.43.211/removeFriends'
            data = {'your_name': remove_friend,'my_name': current_user.username}
            response = requests.post(url, data)
            json_response_remove = response.json()
            print(json_response_remove)
    url = 'http://52.69.43.211/showFriends'
    data = {'my_name': current_user.username}
    response = requests.post(url, data)
    json_response = response.json()
    if(json_response == None):
        json_response = []
    print(json_response)
    username = current_user.username
    return render_template('add_friends.html', json=json_response, username=username)

if __name__ == '__main__':
    with app.app_context():
        if int(flag) == 2:
            db.create_all()
    app.run(host='0.0.0.0', port='80', debug=True)
