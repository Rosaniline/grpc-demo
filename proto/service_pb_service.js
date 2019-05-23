// package: 
// file: service.proto

var service_pb = require("./service_pb");
var person_pb = require("./person_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var HuaFu = (function () {
  function HuaFu() {}
  HuaFu.serviceName = "HuaFu";
  return HuaFu;
}());

HuaFu.GetPerson = {
  methodName: "GetPerson",
  service: HuaFu,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.getPersonRequest,
  responseType: person_pb.Person
};

exports.HuaFu = HuaFu;

function HuaFuClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

HuaFuClient.prototype.getPerson = function getPerson(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(HuaFu.GetPerson, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.HuaFuClient = HuaFuClient;

