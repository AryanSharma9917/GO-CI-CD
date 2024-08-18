# CI CD pipeline for a go web application

### Run & Test application locally
- Clone repo
```
git clone https://github.com/AryanSharma9917/GO-CI-CD.git
```
- Build application
```
cd go-ci-cd
go build .
```
- Run the Tests
```
go test -v 
```
- Run application
```
./go-ci-cd
```
- Access application
```
curl localhost:8080/greet
curl localhost:8080/hello/Aryan
```
## Note: All the yml files are present inside `/ymls` directory.

### Pipeline (CI)

![fig-1](https://github.com/aryansharma9917/go-ci-cd/blob/main/screenshot/pipeline.png)


- Clone this repo 
```
git clone https://github.com/AryanSharma9917/go-ci-cd.git
```
- Login to Openshift cluster: 
```
oc login --token=<TOKEN> --server=<SERVER_URL>
```
- Create new project/namespace in Openshift cluster 
```
oc new-project go-app-pipeline
```
- Install Red Hat OpenShift Pipelines from Operators --> Operator Hub
- Run the following command to see the pipeline service account $
```
oc get serviceaccount pipeline
```

- Create tasks
```
oc apply -f ymls/tasks 
```
- Create pipeline
```
oc apply -f ymls/pipeline 
```
- Check the created pipeline and tasks
```
oc get pipeline,task 
```
- Run the pipeline
```
tkn pipeline start go-pipeline -w name=workspace-test,volumeClaimTemplateFile=https://raw.githubusercontent.com/AryanSharma9917/go-ci-cd/main/ymls/pvc/pvc.yml --use-param-defaults
```
- Check the pipelinerun
```
tkn pipelinerun ls
```
- Check the log of the pipelinerun
```
tkn pipelinerun logs <pipelinerun_name>
```
- Check the imagestream created by the latest pipelinerun
```
oc get is
```

### CD (ArgoCD)
- To install Openshift Gitops Operator follow this official doc: https://github.com/redhat-developer/gitops-operator/blob/master/docs/OpenShift%20GitOps%20Usage%20Guide.md
- Login to argo-cd console
- Create argocd application 
```
oc apply -f ymls/argocd/app.yml -n openshift-gitops
```
- ArgoCD Application

![fig-2](https://github.com/AryanSharma9917/go-ci-cd/blob/main/screenshot/argo-app.png)

- Granting permissions to a specific serviceaccount through role assignment to deploy the go application in namespace `go-app-pipeline`
```
oc adm policy add-role-to-user admin system:serviceaccount:openshift-gitops:openshift-gitops-argocd-application-controller -n go-app-pipeline
```
- Check the status of the deployed argocd application 
```
oc get application -n openshift-gitops
```

### Access the deployed application in Openshift

- Get all the resources in the namespace `go-app-pipeline`
```
$ oc get all -n go-app-pipeline

NAME                                                                 READY   STATUS      RESTARTS   AGE
pod/go-pipeline-run-kdtgn-build-and-push-to-openshift-registry-pod   0/1     Completed   0          31m
pod/go-pipeline-run-kdtgn-git-clone-pod                              0/1     Completed   0          32m
pod/go-pipeline-run-kdtgn-go-test-pod                                0/1     Completed   0          32m
pod/go-web-app-deployment-98bb59849-8prpm                            1/1     Running     0          35m

NAME                     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/go-web-app-svc   ClusterIP   172.30.130.60   <none>        8080/TCP   39m

NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/go-web-app-deployment   1/1     1            1           39m

NAME                                               DESIRED   CURRENT   READY   AGE
replicaset.apps/go-web-app-deployment-58c48f68db   0         0         0       39m
replicaset.apps/go-web-app-deployment-98bb59849    1         1         1       35m

NAME                                        IMAGE REPOSITORY                                                                                                  TAGS                                          UPDATED
imagestream.image.openshift.io/go-web-app   default-route-openshift-image-registry.apps.cndsno3.cee.ral3.lab.eng.rdu2.redhat.com/go-app-pipeline/go-web-app   latest,3f93d31c-dfdd-4ca5-9a8e-3a9a86eb9de4   28 minutes ago

NAME                                        HOST/PORT                                                                        PATH     SERVICES         PORT   TERMINATION   WILDCARD
route.route.openshift.io/go-web-app-route   go-web-app-route-go-app-pipeline.apps.cndsno3.cee.ral3.lab.eng.rdu2.redhat.com   /greet   go-web-app-svc   8080                 None

```

- Access the application using curl
```
$ curl go-web-app-route-go-app-pipeline.apps.cndsno3.cee.ral3.lab.eng.rdu2.redhat.com/greet 
        
Hello world!

```
