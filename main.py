from concurrent import futures
import logging
import classifier
from breed_formatter import *
from multiprocessing.pool import ThreadPool

pool = ThreadPool(processes=1)

import grpc
import classifier_pb2
import classifier_pb2_grpc

class BreedClassifier(classifier_pb2_grpc.BreedClassifierServiceServicer):
    
    def RecognizeBreed(self, request, context):
        async_breeds = pool.apply_async(classifier.Xception_predict_breed, (request.path, True))
        breeds = async_breeds.get()
        # Give a formatted breed (rus)
        breeds = [breed_format(breed) for breed in breeds]
        print('breeds: %s' % breeds)
        grpc_breeds = classifier_pb2.Breed()
        grpc_breeds.name.extend(breeds)
        return grpc_breeds
    
def serve():
    breed_classifier = BreedClassifier()
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
    classifier_pb2_grpc.add_BreedClassifierServiceServicer_to_server(breed_classifier, server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    print('*** Breed classificator is ready ***')
    serve()