#!/usr/bin/env bash
set -u -o -x -v -e

export GOPATH=$ROCKET_WORKSPACE
add-apt-repository ppa:ubuntu-lxc/lxd-stable
apt-get update
apt-get install -y golang cmake
go get github.com/codegangsta/cli
cd $HOME
wget -O libgit2-0.24.1.tar.gz https://github.com/libgit2/libgit2/archive/v0.24.1.tar.gz
tar -xzvf libgit2-0.24.1.tar.gz
cd libgit2-0.24.1 && mkdir build && cd build
cmake -DTHREADSAFE=ON -DBUILD_CLAR=OFF -DCMAKE_C_FLAGS=-fPIC -DCMAKE_BUILD_TYPE="RelWithDebInfo" -DCMAKE_INSTALL_PREFIX=/usr/local .. && make && sudo make install
ldconfig
cd $ROCKET_WORKSPACE
go get gopkg.in/libgit2/git2go.v24
go get github.com/plouc/go-gitlab-client
make build
cp -r build/ $ROCKET_ARTIFACTS
