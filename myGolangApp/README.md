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
        
# Kubernetes Section

## Get Credentials
  * gcloud container clusters get-credentials la --zone=us-central1-c
  * kubectl get services
  * kubectl get deployments
  * kubectl delete pod myapp-deployment-xxxx
  * kubectl apply -f myapp-deployment.yaml --record
  
## Custom Output
  * kubectl get pods -o=custom-columns=NODES_LIST:.spec.nodeName,DEPLOYMENTS_NAME:.metadata.name
  
## Making 1 node in not ready state to make pods not to be scheduled
  * kubectl taint nodes gke-la-default-pool-e91d65ad-69pb key=value:NoSchedule

## Making 1 node in ready state to make pods to be scheduled
  * kubectl taint nodes gke-la-default-pool-e91d65ad-69pb key:NoSchedule-

## Rollout Status
  * kubectl rollout status deployment.v1.apps/myapp-deployment
  * kubectl rollout history deployment.v1.apps/myapp-deployment
  * kubectl rollout undo deployment.v1.apps/myapp-deployment  