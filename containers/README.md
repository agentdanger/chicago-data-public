# chicago-data | containers

This directory contains images for each container that will house a single microservice in our full-stack data application. We'll also deploy these microservices using continuous integration with Docker Hub.

Microservice images will be registered with Docker Hub in private registries.
Source code and continuous integration will be housed on Github.com in this monorepo.
Deployment and services will be provided by AWS Elastic Kubernetes Services using EC2 as our compute engine.

## Connection details:

Run the following command to configure kubectl to use our EKS clusters:

	aws eks update-kubeconfig --region <region-code> --name <cluster-name>

Run the following command to configure kubectl to use our local cluster:

	minikube start
