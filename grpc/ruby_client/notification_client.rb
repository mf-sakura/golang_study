require 'grpc'
require './proto/notification_services_pb'

include Sample

def main
  stub = Notifier::Stub.new('127.0.0.1:5502', :this_channel_is_insecure)

  pereodic_hello_request = PereodicHelloRequest.new(first_name: "Roger", last_name: "Federer")
  resp = stub.pereodic_hello(pereodic_hello_request)
  resp.each do |r|
    puts r.message
  end
end

main
