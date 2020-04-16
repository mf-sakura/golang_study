setup:
	go get -u github.com/golang/protobuf/protoc-gen-go
	gem install grpc-tools


gen:
	# https://grpc.io/docs/tutorials/basic/go/
	protoc proto/*.proto --go_out=plugins=grpc:server;
	protoc proto/*.proto --go_out=plugins=grpc:go_client;
	# https://grpc.io/docs/tutorials/basic/ruby/
	grpc_tools_ruby_protoc --ruby_out=ruby_client --grpc_out=ruby_client proto/*.proto
