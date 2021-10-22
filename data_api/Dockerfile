FROM python:3.8.8

RUN apt-get update

WORKDIR /app
COPY requirements.txt /app

RUN pip install --upgrade pip
RUN pip install -r requirements.txt

EXPOSE 5000

ENV FLASK_APP '/app/app.py'
ENV FLASK_DEBUG 1

CMD ["python", "app.py"]