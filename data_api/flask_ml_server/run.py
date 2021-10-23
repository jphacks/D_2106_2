from flask import Flask, jsonify
from flask_cors import CORS

from api import api

app = Flask(__name__)
# cors = CORS(app, resources={r"/api/*": {"origin": ""}})
# cors = CORS(app, resources={r"/api/*": {"origin": "localhost"}})
cors = CORS(app)

app.register_blueprint(api)

if __name__ == '__main__':
    app.run(host="0.0.0.0", port="8080", debug=True)
