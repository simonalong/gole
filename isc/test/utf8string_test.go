package test

import (
	"testing"

	. "github.com/simonalong/gole/isc"
)

func TestUTF8String(t *testing.T) {
	var s = ISCUTF8String("UTF8字符串")

	t.Logf("len = %d\n", s.Length())
	idx := s.IndexOf(ISCUTF8String("集")) // 2
	t.Logf("indexOf(\"集\") = %d\n", idx)

	ss := s.Insert(3, ISCUTF8String("xyz"))
	t.Logf("%v\n", ss) // xyzUTF8字符串

	sss := ss.Delete(3, 3)
	t.Logf("%v\n", sss) // UTF8字符串

}
