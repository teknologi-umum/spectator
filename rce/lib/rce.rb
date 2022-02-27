# frozen_string_literal: true

require "grpc"
require "dotenv/load"
require_relative "proto/rce/rce_services_pb"

# RceServer is an implementation for the gRPC server service
class RceServerImpl < Rce::CodeExecutionEngineService::Service
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
      language: _request['language'],
      version: "3.1.1",
      stdout: "Hello world",
      stderr: "",
      output: "Hello world",
      exitCode: 0
    )
  end

  def ping(_request, _call)
    Rce::PingResponse.new(message: "OK")
  end
end

if __FILE__ == $PROGRAM_NAME
  port = ENV["PORT"] || "50052"

  server = GRPC::RpcServer.new
  server.add_http2_port("0.0.0.0:#{port}", :this_port_is_insecure)

  puts("Running insecurely on 0.0.0.0:#{port}")
  puts("implemented? #{server.implemented?}")
  server.handle(RceServerImpl.new)
  server.run_till_terminated_or_interrupted([1, "int", "SIGQUIT"])
end
