# middleware 中间件
中间件技术将业务和非业务代码功能进行解耦。<br>
通过一个或多个函数对 handler 进行包装，返回一个包括了各个中间件逻辑的函数链。<br>

**比如开源界很火的 gin 这个框架，就专门为用户贡献的 middleware 开了一个仓库：<br>**
https://github.com/gin-gonic/contrib


