import numpy as np
import random


def fix_seed(seed):
    random.seed(seed)
    np.random.seed(seed)


SEED = 42
fix_seed(SEED)


def produce_random_gps():
    n = 10

    max_ = 0.3
    min_ = -max_

    longitude = np.random.normal(loc=0, scale=max_, size=n)
    latitude = np.random.normal(loc=0, scale=max_, size=n)
    gps = np.zeros((n, 2))
    gps[:, 0] = longitude
    gps[:, 1] = latitude

    return gps


def produce_gps_request():
    gps = produce_random_gps()
    gps_request = {"gps_data": []}
    for i, g in enumerate(gps):
        gps_data_ = {
            "gps_id": -1,
            "longitude": -1,
            "latitude": -1,
        }
        gps_data_["gps_id"] = i + np.random.randint(500)
        gps_data_["longitude"] = g[0] * 20 + 130
        gps_data_["latitude"] = g[1] * 15 + 30
        gps_request["gps_data"].append(gps_data_)
    return gps_request
