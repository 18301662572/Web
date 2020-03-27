# Database 和数据库打交道

官方 database/sql 库:http://go-database-sql.org/<br>

1.database/sql 官方库
```text
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
```

2.提高生产效率的 ORM 和 SQL Builder
```text
对象关系映射（英语：Object Relational Mapping，简称ORM，或O/RM，或O/R mapping），
是一种程序设计技术，用于实现面向对象编程语言里不同类型系统的数据之间的转换。
从效果上说，它其实是创建了一个可在编程语言里使用的“虚拟对象数据库”。
最为常见的 ORM 实际上做的是从 db -> 程序的 class / struct 这样的映射。
```

3.ORM 与 SQL Builder
```text
当然，我们不能否认 ORM 的进步意义，它的设计初衷就是为了让数据的操作和存储的具体实现所剥离。
但是在上了规模的公司的人们渐渐达成了一个共识，由于隐藏重要的细节，ORM 可能是失败的设计。
其所隐藏的重要细节对于上了规模的系统开发来说至关重要。
相比 ORM 来说，sql builder 在 sql 和项目可维护性之间取得了比较好的平衡。
首先 sql builer 不像 ORM 那样屏蔽了过多的细节，其次从开发的角度来讲，
sql builder 简单进行封装后也可以非常高效地完成开发

说白了 sql builder 是 sql 在代码里的一种特殊方言，
如果你们没有DBA但研发有自己分析和优化 sql 的能力，或者你们公司的 dba 对于学习这样一些 sql 的方言没有异议。
那么使用 sql builder 是一个比较好的选择，不会导致什么问题。

另外在一些本来也不需要DBA介入的场景内，使用 sql builder 也是可以的，
例如你要做一套运维系统，且将 mysql 当作了系统中的一个组件，系统的 QPS 不高，查询不复杂等等。

一旦你做的是高并发的 OLTP 在线系统，且想在人员充足分工明确的前提下最大程度控制系统的风险，
使用 sql builder 就不合适了。
```
#### orm 类似于 EF+ Lambda表达式
```text
o := orm.NewOrm()
num, err := o.QueryTable("cardgroup").Filter("Cards__Card__Name", cardName).All(&cardgroups)
```

#### 操作sql builder 类似 linq to sql 
```text
where := map[string]interface{} {
    "order_id > ?" : 0,
    "customer_id != ?" : 0,
}
limit := []int{0,100}
orderBy := []string{"id asc", "create_time desc"}
orders := orderModel.GetList(where, limit, orderBy)
```
