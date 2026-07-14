# tgbotkit client

[![Go Report Card](https://goreportcard.com/badge/github.com/tgbotkit/client)](https://goreportcard.com/report/github.com/tgbotkit/client)
[![tests](https://github.com/tgbotkit/client/actions/workflows/tests.yml/badge.svg)](https://github.com/tgbotkit/client/actions/workflows/tests.yml)
[![version](https://img.shields.io/github/v/release/tgbotkit/client?sort=semver)](https://github.com/tgbotkit/client/releases)
[![license](https://img.shields.io/github/license/tgbotkit/client)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/tgbotkit/client.svg)](https://pkg.go.dev/github.com/tgbotkit/client)

Go client for the Telegram Bot API, generated from the official spec. This project follows an **API-first** approach, using an OpenAPI specification to ensure consistency and correctness.

Package docs: https://pkg.go.dev/github.com/tgbotkit/client

## Supported Telegram Bot API version

| Package Version | API Version | API Release Date |
|-----------------| --- | --- |
| latest, 0.9.0    | [10.2](https://core.telegram.org/bots/api#july-14-2026) | July 14, 2026 |
| 0.8.0            | [10.1](https://core.telegram.org/bots/api#june-11-2026) | June 11, 2026 |
| 0.7.0            | [10.0](https://core.telegram.org/bots/api#may-8-2026) | May 8, 2026 |
| 0.6.0            | [9.6](https://core.telegram.org/bots/api#april-3-2026) | April 3, 2026 |
| 0.5.0, 0.5.1    | [9.5](https://core.telegram.org/bots/api#march-1-2026) | March 1, 2026 |
| <= 0.3.2        | [9.3](https://core.telegram.org/bots/api#december-31-2025) | December 31, 2025 |
| <= 0.2.0        | [9.2](https://core.telegram.org/bots/api#august-15-2025) | August 15, 2025 |

## Tooling

This client is generated using:

- https://github.com/metalagman/tgbotspec to fetch/build the OpenAPI spec
- https://github.com/oapi-codegen/oapi-codegen to generate the Go client

Generation is wired via `go generate` in `generate.go` and uses `cfg.yaml`.

## Usage

See `example_test.go` for a minimal usage example.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
