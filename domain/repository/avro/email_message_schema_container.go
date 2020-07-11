// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     email-message.json
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewEmailMessageSchemaWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewEmailMessageSchema()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type EmailMessageSchemaReader struct {
	r io.Reader
	p *vm.Program
}

func NewEmailMessageSchemaReader(r io.Reader) (*EmailMessageSchemaReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewEmailMessageSchema()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &EmailMessageSchemaReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r EmailMessageSchemaReader) Read() (*EmailMessageSchema, error) {
	t := NewEmailMessageSchema()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
