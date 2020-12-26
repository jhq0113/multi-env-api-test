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


