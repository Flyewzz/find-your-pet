import tensorflow as tf
import numpy as np
import tensorflow.keras.backend as K
from glob import glob
from extract_bottleneck_features import *
from tqdm import tqdm

from io import BytesIO
import base64
from PIL import Image

import time

DATA_PATH = '/home/gamma/projects/find-your-pet/breeds'

# load list of dog names
dog_names = [item[20:-1] for item in sorted(glob(DATA_PATH + "/train/*/"))]


def path_to_tensor(img_path, is_base64):
    # loads RGB image as PIL.Image.Image type
    img = None
    if is_base64:
        img = Image.open(BytesIO(base64.b64decode(img_path))).resize((224, 224))
    else:
        img = tf.keras.preprocessing.image.load_img(img_path,
                                                target_size=(224, 224))
    # convert PIL.Image.Image type to 3D tensor with shape (224, 224, 3)
    x = tf.keras.preprocessing.image.img_to_array(img)
    # convert 3D tensor to 4D tensor with shape (1, 224, 224, 3) and return 4D tensor
    return np.expand_dims(x, axis=0)


def paths_to_tensor(img_paths):
    list_of_tensors = [path_to_tensor(img_path) for img_path in tqdm(img_paths)]
    return np.vstack(list_of_tensors)


# model creation
def create_model():
    # it was gotten from train_Xception.shape[1:]
    input_shape = (7, 7, 2048)
    Xception_model = tf.keras.models.Sequential()
    Xception_model.add(
        tf.keras.layers.GlobalAveragePooling2D(input_shape=(input_shape)))
    Xception_model.add(tf.keras.layers.Dense(133, activation='softmax'))
    # some info about model
    print(Xception_model.summary())

    Xception_model.compile(loss='categorical_crossentropy', optimizer='rmsprop',
                           metrics=['accuracy'])
    Xception_model.load_weights('Saved_Models/weights.best.Xception.hdf5')
    return Xception_model


def Xception_predict_breed(img_path, is_base64):
    model = create_model()
    # extract the bottle neck features
    bottleneck_feature = extract_Xception(path_to_tensor(img_path, is_base64))
    # get a vector of predicted values
    predicted_vector = model.predict(bottleneck_feature)
    top_predictions = np.argpartition(predicted_vector.flatten(), -4)[-3:]
    K.clear_session()
    return [dog_names[i] for i in top_predictions]

if __name__ == '__main__':
    for i in range(4):
        print(Xception_predict_breed('./pul.jpg'))
        time.sleep(5)