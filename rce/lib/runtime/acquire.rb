require 'toml'
require_relative 'runtime'

module Runtime
  class Acquire
    def acquire
      runtimes = []

      Dir.entries("./packages") do |dir|
        config_file = File.read("./packages/#{dir}/config.toml")
        config = TOML::Parser.new(config_file).parsed

        runtime = Runtime.new(
          config['language'],
          config['version'],
          config['extension'],
          config['compiled'],
          config['run_command'],
          config['build_command'])

        runtimes.append(runtime)
      end

      runtimes
    end
  end
end
