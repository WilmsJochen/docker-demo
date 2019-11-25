# Ex1: Build your own docker container and push it to the cloud.

Good you found this git repo to start the exercise from.
 
##Choose your language

In this project you can find some basic examples for different programming languages.
Needless to say that this list of languages can be expanded. 

## google cloud
To build this docker images, i set up a google account to log in to the google cloud console.
Please follow theses steps:

1) go to google cloud console: https://console.cloud.google.com/

2) login as kubernetesTalk user with the credentials from the ppt.
    
   **_NOTE:_**  If you use the chrome browser you should probably add the kubernetestalk user in your browser.

3) use the navigation menu to navigate to the Kubenetes engine
4) Connect to the cluster in a new terminal.
5) create a folder with your own name:

```
mkdir YOUR_NAME
```
6) Change directory to that folder.
```
cd YOUR_NAME
```

From this terminal we can execute all commands for the next exercises.

##instructions
### checkout Git repo
To use our code on the virtual machine, we need to clone our git repo.

1) clone the git repo
```
git clone https://github.com/WilmsJochen/ex1.git
```
2) change your directory to that repo
```
cd ex1
```
3) have a look on all the folders and files in that directory.
 Verify all the directories and files corresponds to the git repo.
 ```
ls
 ```

### build your first docker image
In this step we are going to build a docker image and push it to a docker repo.

1) choose a programming language and functional folder (Server/Job). (You can repeat these steps later with another folder/code)
```
cd Go/Job
```

2) build your docker image and name the image to your own name.
```
docker build . --tag YOUR_NAME
```
3) check your image and those of your colleagues.
```
docker images
```
4) tag your image to give it a proper version number
```
docker tag YOUR_NAME YOUR_NAME:V1.0
```
5) check your tagged image.
```
docker images
```
6) tag your image again with the docker registry as prefix. 
This is needed to push your local docker image to the registry where everyone can access it.
```
docker tag YOUR_NAME:V1.0 eu.gcr.io/kubernetes-talk-259721/YOUR_NAME:V1.0
```

Congratulations! you've managed to create your first docker image and make it publicly accessible in a docker registry.

