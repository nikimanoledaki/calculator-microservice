[![Codefresh build status](https://g.codefresh.io/api/badges/pipeline/zabon/calculator-microservice?key=eyJhbGciOiJIUzI1NiJ9.NWVmYjk4MGM5Zjg4MTkzOTRjZTkzM2Q0.qIGEzYTOB3eZyFH-SLjUYJJzjue2FGMHoGEnJ9h11mw&type=cf-1)](https%3A%2F%2Fg.codefresh.io%2Fpipelines%2Fcalculator-microservice%2Fbuilds%3Ffilter%3Dtrigger%3Abuild~Build%3Bpipeline%3A5efb9893e8d6bb7c8b1aa55c~calculator-microservice)

# Calculator Microservice

This is a simple microservice that can be used to do basic arithmetic operations. It has a client and a server that implement two methods from a gRPC interface, one to add two `int32` values and one to find the average of two `float32` values. The client is CLI-based and prints out the result of the calculation.

## 1. Implementation

### Start the server and the client

To set up the server, pull the container image from Docker Hub. Then, run the container to start the server, which listens on port `9092` by default.

```
docker pull niki2401/calculator-microservice
docker run -d -p 9092:9092  niki2401/calculator-microservice
```

If for some reason this does not work for you, you can also run the server locally.

```
go run cmd/server/main.go
```

Then, in another terminal, build the client.

```
go build -o client cmd/client/main.go
```

Now you can use the CLI client to find the sum of two integers or the average of two floats!

```
./client sum 3 8
./client average 9.5 10
```

### Testing

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

[grpcurl](https://github.com/fullstorydev/grpcurl) was used a lot in the making of this service. It is a very easy to use open-source tool that can be used to do manual feature tests to test a gRPC server. It was particularly useful while building the server, before there was a client to test that the server functioned properly.

Here is how to install it if you would like to try it out:

```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

To use it, run the server locally in one terminal, and then run the following command in another terminal.

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

It is also a pretty neat way to get more information about specific methods and messages available in the gRPC interface.

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

Many of the requirements of the 12factor app methodology can be satisfied through the use of Go, containers (runc, containerd), Docker, and Kubernetes. Most of the 12 points are "checked" by this service. Points 4 (Backing Service) and 12 (Admin Processes) are a bit out of scope, but this service fares well otherwise!

**1 Codebase:** The code is managed with Git and an image of it is also created with Docker, which can be pulled from Docker Hub. The container image is also specified in the Kubernetes Deployment for the service.

**2 Dependencies:** All dependencies are explicitly declared in `go.mod` and locked in `go.sum`. They have also been copied into the container through the Dockerfile.

**3 Config:** `ENV` variables are set in the Dockerfile. A ConfigMap is not provided with this service but it could be used to fulfill this requirement.

**5 Build, release, run:** Running one or multiple instances of this service can be done by implementing the Deployment and the ReplicaSet provided in the `Deploy/Kubernetes` folder.

**6 Processes:** Some might argue that the `StatefulSet` config might breach this point because it stores some state, but there is no state actually being stored in this service since its implementation is very ephemereal.

**7 Port binding:** The port of this service can be exported because of the `Service` of type `LoadBalancer` that applies to this application.

**8 Concurrency & 9 Disposability:** Kubernetes makes it easy to scale the service by making replicas and the containers themselves are ephemeral and easy to reproduce and to quickly spin up.

**10 Dev/prod parity:** The use of lightweight virtualization such as containers minimizes the gap between development and production because the container runtime makes it easy to run environment variables, tools and services in isolation.

**11 Logs:** This service relies heavily on printing things to `stdout` and relies on it heavily for error-handling as well.

### Prove how it fits and uses the best cloud native understanding

### How would you expand on this service to allow for the use of an eventstore?

### How would this service be accessed and used from an external client from the cluster?

According to the Kubernetes [docs](https://kubernetes.io/docs/tutorials/hello-minikube/#create-a-service), resources are only accessible by their internal IP address within the Kubernetes cluster and must be exposed as a Service in order to be accessed and used from an external client from the cluster, meaning outside of the Kubernetes virtual network. In this case, a Service of type `LoadBalancer` applies to our `Calculator` app to expose `port 9092` internally and externally of our cluster.
