require "grit"

class GitlabLocal
	attr_reader :repository

	def initialize
		@repository = Grit::Repo.new(`git rev-parse --show-toplevel`.chomp)
	end
end