# BIGER 交易所 OPEN API

BIGER OPEN API 提供两种API， 1. Rest API 用于操作用户账户和订单，2. WebSocket API 用于获取行情信息，主要功能如下：

* WebSocket API: 获取市场行情
* REST API: 查询账户信息, 可用金额和冻结金额
* REST API: 执行买入、卖出、撤单和查询挂单命令


# REST API 简介
BIGER 的REST API URL为 https://pub-api.biger.in , 在使用REST API 操作订单时，需要使用签名认证，以保证通信安全。通过REST API操作以下功能
* 行情的获取
* 账户操作
* 交易


## 签名认证
### 安全认证
基于安全考虑，除行情外的REST API均需要访问令牌，即Access Token.一个账户可以申请多个Access Token, 以满足每个APP使用独立的Access Token。Access Token需单独申请，请联系管理员. 申请时需提供公钥和Access Token的过期日, 还有用API的IP地址，目前支持的公钥是RSA。

### 请求头

非行情API均需以下三个请求头

`名字` | `值`
----------------- | -----------------------------------------
UCEX-ACCESS-TOKEN | 申请后获取的Access Token
UCEX-REQUEST-EXPIRY | 此请求的过期时间，Unix epoch millisecond 
UCEX-REQUEST-HASH | 由请求参数和私钥计算出来的签名

### 签名运算
用SHA256进行签名，签名计算的字符串由以下四部分连接组成
* query string
* 请求方法
* UCEX-REQUEST-EXPIRY的值
* 请求体

`示例`
```
{ 
GET /exchange/someEndpoint?someKey=someValue&anotherKey=anotherValue
HOST:xxxx
UCEX-REQUEST-EXPIRY: 999999999999999
UCEX-ACCESS-TOKEN: myAccessToken
UCEX-REQUEST-HASH: c8owjqPSnY4mgFK8IHTk+1S+zhaEaAdoS6tJvr+o5FJFLymMyedOC6xJL9vCmVHALgXm+1mwF+0z1ZHVyJDKrdptZIfXis1tswBtt0v4k69ADYBlZkpLAhCpf0s55OQ18BbhGsrWpjm2kLtPEsPY3hvsh5nqWQQfJRAMzWFmg/8hnNa3MvWJLpZexFOYRLzmTdqthhKlw8pOvuE4pURbe27OLS4lINwY+0ck1DGINRE4/UtH+kYK3AAQq8CE/mSnWVNrIBFpYAe0frEZDluYppnuVXs3IGIQelR3RPqyYY5bfdccHVU8yBBaACRWZMTnvbdQW3TOSV/ccojaHEHBJA==
}
```

UCEX-REQUEST-HASH 计算公式如下: 
```
{
Base64Encode(RSAEncrypt(myPrivateKey, SHA256(utf8ToBytes(“someKey=someValue&anotherKey=anotherValueGET999999999999999”)))
}
```

## API列表

#### 查询钱包
路径：/exchange/accounts/list/accounts
方法: GET
示例：
```
路径	/exchange/accounts/list/accounts
返回	{
	"result"		: 	"Success",
 	"code"		: 	200,
 	"msg"		: 	"Success",
 	"data"		: 
	[
    	{
      		"coinCode"				:	 101,
      		"coinName"				:	 "BTC",
     		"balance"					:	 "9945.41972572",
		    "balanceUpdateTime"		: 	1530520590125,
		    "lockedAmount"			: 	"0",
		    "availBalance"				: 	"9945.41972572",
		    "lockedAmountUpdateTime"	: 	1530520592901
},
…
]
}
```

#### 查询指定单
路径：		/exchange/orders/get/{orderId}
方法: 		GET
示例：
路径	/exchange/orders/get/43960eab-d040-4eca-a4cd-bb20473e9960
返回	

```
{
"result":		"Success",
	"code":		200,
	"msg":		"Success",
	"data":
	{
		"orderId"			:"43960eab-d040-4eca-a4cd-bb20473e9960",
		"clientOrderId"		:11741743498134528,
		"side"				:"SELL",
		"symbol"				:"LTCUSDT",
		"baseCurrencyCode"	:103,
		"orderType"			:"LIMIT",
		"orderState"			:"FILLED",
		"price"				:"56.79",
		"orderQty"			:"1.08751",
		"filledQty"			:"1.08751",
		"totalPrice"			:"61.7596929",
		"dealPrice"			:"56.79",
		"completeTime"		:1537160398445,
		"createTime"			:1537160398408,
		"updateTime"			:1537160398448,
		"rejectReason"		:null
	}
}
```

