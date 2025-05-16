package config

type FileInterface interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	WriteString(s string) (n int, err error)
	Name() string
	Close() error
}
