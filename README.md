# Building and Running Applications on Kubernetes
## Startup Slam 2018

### 0. Set Up Resources (5-15 min)
Make sure Docker is installed: https://store.docker.com/search?type=edition&offering=community

You will need some kind of cluster and container registry. Many systems (Azure, AWS, Minikube, etc) will do. We'll use Google Cloud in this demo. Create an account here: https://console.cloud.google.com/freetrial

Ensure you have the required SDK and kubectl. For Google Cloud:

* Set up the SDK: https://cloud.google.com/sdk/
* TODO initialize
* TODO add kubectl
* gcloud auth configure-docker

### 1. Build and push container images (3 min)
TODO note about public image

`cd <git directory>`

Build the container.
`docker build -t demo ./app`

Tag the container.
`docker tag demo us.gcr.io/[GCLOUD PROJECT ID]/demo:1`

Push the container to the registry.
`docker push us.gcr.io/[GCLOUD PROJECT ID]/demo:1`

## 2. Create a Kubernetes cluster (2 min)
We only need a cluster with a small quantity of resources. A 2-thread machine is recommended, as Kubernetes itself will also require resources.

https://console.cloud.google.com/kubernetes

TODO defaults

## 3. Create the deployment (1 min)
`kubectl create -f k8s/deployment.yml`
