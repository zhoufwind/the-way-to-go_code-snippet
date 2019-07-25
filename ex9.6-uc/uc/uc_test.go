package uc

import "testing"

// 定义测试数据类型
type ucTest struct {
    in, out string
}

// 定义测试数据，以及测试结果
var ucTests = []ucTest {
    ucTest{"abc", "ABC"},
    ucTest{"cvo-az", "CVO-AZ"},
    ucTest{"Antwerp", "ANTWERP"},
}

// 定义测试方法，判断执行方法后是否与测试结果一致
func TestUC(t *testing.T) {
    for _, ut := range ucTests {
        uc := UpperCase(ut.in)
        if uc != ut.out {
            t.Errorf("UpperCase(%s) = %s, must be %s.", ut.in, uc, ut.out)
        }
    }
}

/*
$ go install github.com/zhoufwind/the-way-to-go_code-snippet/ex9.6-uc/uc
$ go test
PASS
ok  	github.com/zhoufwind/the-way-to-go_code-snippet/ex9.6-uc/uc	0.008s
*/
