# Git::Gitlab

[日本語](https://github.com/numa08/git-gitlab/blob/master/README_ja.md)

Command line tool for [Gitlab](https://www.gitlab.com/)

## Installation

Add this line to your application's Gemfile:

    gem 'git-gitlab'

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install git-gitlab

## Usage

### Set up

Config some setting

	git config --global gitlab.url http://gitlab.example.com/
	git config --gloabl gitlab.token GITLAB_SECRET_TOKEN
	git config gitlab.projectid GITLAB_PROJECTID

### Let' run and Confirm setting

	git gitlab

### Create Merge Request

	git gitlab merge SOURCE_BRANCH TARGET_BRANCH --assign USER_NAME

### Show Issue by ID

	git gitlab issue ISSUE_ID

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
