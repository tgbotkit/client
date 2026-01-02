package client

//go:generate go tool github.com/metalagman/tgbotspec/cmd/tgbotspec --merge-union-types -o api/openapi.yaml
//go:generate go tool oapi-codegen -config cfg.yaml api/openapi.yaml
