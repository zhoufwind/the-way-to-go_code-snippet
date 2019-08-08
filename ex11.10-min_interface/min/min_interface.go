// Exercise 11.10
// Analogous to the Sorter interface we developed in § 11.7, make a Miner interface with the necessary
// operations, and a function Min which has as a parameter a variable which is a collection of type
// Miner and which calculates and returns the minimum element in that collection.

// 仿照11.7中开发的 Sorter 接口，创建一个 Miner 接口并实现一些必要的操作。函数 Min 接受一个 Miner 类
// 型变量的集合，然后计算并返回集合中最小的元素。

package min

type Miner interface {
    Len() int
    ElemIx(ix int) interface{}
    Less(i, j int) bool
    Swap(i, j int)
}

func Min(data Miner) interface{} {
    min := data.ElemIx(0)
    for i := 1; i < data.Len(); i++ {
        if data.Less(i, i-1) {
            min = data.ElemIx(i)
        } else {
            data.Swap(i, i-1)
        }
    }
    return min
}

type IntArray []int

func (p IntArray) Len() int	{ return len(p) }
func (p IntArray) ElemIx(ix int) interface{}	{ return p[ix] }
func (p IntArray) Less(i, j int) bool	{ return p[i] < p[j] }
func (p IntArray) Swap(i, j int)	{ p[i], p[j] = p[j], p[i] }

type StringArray []string

func (p StringArray) Len() int	{ return len(p) }
func (p StringArray) ElemIx(ix int) interface{}	{ return p[ix] }
func (p StringArray) Less(i, j int) bool	{ return p[i] < p[j] }
func (p StringArray) Swap(i, j int)	{ p[i], p[j] = p[j], p[i] }
