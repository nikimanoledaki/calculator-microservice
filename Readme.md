[![Codefresh build status](https://g.codefresh.io/api/badges/pipeline/zabon/calculator-microservice?key=eyJhbGciOiJIUzI1NiJ9.NWVmYjk4MGM5Zjg4MTkzOTRjZTkzM2Q0.qIGEzYTOB3eZyFH-SLjUYJJzjue2FGMHoGEnJ9h11mw&type=cf-1)](https%3A%2F%2Fg.codefresh.io%2Fpipelines%2Fcalculator-microservice%2Fbuilds%3Ffilter%3Dtrigger%3Abuild~Build%3Bpipeline%3A5efb9893e8d6bb7c8b1aa55c~calculator-microservice)

# Calculator Microservice

This is a simple microservice to do basic arithmetic operations. It has a gRPC client and server that implement two methods, one to add two `int32` values and one to find the average of two `float32` values. It has a CLI-based client that prints out the result of the calculations.

## 1. Implementation

### Start the server and client

To set up the server, pull the container image from Docker Hub and then start the server on port `9092`.

```
docker pull niki2401/calculator-microservice
docker run -d -p 9092:9092  niki2401/calculator-microservice
```

If this does not work for you, you can set up the server locally. The server will isten to port `9092` by default.

```
go run cmd/server/main.go
```

Then build the client in another terminal.

```
go build -o client cmd/client/main.go
```

You can use the CLI client to find the sum of two integers or the average of two floats!

```
./client sum 3 8
./client average 9.5 10
```

### Testing

#### Unit Tests

Unit tests use the `Ginkgo`/`Gomega` dependecies and can be run with the following command if ginkgo is installed on your machine and can be found on your `$PATH`. Add the `-cover` flag to run the tests using Go's code coverage tool.

To install ginkgo and gomega:

```
 go get github.com/onsi/ginkgo/ginkgo
 go get github.com/onsi/gomega/...
```

To run the tests:

```
ginkgo -r -cover
```

The following commands can be used to inspect the test coverage in depth, using the test coverage file for `Client` as an example.

```
go tool cover -func=client.coverprofile
go tool cover -html=client.coverprofile
```

#### grpcurl

[grpcurl](https://github.com/fullstorydev/grpcurl) is a very nice open-source go-based tool to do manual feature tests with sending and receiving requests. It was particularly useful while building the server, before there was a client to test that the server functioned properly.

To install it:

```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

To use it:

```
$ grpcurl --plaintext -d '{"FirstNumber": 1, "SecondNumber": 5}' localhost:9092 Calculator.GetSum

// Client view
{
  "Result": 6
}

// Server view
2020-06-30T19:41:53.038+0100 [INFO]  Handle GetSum: firstNumber=1 secondNumber=5
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

It is also pretty neat to get more information about specific methods and messages available in the gRPC interface.

```
$ grpcurl --plaintext localhost:9092 describe Calculator
Calculator is a service:
service Calculator {
  rpc GetAverage ( .AverageRequest ) returns ( .AverageResponse );
  rpc GetSum ( .SumRequest ) returns ( .SumResponse );
}
```

```
$ grpcurl --plaintext localhost:9092 describe Calculator.GetAverage
Calculator.GetAverage is a method:
rpc GetAverage ( .AverageRequest ) returns ( .AverageResponse );
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

**Prove how it aligns to 12factor app best practices**</br>

**Prove how it fits and uses the best cloud native understanding**</br>
**How would you expand on this service to allow for the use of an eventstore?**</br>
**How would this service be accessed and used from an external client from the cluster?** </br>
   According to the Kubernetes [docs](https://kubernetes.io/docs/tutorials/hello-minikube/#create-a-service), resources are only accessible by their internal IP address within the Kubernetes cluster. For resources to be accessed and used from an external client from the cluster, meaning outside of the Kubernetes virtual network, the Pod that has the Service container must be exposed as a Service.

   In this case, a `NodePort` type of Service applies to our `Calculator` app to exposes `port 9092` internally andexternally from our cluster. 
