require File.dirname(__FILE__) + '/spec_helper'
require "git/gitlab"

describe GitlabKernel, "create new mergerequest" do
	
	before do
		@gitlab = GitlabKernel.new

		@test_config = {:projectid => 8,
			   :project_path => "numa08/git-gitlab",
			   :gitlab_url => "http://vagrant-ubuntu-precise-64/"}
	end

	it "should return mergerquest url" do
		actual = @gitlab.create_merge_request("test", nil, "tem", "master")
		expect(actual).to include("#{@test_config[:gitlab_url]}#{@test_config[:project_path]}/merge_requests")
	end	
end