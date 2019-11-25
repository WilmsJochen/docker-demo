from flask import Flask

app = Flask(__name__)


@app.route('/')
def hello():
    print('Http request Received')
    return "Python says: wE lOvE KUbErNetEs"

if __name__ == "__main__":
    app.run(port=8080)



