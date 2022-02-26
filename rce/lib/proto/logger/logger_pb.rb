# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: logger.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("logger.proto", :syntax => :proto3) do
    add_message "logger.EmptyResponse" do
    end
    add_message "logger.EmptyRequest" do
    end
    add_message "logger.LogData" do
      optional :request_id, :string, 1
      optional :application, :string, 2
      optional :message, :string, 3
      proto3_optional :level, :enum, 4, "logger.Level"
      proto3_optional :environment, :enum, 5, "logger.Environment"
      proto3_optional :language, :string, 6
      map :body, :string, :string, 7
      proto3_optional :timestamp, :int64, 8
    end
    add_message "logger.LogRequest" do
      optional :access_token, :string, 1
      optional :data, :message, 2, "logger.LogData"
    end
    add_message "logger.Healthcheck" do
      optional :status, :string, 1
    end
    add_message "logger.ReadLogRequest" do
      proto3_optional :level, :enum, 1, "logger.Level"
      proto3_optional :request_id, :string, 2
      proto3_optional :application, :string, 3
      proto3_optional :timestamp_from, :int64, 4
      proto3_optional :timestamp_to, :int64, 5
    end
    add_message "logger.ReadLogResponse" do
      repeated :data, :message, 1, "logger.LogData"
    end
    add_enum "logger.Level" do
      value :DEBUG, 0
      value :INFO, 1
      value :WARNING, 2
      value :ERROR, 3
      value :CRITICAL, 4
    end
    add_enum "logger.Environment" do
      value :UNSET, 0
      value :DEVELOPMENT, 1
      value :TESTING, 2
      value :STAGING, 3
      value :PRODUCTION, 4
    end
  end
end

module Logger
  EmptyResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.EmptyResponse").msgclass
  EmptyRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.EmptyRequest").msgclass
  LogData = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.LogData").msgclass
  LogRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.LogRequest").msgclass
  Healthcheck = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.Healthcheck").msgclass
  ReadLogRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.ReadLogRequest").msgclass
  ReadLogResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.ReadLogResponse").msgclass
  Level = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.Level").enummodule
  Environment = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("logger.Environment").enummodule
end
