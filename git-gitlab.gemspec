# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'git/gitlab/version'

Gem::Specification.new do |spec|
  spec.name          = "git-gitlab"
  spec.version       = Git::Gitlab::VERSION
  spec.authors       = ["numa08"]
  spec.email         = ["n511287@gmail.com"]
  spec.description   = "Gitlab Command line interface"
  spec.summary       = "git-gitlab command line"
  spec.homepage      = "https://github.com/numa08/git-gitlab"
  spec.license       = "MIT"

  spec.files         = `git ls-files`.split($/)
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.test_files    = spec.files.grep(%r{^(test|spec|features)/})
  spec.require_paths = ["lib"]

  spec.add_development_dependency "bundler", "~> 1.3"
  spec.add_development_dependency "rake"
  spec.add_development_dependency "rspec", "~> 2.14.1"
  spec.add_dependency "thor" , "~> 0.19.1"
  spec.add_dependency "gitlab", "~> 3.0.0"
  spec.add_dependency "grit", "~> 2.5.0"
end
