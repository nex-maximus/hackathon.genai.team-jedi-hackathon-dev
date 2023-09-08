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
   
     mediaPath = data["MediaPath"]
     modelPath = data["ModelPath"]

     pipelineStatus = sound_classification.classify(input=mediaPath, model=modelPath, device="CPU", sample_rate=16000)
     print(pipelineStatus)
     
     if pipelineStatus == True:
      return "Success"
     else:
      return "Failure"
    except Exception as e: 
     print(str(e))
     return "Failure"