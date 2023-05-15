# ptr
放些指针运算的辅助函数

# 获取这个类型的一级指针并赋值
```go
p := ptr.New(1)
p2 := ptr.New(1.12)
```
# 获取这个类型的二级指针并赋值
```go
p := ptr.New2(1)
p2 := ptr.New2(1.12)
```
# 获取这个类型的三级指针并赋值
```go
p := ptr.New3(1)
p2 := ptr.New3(1.12)
```

# 指针相加
```go
type test struct {
    A int
    B int
}

var t test
pb := ptr.Add(&t, 8)
*pb = 3
fmt.Printf("%d\n", t.B)
```

# 指针相加
```go
s := []int{1, 2, 3}
p3 := ptr.AddOffset(&s[0], 8, 2)
fmt.Printf("%d\n", *p3)

```
# 指针相减
```go
type test struct {
    A int
    B int
}

var t test
t.A = 3
pa := ptr.Sub(&t.B, 8)
fmt.Printf("%d\n", *pa)
```
