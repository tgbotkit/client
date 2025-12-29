# tgbotkit client

Go client for the Telegram Bot API, generated from the official spec. This project follows an **API-first** approach, using an OpenAPI specification to ensure consistency and correctness.

Package docs: https://pkg.go.dev/github.com/tgbotkit/client

## Supported Telegram Bot API version

- [Telegram Bot API 9.2 (August 15, 2025)](https://core.telegram.org/bots/api "Telegram Bot API")

## Tooling

This client is generated using:

- https://github.com/metalagman/tgbotspec to fetch/build the OpenAPI spec
- https://github.com/oapi-codegen/oapi-codegen to generate the Go client

Generation is wired via `go generate` in `generate.go` and uses `cfg.yaml`.

## Usage

See `example_test.go` for a minimal usage example.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
