package exporter

type Exporter interface {
	Export() error
}
