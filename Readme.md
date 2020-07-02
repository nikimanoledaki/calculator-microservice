[![Codefresh build status](https://g.codefresh.io/api/badges/pipeline/zabon/calculator-microservice?key=eyJhbGciOiJIUzI1NiJ9.NWVmYjk4MGM5Zjg4MTkzOTRjZTkzM2Q0.qIGEzYTOB3eZyFH-SLjUYJJzjue2FGMHoGEnJ9h11mw&type=cf-1)](https%3A%2F%2Fg.codefresh.io%2Fpipelines%2Fcalculator-microservice%2Fbuilds%3Ffilter%3Dtrigger%3Abuild~Build%3Bpipeline%3A5efb9893e8d6bb7c8b1aa55c~calculator-microservice)

# Calculator Microservice

This is a simple microservice that can be used to do basic arithmetic operations. It has a client and a server that implement two methods from a gRPC interface, one to add two `int32` values and one to find the average of two `float32` values. The client is CLI-based and prints out the result of the calculation.

## 1. Implementation

### Technologies

- Testing: Ginkgo, Gomega, grpcurl

- Code Quality: Golint

- Containerisation: Docker

- CI/CD: Codefresh

- Container orchestration: Kubernetes, kind

### Use

To set up the server, pull the container image from Docker Hub. Then, run the container to start the server, which listens on port `9092` by default.

```
docker pull niki2401/calculator-microservice
docker run -d -p 9092:9092  niki2401/calculator-microservice
```

If for some reason this does not work for you, you can also run the server locally.

```
go run cmd/server/main.go
```

Once the server is running, build the client in another terminal.

```
go build -o client cmd/client/main.go
```

Now you can use the CLI client to find the sum of two integers or the average of two floats!

```
./client sum 3 8
./client average 9.5 10
```

### Test

#### Unit and Integration Tests

To run the tests, you will need to have `Ginkgo` and `Gomega`, which can be installed with the following commands.

```
 go get github.com/onsi/ginkgo/ginkgo
 go get github.com/onsi/gomega/...
```

The tests are located in the `Test` folder and can be run with the following command.

```
ginkgo -r test
```

#### Feature tests

[grpcurl](https://github.com/fullstorydev/grpcurl) was used a lot in the making of this service. It is a neat open-source tool that can be used to do manual feature tests to test a gRPC server.

It can be installed like this:

```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

If you would like to use it, run the server locally in one terminal (by running the Docker image or running it locally), and then enter the following command in another terminal.

```
// Client terminal
$ grpcurl --plaintext -d '{"FirstNumber": 1, "SecondNumber": 5}' localhost:9092 Calculator.GetSum

// Server view
2020-06-30T19:41:53.038+0100 [INFO]  Handle GetSum: firstNumber=1 secondNumber=5

// Client view
{
  "Result": 6
}
```

```
$ grpcurl --plaintext -d '{"FirstNumber": 1.0, "SecondNumber": 2.0}' localhost:9092 Calculator.GetAverage

// Client view
{
  "Result": 1.5
}

// Server view
2020-06-30T19:40:28.249+0100 [INFO]  Handle GetAverage: firstNumber=1 secondNumber=2
```

Here is more information about specific the methods of the service's gRPC interface:

```
$ grpcurl --plaintext localhost:9092 describe Calculator
Calculator is a service:
service Calculator {
  rpc GetAverage ( .AverageRequest ) returns ( .AverageResponse );
  rpc GetSum ( .SumRequest ) returns ( .SumResponse );
}
```

```
$ grpcurl --plaintext localhost:9092 describe .AverageRequest
AverageRequest is a message:
message AverageRequest {
  float FirstNumber = 1;
  float SecondNumber = 2;
}
```

## 2. Documentation

### Prove how it aligns to 12factor app best practices

Many of the requirements of the 12factor app methodology can be satisfied through the use of Go, containers (runc, containerd), Docker, and Kubernetes. The list below illustrates how many of the 12 points are "checked" by this service. Points 4 (Backing Service) and 12 (Admin Processes) are a bit out of scope, but this service fares well otherwise!

**1 Codebase:** The code is on Github and its image can be pulled from Docker Hub, which is the same as the one specified in the Kubernetes Deployment for the service.

**2 Dependencies:** All dependencies are explicitly declared in `go.mod` and locked in `go.sum`. They are copied into the container, which is declared in the Dockerfile.

**3 Config:** `ENV` variables are set in the Dockerfile. A ConfigMap is not included in this service but it could be used to fulfill this requirement.

**5 Build, release, run:** Running one or multiple instances of this service can be done by implementing the `Deployment` and the `ReplicaSet` provided in the `Deploy/Kubernetes` folder.

**6 Processes:** There is no state actually being stored in this service and its implementation is very ephemereal.

**7 Port binding:** The port of this service can be exported because of the `Service` of type `LoadBalancer` that applies to the service.

**8 Concurrency & 9 Disposability:** Kubernetes makes it easy to scale the service by making replicas. The containers themselves are ephemeral and easy to reproduce and to quickly spin up.

**10 Dev/prod parity:** The use of lightweight virtualization through containers minimizes the gap between development and production because the container runtime makes it easy to run environment variables, tools and services in isolation.

**11 Logs:** This service relies heavily on printing things to `stdout` and relies on it heavily for error-handling as well.

### Prove how it fits and uses the best cloud native understanding

Official definition of "Cloud Native" from the Cloud Native Computing Foundation - **[CNCF Cloud Native Definition v1.0](https://github.com/cncf/foundation/blob/master/charter.md#1-mission-of-the-cloud-native-computing-foundation)**

> Cloud native technologies empower organizations to build and run scalable applications in modern, dynamic environments such as public, private, and hybrid clouds. Containers, service meshes, microservices, immutable infrastructure, and declarative APIs exemplify this approach.
> </br></br>
> These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.

This service fits and uses the best cloud native understanding given that it checks off most of the 12factor requirements. It implements microservice best practices such as containerisation, Docker, and a gRPC interface. It can run independently and is easy to deploy within an orchestrator such as Kubernetes, so it can benefit from the monitoring and configuration capabilities of Kubernetes, which makes this service resilient, observable, and easy to replicate.

### How would you expand on this service to allow for the use of an eventstore?

The server of the calculator service could be a "consumer" of an eventstore stream so that it can [read events](https://eventstore.com/docs/getting-started/reading-subscribing-events/index.html?tabs=tabid-6%2Ctabid-dotnet-client%2Ctabid-dotnet-read-event%2Ctabid-create-sub-dotnet) by subscribing to an event stream. Events could be written to the stream by a client of the calculator service, which would mean that the calculator server and the calculator client share an event store database. This way, the event store could be a broker for the transactions between the client and the server. The eventstore would need the capability to receive `GET` and `POST` requests to its HTTP API interface so, within a Kubernetes cluster, it would have to be matched with a `Service` that exposes a `PORT` for it. There are different [subscription types](https://eventstore.com/docs/getting-started/reading-subscribing-events/index.html?tabs=tabid-6%2Ctabid-dotnet-client%2Ctabid-10%2Ctabid-create-sub-dotnet#subscription-types) that could be used by the service depending on the use case.

### How would this service be accessed and used from an external client from the cluster?

According to the Kubernetes [docs](https://kubernetes.io/docs/tutorials/hello-minikube/#create-a-service), resources are only accessible by their internal IP address within the Kubernetes cluster and must be exposed as a Service in order to be accessed and used from an external client from the cluster, meaning outside of the Kubernetes virtual network. In this case, a Service of type `LoadBalancer` applies to our `Calculator` app to expose `port 9092` internally and externally of our cluster.
