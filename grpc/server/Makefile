build:
	go build -o bin/grpc_server_example

hello:
	grpcurl -plaintext -d '{"first_name": "yuto", "last_name": "sakura"}' localhost:5502 sample.Greeter/SayHello
pereodic_hello:
	grpcurl -plaintext -d '{"first_name": "yuto", "last_name": "sakura"}' localhost:5502 sample.Notifier/PereodicHello
list_services:
	grpcurl -plaintext localhost:5502 list
describe_service:
	grpcurl -plaintext localhost:5502 describe  sample.Greeter.SayHello
describe_message:
	grpcurl -plaintext localhost:5502 describe sample.HelloRequest
