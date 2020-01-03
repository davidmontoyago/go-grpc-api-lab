/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.api = require('./streaming-service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.EventStreamingServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.EventStreamingServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.EventRequest,
 *   !proto.api.Event>}
 */
const methodDescriptor_EventStreamingService_GetEventStream = new grpc.web.MethodDescriptor(
  '/api.EventStreamingService/GetEventStream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.api.EventRequest,
  proto.api.Event,
  /**
   * @param {!proto.api.EventRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Event.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.EventRequest,
 *   !proto.api.Event>}
 */
const methodInfo_EventStreamingService_GetEventStream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.Event,
  /**
   * @param {!proto.api.EventRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.Event.deserializeBinary
);


/**
 * @param {!proto.api.EventRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.api.Event>}
 *     The XHR Node Readable Stream
 */
proto.api.EventStreamingServiceClient.prototype.getEventStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/api.EventStreamingService/GetEventStream',
      request,
      metadata || {},
      methodDescriptor_EventStreamingService_GetEventStream);
};


/**
 * @param {!proto.api.EventRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.api.Event>}
 *     The XHR Node Readable Stream
 */
proto.api.EventStreamingServicePromiseClient.prototype.getEventStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/api.EventStreamingService/GetEventStream',
      request,
      metadata || {},
      methodDescriptor_EventStreamingService_GetEventStream);
};


module.exports = proto.api;

