require "gitlab"
require "grit"
require "git/error"


class GitlabApi::ApiClient

  module Issue
    PER_PAGE = 100
    PAGE = 1

    def issue(id)
      pid = project_id

      issues = all_issues(pid, PAGE, PER_PAGE)

      specfied_issue = issues.select { |issue|
        issue.iid == id.to_i
      }.first

      if specfied_issue == nil
        raise GitlabApi::Error::IssueNotFound, "Issue not find"
      else
        specfied_issue
      end
    end

    def issues(with_closed)
      pid = project_id

      issues = all_issues(pid, PAGE, PER_PAGE)
      if with_closed
        issues
      else
        issues.select { |i|
          i.state == "opened"
        }
      end
    end

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

  end

end
