# Steps for Docker Image

* docker build -t myapp .
* docker images
* docker history myapp
    * It shows all the history for the docker Image
* docker run -d -p 8888:8888 myapp
* docker ps
* curl http://localhost:8888
* Syntax for Docker Push with Tag
    * docker tag <app_name> <RegistryName>/<Project_Name>/<App_Name>:<Version/Tag>
        * docker tag myapp gcr.io/playground-s-11-82a90540/myapp:limeGreen
* docker push syntax
    * docker push <RegistryName>/<Project_Name>/<App_Name>:<Version/Tag>
        * docker push gcr.io/playground-s-11-82a90540/myapp:limeGreen