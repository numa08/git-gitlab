# git + lab = gitlab

[![Build Status](https://travis-ci.org/numa08/git-gitlab.svg?branch=develop)](https://travis-ci.org/numa08/git-gitlab)

lab is a command line tool that wrap `git` in order to extend it with extra features and commands that make working with GitLab easier.


~~~sh
$ git lab clone numa08/dejiko

# extends to:
$ git clone git://${gitlab.url}/numa08/dejiko.git
~~~

## Installation

### Mac OS X

Install go-lang (if not already installed)

~~~sh
$ brew install go
~~~

If go-lang is installed, install libgit2

~~~sh
$ brew install libgit2
~~~

### Fedora

Install go-lang (if not already installed)

~~~sh
$ yum install go
~~~

If go-lang is installed, install libgit2-devel

~~~sh
$ yum install libgit2-devel
~~~

### Final steps

Install the required go packages

~~~sh
$ go get gopkg.in/libgit2/git2go.v23
$ go get github.com/plouc/go-gitlab-client
$ go get github.com/codegangsta/cli
~~~

Execute from within the git-gitlab repository clone

~~~sh
$ make build
~~~

Finally, move `build/git-lab` somewhere in $PATH.

#### Note

OSX El Capitan users may encounter the following error

~~~sh
could not determine kind of name for C.GIT_CHECKOUT_SAFE_CREATE
~~~

when running

~~~sh
$ go get gopkg.in/libgit2/git2go.v22
~~~

This is because libgit v0.23.x was installed. You need 0.22.x.
See https://github.com/libgit2/git2go/issues/181#issuecomment-129849557 for more information.

### Archlinux

There is an aur package called git-gitlab.

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
$ git lab clone numa08/dejiko
> git clone git://{gitlab.url}/numa08/dejiko

$ git lab clone -p numa08/dejiko
> git clone git@{gitlab.url}/:numa08/dejiko
~~~

### merge request

~~~sh
$ git lab merge-request -b basebranch -h headbranch
>open text editor to edit title and body
>open pull request on GitLab

$ git merge-request -b forked:branch -h origin:branch -m "Fix issue #xxx"
~~~

### git merge

~~~sh
$ git lab merge [MERGE_REQUEST_ID]
~~~

### git show

~~~sh
$ git lab show -- issue/10
> open http://gitlab.example.com/NAMESPACE/PROJECT/issues/10
~~~
