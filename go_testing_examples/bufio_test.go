package go_testing_examples

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
)

func TestBufio(t *testing.T) {
	var pkg Package
	pkg.Version[0] = 'V'
	pkg.Version[1] = 1
	pkg.Data = []byte("ABCDEFGHIJK")
	pkg.Datalen = int16(len(pkg.Data))
	fmt.Println(&pkg)

	var buf bytes.Buffer
	pkg.Pack(&buf)

	var pkg1 Package
	pkg1.Unpack(&buf)
	fmt.Println(&pkg1)
}

// 自定义协议的组包和拆包
type Package struct {
	Version [2]int8
	Datalen int16
	Data    []byte
}

func (p *Package) String() string {
	return fmt.Sprintf("Version:%d DataLen:%d Data:%s",
		p.Version, p.Datalen, p.Data)
}

func (p *Package) Pack(w io.Writer) {
	binary.Write(w, binary.BigEndian, p.Version)
	binary.Write(w, binary.BigEndian, p.Datalen)
	binary.Write(w, binary.BigEndian, p.Data)
}

func (p *Package) Unpack(r io.Reader) {
	binary.Read(r, binary.BigEndian, &p.Version)
	binary.Read(r, binary.BigEndian, &p.Datalen)
	if p.Datalen > 0 {
		p.Data = make([]byte, p.Datalen)
	}
	binary.Read(r, binary.BigEndian, &p.Data)
}
