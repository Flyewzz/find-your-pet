from concurrent import futures
import logging

import classifier
from breed_formatter import *
import json
from flask import Flask, request

app = Flask(__name__)

@app.route('/breed', methods=['POST'])
def index():
    img = request.json['picture']
    breed = classifier.Xception_predict_breed(img)
    print('breed: %s' % breed)
    # Give a formatted breed (eng)
    breed = breed_format(breed)
    # Only one breed right now (*TODO* FIX IT)
    return json.dumps([breed], ensure_ascii=False).encode('utf8')

if __name__ == '__main__':
    logging.basicConfig()
    print('*** Breed classificator is ready ***')
    app.run(host='0.0.0.0', port=5000, debug=True)