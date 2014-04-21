require "gitlab"
require "grit"
require "git/error"

class GitlabApi::ApiClient
	
	module Authorize
		def authorize
			@client.user
		end
	end
	
end