## [0.5.1](https://github.com/immobiliare/peephole/compare/0.5.0...0.5.1) (2025-05-05)


### Bug Fixes

* **config:** replace unmaintained gopkg.in/yaml.v2 with github.com/goccy/go-yaml ([7aa755a](https://github.com/immobiliare/peephole/commit/7aa755a438f38f36e9450a149ad8f726f145b83e))
* **linter:** remove deprecated linter configurations and update golangci-lint settings ([c969af9](https://github.com/immobiliare/peephole/commit/c969af9434b7a99f675f71a37266379fdfbc69b1))
* **mold:** replace xujiajun/nutsdb with nutsdb/nutsdb ([abbe115](https://github.com/immobiliare/peephole/commit/abbe115405be4d343e27e72555ed1b9565d2352e))
* **oci:** add missing provenance option to docker build-push action ([d079474](https://github.com/immobiliare/peephole/commit/d0794749eb7d811fe171127d67fee74c663b9c7b))
* **oci:** add missing sbom option to docker build-push action ([cd801e5](https://github.com/immobiliare/peephole/commit/cd801e535b94c0c24d2f4138e8ff874d44dcd436))
* **oci:** ignore hadolint for DL3007 ([0c35f87](https://github.com/immobiliare/peephole/commit/0c35f87bdfef6a03f85d7f08e46dd86b0cd96954))
* **oci:** update dockerfile to use chainguard distroless base images ([7a87acd](https://github.com/immobiliare/peephole/commit/7a87acd8b142955b0ff784c4ae96b51c6e5b6f71))

## [0.5.0](https://github.com/immobiliare/peephole/compare/0.4.7...0.5.0) (2024-01-22)


### Features

* **ci:** generate release from github action and update changelog ([2381a40](https://github.com/immobiliare/peephole/commit/2381a40abac1dfcc8c0b5f02efce7a94babab4a1))
* **ci:** introduce a conventional commits linter ([8f011fe](https://github.com/immobiliare/peephole/commit/8f011fe711d04a6b0e7392c6f748458b096a244f))
* **ci:** introduce dependabot ([ebab5a5](https://github.com/immobiliare/peephole/commit/ebab5a5eef7e85a0d177f0e5a1d34c19d8216814))


### Bug Fixes

* **ci:** drop broken scan step ([8582902](https://github.com/immobiliare/peephole/commit/858290248e5e55ac60b93429f4fd2e120f847217))
* **ci:** publish with a custom testing command ([16aaeaa](https://github.com/immobiliare/peephole/commit/16aaeaa3fdf1a47241382d17b1be56bc60b16b2e))
* **errors:** log when a deferred transaction rollback does not work ([7d51e71](https://github.com/immobiliare/peephole/commit/7d51e71bf3a2935cdc1bbc12a715789e0790df08))
* **kiosk:** replace packd with a new static handler ([ba11e11](https://github.com/immobiliare/peephole/commit/ba11e116f9274d8dc86690669f32b313e7a10c16))
* **kiosk:** replace packr with embed ([b57c964](https://github.com/immobiliare/peephole/commit/b57c96465206f461ae8cc0416aa1aef40cee93a5))
* **linters:** replace old linters with their updated counterparts ([ac0fa28](https://github.com/immobiliare/peephole/commit/ac0fa28d4660067f3829252627eb373a67f7a83c))
* **syntax:** ioutil has been replaced by io && os ([61f4f66](https://github.com/immobiliare/peephole/commit/61f4f6601926ba6ece8daa61dc8a19e5d13c2e18))

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.4.7] - 2021-11-09
### Changed
- Reduce container size by switching to Alpine
### Fixed
- Relogin on 401 on Master reconnection

## [0.4.6] - 2021-10-20
### Fixed
- Reconnecting to lost Salt master

## [0.4.5] - 2021-10-14
### Changed
- Slugify Mold event IDs

## [0.4.4] - 2021-10-14
### Changed
- Simplify spying retry mechanism

## [0.4.3] - 2021-10-12
### Changed
- Enforce Kiosk show button background to prevent text overlap

## [0.4.2] - 2021-10-12
### Added
- CI docker build
### Fixed
- Filter by success field

## [0.4.1] - 2021-10-11
### Fixed
- Minor Mold typos

## [0.4.0] - 2021-10-11
### Changed
- Drop Mold channel-based return logic
- Greatly improve query performance
- Use job timestamp as Mold timestamp
- Use Jid-inverse as record ID (to fix ordering by default)
### Fixed
- Count by passing filter
- Kiosk default page

## [0.3.2] - 2021-09-06
### Changed
- Interleave spying errors from retry calls

## [0.3.1] - 2021-09-03
### Fixed
- Docker build

## [0.3.0] - 2021-09-03
### Added
- Support for events list paging
- Show events outline only on events list
- Filter by job arg
- Wrap build with packr to serve static files
- Kiosk GZIP compression
- Kiosk dialog animation
- Kiosk jobs failures highlight
- Filter by status
- Kiosk responses minification (HTML/JS/CSS)
- Kiosk responses cache-control headers
### Changed
- Rewrite Kiosk datetime formatting
- Use dedicated event field for job args
- Bring everything by default on event outline
### Fixed
- Re-establish connection on Salt master disconnection
- DB order on select
- Disable Kiosk basic auth
- Corner case on adding events to Kiosk via EventSource
- Return 404 on event not found

## [0.2.4] - 2021-08-31
### Changed
- Use retcode on state parsing

## [0.2.3] - 2021-08-31
### Fixed
- Properly handle of connreset errors while spying

## [0.2.2] - 2021-08-30
### Fixed
- Events add when already filtering

## [0.2.1] - 2021-08-30
### Changed
- Debug environment key management

## [0.2.0] - 2021-08-30
### Added
- Support for debug option
### Fixed
- Docker container generation

## [0.1.0] - 2021-08-25
### Added
- First dump

[Unreleased]: https://github.com/immobiliare/peephole/compare/0.4.7...HEAD
[0.4.7]: https://github.com/immobiliare/peephole/releases/tag/0.4.7
[0.4.6]: https://github.com/immobiliare/peephole/releases/tag/0.4.6
[0.4.5]: https://github.com/immobiliare/peephole/releases/tag/0.4.5
[0.4.4]: https://github.com/immobiliare/peephole/releases/tag/0.4.4
[0.4.3]: https://github.com/immobiliare/peephole/releases/tag/0.4.3
[0.4.2]: https://github.com/immobiliare/peephole/releases/tag/0.4.2
[0.4.1]: https://github.com/immobiliare/peephole/releases/tag/0.4.1
[0.4.0]: https://github.com/immobiliare/peephole/releases/tag/0.4.0
[0.3.2]: https://github.com/immobiliare/peephole/releases/tag/0.3.2
[0.3.1]: https://github.com/immobiliare/peephole/releases/tag/0.3.1
[0.3.0]: https://github.com/immobiliare/peephole/releases/tag/0.3.0
[0.2.4]: https://github.com/immobiliare/peephole/releases/tag/0.2.4
[0.2.3]: https://github.com/immobiliare/peephole/releases/tag/0.2.3
[0.2.2]: https://github.com/immobiliare/peephole/releases/tag/0.2.2
[0.2.1]: https://github.com/immobiliare/peephole/releases/tag/0.2.1
[0.2.0]: https://github.com/immobiliare/peephole/releases/tag/0.2.0
[0.1.0]: https://github.com/immobiliare/peephole/releases/tag/0.1.0
