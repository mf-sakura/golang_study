# frozen_string_literal: true

require './proto/helloworld_services_pb.rb'
require './proto/notification_services_pb.rb'

HOST = '127.0.0.1:5502'

hello_client = Sample::Greeter::Stub.new(HOST, :this_channel_is_insecure)
hello_request = Sample::HelloRequest.new(first_name: 'Alan', last_name: 'Turing')
puts "Message: #{hello_client.say_hello(hello_request).message}"

notification_client = Sample::Notifier::Stub.new(HOST, :this_channel_is_insecure) 
notification_request = Sample::HelloRequest.new(first_name: 'Alan', last_name: 'Turing')
notification_resopnses = notification_client.pereodic_hello(notification_request)
notification_resopnses.each do |response|
  puts "Periodic: #{response.message}"
end
