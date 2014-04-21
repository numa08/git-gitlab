require File.dirname(__FILE__) + '/spec_helper'
require "git/gitlab"

describe GitlabKernel, "get issue" do
	
	before do
		@gitlab = GitlabKernel.new
	end

	it "should return issue title" do
		actual = @gitlab.issue(1).title
		actual.should == "the issue"
	end	
end