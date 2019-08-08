// Exercise 11.7
// Analogous as in § 11.7 and Listings 11.3/4 define a package float64 and in it a type Float64Array,
// and let this type implement the Sorter interface to sort an array of float64.
// Also provide the following methods:
// - a NewFloat64Array-method for making a variable with 25 elements (see § 10.2)
// - a List()-method returning a string for pretty-printing the array, wrap this in a
// String()-method so that you don’t have to call List() explicitly (see § 10.7)
// - a Fill() method for creating such an array with 10 random floats (see § 4.5.2.6)
// In the main-program create a variable of that type, sort it and test that.

// 类似11.7和示例11.3/4，定义一个包 float64，并在包里定义类型 Float64Array，然后让它实现 Sorter 接口用来对float64 数组进行排序。
// 另外提供如下方法：
// NewFloat64Array()：创建一个包含25个元素的数组变量（参考10.2）
// List()：返回数组格式化后的字符串，并在 String() 方法中调用它，这样就不用显式地调用 List() 来打印数组（参考10.7）
// Fill()：创建一个包含10个随机浮点数的数组（参考4.5.2.6）
// 在主程序中新建一个此类型的变量，然后对它排序并进行测试。

package float64

import (
    "fmt"
    "math/rand"
    "time"
)

type Sorter interface {
    Len()	int
    Less(i, j int)	bool
    Swap(i, j int)
}

func Sort(data Sorter) {
    for pass := 1; pass < data.Len(); pass++ {
        for i := 0; i < data.Len() - pass; i++ {
            if data.Less(i+1, i) {
                data.Swap(i, i+1)
            }
        }
    }
}

func IsSorted(data Sorter) bool {
    n := data.Len()
    for i := n-1; i > 0; i-- {
        if data.Less(i, i-1) {
            return false
        }
    }
    return true
}

type Float64Array []float64

func (p Float64Array) Len() int { return len(p) }
func (p Float64Array) Less(i, j int) bool { return p[i] < p[j] }
func (p Float64Array) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func NewFloat64Array() Float64Array {
    return make([]float64, 25)
}

func (p Float64Array) Fill(n int) {
    rand.Seed(int64(time.Now().Nanosecond()))
    for i := 0; i < n; i++ {
        p[i] = 100 * (rand.Float64())
    }
}

func (p Float64Array) List() string {
    s := "{ "
    for i := 0; i < p.Len(); i++ {
        if p[i] == 0 {
            continue
        }
        s += fmt.Sprintf("%3.1f ", p[i])
    }
    s += " }"
    return s
}

func (p Float64Array) String() string {
    return p.List()
}
