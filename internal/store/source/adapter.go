package source

type SourceConnector[T any] interface {
	Connect() (*T, error)
}
