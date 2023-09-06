########################################################################
 # Copyright (c) Intel Corporation 2023
 # SPDX-License-Identifier: BSD-3-Clause
########################################################################

#import os
import subprocess

import numpy as np
import bentoml
from bentoml.io import NumpyNdarray
from bentoml.io import Text
import subprocess
import json
import sound_classification
import logging

svc = bentoml.Service("sound_classification", runners=[])

@svc.api(input=Text(), output=Text())
def classify(text: str) -> str:

    try:
     print("Input message received by the bentoml service: ",text)
     
     # Call sound classification
     #os.system('python sound_classification.py -str '+ '"'+str(transcription)+'"   -io'+' "'+ str(args.inputo)+' " ')
     #os.system("python3 sound_classification.py -i ../../media/ak47s_gun_sound_mono.wav \
     # -m ../../models/aclnet/aclnet/aclnet_des_53.xml --sample_rate 16000 -d CPU")
     
     completedProc = subprocess.run(['python3','sound_classification.py','-i','../../media/ak47s_gun_sound_mono.wav', \
      '-m','../../models/aclnet/aclnet/aclnet_des_53.xml','--sample_rate','16000','-d','CPU'], capture_output=True)

     # Print the exit code.
     print(completedProc.returncode)
     output = completedProc.stdout
     print("out ", output)
     
     ''' result = image_classification.classify(grpc_address=data["GatewayIP"], grpc_port=9001, input_name="0", output_name="1463", images_list="image_classification/input_images.txt")
     print("json result returned from pipeline: ",result)

     resultList = result
     
     if resultList["Status"] == "PipelineComplete":
      helper.send_pipeline_inference_results(data["JobUpdateUrl"],resultList, data["GatewayIP"])
      helper.send_pipeline_status(data["PipelineStatusUrl"], resultList["Status"], data["GatewayIP"])
      return resultList["Status"]
     else:
       raise Exception("Pipeline completed, but failed")   
    except Exception as e:
      try:
       print("Error occurred while handling the service: "+ str(e))
       infereneceResults = {"Status": "PipelineFailed"}
       helper.send_pipeline_inference_results(data["JobUpdateUrl"], infereneceResults, data["GatewayIP"])
       helper.send_pipeline_status(data["PipelineStatusUrl"], "PipelineFailed", data["GatewayIP"])
      finally:
       print(str(e))
       return "PipelineFailed" '''
     
     if completedProc.returncode == 0:
      return "Success" + str(output)
     else:
      return "Failure"
    except Exception as e: 
     print(str(e))
     return "Failure"