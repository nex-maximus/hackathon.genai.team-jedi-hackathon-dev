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

svc = bentoml.Service("sound_classification", runners=[])

@svc.api(input=Text(), output=Text())
def classify(text: str) -> str:

    data = json.loads(text)
    try:
     print("Input message received by the bentoml service: ",data)
   
     media_path = data["MediaPath"]
     model_path = data["ModelPath"]
     label_path = data["LabelPath"]

     inference_results = sound_classification.classify(input=media_path, model=model_path, sample_rate=16000, device="CPU", labelsFile=label_path )
     print("inference_results:" + str(inference_results))
     
     if inference_results != None:
      return "Success, inference_results: " + str(inference_results)
     else:
      return "Failure"
    except Exception as e: 
     print(str(e))
     return "Failure"