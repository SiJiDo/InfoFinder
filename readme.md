
### step 1
katana 爬取页面

### step 2
正则匹配爬取到的页面内容，如果有url的正则则放入url列表变量中，正则信息可以自定义到config.yaml文件中
正则可以是邮箱，secretkey，secretid，password之类的

### step 3
从url列表变量中提取单个路径比如/api /a 这类的路径，进行与其他路径拼接，比如有个 /customer/info, 拼接后就是/api/customer/info
然后组成新的列表字典
```
/api
/a
/customer/info
/api/customer/info
/a/customer/info
```


