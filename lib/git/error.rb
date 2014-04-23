module GitlabApi
	module Error
		class MergeRequestError < StandardError; end
		class IssueNotFound < StandardError; end
		class ProjectIdNotFound < StandardError; end
		class ReviewError < StandardError; end
	end
end