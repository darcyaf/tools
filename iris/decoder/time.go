package decoder

import (
	"reflect"
	"strconv"
	"time"
	_ "unsafe"

	"github.com/iris-contrib/schema"
	_ "github.com/iris-contrib/schema"
)

//go:linkname formDecoder github.com/iris-contrib/schema.formDecoder
//go:nosplit
var formDecoder *schema.Decoder

type Time time.Time

func (n Time) MarshalJSON() ([]byte, error) {
	s := time.Time(n).Format("2006-01-02 15:04:05")
	s = strconv.Quote(s)
	return []byte(s), nil
}
func (n Time) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("2006-01-02 15:04:05", string(b))
	n = Time(t)
	return err
}
func (n Time) Converter(s string) reflect.Value {
	if t, err := time.Parse("2006-01-02 15:04:05", s); err == nil {
		return reflect.ValueOf(Time(t))
	}
	return reflect.Value{}
}
func RegisterNormalTimeDecoder() {
	var nt = Time{}
	formDecoder.RegisterConverter(nt, nt.Converter)
}
