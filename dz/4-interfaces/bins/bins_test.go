package bins_test

import (
	"reflect"
	"struct/bins"
	"testing"
)

func TestCreateBin(t *testing.T) {
	binReq := `{
      "id": "68121cb08960c979a5905b77",
      "name": "check",
      "private": true,
      "createdAt": "2025-04-30T12:50:56.264Z"
    }`

	bin := bins.CreateBin(binReq)
	typ := reflect.TypeOf(bin)

	if typ.Elem().Name() != "Bin" {
		t.Errorf("bin не указывает на тип Bin")
	}
}
