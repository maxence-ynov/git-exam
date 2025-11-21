#!/usr/bin/env bash

git clone $1 /tmp/git_exam
cd /tmp/git_exam
git checkout feature-audio-support
echo "func tacos(){}" >> main.go
git add . && git commit -m "add : func tacos (to dev)" && git push
rm -rf /tmp/git_exam