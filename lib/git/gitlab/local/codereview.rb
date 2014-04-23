require "git/gitlab/version"
require "git/error"
require "grit"

class GitlabLocal
	
	module CodeReview
		
		def review(title ,source, target, remote)
			status = @repository.status
			is_modified = status.added.count > 0  ||
						  status.changed.count >0 ||
					  	  status.deleted.count >0
	  	    if is_modified
	  	    	raise GitlabApi::Error::ReviewError, `git status`
	  	    end

  	    	puts(`git fetch #{remote}`)
  	    	puts(`git checkout -b #{title} #{remote}/#{target}`)
  	    	puts(`git merge --no-ff --no-commit #{remote}/#{source}`)
		end
	end
end