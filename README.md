# Ex1: Build your own docker container and push it to the cloud.

Good you found this git repo!
 
## Choose your language

In this project you can find some basic examples for different programming languages. Have a quiick look on all the different dockerfiles and source code in these folders.
Needless to say that this list of languages can be expanded. 

## google cloud
To build this docker images, I set up a google account to log in to the google cloud console.
Please follow theses steps:

1) Go to the google cloud console: https://console.cloud.google.com/

2) Login as kubernetesTalk user with the credentials from the ppt.
    
   **_NOTE:_**  If you use the chrome browser you should probably add the kubernetestalk user in your browser.

3) Use the navigation menu to navigate to the Kubenetes engine
4) Connect to the cluster in a new terminal by pushing `connect` button
5) Create a folder with your own name:

 ```
 mkdir YOUR_NAME
 ```
6) Change directory to that folder.
 ```
 cd YOUR_NAME
 ```

From this terminal we can execute all commands for the next exercises.

##Instructions
### Checkout Git repo
To use our code on the virtual machine, we need to clone our git repo.

1) Clone the git repo
 ```
 git clone https://github.com/WilmsJochen/ex1.git
 ```
2) Change your directory to that repo
 ```
 cd ex1
 ```
3) Have a look on all the folders and files in that directory.
 Verify all the directories and files corresponds to the git repo.
  ```
 ls
  ```

### Build your first docker image
In this step we are going to build a docker image and push it to a docker repo. For the next exercise it would be better to first build a server. If you have time left, you can always come back to here to build the job code.

1) Choose a programming language you are the most familiar with.
 ```
 cd Go/Server
 ```

2) Modify the message that the server will return to a personal message. This is the moment to send a love letter to a colleague :wink.

2) Build your docker image and name the image to your own name.
 ```
 docker build . --tag YOUR_NAME
 ```

3) Check your docker image.
 ```
 docker images
 ```

4) Tag your image to give it a proper version number
 ```
 docker tag YOUR_NAME YOUR_NAME:V1.0
 ```
5) Check your tagged image.
 ```
 docker images
 ```
6) Tag your image again with the docker registry as prefix. 
This is needed to push your locally created docker image to a registry where everyone can access it.
 ```
 docker tag YOUR_NAME:V1.0 eu.gcr.io/kubernetes-talk-259721/YOUR_NAME:V1.0
 ```

7) Push this tagged image to the docker registry.
 ```
 docker push eu.gcr.io/kubernetes-talk-259721/YOUR_NAME:V1.0
 ```

Congratulations! You've managed to create your first docker image and make it publicly accessible in a docker registry.

**_STOP:_** Too show you that we still have some issues, we can all run our container on the same VM and look what will happen.
 ```
 docker run YOUR_NAME
 ```
 To stop running this image, you need to execute crtl+C
 
## Tips

This link can be used if you can't access our kubernetes cluster on google cloud. But be aware that you will miss some functionality on this platform. 
https://www.katacoda.com/courses/kubernetes/playground