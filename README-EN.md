# BIGER EXCHANGE OPEN API
Official Documentation for the Biger APIs and Streams


BIGER OPEN API Provides two type APIs， 1. Rest API for Account and Orders and query historic K line data，2. WebSocket API for realtime market data and K-line data：

* WebSocket API: Query market data and K line data
* REST API: Query Account information and balance
* REST API: Execute Order, queyr Orders and cancel orders 
* REST API: query historic K line data
* Temporary websocket auth token exchange

# REST API Introduction
BIGER REST API URL is under: https://pub-api.biger.in , when you use REST API to execute orders, you need to sign your request to make sure the communication safty 
. REST API Provdies following functions:
* Query markets
* Operate Account, e.g. query balance
* Exeucte Orders


## Signing requests
### Token Authentication
To make sure that API commucation is safe，REST API must need Access token apart from Market Data API, every account  can apply multiple Access Token, so that each APP can use different Access Token。
Access Token need to apply in http://biger.in, please provide your public key (RSA), the expire data of Access Token and IP address when you apply Access Token.

To apply for access token, you will also first need to generate you rown RSA key pair, and give us your public key. (keep your private key safe on your own end).

To generate a RSA key pair, you can use a multitude of openly available tools.
 * option 1 - Using openssl via command line - https://rietta.com/blog/2012/01/27/openssl-generating-rsa-key-from-command/
 * option 2 - a few lines of java code - 
```
        KeyPairGenerator g = KeyPairGenerator.getInstance("RSA");
        g.initialize(2048);
        KeyPair p = g.generateKeyPair();
        Files.write(Paths.get("private"), p.getPrivate().getEncoded(), StandardOpenOption.CREATE_NEW);
        Files.write(Paths.get("public"), p.getPublic().getEncoded(), StandardOpenOption.CREATE_NEW);
```

### 请求头

非行情API均需以下三个请求头

`名字` | `值`
----------------- | -----------------------------------------
BIGER-ACCESS-TOKEN | 申请后获取的Access Token
BIGER-REQUEST-EXPIRY | 此请求的过期时间，Unix epoch millisecond 
BIGER-REQUEST-HASH | 由请求参数和私钥计算出来的签名

### 签名运算
用SHA256进行签名，签名计算的字符串由以下四部分连接组成
* query string
* 请求方法
* BIGER-REQUEST-EXPIRY的值
* 请求体

`示例`
```
{ 
GET /exchange/someEndpoint?someKey=someValue&anotherKey=anotherValue
HOST:xxxx
BIGER-REQUEST-EXPIRY: 999999999999999
BIGER-ACCESS-TOKEN: myAccessToken
BIGER-REQUEST-HASH: c8owjqPSnY4mgFK8IHTk+1S+zhaEaAdoS6tJvr+o5FJFLymMyedOC6xJL9vCmVHALgXm+1mwF+0z1ZHVyJDKrdptZIfXis1tswBtt0v4k69ADYBlZkpLAhCpf0s55OQ18BbhGsrWpjm2kLtPEsPY3hvsh5nqWQQfJRAMzWFmg/8hnNa3MvWJLpZexFOYRLzmTdqthhKlw8pOvuE4pURbe27OLS4lINwY+0ck1DGINRE4/UtH+kYK3AAQq8CE/mSnWVNrIBFpYAe0frEZDluYppnuVXs3IGIQelR3RPqyYY5bfdccHVU8yBBaACRWZMTnvbdQW3TOSV/ccojaHEHBJA==
}
```

BIGER-REQUEST-HASH 计算公式如下: 
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
路径：		/exchange/orders/get/orderId/{orderId}
方法: 		GET
示例：
路径	/exchange/orders/get/orderId/43960eab-d040-4eca-a4cd-bb20473e9960
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
----------- | --------------------------------------------------------- | ---------------
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
Note that when creating an order, the order quantity and order price scale(number of decimal points) will be truncated if they were higher than our accepted values.
For example, eg LTC/USDT orders will have the order price truncated to 2 decimal places while order quantity will be truncated to 5 decimal places. See appendix B

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
	"price"		:  "451.29",
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


