# httprouter

### httprouter 通配符
```text
1.:号 (通配符)，可以是非空的，容易产生冲突,panic
    GET /user/info/:name
    POST /user/:id

2.* 号（catchAll）开头的参数（可以是非空的，容易产生冲突,panic）
    只能放在路由的结尾,主要是为了能够使用 httprouter 来做简单的 http 静态文件服务器
    Pattern: /src/*filepath
    /src/                     filepath = ""
    /src/somefile.go          filepath = "somefile.go"
    /src/subdir/somefile.go   filepath = "subdir/somefile.go"

3.httprouter 也支持对一些特殊情况下的回调函数进行定制
    r := httprouter.New()
    r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("oh no, not found"))
    })  
    r.PanicHandler = func(w http.ResponseWriter, r *http.Request, c interface{}) {
        log.Printf("Recovering from panic, Reason: %#v", c.(error))
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(c.(error).Error()))
    }
目前开源界最为流行的 web 框架 gin 使用的就是 httprouter 的变种。
```

### httprouter原理--radix tree，压缩字典树
```text
1.radix tree，压缩字典树;
    程序的局部性较好(一个节点的 path 加载到 cache 即可进行多个字符的对比)，
    从而对 CPU 缓存友好。
2.某个 method 第一次插入的路由就会导致对应字典树的根节点被创建
radix 的节点类型为 *httprouter.node
    1.path: 当前节点对应的路径中的字符串
    2.wildChild: 子节点是否为参数节点，即 wildcard node，或者说 :id 这种类型的节点
    3.nType: 当前节点类型，有四个枚举值: 分别为 static/root/param/catchAll。
        static                   // 非根节点的普通字符串节点
        root                     // 根节点
        param                    // 参数节点，例如 :id
        catchAll                 // 通配符节点，例如 *anyway
    4.indices: 子节点索引，当子节点为非参数类型，即本节点的 wildChild 为 false 时，
    会将每个子节点的首字母放在该索引数组。说是数组，实际上是个 string。        
3.root 节点创建-> 子节点插入->边分裂
4.子节点冲突处理
    在路由本身只有字符串的情况下，不会发生任何冲突。
    只有当路由中含有 wildcard(类似 :id) 或者 catchAll(*) 的情况下才可能冲突。
    原因： children 数组非空且 wildChild 被设置为 false ....                                                                                
```

