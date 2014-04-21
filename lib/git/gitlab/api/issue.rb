require "gitlab"
require "grit"
require "git/error"


class GitlabApi::ApiClient
	
	module Issue
		def issue(id)
			pid = @repository.config["gitlab.projectid"].to_i

			if pid == 0
				raise "Please set 'git config gitlab.projectid ${Gitlab Project id}'"
			end

			per_page = 100
			page = 1

			def all_issues(pid, page, per_page)
				def _issues(list, pid, page, per_page)
					i = @client.issues(pid, :page => page, :per_page => per_page)
					if i.count < per_page
						list + i
					else
						_issues(list + i, pid, page + 1, per_page)
					end

				end

				i = @client.issues(pid, :page => page, :per_page => per_page)
				if i.count < per_page
					i
				else
					_issues(i, pid, page + 1, per_page)
				end
			end

			issues = all_issues(pid, page, per_page)

			specfied_issue = issues.select { |issue|
				issue.iid == id.to_i
			}.first

			if specfied_issue == nil
				raise GitlabKernel::Error::IssueNotFound, "Issue not find"
			else
				specfied_issue
			end
		end
	end
	
end