返回字段说明

字段名 | 描述 | 取值
----------- | ----------------------------------------- | ---------------
orderId | 系统产生的订单标识 | GUID
clientOrderId | 终端显示的订单标识 | 64bit Integer 
orderState | 订单状态 | PENDING: 系统接收了订单，正在处理中 ; NEW: 新订单处理完毕 ; PARTIALLY_FILLED: 部分成交 ; FILLED: 全部成交 ; PENDING_CANCEL: 正在取消 ; CANCELED: 已取消(可能有部分成交) ; REJECTED:  已拒接(拒接原因在rejectReason)
filledQty | 成交数量 | 到目前成交数量
totalPrice | 成交额 | 到目前所有成交额加总，单个成交的成交额=成交数量*成交价格
dealPrice | 成交均价 | totalPrice / filledQty
completeTime | 完成时间 | Unix epoch milliseconds. 仅全成交有此值
updateTime | 更新时间 | Unix epoch milliseconds
rejectReason | 拒接原因 |  String 
side | 买卖 | BUY, SELL
symbol | 交易对 | 参考 Appendix A
orderType | 类型 | LIMIT，暂时只支持限价
orderQty | 数量 |  string 
Price | 单价 | String 



#### 查询当前所有单
路径：	/exchange/orders/current?symbol={ symbol }&side={ side }&offset={ offset=}&limit={ limit }
方法：	GET
请求参数

`参数名` | `必须` | `默认` | `描述` | `取值`
-------| ----- | ------ | ------- | -------------
symbol | 是 | NA | 币对 | 参考 Appendix A
side | 是 | NA | 买卖 | BUY, SELL
offset | 否 | 0 | 起始偏移量  | 用于分页获取
limit | 否 | 20 | 获取数量 | 最大100

示例：
路径	/exchange/orders/current?symbol=LTCUSDT&side=BUY&offset=0&limit=50
返回	

```
{
"result"	:	"Success",
"code"	: 	200,
"msg"	:	"Success",
"data"	:
[
	{
		"orderId"		:	"fdb97848-2034-4638-bd58-dd023e570c3d",
		"clientOrderId":	11362746867123200,
		...
	},
	{
		"orderId"		:	"0f92fb04-ffc7-42f2-9032-b1b6006d3f9d",
		"clientOrderId":	11230322151130112
		...
	}
]
}
```

#### 下单
路径：	/exchange/orders/create
方法: 	POST
请求体

参数名 | 必须? | 默认 | 描述 | 取值
-------| ----- | ------ | ------------- | -------------
symbol | 是 |  | 交易对 | 
side | 是 |  | 买卖 | BUY,  SELL
Price | 是 |  | 交易价格 | unit price - String
orderQty | 是 |  | 交易数量 | String to avoid rounding issues
orderType | 是 |  | 订单类型 | LIMIT

示例：
路径	/exchange/orders/create
请求体	

```
{
	"symbol" 	:  "BCHUSDT",
	"side"   		:  "BUY",
	"price"		:  "451.29"
	"orderQty" 	:  "0.14536",
	"orderType"	:  "LIMIT"
}
```

返回	

```
{
	"result":	"Success",
	"code":	200,
	"msg":	"Success",
	"data":
	{
		"orderId"			:"c0a480e7-211f-4090-8b56-96abee5d32ce",
		"clientOrderId"			:11741744533079040,
		"side"					:"BUY",
		"symbol"					:"BCHUSDT",
		"orderType"				:"LIMIT",
		"orderState"				:"PENDING",
		"price"					:"451.29",
		"orderQty"				:"0.14536",
		"filledQty"				:"0.00000",
		"totalPrice"				:"0.0000000",
		"dealPrice"				:"0",
		"completeTime"			:null,
		"createTime"				:1537160400382,
		"updateTime"				:null,
		"rejectReason"			:null
	}
}
```

2.2.2撤单
路径：		/exchange/orders/cancel/{orderId}
方法:		PUT
示例：
路径	/exchange/orders/cancel/725c0119-114c-471c-a2c5-d5c51d0210dd
返回	

