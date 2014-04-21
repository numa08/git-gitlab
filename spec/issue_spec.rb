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

	it "should get all issues" do
		actual = @gitlab.issues(true).count
		actual.should == 650
	end

	it "should get opened issues" do
		actual = @gitlab.issues(false).count
		actual.should == 649
	end
end