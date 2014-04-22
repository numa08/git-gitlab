require "gitlab"
require "grit"
require "git/error"

class GitlabApi::ApiClient
	
	module Mergerequest
		PER_PAGE = 100
		PAGE = 1

		def create_merge_request(title, assign, source, target = "master")
			pid = project_id

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

		def mergerequests
			pid = project_id

			all_mergerequests(pid, PAGE, PER_PAGE).select { |m|
				m.state == "opened"
			}
		end

		def all_mergerequests(pid, page, per_page)
			def _mergerequests(list, pid, page, per_page)
				m = @client.merge_requests(pid, :page => page, :per_page => per_page)
				if m.count < per_page
					list + m
				else
					_mergerequests(list + m, pid, page +  1, per_page)
				end
			end

			m = @client.merge_requests(pid, :page => page, :per_page => per_page)
			if m.count < per_page
				m
			else
				_mergerequests(m, pid, page + 1, per_page)
			end
		end
	end
end