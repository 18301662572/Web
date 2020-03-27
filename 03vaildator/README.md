# vaildator请求校验
 validator 库：https://github.com/go-playground/validator<br>
 
 基于反射的 validator <br>
 从结构上来看，每一个 struct 都可以看成是一棵树,对这棵 struct 树来进行遍历<br>
 如果基于反射的 validator 真的成为了你服务的性能瓶颈怎么办？现在也有一种思路可以避免反射：<br>
 使用 golang 内置的 parser 对源代码进行扫描，然后根据 struct 的定义生成校验代码。<br>
 我们可以将所有需要校验的结构体放在单独的 package 内。<br>