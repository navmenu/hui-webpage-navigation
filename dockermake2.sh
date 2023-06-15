#!/bin/bash
set -eu

branch=$(git branch --show-current)
commit=$(git log -n 1 --oneline | awk '{print $1}')

echo "branch is ${branch}"
echo "commit is ${commit}"

#date=$(date +%Y%m%d%H%M%S)
#tag="${branch}.${date}.${commit}"
#指定版本是因为其他地方是固定的
tag="alpha.0001"
echo "tag is ${tag}"

docker build --pull --force-rm --no-cache -t "reg.huiwang.io/duotouge/web_navigation:${tag}" -f Dockerfile2 .

# 需要手动登陆dockerhub再执行推送
# 首先执行登陆 docker login reg.huiwang.io
# docker push "reg.huiwang.io/duotouge/web_navigation:${tag}" | awk '/digest/{print $3}'
