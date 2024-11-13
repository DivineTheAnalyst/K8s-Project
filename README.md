# K8s-Project

This is my first extensive Kubernetes project, showcasing much of what I've learned in Kubernetes so far.

**THE OBJECTIVE**

☸️ Deploy and expose 3 different applications:

- **MongoDB**, using a ClusterIP service (internal access only)
- **Mongo Express**, using a NodePort service for external access on a specific port
- **Custom Prime Number Generator App (in Go),** using a LoadBalancer service

☸️ Write and deploy a Horizontal Pod Autoscaler (HPA) manifest to manage CPU load, targeting 50% utilization for the LoadBalancer app.

☸️ Create a restricted Kubernetes user with permissions limited to listing and describing deployments. 

**THE PROCESS**

In this repository you will find 4 different folders;
- RBAC
- go-app
- mongo-express
- mongodb

Each folder contains manifests and configurations for the services along with RBAC roles. I will be talking about each of them extensively in the order I created and deployed them.

**MongoDB**

MongoDB is a very popular open source NoSQL database management program. There are two files in this directory;
- mongoDB-deployment.yaml: This file is the manifest file where I defined the specifications of this app and I used the official Docker image 'Mongo' to pull it. I decided to write the service manifest file in the same deployment file, since they're co-dependent on one another. I exposed with ClusterIP service which is internal and apps defined using this service can only be accessed on the cluster. So to interact with MongoDB on my cluster I used the command

  _kubectl run -it --rm --image=mongo 4.4 --restart=Never mongo-client -- /bin/sh -c "mongo --host mongodb-service --port 27017"_

  ![Screenshot 2024-11-12 172108](https://github.com/user-attachments/assets/a0473f07-5b68-4755-a8a0-b2537abd0f27)

- mongodb-secret.yaml: This is where I wrote my mongo root username and password, which will be needed when connecting to mongo express. The username and password is written in base64. The secret manifest file is very important for security, it enables us not to write the password directly in our deployment manifest. in the environment variables section in the deployment manifest, the offical name for specifying the username- MONGO_INITDB_ROOT_USERNAME, and password-MONGO_INITDB_ROOT_USERNAME are gotten directly from the official image on Docker

**Mongo-Express**

Mongo Express is a lightweight web-based administrative interface for managing MongoDB databases. The two files in this folder are;

- mongo-express.yaml: The deployment and service manifests are in the same file. The NodePort service exposes Mongo Express to be accessed externally. I specified port 30004 for access and used environment variables from mongodb-secret.yaml to connect to MongoDB. Accessing <my-node-ip>:30004 in my browser opened Mongo Express. To see my application I simply pasted '<my-node-ip>:30004' on my browser and I was able to view it.

  Input my credentials
  
![Screenshot 2024-11-12 165454](https://github.com/user-attachments/assets/0c418a52-f39c-4841-ab12-d8150193ace7)

  VIOLA!
  
![Screenshot 2024-11-12 165944](https://github.com/user-attachments/assets/204f29c3-0e8e-4e96-8476-5fb958135727)

- mongodb-configmap.yaml: The config map was used to configure the database server. It directly connects the mongo express to mongodb service. This is specified in the environment variable section in the mongo express deployment manifest. A ConfigMap is an API object used to store configuration files as key-value pairs.

**PRIME NUMBERS GO APP** 

I exposed this app using the LoadBalancer service. I also wanted to practice Horizontal Pod Autoscaling to perform a CPU load test and exceed a utilization target of 50%. I will be going through the 5 files in this directory in tbe order in which I opened it. Let's Gooo;

This lightweight app outputs prime numbers in an infinite loop, useful for CPU stress testing with HPA.

- primenumbers.go: Lists prime numbers continuously to create a CPU load, which I could monitor and scale with HPA.

- Dockerfile: I containerized the app, built it, and pushed it to Docker Hub. The deployment references it as ijeawele/go-primenumbers:latest.

 ![image](https://github.com/user-attachments/assets/4cce7b36-218e-4994-b1d5-bfedf90790b1)
  

- go-primenum-deployment.yaml: This manifest sets CPU and memory limits for the app to test HPA. Two replicas are initially deployed, scaling to a maximum of ten as CPU utilization exceeds 50%.

- go-primenum-service.yaml: The service manifest file for the Go app, using LoadBalancer to expose the app externally. I could access it by copying the external IP from kubectl get svc.

   This is my successfully deployed app viewed on my browser:
  
  ![Screenshot 2024-11-12 200854](https://github.com/user-attachments/assets/a4eacd40-c0ac-4258-bf49-3dc64299485d)
                                              the external IP address is in the search bar

- go-primenums-hpa.yaml: HPA scales pods based on CPU usage. This file sets a minimum of 2 and a maximum of 10 replicas, scaling up as load increases. My app quickly triggered maximum scaling due to its recursive nature, and HPA reached 200%/50% utilization.

  ![Screenshot 2024-11-12 201252](https://github.com/user-attachments/assets/d2d606e1-0e22-44ef-bd0c-512aa9ec7dff)

**ROLE BASED ACCESS CONTROL** *(RBAC)*

Role-Based Access Control (RBAC) in Kubernetes is a security mechanism that restricts access to resources in a cluster based on the roles assigned to users, service accounts, or applications. With RBAC, administrators can define what actions a user or application is allowed to perform within the cluster, ensuring tighter security and control over Kubernetes resources. 

 Here’s what I implemented:

- IAM User Creation: Created an IAM user on AWS and installed AWS IAM Authenticator.

- RBAC Role Definition:
  
  - rbac-role-divine.yaml: Allows the user to list and describe deployments only.
    
  - rbac-rolebinding-divine.yaml: Binds the IAM user to the specified role, ensuring the restricted permissions.

  
**This project helped me implement key Kubernetes concepts, especially multi-service deployment, HPA configuration, and secure access control through RBAC. It’s been a rewarding experience putting my knowledge into practice.**
  
