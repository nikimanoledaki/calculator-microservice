.PHONY: protos

protos: 
	protoc --go_out=plugins=grpc:protos/calculator ./protos/calculator.proto