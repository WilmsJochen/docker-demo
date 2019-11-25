import requests
import os

URL = os.getenv( "URL","http://localhost:8080")
print("Request send to URL: ", URL)

r = requests.get(url = URL)
print(r.text)




