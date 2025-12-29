# tgbotkit client

[![Go Report Card](https://goreportcard.com/badge/github.com/tgbotkit/client)](https://goreportcard.com/report/github.com/tgbotkit/client)
[![tests](https://github.com/tgbotkit/client/actions/workflows/tests.yml/badge.svg)](https://github.com/tgbotkit/client/actions/workflows/tests.yml)
[![version](https://img.shields.io/github/v/release/tgbotkit/client?sort=semver)](https://github.com/tgbotkit/client/releases)
[![license](https://img.shields.io/github/license/tgbotkit/client)](LICENSE)

Go client for the Telegram Bot API, generated from the official spec. This project follows an **API-first** approach, using an OpenAPI specification to ensure consistency and correctness.

Package docs: https://pkg.go.dev/github.com/tgbotkit/client

## Supported Telegram Bot API version

| Package Version | API Version | API Release Date |
|-----------------| --- | --- |
| latest          | [9.2](https://core.telegram.org/bots/api#august-15-2025) | August 15, 2025 |

## Tooling

This client is generated using:

- https://github.com/metalagman/tgbotspec to fetch/build the OpenAPI spec
- https://github.com/oapi-codegen/oapi-codegen to generate the Go client

Generation is wired via `go generate` in `generate.go` and uses `cfg.yaml`.

## Usage

See `example_test.go` for a minimal usage example.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
