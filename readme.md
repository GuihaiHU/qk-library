# qk-library

## 简介
qk-library是一个Golang后端工具库，包含以下几部分：
1. qk-cli 脚手架，可以用来生成Restful相关代码的封装代码，也可以用于和apifox配合，根据apifox里定义好的接口生成代码
2. frame 代码层面的一些工具，以更便捷地写业务代码
3. notify 通知模块，主要用于发送一些报警信息等通知到飞书
4. qfile 主要用于处理文件，比如文件的存储
5. qutil 一些工具函数

   
## qk-cli
### 安装
   wget https://github.com/iWinston/gf-cli/raw/dev/qk && chmod +x qk && ./qk install

### apifox-sync 根据apifox文件同步生成代码
   1. 首先从apifox中导出apifox.json文件
   2. 命令后运行 qk sync <TYPE> ./apifox.json，TYPE的类型支持all，api(包含了service)，define，model，一般情况下，用all即可
   3. 命令运行后，会生成以下结构的文件，以apifox里包含一个Admin管理后台的Cat接口为例子

    ```
    project
    │   README.md
    │   apifox.json   
    │
    └───app
        │
        └───model
        │   │   cat.model.go
        │
        └───system
            │   
            └───admin
                │
                └───define
                │   │   cat.define.go
                │
                └───api
                │   │  
                │   └───internal
                │   │   │   cat.api.go
                │   │   
                │   │   cat.api.go    
                │
                └───service
                │   │  
                │   └───internal
                │   │   │   cat.service.go
                │   │   
                │   │   cat.service.go
                │
                ...
    ...
    ```

    4. 其中api和service都有两个文件，内层文件是不可修改的，每次都由qk-cli生成并覆盖，支持在外层文件中进行重写覆盖实现
    5. define文件只有一个，不支持修改，设计对字段进行修改的时候，请通过apifox修改并重新生成

### apifox-gen 生成apifox格式的json文件快速定义接口
   - 以在Admin管理后台的Cat接口以例子，执行qk gen apifox cat 猫 -s Admin，-s用于表名生成到哪个系统，统一用大驼峰法
   - 执行之后会在apifox文件夹下生成一个json文件，里面是cat的Restful相关代码的封装接口，包括api，service，model和define
   - 打开apifox软件，导入即可
  
## frame

### q
q里主要封装了每个层的通用代码
1. API层：
   - q.AssignParamFormReq(r *ghttp.Request, param interface{}), 这个方法会对param进行赋值，首先会优先调用r.Parse(）方法，从param, query，body里获取参数。之后会根据param结构体的tag标签，从ghttp.Request中获取值。目前支持的是ctx标签，例如通过添加"ctx:User"标签，可以从上下文ctx中获取User的值。ctx标签也支持"ctx.User.Name"的方式获取单个属性。
   - q.Response
    支持q.Response，q.ResponseWithData，q.ResponseWithMeta三种形式
2. Service层：
   - q.Get, q.Post, q.Patch, q.Delete 和 q.List是Restful相关代码的封装, 这些方法会根据传入的param和res生成where和select语句，除了在生成的代码中可以使用外，在自定义的service中也可以使用。
3. Model层：
   - GenSqlByParam(sql *gorm.DB, param interface{}) *gorm.DB，这个方法会根据传入的param的Tag来生成Gorm的Join和Where语句，具体语法和规则如下：
    ```
    type Cat struct {
        Name *string // 没加标签的，默认不添加为条件
        Age *uint `where:">"` // where后面可以加个比较符号，比如>,<,=,like等，如此时Age的值是2，那么产生的sql是"Where age > 2"
        Sex *string `where:"=;Gender"` // where后面可以用分号割开，加第二个参数，自定义字段名，如此时Sex的值是1，那么产生的sql是"Where gender = 1"
        ParentName *string `where:"=;Parent.Name"` // 如果自定义字段名是级联形式，会触发联表操作
    }
    ```
    - GenSqlByRes(sql *gorm.DB, param interface{}) *gorm.DB，这个方法会根据传入的res的Tag来生成Gorm的Join和Select语句，具体语法和规则如下：
    ```
    type Cat struct {
        Name *string // 没有添加标签的，默认select
        Age *uint `select:"_"` // 下划线代表跳过不select，这种情况下，该字段会是零值
        ParentName *string `where:"Parent.Name"` // 这种情况会触发联表操作，且会筛选对应的值
        Parent Parent `select:"Parent"` // 暂不支持Preload，遇到这种情况，可以设置select为_,再多一次查询，或者自行Preload
    }
    ```
