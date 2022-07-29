package pkg

type Auth interface {
	HasReadAccess(apiKey string) bool
	HasWriteAccess(apiKey string) bool
}