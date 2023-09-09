# team-jedi-hackathon-dev
NEX Jedi Team Hackathon Microservice project - Development Repo

1. To build data export run `cd app-data-export` and run `make docker`

2. To run the initial services, including the TIG stack, run `make run`

3. **To run Bentoml pipelines-**
- You can run these commands directly on terminal, but to avoid any package environment issues, run them in a conda environment with python version 3.8 -
- `conda create -n hackathon_env python=3.8`
- `conda activate hackathon_env`
- `cd pipelines/sound_classification_demo`
- `pip install -r requirements.txt`
- Resolve any package issues - use `conda list` to check if all the packages in requirements.txt ar installed
- Build bento service locally - `make build`
- Serve bento locally - `make serve`
- Open swagger api UI at `http://0.0.0.0:3000/`
- Test the post /classify api by providing input text as {"MediaPath": "[PATH]/team-jedi-hackathon-dev/media/ak47s_gun_sound_mono.wav", "ModelPath": "[PATH]/team-jedi-hackathon-dev/models/aclnet/aclnet/aclnet_des_53.xml"}
- If pipelines runs successfully, status "Success" is returned else it is "Failure"
- Check the terminal to view inference output -
  `[0.00-1.00] - 100.00% Class 50
   [1.00-2.00] - 15.64% Class 30
   Metrics report:
	Latency: 21.1 ms`
- Build bento docker - `make docker-build BENTO_TAG=<bento_image_name>:<bento_image_tag>` 
- `<bento_image_name>:<bento_image_tag>` are output of make build, more details can be found in Makefile
- Run bento docker - `make docker-run BENTO_TAG=<bento_image_name>:<bento_image_tag> PROJECT_REPO_PATH=<project_repo_path>`
- `project_repo_path` is complete project repo path eg: /home/nesubuntu20nvda/Neethu/hackathon-team-jedi/team-jedi-hackathon-dev
- Check portainer or terminal if container with name `<bento_image_name>:<bento_image_tag>` is created
- Access the open swagger api UI at `http://0.0.0.0:3000/` and test the post /classify api for status "Success"
- Check container logs for output


# Name of Product
## Overview 
Provide a 2-3 line description of what this product allows the developer to do. 

>NOTE: Keep this section easy to understand. Make it easy for users to
>understand what the input and output is.

-  **Programming Language:** 
-  **Technologies used :** 
>NOTE: List the technologies, frameworks, libraries, and tools you utilized in your microservices project


## Target System Requirements 
-  Disk Space needed 
-  Other Requirements 


## Microservice descriptions:

>NOTE: 
> 1. Discuss the interactions between any other microservices and how they communicate (e.g., RESTful APIs, message queues). 
> 2. Describe how data is stored, managed, and shared

## How It Works 
>NOTE: Provide description, including architecture diagram, of how the product
>works. All diagrams and screenshots must have alt-text and captions.

![Add alt-text description of image here.](images/my-arch-diagram.png)

Figure 1: Architecture Diagram  
 
 

## Get Started  
Provide step-by-step instructions for getting started.

>NOTE: Keep these easy to run. Avoid coding and manual configurations after
>installation. The input, or the configuration of the input, must be included in
>the package.

1. Text.

2. Text with code.

   ```bash
   code snippet
   ```

3. Text with filepath. Go to the ``name`` directory.

   ```bash
   cd name /
   ```

4. Text with code and screenshot.

   ```bash
   code snippet
   ```

   You will see output similar to the following:

   ![A browser window showing the product dashboard.](images/my-dashboard.png)

   Figure 2: Product Dashboard


## Run the Application 
Provide detailed steps for running the application. 

1. Text.

2. Text with code.

   ```bash
   code snippet
   ```

## API Documentation
> If your microservices expose APIs, document each API endpoint, including its purpose, input parameters, expected output, and any authentication/authorization requirements.
> Provide sample API requests and responses for clarity.

## Testing# Name of Product
## Overview 
Provide a 2-3 line description of what this product allows the developer to do. 

>NOTE: Keep this section easy to understand. Make it easy for users to
>understand what the input and output is.

-  **Programming Language:** 
-  **Technologies used :** 
>NOTE: List the technologies, frameworks, libraries, and tools you utilized in your microservices project


## Target System Requirements 
-  Disk Space needed 
-  Other Requirements 


## Microservice descriptions:

>NOTE: 
> 1. Discuss the interactions between any other microservices and how they communicate (e.g., RESTful APIs, message queues). 
> 2. Describe how data is stored, managed, and shared

## How It Works 
>NOTE: Provide description, including architecture diagram, of how the product
>works. All diagrams and screenshots must have alt-text and captions.

![Add alt-text description of image here.](images/my-arch-diagram.png)

Figure 1: Architecture Diagram  
 
 

## Get Started  
Provide step-by-step instructions for getting started.

>NOTE: Keep these easy to run. Avoid coding and manual configurations after
>installation. The input, or the configuration of the input, must be included in
>the package.

1. Text.

2. Text with code.

   ```bash
   code snippet
   ```

3. Text with filepath. Go to the ``name`` directory.

   ```bash
   cd name /
   ```

4. Text with code and screenshot.

   ```bash
   code snippet
   ```

   You will see output similar to the following:

   ![A browser window showing the product dashboard.](images/my-dashboard.png)

   Figure 2: Product Dashboard


## Run the Application 
Provide detailed steps for running the application. 

1. Text.

2. Text with code.

   ```bash
   code snippet
   ```

## API Documentation
> If your microservices expose APIs, document each API endpoint, including its purpose, input parameters, expected output, and any authentication/authorization requirements.
> Provide sample API requests and responses for clarity.

## Testing
> Discuss the testing approach you followed for your microservices.
> Document unit tests, integration tests, and any other types of tests performed.
> Include instructions on how to run the tests.

## Summary and Next Steps 
>Note: Provide 2-3 line description of what the user has successfully done and
>where they should go to as the next step. 




## Troubleshooting 
>Include a section addressing common issues, error handling, and troubleshooting tips.

 
 on how to run the tests.

## Summary and Next Steps 
>Note: Provide 2-3 line description of what the user has successfully done and
>where they should go to as the next step. 




## Troubleshooting 
>Include a section addressing common issues, error handling, and troubleshooting tips.

 
