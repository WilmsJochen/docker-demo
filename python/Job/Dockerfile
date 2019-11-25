FROM python:3
EXPOSE 8080
WORKDIR /usr/src/app
COPY . .
RUN pip install requests

CMD [ "python", "./main.py" ]
