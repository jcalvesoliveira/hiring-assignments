## Assignment changes applied on this service

### Deploy
This service is running in a GKE instance on Google Cloud Platform and can be accessed at `35.223.122.149:50051`

### Redirect
Any document received by this API is sent for calculation at the [statistics-processing](https://github.com/jcalvesoliveira/statistics-processing) micro-service.

### Document or URL
This api is able to receive either a URL for downloading documents or documents bytes.



----------------------------
# Assignment used for Software Engineer position
This task is intended for candidates applying for a Software Engineer position at the Visma Machine Learning team. The assignment is built around the technologies and stack used in the production environments of the team, and the problem is a toy version of some of the tasks we face.


![Interview](./interview-gopher.png)


## The problem

Companies often have a lot of expenses. Each payment of these expenses is in accounting terms called a financial transaction which should be entered into the financial books. But where? This is the job of the accountant to figure out. An expense for paint might go to the maintenance account, and an expense for a taxi ride might go to the account for travel expenses.

In e-conomic it is possible to import your bank statements and use these to create the financial transactions, 
so all you have to do is to import the bank statements, and decide on which account you want to book each expense ..and you’re done.

We have tailored a machine learning solution that can learn how expenses should be booked. We continuously retrain our models with new human annotated expenses.

Sometimes new expenses that are uploaded are very different from expenses that a company has uploaded before. When that happens at a single company,
 it might just mean that they got a new accountant. But when many reports change, we may want to investigate if there are problems with the new data we receive.
 
We have decided that we want to summarize each new set of expenses we receive. We have set up a service that collects
 summary statistics and alerts us when there is significant drift. We are missing a solution that processes new reports and computes those statistics.

A key metric for us are categorical columns and the frequencies of each of the categories in a column. We want to calculate those numbers per CompanyId.

A summary might look like this:
```
,CompanyId,ColumnName,ColumnValue,Count
0,int:a055470,AccountTypeName,Balance,25
1,int:a055470,AccountTypeName,Profit and Loss,75
...
```

To keep up with different document formats, we want the service to automatically exclude columns that have too many different values from the summary table and we want to be able to specify other columns to aggregate by. 


## The sample code
We have provided the code for a simple gRPC API for the statistics service, along with YAML manifests for deploying the API to a Kubernetes cluster. The code currently just echos the expenses sent in the request, no statistics done yet.

### Prerequisites
You need to have a working setup for writing code in Go, building container images with Docker, and deploying to a Kubernetes cluster with eg. `kubectl`

### Local testing
The sample is written in Go, with the API defined by the Protocol Buffer found in `proto`. You can start the API locally by running:

```go run cmd/service/service.go```

And test it out with the client:

```go run cmd/client/client.go```

This will take the document in `test.csv`, send it to the service and and save the output to `out.csv`

### Deploy on Kubernetes
Firstly you need to build the Docker image and push it to some image registry that your cluster has access to. To deploy the service on a Kubernetes cluster use the manifests in `manifests`. The service and deployment in `api.yaml` should spin up the API service and get you an external ip to connect to.

## Your Task
We would like you to extend this system with actual statistics, and make the system more "production ready".

### Part 1 - The mandatory part
Create you own micro-service to do the statistics processing, separate from the API service. You can choose to write it in whatever programming language you are comfortable with. We expect it to:

* Contain an API definition in a separate protobuf file
* Handle requests through gRPC
* Take as part of the request:
    * The document originally sent to the API
    * Parameters for aggregation columns
    * Parameters for computing statistics over columns and columns to exclude

Alongside your code you should also write a manifest YAML to setup a deployment and service. The deployment should specify that your container should be replicated at least in 2 pods, which should be load-balanced through the service.

You will also have to add code in the existing API code in `pkg/api/api.go`, so that it will forward the document to your service. You should not have to change the existing API and protobuf. The API should satisfy the following:

* Be configurable, such that the desired aggregation options should be supplied at startup. How is up to you. You can use 25 values as the cut-off for categorical columns.
* Have an option to supply the API with a URL for downloading documents, instead of documents bytes(This is already present in the protobuf, but not implemented). Whether you want to download the document content in the API and send bytes to your service - or send the URL and download in the service - is up to you.

### Part 2 - The optional extras
Beyond the mandatory part we would like you to extend your service in some way. 

Choose one or more larger task from below:
* Add monitoring stats for number of requests per second, request times etc., to benchmark the performance of your system (Do what you find practical. The API has the basic scaffolding for exporting Prometheus metrics, but you would need to deploy a setup for that)
* Add autoscaling to your service. Write a new client that sends requests in a manner that triggers the autoscaling. __Quick question:__ *What considerations should you have when loadbalancing gRPC requests?*

Lastly, add one or more bonus features. This could be logging, error-handling, cloud storage, authentication, rate-limiting, more aggregation options, or anything you think could be relevant. You decide.
