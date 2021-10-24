from flask import Blueprint, request, abort, jsonify
from gps_clustering.get_sample import produce_gps_request
from gps_clustering.clustering_gps import request2response
import json

api = Blueprint('api', __name__, url_prefix='/api')


@api.route('/get_sample', methods=["GET"])
def get_sample_gps_data():
    response = produce_gps_request()
    return jsonify(response), 200


@api.route('/clustering', methods=["POST"])
def clustering_gps_data():
    request_data = request.data.decode("utf-8")
    request_data = json.loads(request_data)
    response = request2response(request_data)
    return jsonify(response), 201
