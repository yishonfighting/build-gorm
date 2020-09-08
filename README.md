## model builder 一键生成model文件

### 背景：
为方便大家的使用，减少重复model定义，并提高大家的工作效率。
因此大家请使用此工具来进行model文件等等的目录与代码的生成。

### 使用示例：

执行以下命令：
> go run main.go 
>--name 服务目录名称（生成的model所属服务，比如social-micro） 
>--path sql文件目录(相对路径即可) 
>--type 数据表类型@1：常规@2：带后缀
