require './proto/helloworld_services_pb.rb'
require './proto/notification_services_pb.rb'

hello_client = Sample::Greeter::Stub.new('127.0.0.1:5502', :this_channel_is_insecure)
hello_request = Sample::HelloRequest.new(first_name: 'Alan', last_name: 'Turing')
hello_response = hello_client.say_hello(hello_request)
puts "Message: #{hello_response.message}"

notification_client = Sample::Notifier::Stub.new('127.0.0.1:5502', :this_channel_is_insecure) 
notification_request = Sample::HelloRequest.new(first_name: 'Alan', last_name: 'Turing')
notification_resopnses = notification_client.pereodic_hello(notification_request)
notification_resopnses.each do |response|
  puts "Message: #{response.message}"
end
