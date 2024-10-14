# clashmerge

用于将多个的 clash 配置文件合并成一个。也可以实现在某个 clash 配置的基础上进行修改。



### 配置文件格式

```yaml
- operation: merge
  type: url
  url: https://example.com/a.yaml # type 为 url 时生效

- operation: merge # 合并方式，支持 merge 和 replace
  type: data # 数据源类型
  data: # 数据内容，当 type 是 data 时这个字段有效
    num: 2
    list: ["dave"]
    list2: ["eric"]
    obj:
      age: 21
      color: red


```

### 合并逻辑

程序会按配置文件依次处理每项配置。

1. 首先根据 type 和对应的内容值获取配置内容

2. 然后根据 operation 选择使用 merge 方式或者 replace 方式将当前内容与前一项得到的配置内容合并处理。

**merge 方式**

merge 方式下如果 yaml 文件一级键的值是数组，合并时会将新的内容追加到数组后。

如果一级键的值是字典，则新值会追加，旧值会覆盖。如果是其它简单的类型，比如字符串，新值会覆盖旧值。

```yaml
- operation: merge 
  type: data 
  data: 
    num: 2
    list: ["dave"]
    list2: ["eric"]
    obj:
      name: ace
      color: red
- operation: merge
  type: data
  data: 
    num: 3
    new_num: 4
    list: ["dave", "eric"]
    obj:
      name: bob
      color: blue
```

会得到

```yaml
num: 3
new_num: 4
list: ["dave", "eric"]
list2: ["eric"]
obj:
  name: bob
  color: blue
```

**replace 方式**

相同键的值会被覆盖，可以用这种方式来实现删除。

比如

```yaml
- operation: merge 
  type: data 
  data: 
    num: 2
    list: ["dave"]
    list2: ["eric"]
    obj:
      name: ace
      color: red
- operation: replace
  type: data
  data: 
    num: 3
    list: ["dave"]
    obj: {}
```

会得到

```yaml
num: 3
list: ["dave"]
list2: ["eric"]
obj: {}
```

### 部署

docker-compose

```yaml
version: "3"
services:
  clashmerge:
    image: mailth/clashmerge:latest
    network_mode: host
    environment:
      PORT: 8081 # 服务端口
      LOG_LEVEL: error # debug, info, warn, error.
    volumes:
      - "./_data:/data"
```

在 _data 目录里创建 config1.yaml 按上面的配置格式填好，然后访问 https://your-domain?name=config1.yaml 即可以得到合并后的配置文件。
