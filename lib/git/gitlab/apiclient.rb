require "git/gitlab/version"
require "git/error"
require "gitlab"
require "grit"
require "uri"

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

		def project_id
			pid = if @repository.config.keys.include? "gitlab.projectid"
					@repository.config["gitlab.projectid"].to_i
				else
					@repository.config["gitlab.project"]
				end

			if pid == nil
				raise "Please set 'git config gitlab.projectid ${Gitlab Project id}' of git config gitlab.project ${NAMESPACE/PROJECT}"
			end

			begin
				@client.project( URI.encode_www_form_component(pid.to_s)).id
			rescue Gitlab::Error::NotFound => e
				raise GitlabApi::Error::ProjectIdNotFound, "Can not find #{pid} Project."
			end

		end
	end
end