package ico

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jsummers/gobmp"
)

const testImage = "/vagrant_data/wikipedia.ico"

func TestDecodeAll(t *testing.T) {
	data, err := ioutil.ReadFile(testImage)
	if err != nil {
		t.Error(err)
	}
	r := bytes.NewReader(data)
	ic, err := DecodeAll(r)
	if err != nil {
		t.Error(err)
	}

	w, err := os.Create("/vagrant_data/wikipedia.bmp")
	if err != nil {
		t.Error(err)
	}

	options := new(gobmp.EncoderOptions)
	options.SupportTransparency(true)
	err = gobmp.EncodeWithOptions(w, ic.Image[0], options)
	if err != nil {
		t.Error(err)
	}
}
