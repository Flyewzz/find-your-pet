from keras.utils import np_utils
import numpy as np
from glob import glob
from extract_bottleneck_features import *
import keras
from keras.preprocessing import image
from tqdm import tqdm


DATA_PATH = '/home/kotyarich/Dev/tp/breed_classificator/data/'


# load list of dog names
dog_names = [item[20:-1] for item in sorted(glob(DATA_PATH + "train/*/"))]


def path_to_tensor(img_path):
    # loads RGB image as PIL.Image.Image type
    img = image.load_img(img_path, target_size=(224, 224))
    # convert PIL.Image.Image type to 3D tensor with shape (224, 224, 3)
    x = image.img_to_array(img)
    # convert 3D tensor to 4D tensor with shape (1, 224, 224, 3) and return 4D tensor
    return np.expand_dims(x, axis=0)


def paths_to_tensor(img_paths):
    list_of_tensors = [path_to_tensor(img_path) for img_path in tqdm(img_paths)]
    return np.vstack(list_of_tensors)


# model creation
def create_model():
    # it was gotten from train_Xception.shape[1:]
    input_shape = (7, 7, 2048)
    Xception_model = keras.models.Sequential()
    Xception_model.add(keras.layers.GlobalAveragePooling2D(input_shape=(input_shape)))
    Xception_model.add(keras.layers.Dense(133, activation='softmax'))
    # some info about model
    print(Xception_model.summary())

    Xception_model.compile(loss='categorical_crossentropy', optimizer='rmsprop', metrics=['accuracy'])
    Xception_model.load_weights('Saved_Models/weights.best.Xception.hdf5')

    return Xception_model


# creating model
Xception_model = create_model()


def Xception_predict_breed (img_path):
    """takes a path to an image and returns the dog breed"""
    # extract the bottle neck features
    bottleneck_feature = extract_Xception(path_to_tensor(img_path))
    ## get a vector of predicted values
    predicted_vector = Xception_model.predict(bottleneck_feature)
    return dog_names[np.argmax(predicted_vector)]


print(Xception_predict_breed('/home/kotyarich/Dev/tp/breed_classificator/data/test/057.Dalmatian/Dalmatian_04015.jpg'))

