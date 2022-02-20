# frozen_string_literal: true

# User provides the functionality of interacting with Linux users
class User
  # @param [Integer] _upper_limit
  def initialize(_upper_limit = 50)
    @current_id = []
    @upper_limit = 50
  end

  # @return [NilClass]
  def next_available
    unless @current_id.empty?
        return acquire(@current_id.last + 1) if @current_id.last < @upper_limit
        return acquire(0) if @current_id.last == @upper_limit
        raise SecurityError 'How did you ended up here?'
      end
    end

    acquire(0)
  end

  # @param [Integer] user_id
  # @return [Integer]
  def acquire(user_id)
    @current_id.append(user_id)
    user_id
  end

  # @param [Integer] user_id
  def release(user_id)
    @current_id = @current_id.select {|id| id != user_id}
  end
end
