# 接口文档

##1 登录

#### Method: ``` POST  ```
#### Path: ``` v1/login/```

> Argument

| 参数 | 类型 | 必选 | 说明 | 备注 |
| :--- | :--- | :--- |:---| :--- |
| username | string | true | 用户名|zhangsan |
| password | string | true | 密码|123123 |

> Response

- 成功

```json
{
    "code": 0,
    "msg": "",
    "data": {
        "user_id": 1,
        "user_name": zhangsan
    }  
}
```
##2 添加商品

#### Method: ``` POST  ```
#### Path: ``` v1/add/product```

> Argument

| 参数 | 类型 | 必选 | 说明 | 备注 |
| :--- | :--- | :--- |:---| :--- |
| user_id | string | true | 用户ID| |
| product_name | string | true | 商品名称| |

> Response

- 成功

```json
{
    "code":0,
    "msg":"添加商品成功",
    "data":nil
}

```

##3 获取商品

#### Method: ``` GET  ```
#### Path: ``` v1/get/product?product_id=1```

> Argument

| 参数 | 类型 | 必选 | 说明 | 备注 |
| :--- | :--- | :--- |:---| :--- |
| product_id | string | true | 商品ID| |

> Response

| 参数 | 类型 | 必选 | 说明 | 备注 |
| :--- | :--- | :--- |:---| :--- |
| userID | string | true | 用户ID| |
| productID | string | true | 商品ID| |
| productName | string | true | 商品名称| |

- 成功

```json
{
    "code":0,
    "msg":"",
    "data":{
            "userID":1,
            "productID":1,
            "productName":"电脑"
            }
}
