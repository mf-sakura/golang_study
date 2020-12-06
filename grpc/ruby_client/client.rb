# frozen_string_literal: true

require 'grpc'
require './proto/helloworld_services_pb'
require './proto/notification_services_pb'

# 通常のレスポンス
http_client = Sample::Greeter::Stub.new(
  '127.0.0.1:5502',
  :this_channel_is_insecure
)

request_opts = Sample::HelloRequest.new(first_name: 'John', last_name: 'Doe')
http_response = http_client.say_hello(request_opts)
puts http_response

puts '====================='

# Streamなレスポンスを受け取ったとき
stream_client = Sample::Notifier::Stub.new(
  '127.0.0.1:5502',
  :this_channel_is_insecure
)

request_opts = Sample::HelloRequest.new(first_name: '太郎', last_name: '兼進')
stream_request = stream_client.pereodic_hello(request_opts)
stream_request.each { |req| puts req }
