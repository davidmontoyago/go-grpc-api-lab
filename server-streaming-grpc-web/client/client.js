const {EventRequest, EventType} = require('./api/server-streaming/streaming-service_pb.js');
const {EventStreamingServiceClient} = require('./api/server-streaming/streaming-service_grpc_web_pb.js');

var client = new EventStreamingServiceClient('http://localhost:8080');

var request = new EventRequest();
request.setType(EventType.CREATE);

var deadline = new Date();
deadline.setSeconds(deadline.getSeconds() + 3600 * 24)
console.log("deadline", deadline)
metadata = {deadline: deadline.getTime()}

stream = client.getEventStream(request, metadata);

stream.on('data', function(event) {
  console.log(event.getId());
  console.log(event.getDescription());
});
stream.on('status', function(status) {
  console.log(status.code);
  console.log(status.details);
  console.log(status.metadata);
});
stream.on('end', function(end) {
  console.log("stream ended")
});
