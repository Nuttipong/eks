// AWS
// Config
aws configure list

// Profile
aws configure list-profiles
*** https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html

// ECR
aws ecr create-repository –-repository-name demo​
aws ecr describe-repositories --repository-name demo
aws ecr describe-repositories --repository-name demo --profile ecr
aws ecr get-login-password --region eu-central-1 --profile ecr | docker login --username AWS --password-stdin 991654915716.dkr.ecr.eu-central-1.amazonaws.com

// Docker
// Build with Tag 
docker build -t 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:fe1.0 . -f dockerfile.scratch
docker build -t 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:be-foo-v1.0 . \
-f dockerfile.foo.scratch
docker build -t 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:be-bar-v1.0 . \
-f dockerfile.bar.scratch

or
docker tag e47cf4c537d8 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:fe1.0

// Push
docker push 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:fe1.0
docker push 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:be-bar-v1.0
docker push 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:be-foo-v1.0

// Run
docker run -d -p 4001:4001 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:fe1.0
docker run --rm -it -p 4001:4001 1600be4fdfeb

// Exec
docker exec --rm -it 85d3f6e675e4 sh
// Curl inside container
apk add curl
curl -s http://localhost:4001

// kubectl
// When your cluster is ready can check by typing
kubectl cluster-info

// Context
kubectl get namespace
kubectl config get-clusters
kubectl config get-contexts
kubectl config current-context
kubectl config use-context my-cluster-name
kubectl config set-context --current --namespace=dev

// Create Cluster


// Create Deployment
// Imperative
k create deployment web-app --image=991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:fe1.0
// Declarative
k create -f ./frontend.deployment.yml --dry-run
k create -f ./frontend.deployment.yml --save-config
k apply -f ./frontend.deployment.yml
//The kubectl diff command allows you to see a diff between currently running resources, and the changes proposed in the supplied configuration file:
k diff -f ./frontend.deployment.yaml
kubectl patch deploy ./frontend.deployment.yaml -p '{"spec": {"replicas": 4}}'

// Pod
// Exec Pod
k exec web-app-5f76887d87-njfwb -it sh
apk add curl
curl -s http://10.251.10.45:8080
k exec web-app-5f76887d87-njfwb -- curl -s http://10.251.10.45:8080
OR
k run -it --rm curl --restart=Never --image=991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:1.0 -- 172.20.108.108:8080
k run -it --rm curl --restart=Never --image=991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:1.0 -- web-app:8080

// Service
k expose deployment web-app --port 8080 --target-port 8080
k expose deployment web-app --port=8080 --target-port=8080 --type=ClusterIP
k expose deployment web-app --type=NodePort --port 5000 --target-port 5000
k expose deployment web-app --type=LoadBalancer --port 5000 --target-port 5000
echo $(kubectl get service web-app -o jsonpath='{ .spec.clusterIP }')
k edit svc web-app
k apply -f frontend.service.yml
k expose deployment web-app --type=NodePort --port 5000 --target-port 5000​
k get endpoints
for i in backend.deploy*.yml; do k apply -f $i; done

// Scale

// Update

// Ingress
// Create a deployment, scale it to 2 replicas and expose it as a serivce. 
kubectl create deployment web-app --image=991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:v1.0
kubectl scale deployment web-app --replicas=2

// This service will be ClusterIP and we'll expose this service via the Ingress.
kubectl expose deployment web-app --port=80 --target-port=8080 --type=ClusterIP

// Apply ingress
kubectl apply -f ingress-single.yaml

// Get ingress
kubectl get ingress
kubectl get services --namespace ingress-nginx

// Describe
kubectl describe ingress ingress-single

// EC2
// Check vailability zone
aws ec2 describe-availability-zones --region eu-central-1


// Bluegreen
docker-compose build --parallel
docker-compose push --parallel
docker-compose up
docker-compose down
docker-compose run --rm nginx-blue

docker push 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:nginx-blue-v1.0
docker push 991654915716.dkr.ecr.eu-central-1.amazonaws.com/demo:nginx-green-v1.0

k apply -f nginx-blue.deployment.yml
k apply -f nginx-blue.svc.yml

k apply -f nginx-green.deployment.yml
k apply -f nginx-green.svc.yml

# Change Service's selector to green (imperative)
k set selector svc nginx-svc 'role=green'