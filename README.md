# hwtest
Hello world on golang

1. Simple web app on golang which uses *net/http* and *html/template*. Template file has a version number placeholder. One can build it by cloning the repo locally:
```
# https://github.com/tutolmin/hwtest.git
# cd hwtest
# go build -o hwtest hwtest.go
```
Web app binds to port 8080. Check the service is responding:
```
# ./hwtest &
# curl http://localhost:8080
<h1>Hello world: 0.0.4
</h1>
```
2. Github repo has a webhook to notify https://hub.docker.com/r/tutolmin/hwtest on every pull_request and push. Docker hub has an Automated Build configured which uses *Dockerfile* from the repo. Once build is complete one can get a docker image locally and start a container binding port 8080:
```
# docker run -d -p 8080:8080 tutolmin/hwtest
# docker ps
CONTAINER ID   IMAGE             COMMAND      CREATED         STATUS         PORTS                    NAMES
5ebd147f99aa   tutolmin/hwtest   "./hwtest"   5 seconds ago   Up 4 seconds   0.0.0.0:8080->8080/tcp   practical_lovelace
```
Now, that our web application is running in a docker container we can still access it with curl:
```
# curl http://localhost:8080
<h1>Hello world: 0.0.4
</h1>
```
3. Use the following deploy.yml file and run *kubectl create -f deploy.yml*
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deploy
  labels:
    apps: hwtest
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hwtest
  template:
    metadata:
      labels:
        app: hwtest
    spec:
      containers:
      - name: hwtest
        image: tutolmin/hwtest
        ports:
        - containerPort: 8080
```
... and *service.yml* to create a service *kubectl create -f service.yml*
```
kind: Service
apiVersion: v1
metadata:
  name: my-service
spec:
  selector:
    app: hwtest
  ports:
  - protocol: TCP
    port: 32768
    targetPort: 8080
```
As a result two pods will be created running our web application in it.
```
# kubectl get pods
NAME                        READY   STATUS    RESTARTS   AGE
my-deploy-659947555-bp2k4   1/1     Running   0          11h
my-deploy-659947555-pq8gb   1/1     Running   0          11h

# kubectl describe service my-service
Name:              my-service
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          app=hwtest
Type:              ClusterIP
IP Families:       <none>
IP:                10.111.33.253
IPs:               10.111.33.253
Port:              <unset>  32768/TCP
TargetPort:        8080/TCP
Endpoints:         10.244.0.10:8080,10.244.0.9:8080
Session Affinity:  None
Events:            <none>

# curl http://10.111.33.253:32768
<h1>Hello world: 0.0.4
</h1>
```
4. Have not got any experience with OpenShift yet.
5.
  * Version file has a version number in the following format x.y.z When a new version of the web app is pushed to the repo one have to tag it with:
```
git tag -a v0.0.4 -m "version 0.0.4"
git push origin v0.0.4
```
Adding the tag will trigger a build of an additional build at docker hub. There is a build rule to extract version number from the repo tag and build a *tagged* version of the image, ex. *tutolmin/hwtest:version-0.0.4* One can pull the tagged image locally like this:
```
# docker run -d -p 8080:8080 tutolmin/hwtest:version-0.0.4
CONTAINER ID   IMAGE                           COMMAND      CREATED         STATUS         PORTS                    NAMES
e46937744a99   tutolmin/hwtest:version-0.0.4   "./hwtest"   7 seconds ago   Up 5 seconds   0.0.0.0:8080->8080/tcp   confident_dhawan
```
  * I can only think of modifying a *.gitignore* file to include the following:
```
*
!version
!hwtest.go
```
  *
