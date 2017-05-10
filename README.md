[![Build status](https://badge.buildkite.com/ac92706a4a7dd0e5bbd042ea409c5e5fcc24fb279f1f78c0c4.svg)](https://buildkite.com/myob/ops-go-cfn)

# ops-go-cfn
## Helpers using AWS GoSDK

    docker-compose run --rm build
    docker-compose run --rm run

## Prerequisites

Using https://github.com/Masterminds/glide

On Mac: `brew install glide`

Setup project like so:
```
mkdir -p ${HOME}/go/src/github.com/MYOB-Technology
export GOPATH=${HOME}/go
cd ${HOME}/go/src/github.com/MYOB-Technology
git clone git@github.com:MYOB-Technology/ops-go-cfn.git
cd ops-go-cfn
glide install
```

