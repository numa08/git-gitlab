module GitlabApi
	module Error
		class MergeRequestError < StandardError; end
		class IssueNotFound < StandardError; end
		class ProjectIdNotFound < StandardError; end
	end
end