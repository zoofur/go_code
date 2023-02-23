```go
func main() {
 sl := make([]int, 0, 10)
 var appendFunc = func(s []int) {
  s = append(s, 10, 20, 30)
  fmt.Println(s)
 }
 fmt.Println(sl)
 appendFunc(sl)
 fmt.Println(sl)
 fmt.Println(sl[:10])
}
 
// 结果
[]
[10 20 30]
[]
[10 20 30 0 0 0 0 0 0 0] // 为什么输出不一样？
```
1、为什么两次输出s1的内容不一样？

Slice 在经过 makeSlice的转换后为三部分，指向底层数组的指针array、cap、len。

Go 中只有值传递，向 appendFunc 传递的 Slice 虽然为副本但是其 array 指针指向同一个底层数组，对 Slice 进行修改后，其底层数组也进行了修改，但是 s 和 s1 仍为两个不同的 SliceHeader，其 len 的长度并没有被修改，s1 len 为0，s len 为3，而 Slice 的输出内容的数量又是基于 len 的，所以加入下标后才会看到输出的值不一样。
```go
func main() {
    s1 := make([]int, 0, 4)
    fmt.Printf("%p\n", s1)
    var appendFunc = func(s []int) {
        s = append(s, 1, 2, 3)
        fmt.Printf("%p\n", s)
        s = append(s, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
        fmt.Println(s)
        fmt.Printf("%p\n", s)
    }
    fmt.Println(s1)
    appendFunc(s1)
    fmt.Println(s1)
    fmt.Println(s1[:4])
}
 
// 结果
0xc000010200
[]                                    
0xc000010200                          
[1 2 3 10 20 30 40 50 60 70 80 90 100]
0xc0000120e0                          
[]                                    
[1 2 3 0]   
```
2、传入的 slice 在扩容前的地址值一样，扩容后的地址值不一样？

fmt 包（fmt/print.go）对 slice 的输出做了优化，打印 slice 的地址时会直接打印其 array 指针，扩容后的 slice 是新的底层数组地址，所以就不一样了。

```go
func (p *pp) fmtPointer(value reflect.Value, verb rune) {
    var u uintptr
    switch value.Kind() {
    case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
        u = value.Pointer()
    default:
        p.badVerb(verb)
        return
    }
    ...
}
```
3、想要s1也被同步修改生效，要怎么办呢？

想要s1也被同步修改生效，需要传入 Slice 的指针。如果传入的是指针，由下面代码输出结果可以看到 slice 扩容前的地址已经就不一样了，因为“值传递”的存在，其生成了一个指针的副本。
```go
func main() {
    s1 := make([]int, 0, 4)
    fmt.Printf("%p\n", s1)
    var appendFunc = func(s *[]int) {
        *s = append(*s, 1, 2, 3)
        fmt.Printf("%p\n", s)
        *s = append(*s, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
        fmt.Println(s)
        fmt.Printf("%p\n", s)
    }
    fmt.Println(s1)
    appendFunc(&s1)
    fmt.Println(s1)
    fmt.Println(s1[:4])
}
 
// 结果
0xc000010200
[]
0xc000004060
&[1 2 3 10 20 30 40 50 60 70 80 90 100]
0xc000004060
[1 2 3 10 20 30 40 50 60 70 80 90 100] 
[1 2 3 10]
```