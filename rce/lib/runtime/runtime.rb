# frozen_string_literal: true

# Runtime provides the access to a certain runtime
module Runtime
  class Runtime
    # @param [String] language
    # @param [String] version
    # @param [String] extension
    # @param [Boolean] compiled
    # @param [Array<String>] run_command
    # @param [Array<String>] build_command
    # @return [Runtime] a Runtime instance
    def initialize(language, version, extension, compiled, run_command, build_command)
      raise ArgumentError 'Language is required' unless language
      raise ArgumentError 'Version is required' unless version
      raise ArgumentError 'Extension is required' unless extension
      raise ArgumentError 'Compiled is required' unless compiled
      raise ArgumentError 'Run command is required' unless run_command

      if compiled
        raise ArgumentError 'Build command is required' unless build_command
      end

      @language = language
      @version = version
      @extension = extension
      @compiled = compiled
      @run_command = run_command
      @build_command = build_command
    end

    # @return [Array<String>]
    attr_reader :run_command

    # @return [Array<String>]
    attr_reader :build_command

    # @return [String]
    attr_reader :extension

    # @return [Boolean]
    attr_reader :compiled
  end
end
