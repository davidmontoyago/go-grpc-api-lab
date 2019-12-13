const {HelloRequest, HelloResponse} = require('./api/hello-world/hello-service_pb.js');
const {HelloServiceClient} = require('./api/hello-world/hello-service_grpc_web_pb.js');

var client = new HelloServiceClient('http://localhost:8080');

var request = new HelloRequest();
request.setGreeting('World');

client.sayHello(request, {}, (err, response) => {
  console.log(response);
});