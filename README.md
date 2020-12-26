# 多环境测试工具

## 配置文件示例

```yaml
# 环境列表
envirList:
  - name: qa                        #环境名称
    baseUri: 'http://192.168.1.10'  #环境baseUri，支持ip:port、域名
    host: roach.360tryst.com        #host

  - name: gray
    baseUri: 'http://404.360tryst.com:80'
    host: roach.360tryst.com

  - name: product
    baseUri: 'http://404.360tryst.com'
    host: product.360tryst.com

#接口列表
apiList:
  - name: 获取产品列表               #接口名称
    method: get                    #请求方法，默认为get
    path: /test.php                #请求路径
    params: c=product&a=list       #请求参数
    timeout: 3                     #超时时间，默认10秒
    headers:                       #请求头设置
      Content-Type: text/plain
      User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36

  - name: 添加商品
    method: POST
    path: /test.php
    params: c=product&a=create&name=iPhone12&price=6799.00
    timeout: 3
    headers:
      Content-Type: application/x-www-form-urlencoded
      User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36

  - name: 获取商品详情
    path: /test.php
    params: id=15
    headers:
      User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36

  - name: 更新商品
    method: PUT
    path: /test.php?id=15
    params: c=product&a=update&price=6299.00
    timeout: 3
    headers:
      Content-Type: text/plain
      User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36

  - name: 删除商品
    method: DELETE
    path: /test.php?id=15
    params:
    timeout: 3
    headers:
      User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36
```

## 测试结果样例

```yaml
list:
# 接口详情
- api: 获取产品列表 [GET] /test.php params:c=product&a=list
  isSame: false   #多个环境响应body是否一致
  envirList:
  #qa环境响应结果详情
  - envir: qa baseUrl:http://192.168.1.10 host:roach.360tryst.com
    httpCode: -1   # -1代表请求失败
    body: ""       #响应body

  #gray环境响应结果详情
  - envir: gray baseUrl:http://404.360tryst.com:80 host:roach.360tryst.com
    httpCode: 200
    body: '{"method":"GET","queryArgs":{"c":"product","a":"list"},"formArgs":[],"requestBody":""}'
  
  #product环境响应结果详情
  - envir: product baseUrl:http://404.360tryst.com host:product.360tryst.com
    httpCode: 200
    body: '{"method":"GET","queryArgs":{"c":"product","a":"list"},"formArgs":[],"requestBody":""}'

- api: 添加商品 [POST] /test.php params:c=product&a=create&name=iPhone12&price=6799.00
  isSame: false
  envirList:
  - envir: qa baseUrl:http://192.168.1.10 host:roach.360tryst.com
    httpCode: -1
    body: ""
  - envir: gray baseUrl:http://404.360tryst.com:80 host:roach.360tryst.com
    httpCode: 200
    body: '{"method":"POST","queryArgs":[],"formArgs":{"c":"product","a":"create","name":"iPhone12","price":"6799.00"},"requestBody":"c=product\u0026a=create\u0026name=iPhone12\u0026price=6799.00"}'
  - envir: product baseUrl:http://404.360tryst.com host:product.360tryst.com
    httpCode: 200
    body: '{"method":"POST","queryArgs":[],"formArgs":{"c":"product","a":"create","name":"iPhone12","price":"6799.00"},"requestBody":"c=product\u0026a=create\u0026name=iPhone12\u0026price=6799.00"}'

- api: 获取商品详情 [GET] /test.php params:id=15
  isSame: false
  envirList:
  - envir: qa baseUrl:http://192.168.1.10 host:roach.360tryst.com
    httpCode: -1
    body: ""
  - envir: gray baseUrl:http://404.360tryst.com:80 host:roach.360tryst.com
    httpCode: 200
    body: '{"method":"GET","queryArgs":{"id":"15"},"formArgs":[],"requestBody":""}'
  - envir: product baseUrl:http://404.360tryst.com host:product.360tryst.com
    httpCode: 200
    body: '{"method":"GET","queryArgs":{"id":"15"},"formArgs":[],"requestBody":""}'

- api: 更新商品 [PUT] /test.php?id=15 params:c=product&a=update&price=6299.00
  isSame: false
  envirList:
  - envir: qa baseUrl:http://192.168.1.10 host:roach.360tryst.com
    httpCode: -1
    body: ""
  - envir: gray baseUrl:http://404.360tryst.com:80 host:roach.360tryst.com
    httpCode: 200
    body: '{"method":"PUT","queryArgs":{"id":"15"},"formArgs":[],"requestBody":"c=product\u0026a=update\u0026price=6299.00"}'
  - envir: product baseUrl:http://404.360tryst.com host:product.360tryst.com
    httpCode: 200
    body: '{"method":"PUT","queryArgs":{"id":"15"},"formArgs":[],"requestBody":"c=product\u0026a=update\u0026price=6299.00"}'

- api: '删除商品 [DELETE] /test.php?id=15 params:'
  isSame: false
  envirList:
  - envir: qa baseUrl:http://192.168.1.10 host:roach.360tryst.com
    httpCode: -1
    body: ""
  - envir: gray baseUrl:http://404.360tryst.com:80 host:roach.360tryst.com
    httpCode: 200
    body: '{"method":"DELETE","queryArgs":{"id":"15"},"formArgs":[],"requestBody":""}'
  - envir: product baseUrl:http://404.360tryst.com host:product.360tryst.com
    httpCode: 200
    body: '{"method":"DELETE","queryArgs":{"id":"15"},"formArgs":[],"requestBody":""}'
```

## 使用方法

> 1.修改cases.yml

> 2.执行测试

```bash
#mac操作系统执行
./mac

#linux操作系统执行
./linux

#windows操作系统执行
./test.exe

#指定配置文件和测试结果输出文件
./mac -c ./roach.360tryst.com.yml -o ./roach-result.yml
```



