// Author By ChaiHanLin. Date 2019-07-15.
// Contact me <ChaiHanLin@blued.com>
package fmtsize

import "testing"

func TestByteSize_String(t *testing.T) {
	t.Log(1024, ByteSize(1024*1024))
}
