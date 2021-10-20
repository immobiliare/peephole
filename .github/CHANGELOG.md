# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
### Minor Mold typos

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

[Unreleased]: https://github.com/immobiliare/peephole/compare/0.4.6...HEAD
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
