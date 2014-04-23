require "git/gitlab/version"
require "git/gitlab/apiclient"
require "git/gitlab/api/issue"
require "git/gitlab/api/authorize"
require "git/gitlab/api/mergerequest"
require "git/gitlab/gitlablocal"
require "git/gitlab/local/codereview"
require "git/error"
require "gitlab"
require "grit"

class GitlabKernel < GitlabApi::ApiClient
	include GitlabApi::ApiClient::Issue
	include GitlabApi::ApiClient::Authorize
	include GitlabApi::ApiClient::Mergerequest
end

class GitlabLocalRepository < GitlabLocal
	include GitlabLocal::CodeReview
end