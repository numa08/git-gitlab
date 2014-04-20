module GitlabApi
	module Error
		class MergeRequestError < StandardError; end
		class IssueNotFound < StandardError; end
	end
end