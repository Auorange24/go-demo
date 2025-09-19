# go-reflect


## 变量的内在机制

## 反射

## reflect包

**如何理解Go运行中的值**

Go中值在运行的时候都会被编译成`reflect.Type`和`reflect.Value`结构体。

所谓的反射就是获取某个值的这两个对应的结构体。

### Typeof

**Typeo函数结构**

- 传入一个值，得到这个值对应的`Type`结构体。

```go
func TypeOf(i any) Type {
	return toType(abi.TypeOf(i))
}
```


# Valuof

# StructFIeld