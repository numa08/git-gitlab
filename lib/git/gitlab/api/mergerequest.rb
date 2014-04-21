require "gitlab"
require "grit"
require "git/error"

class GitlabApi::ApiClient
	
	module Mergerequest
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
				raise GitlabApi::Error::MergeRequestError, "Failed Create Merge Request"
			end

			project_url = @client.project(pid).web_url
			mergerequest_url = project_url + "/merge_requests/" + mergerequest.iid.to_s

			mergerequest_url
		end
	end
end