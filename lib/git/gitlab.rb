require "git/gitlab/version"
require "gitlab"
require "grit"

module GitlabApi
	API_VERSION = "v3"
	USER_AGENT = "nyonyonyonyonyo! Gitlab nyo!!"
	attr_reader :client

	def initialize
		config = Grit::Repo.new(`git rev-parse --show-toplevel`.chomp).config

		url = config["gitlab.url"]
		if url == nil
			raise "Plase set 'git config gitlab.url ${Gitlab URL}'"
		end

		token = config["gitlab.token"]
		if token == nil
			raise "Please set 'git config gitlab.token ${Gitlab Token}'"
		end

		Gitlab.configure do |config|
		  config.endpoint       = "#{url}/api/#{API_VERSION}"
		  config.private_token  = token
		  config.user_agent     = USER_AGENT
		end

		@client = Gitlab.client
	end

	def authorize
		user = @client.user
		puts("You are #{user.username}")		
		puts("Gitlab is #{Gitlab.endpoint}")
	end
end


class GitlabKernel
	include GitlabApi
end