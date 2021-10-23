import numpy as np
from sklearn.mixture import GaussianMixture
import random


def fix_seed(seed):
    random.seed(seed)
    np.random.seed(seed)


SEED = 42
fix_seed(SEED)


def request2numpy(gps_request):
    gps_data = gps_request["gps_data"]
    gps_data_numpy = np.zeros((len(gps_data), 3))
    for i, g in enumerate(gps_data):
        gps_data_numpy[i, 0] = g["gps_id"]
        gps_data_numpy[i, 1] = g["longitude"]
        gps_data_numpy[i:, 2] = g["latitude"]
    return gps_data_numpy


def produce_fake_gps_data(gps_data):
    n = len(gps_data)
    fake_n = 50
    new_gps = np.zeros((n * fake_n, 2))
    # new_gps = np.array((n , 2))
    for i in range(n):
        new_x = np.random.normal(
            loc=(gps_data[i, 1:3]), scale=(1, 1), size=(fake_n, 2))
    #     new_gps = np.concatenate([new_gps, new_x])
        new_gps[i*fake_n: (i+1)*fake_n, :] = new_x
    return new_gps


def request2response(request_gps_data):
    gps_data = request2numpy(request_gps_data)
    n = len(gps_data)
    new_gps = produce_fake_gps_data(gps_data)
    gps = gps_data[:, 1:3]

    # Gaussian Mixture Model
    best_bic = np.inf
    best_n_components = -1
    for i in range(n+1):
        if i == 0:
            continue
        gmm = GaussianMixture(
            n_components=i, covariance_type="full", random_state=SEED).fit(new_gps)
        bic = gmm.bic(new_gps)
        if bic < best_bic:
            best_bic = bic
            best_n_components = i

    gmm = GaussianMixture(n_components=best_n_components,
                          random_state=SEED).fit(new_gps)

    label = gmm.predict(gps)
    used_label = list(set(label))
    label2coo = {}
    for i in used_label:
        mean = gps[label == i].mean(axis=0)
        label2coo[i] = mean

    response_gps_data = {"cluster_data": []}
    for i in used_label:
        gps_data_ = {
            "cluster_id": -1,
            "gps_id_belongs_to": [],
            "mean_longitude": -1,
            "mean_latitude": -1
        }
        gps_data_["cluster_id"] = int(i)
        gps_data_["gps_id_belongs_to"] = [
            int(gps_id) for gps_id in gps_data[label == i][:, 0]]
        gps_data_["mean_longitude"] = float(label2coo[i][0])
        gps_data_["mean_latitude"] = float(label2coo[i][1])
        response_gps_data["cluster_data"].append(gps_data_)

    return response_gps_data
