require "git/gitlab/version"
require "git/error"
require "gitlab"
require "grit"

module GitlabApi
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

	def authorize
		user = @client.user
		puts("You are #{user.username}")		
		puts("Gitlab is #{Gitlab.endpoint}")
	end

	def create_merge_request(title, assign, source, target = "master")
		pid = @repository.config["gitlab.projectid"].to_i

		if pid == 0
			raise "Please set 'git config gitlab.projectid ${Gitlab Project id}'"
		end

		mr_title = if title == nil
			@repository.head.name
		else
			title
		end

		mr_source = if source == nil
			@repository.head.name
		else
			source
		end

		assignee_id = if assign == nil
			0
		else
			@client.users.select { |user|
				user.username == assign
			}[0].id
		end

		begin
			mergerequest = if assignee_id > 0
				@client.create_merge_request(pid, mr_title, :source_branch => mr_source, :target_branch => target, :assignee_id => assignee_id)
			else
				@client.create_merge_request(pid, mr_title, :source_branch => mr_source, :target_branch => target)
			end
		rescue Gitlab::Error::NotFound => e
			raise Error::MergeRequestError, "Failed Create Merge Request"
		end

		project_url = @client.project(pid).web_url
		mergerequest_url = project_url + "/merge_requests/" + mergerequest.id.to_s

		mergerequest_url
	end
end

class GitlabKernel
	include GitlabApi
end