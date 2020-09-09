#!/bin/bash

set -e

export SMOKE_DIR="$( pwd -P )"

cd test
. ./smoke.common.sh
trap cleanup EXIT

echo SMOKE_DIR=$SMOKE_DIR

downloadFootloose
generateKey
createCluster
generateYaml

./footloose ssh root@manager0 "cd /launchpad/test; DISABLE_TELEMETRY=true ACCEPT_LICENSE=true ../bin/launchpad --debug apply"