# REST K-line query API
 REST API https://biger.in/md/kline is dedicated to K-line history query. Please use WebSocket API for real-time K-Line subscription/query.

## 语法

Param | Required | Type | Description  
------ | ------ | ------ | ------------------------------------------------------
symbol | Yes | String | coin pair symbol, eg. BTCUSDT
period / interval | Yes | String | K-line timeframe. Possible values：1min，5min，15min，30min，60min，1day，1mon，1week，60，300，900，1800，3600，86400，604800, 2592000
start_time  | 	No | Integer | time in seconds since epoch. eg. 1543274801. The default value is the start time of last 200 K-lines，
end_time | No | Integer | time in seconds since epoch. eg. 1543274801. The default value is current time.

### HTTP request URL syntax
```
https://biger.in/md/kline?id=0&symbol=<symbol>&start_time=<timestamp>&end_time=<timestamp>&period=<period>

```

### HTTP response syntax
```
{“error":null,"id":0,"result":[
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
]}
```

### Sample
```
Request: 
https://biger.in/md/kline?id=0&symbol=BTCUSDT&start_time=1543274801&end_time=1543374801&period=1day
Response: 
{“error":null,"id":0,"result":[
[1543190400,”4394","3863.05","4394","3701.72","1809.258054","7117136.76413459","BTCUSDT"],
[1543276800,”3862.7","3875.11","3939.02","3686.59","1597.117575","6097170.88594629","BTCUSDT"],
[1543363200,”3909.69","4262.39","4389.04","3887.99","1734.877599","7166445.63528313","BTCUSDT"]
]}
```



# Websocket API
 Websocket API URL为 wss://www.biger.in/ws , 通过websocket API可以获取市场数据。

## Temporary token exchange
Some websocket APIs require you to authenticate using a temporary API token. To retrieve this temporary API token, you need to request
via a HTTP POST to https://pub-api.biger.in/tokens/exchange
with HTTP header BIGER-ACCESS-TOKEN-FOR-EXCHANGE where the value is your access token.

You should get a HTTP 200 response (if not then check your access token or contact us) that looks like the below-
```
{
    "code": 200,
    "result": "0d4fCb6YHhhOg6QsTyhydOLqISfF8V8aU8CV39w1BjeBXMv9oHiKrAcsRkpasmrRNh/LJzoEf/Ah4ul/ELnmKg0z/jvJ3DOhnsPO16UhW2LC6+Gw1EHh5bpMQx1AVeMjDAZZ9fMCJe52lbwvV6QaresUtez8tJFrvIfoL/APVX0wt60Ze54Gu0lCOVTUoYLHlOopBg+Vrrzxm5vtsSSG32Ivm2zr2vQ7ydxhiutpXwA4CXUfT60QBo0cU0l6UL9yd/dPnB/UXQ7PIveoQzb7/kxJ8dIeykxSVbkVN7q0JL9psKDGqn//UmkGui5huvIWlJuun2RAKujZna5uMdVW1aRObt8nSxjJey1AXThaW6AWnObre1h49l1MHn+qf+I6StJiUOljPKL0gbdvOGMXlsiMRNdxnvDeJuwWghiFByINYmGvp1BrYb7Ipe7Ja38YRMdidd3Z7TvXUKIj7iv5BWL0fNO+OGeXpuWQOelP5rhyeOwvra2yRPzrUMkUnuZGrrpjQpQvqmiGpkPvdCyLYsjUhaCpRRwAcGbtw+yN+SY=",
    "encryptedKey": "Ukb6CLSg5g0Ey4iKZUeVq/HcNacXKwjGC+8UCwMAdej7V+V7Xdp4yE4drYV5YPJu/fr/nVtVWVogfLMKF9sHMpPU6KDZeFsGZlsciTDnf3uDcS5b7mgpsap6DU38rxE7+20GiWQf5TUTIcJ23lI9oRZSE9ooU5NCDgHtQsrshIP1HiI4+iACC9WiLOqo9zESgFsRr9I7ICjQNM7sFjw4NsCLurJdFaFdC79vjMruq6DpcnkWRbLysQFqRWxBQsxAkXB1i1FMeU3McTdKkEWyuOKwpLBXDm9VKlauS7VKOgWEPQ+mUeiPi6KwHBhtIGzbJ8glCAsxVyQ+j06KuxajRg=="
}
```
Now perform the following steps
  * base64 decode the value of 'encryptedKey', and then use your private key to decrypt the result to obtain AES secret key
  * base64 decode the value of 'result', and then perform AES decryption using the AES secret key on the result to obtain your temporary token

