![Peephole](https://github.com/immobiliare/peephole/blob/master/.github/immobiliare-labs.png)

# Peephole

[![pipeline status](https://github.com/immobiliare/peephole/actions/workflows/test.yml/badge.svg)](https://github.com/immobiliare/peephole/actions/workflows)

Peephole is a web-based events explorer of SaltStack transactions.

It can be used to watch for events on several SaltStack master nodes by GETting _from the_ `/events` API endpoint.
As soon as events are watched, they get persisted into a local (auto-expiring and configurable) cache that allows to explore, filter and correlate events for a variable time interval.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Changelog](#changelog)
- [Contributing](#contributing)
- [Issues](#issues)

## Install

Either install using Docker:

```bash
docker pull ghcr.io/immobiliare/peephole:latest
```

Or build it yourself:

```bash
git clone https://github.com/immobiliare/peephole.git
cd peephole
make build
```

## Usage

First off, create a configuration file `config.yml` accordingly.
The sample [example.yml](../example.yml) file can be used as a reference.

Either run it with Docker:

```bash
docker run -v ${PWD}/config.yml:/etc/peephole \
  -p 8080:8080 ghcr.io/immobiliare/peephole:latest
```

Or using the manually built binary file:

```bash
./peephole -c config.yml
```

Clearly, for each SaltMaster defined in the configuration file, ensure to enable [`salt-api`](https://docs.saltproject.io/en/latest/ref/cli/salt-api.html) capabilities.

## Changelog

See [changelog](./CHANGELOG.md).

## Contributing

See [contributing](./CONTRIBUTING.md).

## Issues

You found a bug or need a new feature? Please <a href="https://github.com/immobiliare/peephole/issues/new" target="_blank">open an issue.</a>
