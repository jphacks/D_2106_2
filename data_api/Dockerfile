FROM ubuntu:18.04

WORKDIR /app

RUN apt update
RUN apt install -y python3 curl python3-distutils
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
RUN python3 get-pip.py
RUN rm get-pip.py

RUN pip3 install numpy
RUN pip3 install scikit-learn
RUN pip3 install Flask
RUN pip3 install flask-cors

CMD [ "python3", "run.py" ]