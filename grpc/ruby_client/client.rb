# frozen_string_literal: true

require 'grpc'
require './proto/helloworld_services_pb'

client = Sample::Greeter::Stub.new(
  '127.0.0.1:5502',
  :this_channel_is_insecure
)

opts = { first_name: 'John', last_name: 'Doe' }
response = client.say_hello(Sample::HelloRequest.new(opts))
puts response
