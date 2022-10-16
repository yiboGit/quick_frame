/**
* @Author: yibo_LastDay
* @Date: 2022/10/16 11:40
 */

package fmtsize

import "testing"

func TestByteSize_String(t *testing.T) {
	t.Log(1024, ByteSize(1024*1024))
}
