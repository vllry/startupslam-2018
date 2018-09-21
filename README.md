# Building and Running Applications on Kubernetes
## Startup Slam 2018

### 0. Set Up Resources (5-15 min)

You will need some kind of Kubernetes, cluster and container registry. Many systems (Azure, AWS, Minikube, etc) will do. We'll use Google Cloud in this demo. Create an account here: https://console.cloud.google.com/freetrial

Ensure you have the required SDK and kubectl. For Google Cloud:

* Set up the SDK: https://cloud.google.com/sdk/
* `gcloud init`
* `gcloud components install kubectl`

**YOU MAY WANT TO SKIP DOCKER ON UVIC WIFI**

* Install Docker: https://store.docker.com/search?type=edition&offering=community
* `gcloud auth configure-docker`

### 1. Build and push container images (3 min on fast internet)
**As noted above**, you may not want to do the Docker steps yourself on UVic wifi (it requires substantial uploading and downloading). I have a public repo with `us.gcr.io/vallery-demo/demo`.

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
`kubectl apply -f k8s/deployment.yml`

The [GUI](https://console.cloud.google.com/kubernetes/workload), or `kubectl get deployments` and `kubectl get pods` should show the deployment spinning up. 

## 4. Create the service (1 min)
`kubectl apply -f k8s/service.yml`

## 5. Create the ingress (5 min)
You will need to customize the ingress if not using the Google Load Balancer.

`kubectl apply -f k8s/ingress.yml`

The ingress may take several minutes to be created. You can check the load balancer [GUI](https://console.cloud.google.com/net-services/loadbalancing) or `kubectl describe ingress app` to check on progress.

The GUI and `kubectl describe ingress app | grep Address | awk -v N=2 '{print $N}'` will show the IP address. This address is ephemeral, normally an address is reversed in advance and specified as an ingress annotation.

## 6. Create the autoscaler (1 min)

`kubectl apply -f ./k8s/autoscaler.yml`
