package exporter

type Exporter interface {
	export() error
}
