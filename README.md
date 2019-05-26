## prchecker

[![CircleCI](https://circleci.com/gh/go-go-megaman/prchecker.svg?style=svg)](https://circleci.com/gh/go-go-megaman/prchecker)

The prchecker package can fetch pull requests with open status from specific repositories. It can also filter pull requests by the author.

### Installation

Execute `go get` command.

```bash
go get -u github.com/go-go-megaman/prchecker
```

And create [YAML configuration](#configuration) in where you want to create. If you don't specified `--config`, the package read configuration form $HOME/.prchecker.yaml.

### Usage

```text
prchecker fetches pull requests from Github repositories.

Usage:
  prchecker [command]

Available Commands:
  help        Help about any command
  run         Display list of pull requests.

Flags:
      --config string   config file (default is $HOME/.prchecker.yaml)
  -h, --help            help for prchecker

Use "prchecker [command] --help" for more information about a command.
```

### Configuration

#### token

Use token you got from Github.

```yaml
tpken: TYPE_YOUR_TOKEN
```

#### repositories

Set repositories in the format.

```yaml
repositories:
    - TYPE_YOUR_ACCOUNT_NAME_OR_ORGANIZATION_NAME/TYPE_YOUR_REPOSITORY_NAME
```

#### authors

The authors configuration make fetched pull requests filter by setting values. This is optional configuration. If you didn't set this configuration, prchecker will not filter pull requests by author.

```yaml
authors:
    - TYPE_PULL_REQUEST_AUTHOR
```