```
{
	"result"		: "Success",
	"code"   	: 200,
	"msg"    	: "Success"
}
Or in cases of error	{
	"result" 		: "Error",
	"code"   	: 99506,
	"msg"    	: "Some error message"
}

```

常见错误	
* order.not.exist – 给定订单不存在或者在PENDING状态
* order.update.error.user.mismatch – 不是此用户订单
* order.update.error.cancelled – 此此订单已撤销
* order.cancel.failed.wrong.state – 此订单不可撤销，非NEW和 PARTIALLY_FILLED状态
* ORDER CANCEL FAILURE PENDING ENGINE – 此订单不在订单簿中，可能已成交
* The system is busy, please try again later – 系统繁忙，请重试




# Websocket API
 Websocket API URL为 wss://www.biger.in/ws , 通过websocket API可以获取市场数据。

## 系统接口
### 心跳请求
客户端需定时向系统发送心跳请求以确认网络和系统状态正常。正常情况下，系统会立即回复Pong消息。系统超出30秒没有收到客户端的心跳请求，将关闭客户端网络链接。

语法
```
{
"method"	: "server.ping",
"params"	: [],
"id"		: <id>
}

```

### 示例
```
请求: {"method": "server.ping", "params": [], "id": 1516681178}
返回: {"result": "pong", "error": null, "id": 1516681178}
```

### 查询系统时间
获取当前系统时间，回复时间从Epoch开始计算起，单位为秒。本文以下所有涉及时间的参数以及回复内容均为Epoch时间。建议客户端用此时间作为与系统交互的时间基准。

语法
```
{
  "method"	: "server.time",
  "params"	: [],
  "id"		: <id>
}
```

示例

```
请求: {"method": "server.time", "params": [], "id": 1516681178}
返回: {"result": 1520437025, "error": null, "id": 1516681178}
```


### K线接口
K线间隔参数可设置为以下之一：60（1分钟），300（5分钟）， 600（10分钟），900（15分钟），1800（30分钟），3600（1小时），14400（4小时），86400（1天），604800（1周）， 2592000（1月）。

#### 查询K线
K线查询数量最多可以同时请求2500条，如果超出范围系统将返回参数错误。

语法
```
{
  "method"	: "kline.query",
  "params"	: [<market>, <start_time>, <end_time>, <interval>],
  "id"		:      <id>
}
```

参数 | 数据类型 | 描述
-------| -------| ---------
market | String | 交易品种
start_time | Integer	开始时间
end_time | Integer | 结束时间
Interval | Integer | K线间隔

回复语法:
```
"result": [
    [
        1492358400, 时间
        "7000.00",  开盘价
        "8000.0",   收盘价
        "8100.00",  最高价
        "6800.00",  最低价
        "1000.00"   成交量
        "123456.00" 成交额
        "BTCUSDT"   交易品种
    ]
    ...
]
```

示例
```
> {"method": "kline.query", "params": ["BTCBCH", 1520432255, 1520433255, 900], "id": 1516681178}
< {
"result": 
[
        [
            1520432100,
            "8093",
            "8008",
            "8093",
            "8008",
            "45",
            "361758",
            "BTCUSDT"
        ],
        [
            1520433000,
            "8089",
            "8079",
            "8089",
            "8021",
            "57",
            "459239",
            "BTCUSDT"
        ]
    ],
    "error"	: null,
    "id"		: 1516681178
}
```

