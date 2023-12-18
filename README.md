# Exercise: Build your own docker container and push it to the cloud.

Good you found this git repo!

**_IMPORTANT:_** Make sure to always replace YOUR_NAME with a lowercase name and use this name consistently. 
If there are people with the same name in this group, please use your nicknames to avoid confusion.

**_MOST IMPORTANT:_** please ask questions whenever you want.
 
## Choose your language

In this project you can find some basic examples for different programming languages. Have a quick look on all the different dockerfiles and source code in these folders.
Needless to say that this list of languages can be expanded. 

## Environment
You can execute this exercise on all machines, but if you do not have all binaries installed, you could use the playgrounds of following website.
https://killercoda.com/playgrounds/scenario/kubernetes

## Instructions
### Checkout Git repo
To use our code on the virtual machine, we need to clone this git repo.

1) Clone the git repo
 ```
 git clone https://github.com/WilmsJochen/docker-demo.git
 ```
2) Change your directory to the repo
 ```
 cd docker-demo
 ```
3) Have a look in all the folders and files.
 Verify all the directories and files corresponds to this git repo.
  ```
 ls
  ```

### Build your first docker image
In this step we are going to build a docker image and push it to a docker repo. We will start with the server code as we will deploy this to kubernetes in the next exercises. If you have time left at the end of the session you can come back to this exercise and build the job code as well.

1) Choose the programming language you are the most familiar with.
 ```
 cd nodeJs/server
 ```

2) Modify the message that the server will return to a funny message.
 ```
 nano main.py
 nano main.go 
 nano index.js
 ```
**_NOTE:_**  To exit nano, press crtl+x -> press y (yes) -> press enter to save it in the same file.

2) Build your docker image and name the image to your own name.
 ```
 docker build . --tag YOUR_NAME
 ```

**_NOTE:_** When you work on one of the newer macbooks with silicon processors, you have to set the platform while building the docker image. 
 ```
 docker build --platform linux/amd64 . --tag YOUR_NAME
 ```


3) Check your docker images.
 ```
 docker images
 ```

4) Tag your image to give it a proper version number
 ```
 docker tag YOUR_NAME YOUR_NAME:V1.0.0
 ```

5) Check your tagged image.
 ```
 docker images
 ```
6) Tag your image again with the docker registry as prefix. 
This is needed to push your locally created docker image to a docker registry. Take a look at dockerhub to create your own free registry.
 ```
 docker tag YOUR_NAME:V1.0 REGISTRY_NAME/YOUR_NAME:V1.0.0
 ```

7) Push this tagged image to the docker registry.
 ```
 docker push REGISTRY_NAME/YOUR_NAME:V1.0.0
 ```

Congratulations! You've managed to create your first docker image and make it accessible in a docker registry. 