Note that the temporary token is only valid for 30 seconds.

## System APIs
### Heartbeat request
To keep a websocket session live, client is reuqired to send ping request periodically to biger. And Biger market data service (Biger MD in short) will respond a Pong message immediately to help client identify session states. A session would be closed if Biger MD failed to receive a ping message in 30s.

Syntax
```
{
"method"	: "server.ping",
"params"	: [],
"id"		: <id>
}

```

### Sample
```
Request: {"method": "server.ping", "params": [], "id": 1516681178}
Response: {"result": "pong", "error": null, "id": 1516681178}
```

### Server time query
Get the current time of Biger MD. It is in seconds since epoch. It is suggested that client use the value to keep in sync with Biger MD.
Please note all time related parameters in websocket API request and responses are all in seconds since epoch.

Syntax
```
{
  "method"	: "server.time",
  "params"	: [],
  "id"		: <id>
}
```

Sample
```
Request: {"method": "server.time", "params": [], "id": 1516681178}
Response: {"result": 1520437025, "error": null, "id": 1516681178}
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
------- | -----------------------------------------------------------------------
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

=======
## Appendix B - accepted order quantity and order price scale
Note that when creating an order, the order quantity and order price scale(number of decimal points) will be truncated if they were higher than our accepted values.
For example, eg LTC/USDT orders will have the order price truncated to 2 decimal places while order quantity will be truncated to 5 decimal places.

Symbol | Price Scale | Qty Scale
--- | --- | ---
ETHBTC | 6 | 3
BCHBTC | 5 | 3
LTCBTC | 6 | 3
BTCUSDT | 2 | 6
ETHUSDT | 2 | 5
BCHUSDT | 2 | 5
LTCUSDT | 2 | 5
BCHETH | 8 | 8
LTCETH | 5 | 3

## Appendix C - Check validity of access token
Once given your access token, you can try the sample code below to see if your key pair and tokens are set up correctly.
```
import javax.crypto.Cipher;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import java.security.KeyFactory;
import java.security.MessageDigest;
import java.security.SecureRandom;
import java.security.spec.PKCS8EncodedKeySpec;
import java.util.Base64;

public class TokenValidityCheck {
    public static void main(String[] args) throws Exception {
        HttpClient c = HttpClient.newBuilder().version(HttpClient.Version.HTTP_1_1).build();
        long now = System.currentTimeMillis();
        long expiry = now + 5000;
        HttpRequest req = HttpRequest.newBuilder()
                .GET()
                .uri(new URI("https://pub-api.biger.in/exchange/orders/current"))
                .header("BIGER-ACCESS-TOKEN", "myAccessToken")
                .header("BIGER-REQUEST-EXPIRY", expiry + "")
                .header("BIGER-REQUEST-HASH", hash(("GET" + expiry).getBytes(StandardCharsets.UTF_8)))
                .header("Accept", "application/json")
                .build();
        c.sendAsync(req, HttpResponse.BodyHandlers.ofString())
                .thenApply(HttpResponse::body)
                .thenAccept(System.out::println)
                .join();
    }

    private static String hash(byte[] payload) throws Exception {
        Cipher c = Cipher.getInstance("RSA");
        c.init(Cipher.ENCRYPT_MODE, KeyFactory.getInstance("RSA").generatePrivate(new PKCS8EncodedKeySpec(ClassLoader.getSystemResourceAsStream("private").readAllBytes())), new SecureRandom());

        return Base64.getEncoder().encodeToString(c.doFinal(MessageDigest.getInstance("SHA-256").digest(payload)));
    }

}
```
You need to provide the resource 'private' which is your private key in DER format as well as replace myAccessToken with your actual access token.
