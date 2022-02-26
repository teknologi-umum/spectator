require 'toml'

directories = Dir.entries("./packages/")
directories.each { |dir|
  if dir == "." || dir == ".."
    next
  end

  # Check if the install.sh file exists
  exists = File.exist? "./packages/#{dir}/install.sh"
  if exists
    Open3.popen3("./packages/#{dir}/install.sh") do |stdin, stdout, stderr, wait_thr|
      if wait_thr.value != 0
        puts(stderr)
        exit(1)
      end

      puts(dir + " installed!")
    end

    # Read the config.toml file
    config_exists = File.exists? "./packages/" + dir + "/config.toml"
    unless config_exists
      throw "Config file for #{dir} does not exists!"
    end

    config_file = File.read("./packages/#{dir}/config.toml")
    config = TOML::Parser.new(config_file).parsed

    # Run test files
    if config["compiled"] == true
      Open3.popen3([*config["build_command"].slice(0, config["build_command"].length - 1), config["test_file"]]) do |stdin, stdout, stderr|
        if stderr
          puts("Error compiling #{config["test_file"]}")
          puts(stderr)
          exit(1)
        end
      end

      Open3.popen3(config["run_command"]) do |stdin, stdout, stderr|
        if stderr
          puts("Error running #{config["test_file"]}")
          puts(stderr)
          exit(1)
        end

        if stdout != "Hello world!"
          puts("Expecting \"Hello world!\", got #{stdout}")
          exit(1)
        end
      end
    else
      Open3.popen3([*config["run_command"].slice(0, config["run_command"].length - 1), config["test_file"]]) do |stdin, stdout, stderr|
        if stderr
          puts("Error running #{config["test_file"]}")
          puts(stderr)
          exit(1)
        end

        if stdout != "Hello world!"
          puts("Expecting \"Hello world!\", got #{stdout}")
          exit(1)
        end
      end
    end
  end
}
