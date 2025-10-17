package main
var p1 *int
var p2 *string

i := 1
s := "Hello"
// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
p1 = &i
p2 = &s

p3 := &p2
fmt.Println(p1)
fmt.Println(p2)
fmt.Println(p3)

//使用指针访问值
var p1 *int
i := 1
p1 = &i
fmt.Println(*p1 == i)
*p1 = 2
fmt.Println(i)

a := 2
var p *int
fmt.Println(&a)
p = &a
fmt.Println(p, &a)

// 修改指针指向的值
var pp **int
pp = &p
fmt.Println(pp, p)
**pp = 3
fmt.Println(pp, *pp, p)
fmt.Println(**pp, *p)
fmt.Println(a, &a)

//指针、unsafe.Pointer 和 uintptr
a := "Hello, world!"
upA := uintptr(unsafe.Pointer(&a))
upA += 1

c := (*uint8)(unsafe.Pointer(upA))
fmt.Println(*c)