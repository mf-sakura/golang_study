require 'grpc'
require './proto/helloworld_services_pb'

include Sample

def main
  stub = Greeter::Stub.new('127.0.0.1:5502', :this_channel_is_insecure)

  hello_request = HelloRequest.new(first_name: "Roger", last_name: "Federer")
  hello_reply = stub.say_hello(hello_request)
  puts hello_reply.message
end

main
