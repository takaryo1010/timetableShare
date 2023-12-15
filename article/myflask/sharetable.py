from flask import Flask, request, redirect, url_for, render_template, flash
from flask_sqlalchemy import SQLAlchemy
from flask_login import LoginManager, UserMixin, login_user, logout_user, login_required, current_user
from werkzeug.security import generate_password_hash, check_password_hash
import requests
import sys
@app.route('/')
@login_required
def share_index():
    
    # サーバーのエンドポイントURLを設定
    
    data = {'my_name':current_user.username}
    friends = []
    friendsTimetables = {}
    response = requests.post('http://52.69.43.211/showFriends', data=data)
    json_response = response.json()
    friends.append(current_user.username)
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
    
    return render_template('timetable_sharing.html', friendstimetable=friendsTimetables)