####  订阅K线
订阅成功之后，系统在发现数据变化时会及时推送最新的一到两根K线。
语法
```
{
  "method"	: "kline.subscribe",
  "params"	: [<market>, <interval>],
  "id"		:      <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String | 交易品种
interval | Integer | K线间隔

示例
```
> {"method": "kline.subscribe", "params": ["BTCUSDT", 900], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
< {"method": "kline.update", "id": null, "params": [[1520436600, "8040", "8040", "8040", "8040", "9", "72360", "BTCUSDT"]]}

```

####  取消K线订阅
语法
```
{
  "method"	: "kline.unsubscribe",
  "params"	: [<market>],
  "id"		: <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String | 交易品种。若参数为空，即取消所有K线订阅。

示例
```
> {"method": "kline.unsubscribe", "params": [], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
```

### 最新报价接口
#### 查询最新报价
语法
```
{
  "method"	: "price.query",
  "params"	: [<market>],
  "id"		:  <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String | 交易品种

示例
```
> {"method": "price.query", "params": ["BTCUSDT"], "id": 1516681178}
< {
    "result"	: "8074.00000000",
    "error"	: null,
    "id"		: 1516681178
  }
```

#### 订阅最新报价
订阅成功之后，系统在发现数据变化时会及时推送最新报价。

语法
```
{
  "method"	: "price.subscribe",
  "params"	: [<market>],
  "id"		:  <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market| String | 交易品种

示例
```
> {"method": "price.subscribe", "params": ["BTCUSDT"], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
< {"method": "price.update", "id": null, "params": ["BTCUSDT", "8050"]}
```

#### 取消最新报价订阅
语法
```
{
  "method"	: "price.unsubscribe",
  "params"	: [<market>],
  "id"		:      <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String | 交易品种。若参数为空，即取消所有报价订阅。

示例
```
> {"method": "price.unsubscribe", "params": [], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
```

### 市场成交数据接口
#### 查询逐笔成交历史

支持查询最多100条历史成交数据查询。
语法
```
{
  "method"	: "deals.query",
  "params"	: [“<market>”, “<limit>”, “<last_id>”],
  "id"		:      <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String | 交易品种
last_id | String | 上次查询返回的最新成交ID



示例

```
> {"method": "deals.query", "params": ["BTCUSDT", 3, 0], "id": 1516681178}
< {
    "result": [
        {
            "price"	: "8056",
            "time"	: 1520438100.3066709,
            "id"		: 1759,
            "amount"	: "3",
            "type"	: "buy"
        },
        {
            "price"	: "8007",
            "time"	: 1520438000.2892129,
            "id"		: 1758,
            "amount"	: "9",
            "type"	: "buy"
        },
        {
            "price"	: "8050",
            "time"	: 1520437900.2736571,
            "id"		: 1757,
            "amount"	: "6",
            "type"	: "buy"
        }
    ],
    "error"	: null,
    "id"		: 1516681178
}
```

#### 逐笔成交数据订阅
订阅成功之后，系统在发现数据变化时会及时推送成交数据。

语法
```
{
"method"	: "deals.subscribe",
"params"	: [“<market>”],
"id"		: <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String  | 交易品种

示例

```
> {"method": "deals.subscribe", "params": ["BTCUSDT"], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
< {"method": "deals.update", "id": null, "params": ["BTCUSDT", [{"price": "8044", "type": "buy", "time": 1520438400.361028, "amount": "2", "id": 1762}, {"price": "8078", "type": "buy", "time": 1520438300.341769, "amount": "9", "id": 1761}, {"price": "8076", "type": "buy", "time": 1520438200.324909, "amount": "10", "id": 1760}, {"price": "8056", "type": "buy", "time": 1520438100.3066709, "amount": "3", "id": 1759}, {"price": "8007", "type": "buy", "time": 1520438000.2892129, "amount": "9", "id": 1758}, {"price": "8050", "type": "buy", "time": 1520437900.2736571, "amount": "6", "id": 1757}, {"price": "8074", "type": "buy", "time": 1520437800.257802, "amount": "1", "id": 1756}, {"price": "8014", "type": "buy", "time": 1520437700.239372, "amount": "4", "id": 1755}, {"price": "8054", "type": "buy", "time": 1520437600.223423, "amount": "9", "id": 1754}, {"price": "8049", "type": "buy", "time": 1520437500.2082629, "amount": "8", "id": 1753}, {"price": "8002", "type": "buy", "time": 1520437400.1939909, "amount": "2", "id": 1752}, {"price": "8000", "type": "buy", "time": 1520437300.1761429, "amount": "5", "id": 1751}, {"price": "8002", "type": "buy", "time": 1520437200.1584849, "amount": "10", "id": 1750}, {"price": "8065", "type": "buy", "time": 1520437100.142282, "amount": "5", "id": 1749}, {"price": "8099", "type": "buy", "time": 1520437000.1258199, "amount": "5", "id": 1748}, {"price": "8009", "type": "buy", "time": 1520436900.1072299, "amount": "6", "id": 1747}, {"price": "8066", "type": "buy", "time": 1520436800.0908389, "amount": "10", "id": 1746},
...
```


#### 取消逐笔成交数据订阅
语法
```
{
  "method"	: "deals.unsubscribe",
  "params"	: [<market>],
  "id"		:      <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market	String	交易品种。若参数为空，即取消所有成交订阅。



示例
```
> {"method": "deals.unsubscribe", "params": [], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}

```

### 市场深度数据接口
#### 查询最新市场深度
语法
```
{
  "method"	: "depth.query",
  "params"	: [“<market>”, <limit>, <interval>],
  "id"		: <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String | 交易品种
limit | Integer | 盘口深度
interval | String | 盘口报价精度。“0”为最大精度。可选精度为：    “0”，”0.1", “0.01", “0.001", “0.0001", “0.00001", “0.000001", “0.0000001", "0.00000001"


示例
```
> {"method": "depth.query", "params": ["BTCUSDT", 10, "0"], "id": 1516681178}
< {
    "error"	: null,
"result": 
{
        "asks": [],
        "bids": [
            [
                "803.25",
                "0.063"
            ],
            [
                "803.2",
                "0.051" 
            ],
            [
                "803.15",
                "0.057"
            ],
            [
                "803.1",
                "0.083"
            ],
            [
                "803.05",
                "0.077"
            ],
            [
                "803",
                "0.074"
            ],
            [
                "802.95",
                "0.119"
            ],
            [
                "802.9",
                "0.085"
            ],
            [
                "802.85",
                "0.05"
            ],
            [
                "802.8",
                "0.054"
            ]
    },
    "id": 1516681178
}
```

#### 订阅市场深度
订阅成功之后，系统在发现数据变化时会及时推送深度数据。深度数据更新中的布尔变量为true时，即为全推数据，若是false，即为变化推送。一般情况下，系统只推送深度变化数据，即深度数据的增加和修改，减少（当数量为0时即为删除该档数据）。系统在每超过一分钟之后，有一次全推的数据。

语法
```
{
  "method"	: "depth.subscribe",
  "params"	: [“<market>”, <limit>, <interval> ],
  "id"		: <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market | String | 交易品种
limit | Integer | 盘口深度
interval | String | 盘口报价精度。“0”为最大精度。可选精度为：    “0”，”0.1", “0.01", “0.001", “0.0001", “0.00001", “0.000001", “0.0000001", "0.00000001"


示例
```
> {"method": "depth.subscribe", "params": ["BTCUSDT", 10, "0"], "id": 1516681178}
< {"error": null, "result": {"status": "success"}, "id": 1516681178}
< {"method": "depth.update", "params": [true, {"asks": [], "bids": []}, "BTCUSDT"], "id": null}
```

#### 取消市场深度订阅
语法
```
{
"method"	: "depth.unsubscribe",
"params"	: [],
"id"		:      <id>
}
```

参数 | 数据类型 | 描述
------- | ------- | --------
market |	String | 交易品种。若参数为空，即取消所有深度订阅。



示例
```
> {"method": "depth.unsubscribe", "params": [], "id": 1516681178}
< {"error": null, "result": {"status": "success"}, "id": 1516681178}
```


#### 错误处理
当接口调用失败时，系统会返回表示错误的应答。
语法
```
{
  	“error"	: 
{
    		“code"		: <code>,
    		“message"	: "<message>"
  	},
  "id"		:1516681178,
  "result"	:null
}
```

`参数` | `数据类型` | `描述`
------- | ------- | --------
code | Integer | 错误码，具体内容参考下面说明
message | String | 错误信息，具体内容参考下面说明

错误代码说明

`错误代码` | `错误原因`
------- | ------- | --------
6001 | 任何的参数错误，系统返回返回参数错误代码
6005 | 系统回复超时。若是系统内部由于某种原因无法在5秒钟内正常回复数据，系统将返回此错误代码。
6012 | 用户验证失败。若是用户Access token验证失败，系统将会返回此错误代码。
6013 | 服务繁忙。若是客户端与系统之间网络传输缓慢，系统将会丢掉过时数据，并且返回此服务繁忙错误代码。
6014 | 请求频率超出限制。若是过度发送请求，系统将返回此错误代码并且延迟回复数据。
6015 | 订阅数超出请求。若是订阅数量超出超出系统限制，系统将会返回此错误代码并且关掉网络链接。

## Appendix A - Symbol list
* LTCUSDT
* ETHUSDT
* BTCUSDT
* BCHUSDT
* LTCETH
* BCHETC
* ETHBTC
* LTCBTC
* BCHBTC
