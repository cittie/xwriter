## 概述
- xwriter是一个中间件，用于将数据生成为xlsData文件，或将xlsData写入xlsx/xlsm文件
- 选用序列化/反序列化方式：
	- protobuf
		- 跨多种语言，尤其是前端JavaScript和C#，方便之后做可视化界面
		- 同时支持生成和读取json
		- 方便后续使用grpc拓展协议
	- json
		- 可以由protobuf生成
		- 可读性好，便于修改纠错
- package分为3部分，util共用
    - protocol xlsData结构和grpc协议定义
        - 使用proto文件定义数据结构
        - 通过bash文件生成pb代码
    - writer 将xlsData写入xls文件
        - 需要使用命令行方式，可以通过main编译后独立运行
    - parser 生成和保存xlsData
        - 提供了默认的结构XlsDataParser和XlsDataCell
        - 同时也可以通过接口IXlsDataParser和IXlsDataCell实现同样功能
        
## 可执行文件
- 纯命令行方式，参数如下：
    - data (d) 数据文件所在目录
	- src (s) xls文件所在目录
	- tar (t) 写入数据后xls文件存储目录
 	- ext (e) 数据文件的类型，json或bin
- 示例 ```xwrite -d data -s a3k -t xls -e json```
- 按名称读取数据
- 自动删除Sheet末尾的空行

## 数据
- 写入实现
	- 匹配id和对应的序号，写入内容
	- 譬如 ```"id": "TestID", "id_idx": 1,"content": "TestContent"``` 会搜索第2个ID为"TestID"的列，并将"TestContent"写入对应单元格
- 新文件会保存到指定目录，不会影响原文件
