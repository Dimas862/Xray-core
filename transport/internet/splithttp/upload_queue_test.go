package splithttp_test

import (
	"testing"

	"github.com/Dimas862/xray-core/common"
	. "github.com/Dimas862/xray-core/transport/internet/splithttp"
)

func Test_regression_readzero(t *testing.T) {
	q := NewUploadQueue(10)
	q.Push(Packet{
		Payload: []byte("x"),
		Seq:     0,
	})
	buf := make([]byte, 20)
	n, err := q.Read(buf)
	common.Must(err)
	if n != 1 {
		t.Error("n=", n)
	}
}

