package config_test

import (
	"bytes"
	"reflect"
)

// --- HELPER FUNCTIONS
func isEmptyStruct(s interface{}) bool {
	typ := reflect.TypeOf(s)
	if typ.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		zeroValue := reflect.Zero(field.Type)
		fieldValue := reflect.ValueOf(s).Field(i)
		if !reflect.DeepEqual(fieldValue.Interface(), zeroValue.Interface()) {
			return false
		}
	}

	return true
}

// --- HELPER STRUCTS
type FileBuffer struct {
	name   string
	buffer *bytes.Buffer
}

func NewFileBuffer(name string) *FileBuffer {
	return &FileBuffer{
		name:   name,
		buffer: &bytes.Buffer{},
	}
}

func (m *FileBuffer) Read(p []byte) (int, error) {
	return m.buffer.Read(p)
}

func (m *FileBuffer) Write(p []byte) (int, error) {
	return m.buffer.Write(p)
}

func (m *FileBuffer) WriteString(s string) (int, error) {
	return m.buffer.WriteString(s)
}

func (m *FileBuffer) Name() string {
	return m.name
}

func (m *FileBuffer) Close() error {
	return nil
}
