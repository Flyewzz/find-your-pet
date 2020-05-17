# from concurrent import futures
import logging

import classifier
from breed_formatter import *
import json
from flask import Flask, request
from multiprocessing.pool import ThreadPool
pool = ThreadPool(processes=1)

app = Flask(__name__)

@app.route('/breed', methods=['POST'])
def index():
    img = request.json['picture']
    async_breeds = pool.apply_async(classifier.Xception_predict_breed, (img, True))
    breeds = async_breeds.get()
    print('breeds: %s' % breeds)
    # Give a formatted breed (rus)
    breeds = [breed_format(breed) for breed in breeds]
    return json.dumps(breeds, ensure_ascii=False).encode('utf8')



if __name__ == '__main__':
    logging.basicConfig()
    print('*** Breed classificator is ready ***')
    app.run(host='0.0.0.0', port=5000, debug=False)
