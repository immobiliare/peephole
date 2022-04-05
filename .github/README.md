# Peephole <a href="#peephole"><img align="left" width="100px" src="https://github.com/immobiliare/peephole/blob/master/kiosk/assets/static/peephole.png"></a>

[![pipeline status](https://github.com/immobiliare/peephole/actions/workflows/test.yml/badge.svg)](https://github.com/immobiliare/peephole/actions/workflows/test.yml)

Peephole is a web-based events explorer of [SaltStack](https://github.com/saltstack/salt) transactions.

It can be used to watch for events on several [SaltStack](https://github.com/saltstack/salt) master nodes by GETting _from the_ `/events` API endpoint.
As soon as events are watched, they get persisted into a local (auto-expiring and configurable) cache that allows to explore, filter and correlate events for a variable time interval.

[![Peephole homepage](https://github.com/immobiliare/peephole/blob/master/.github/sample-1.png)](#peephole)

[![Peephole detail](https://github.com/immobiliare/peephole/blob/master/.github/sample-2.png)](#peephole)

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Changelog](#changelog)
- [Contributing](#contributing)
- [Issues](#issues)
- [Powering](#powering)

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

## Powering

What led to the use of Go as the engine for Peephole is mainly the flow of the app itself: fetching data externally from possibily several sources concurrently, to feed a cache with, as well as a pool of several clients.
Go's flexibility on intra-process channel-based communication between goroutines is a very solid paradigm to base the whole app flow on.
That is very crucial in the context for which it's been designed, which is huge.

In fact, Peephole is currently used internally in the SRE / systems teams at Immobiliare to keep an eye on and soften the burden of administering a pool of servers of the order of thousands, fetching events from tens of Salt masters.

**If you feel like using it as well, please, let us know!**
