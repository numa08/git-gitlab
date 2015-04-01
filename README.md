# git + lab = gitlab

lab is a command line tool that wrap `git` in order to extend it with extra features and commands that make working with GitLab easier.


~~~sh
$ git clone numa08/dejiko

# extends to:
$ git clone git://${gitlab.url}/numa08/dejiko.git
~~~

## Installation

TODO

## Configure

Set the value of Gitlab-url and token.

~~~sh
$ git config --global gitlab.url http://gitlab.example.com

$ git config --global gitlab.token GITLAB_SECRET_TOKEN
~~~

then, set gitlab namespace and project.

~~~sh
$ git config gitlab.project [NAMESPACE]/[PROJECT]
~~~

## Commands

### git clone

~~~sh
$ git clone numa08/dejiko
> git clone git://{gitlab.url}/numa08/dejiko

$ git clone -p numa08/dejiko
> git clone git@{gitlab.url}/:numa08/dejiko
~~~

### merge request

~~~sh
$ git merge-request -b basebranch -h headbranch
>open text editor to edit title and body
>open pull request on GitLab

$ git merge-request -b forked:branch -h origin:branch -m "Fix issue #xxx"
~~~

### git merge

~~~sh
$ git merge [MERGE_REQUEST_ID]
~~~

### git show

~~~sh
$ git show -- issue/10
> open http://gitlab.example.com/NAMESPACE/PROJECT/issues/10
~~~