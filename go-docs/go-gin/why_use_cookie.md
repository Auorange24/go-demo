# Cookie

## Cookie intro

**why need Cookie**

- HTTP is non-status, so every single request is independent.
- For Service, every request is new.

**Cookie is what**

- Cookie store little information.
- Cookie Created By Server and then send it to client.
- When client send next request, add cookie in http request header.


## Go use Cookie

**Cookie Struct**

- `Cookie struct` is declared in go library `net/http`. 

```go
type Cookie struct {
    Name       string // cookie name
    Value      string // cookie value
    Path       string // cookie url
    Domain     string // cookie domain
    Expires    time.Time // cookie expire time
    RawExpires string 
    // MaxAge=0表示未设置Max-Age属性
    // MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
    // MaxAge>0表示存在Max-Age属性，单位是秒
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // 未解析的“属性-值”对的原始文本
}
```
