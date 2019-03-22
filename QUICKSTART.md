# Quick Start

## install dependencies

1. Setup AWS credential profile, with a file such as `~/.aws/credentials` like:

        [yourprofileid]
        aws_access_key_id = ...
        aws_secret_access_key = ...

2. tweak `up.json`, especially the `profile`, `regions`, `stages.*.domain` attributes
3. install Go and set `GOPATH` environment variable
4. `curl -sf https://up.apex.sh/install | sh`
5. `up upgrade`
6. `go get`

## test run

    up start -o

## staging deploy

    up

## production deploy

	up deploy production

## subsequent redeploys

    git pull
	up start -o
	up
	up deploy production

## notes

* additional steps will be necessary if DNS is hosted elsewhere.
* additional steps will be necessary to receive the confirmation link
  email for certificate generation.
