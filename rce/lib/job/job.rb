# frozen_string_literal: true

# Job provides the priming, execution, and result of a code
module Job
  class Job
    # @param [User] user
    # @param [Runtime] runtime
    # @param [String] code
    # @param [Integer] timeout
    # @param [Integer] memory_limit
    def initialize(user, runtime, code, timeout = 5000, memory_limit = 128 * 1024)
      raise ArgumentError 'User is required' unless user
      raise ArgumentError 'Runtime is required' unless runtime
      raise ArgumentError 'Code is required' unless code

      # Acquire user
      @user_id = user.next_available
      @compiled = runtime.compiled
      @extension = runtime.extension
      @build_command = runtime.build_command
      @run_command = runtime.run_command
      @code = code
      @args = args
      @timeout = timeout
      @memory_limit = memory_limit
    end

    def execute
      if @compiled
        script_path = create_file
        compiled_path = compile(script_path)
        result = safe_call(compiled_path)

      end
    end

    # create_file just writes a file into the filesystem and return the path.
    #
    # the owner would be the current user id with the group int of 4505.
    # the file mode is always 555.
    # @return [String (frozen)]
    def create_file
      filepath = "/code/#{@user_id}/code.#{@extension}"
      out_file = File.new(filepath)
      out_file.puts(@code)
      out_file.chown(@user_id, 4505)
      out_file.chmod(555)
      out_file.close

      filepath
    end

    # @param [String] path
    # @return [String]
    def compile(path) end

    def safe_call(path)
      cmd_builder = StringIO.new
      cmd_builder << 'nosocket'
      cmd_builder.read

      # return [stdout, stderr, output, exit_code]
    end
  end
end
