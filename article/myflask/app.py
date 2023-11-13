from flask import Flask, render_template

app = Flask(__name__)


@app.route('/')
def index():
    # ここで講義のデータを取得する
    lectures = {
        "Monday": ["微積文法の応用", "", "プログラミング (C/C++)", "プログラミング (C/C++)", "人工知能"],
        "Tuesday": ["西洋近現代史", "科学英語", "", "", "最適化"],
        "Wednesday": ["", "時事英語", "", "" ,"データ構造とアルゴリズム 2"],
        "Thursday": ["", "", "", "", ""],
        "Friday": ["", "", "プロジェクト", "データベース", ""],
    }
    ## return render_template('lectures.html', lectures=lectures)
    return render_template('lectures.html', timetable=lectures)


if __name__ == '__main__':
    app.run(debug=True)
