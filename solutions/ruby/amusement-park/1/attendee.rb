class Attendee
  attr_reader :height
  attr_reader :pass_id
  def initialize(height)
    @height = height
    @pass_id = nil
  end

  # def height
  #   raise 'Implement the Attendee#height method'
  # end

  # def pass_id
  #   raise 'Implement the Attendee#pass_id method'
  # end

  def issue_pass!(pass_id)
    @pass_id = pass_id
  end

  def revoke_pass!
    @pass_id = nil
  end
end
