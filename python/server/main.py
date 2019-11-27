from flask import Flask

app = Flask(__name__)


@app.route('/')
def healthy():
    print('Http request Received')
    return "Python says: wE lOvE KUbErNetEs"

@app.route('/<any>')
def hello(any):
    print('Http request Received')
    return "Python says: wE lOvE KUbErNetEs"

if __name__ == "__main__":
    app.run(port=8080)



