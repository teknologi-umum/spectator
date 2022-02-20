# frozen_string_literal: true

# typed: true

require "grpc"
require "dotenv/load"
require_relative "proto/rce_services_pb"

# RceServer is an implementation for the gRPC stub
class RceServer < Rce::CodeExecutionEngineService::Service
  def list_runtimes(_request, _call)
    Rce::Runtimes.new(runtime: [Rce::Runtime.new(
      language: "Ruby",
      version: "3.1.1",
      aliases: ["rb"],
      compiled: false
    )])
  end

  def execute(_request, _call)
    Rce::CodeResponse.new(
      language: "Ruby",
      version: "3.1.1",
      stdout: "Hello world",
      stderr: "",
      output: "Hello world",
      code: 0
    )
  end

  def ping(_request, _call)
    Rce::PingResponse.new(message: "OK")
  end
end

if __FILE__ == $PROGRAM_NAME
  server = GRPC::RpcServer.new
  server.add_http2_port("0.0.0.0:50051", :this_port_is_insecure)
  server.handle(RceServer.new)
  server.run_till_terminated_or_interrupted([1, "int", "SIGQUIT"])
end
