require "git/gitlab/version"
require "gitlab"
require "grit"


module GitlabApi
	class ApiClient
		API_VERSION = "v3"
		USER_AGENT = "nyonyonyonyonyo! Gitlab nyo!!"
		attr_reader :client, :repository

		def initialize
			@repository = Grit::Repo.new(`git rev-parse --show-toplevel`.chomp)
			config = @repository.config

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
	end
end