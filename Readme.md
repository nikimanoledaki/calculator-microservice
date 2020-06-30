# Calculator Microservice

A simple microservice to do basic arithmetics.

# How to run the gRPC server and client

First run the server:

```
go run cmd/server/main.go
```

Then build the client in another terminal:

```
go build -o client cmd/client/main.go
```

Now you can use the CLI client to find the sum or average of two numbers!

```
./client sum 3 8
./client average 9.5 10
```

# Running the tests

The unit and integration tests are written with `Ginkgo` and `Gomega`.

```
ginkgo -r
```

https://github.com/fullstorydev/grpcurl
