########################################################################
 # Copyright (c) Intel Corporation 2023
 # SPDX-License-Identifier: BSD-3-Clause
########################################################################

import sys
import subprocess

import numpy as np
import bentoml
from bentoml.io import NumpyNdarray
from bentoml.io import Text
import subprocess
import json
import sound_classification
import requests

svc = bentoml.Service("sound_classification", runners=[])

@svc.api(input=Text(), output=Text())
def classify(text: str) -> str:

    data = json.loads(text)
    try:
     print("Input message received by the bentoml service: ",data)
   
     media_path = data["MediaPath"]
     model_path = data["ModelPath"]
     label_path = data["LabelPath"]
     grpc_address = data["GatewayIP"]
     grpc_port = data["Port"]

     inference_results = sound_classification.classify(input=media_path, model=model_path, sample_rate=16000, grpc_address=grpc_address, grpc_port=grpc_port,device="CPU", labelsFile=label_path )
     print("inference_results:" + str(inference_results))
     
     if inference_results != None:
      send_pipeline_inference_results("http://app-sample-service:59741/api/v1/data", inference_results)
      return "Success, inference_results: " + str(inference_results)      
     else:
      return "Failure"
    except Exception as e: 
     print(str(e))
     return "Failure"

def send_pipeline_inference_results(url, json_data):
    try:
        # Convert the Python dictionary to a JSON string
        json_payload = json.dumps(json_data)
        
        # Define headers with the content type
        headers = {'Content-Type': 'application/json'}
        
        # Make the POST request
        response = requests.post(url, data=json_payload, headers=headers)
        
        return response
    
    except Exception as e:
        print(f"An error occurred: {str(e)}")
        return None  # Return None in case of an error
