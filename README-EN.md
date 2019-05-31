# BIGER EXCHANGE OPEN API
Official Documentation for the Biger APIs and Streams


BIGER OPEN API Provides two type APIs, 
1. Rest API for Account and Orders and query historic K line data
2. WebSocket API for realtime market data and K-line data：

* WebSocket API: Query market data and K line data
* REST API: Query Account information and balance
* REST API: Execute Order, queyr Orders and cancel orders 
* REST API: query historic K line data
* Temporary websocket auth token exchange

# OFFICIAL JAVA SDK FOR BIGER EXCHANGE OPEN API
We do provide an open-source official java sdk at https://github.com/biger-exchange/biger-client, please feel free to check it out.

# REST API Introduction
BIGER REST API is at https://pub-api.biger.pro which requrires IP access. We also provide two public market data API under 	https://www.biger.pro which doesn't require IP access.

When you use REST API to execute requests that require authentication, you need to sign your request so that we can authenticate you, as well as ensure that the request was not tempered by a middleman.
. REST API Provdies following functions:
* Query markets
* Operate Account, e.g. query balance
* Exeucte Orders


## Signing requests
### Token Authentication
To sign your http requests, you need an access token. Please apply for one at https://biger.pro. You will need to provide your public key (RSA), the expiry date of requested access Token(else we wil give you 1 year by default) and IP address for whitelisting purposes.

To generate a RSA key pair, you can use a multitude of openly available tools.
 * option 1 - Use our java sdk to generate the rsa key pair. See https://github.com/biger-exchange/biger-client and https://github.com/biger-exchange/biger-client-example/blob/master/src/main/java/com/biger/client/examples/GenerateKeyPair.java
 * option 2 - Using openssl via command line - https://rietta.com/blog/2012/01/27/openssl-generating-rsa-key-from-command/
 * option 3 - a few lines of java code - 
```
        KeyPairGenerator g = KeyPairGenerator.getInstance("RSA");
        g.initialize(2048);
        KeyPair p = g.generateKeyPair();
        Files.write(Paths.get("private.der"), p.getPrivate().getEncoded(), StandardOpenOption.CREATE_NEW);
        Files.write(Paths.get("public.der"), p.getPublic().getEncoded(), StandardOpenOption.CREATE_NEW);
```

### Request headers

Rest API endpoints that require authentication (generally those to deal with orders) need to be signed and filled with specific http request headers.

HTTP header name | Value
----------------- | -----------------------------------------
BIGER-ACCESS-TOKEN | Access token that we provided you
BIGER-REQUEST-EXPIRY | Time of expiry of this request, Unix epoch millisecond (we will reject requests that arrive later than this time)
BIGER-REQUEST-HASH | Calculated hash value(signature) that you have to compute

### Generating BIGER-REQUEST-HASH
First, use SHA256 digest on the string that is formed by concatenating the values of
* query string
* http request method
* BIGER-REQUEST-EXPIRY
* request body

`Example`
```
GET /exchange/someEndpoint?someKey=someValue&anotherKey=anotherValue
HOST:xxxx
BIGER-REQUEST-EXPIRY: 999999999999999
BIGER-ACCESS-TOKEN: myAccessToken
BIGER-REQUEST-HASH: c8owjqPSnY4mgFK8IHTk+1S+zhaEaAdoS6tJvr+o5FJFLymMyedOC6xJL9vCmVHALgXm+1mwF+0z1ZHVyJDKrdptZIfXis1tswBtt0v4k69ADYBlZkpLAhCpf0s55OQ18BbhGsrWpjm2kLtPEsPY3hvsh5nqWQQfJRAMzWFmg/8hnNa3MvWJLpZexFOYRLzmTdqthhKlw8pOvuE4pURbe27OLS4lINwY+0ck1DGINRE4/UtH+kYK3AAQq8CE/mSnWVNrIBFpYAe0frEZDluYppnuVXs3IGIQelR3RPqyYY5bfdccHVU8yBBaACRWZMTnvbdQW3TOSV/ccojaHEHBJA==
```

Then, use RSA and your private key to encrypt the obtained value, and finally encode using base64 to produce the required BIGER-REQUEST-HASH value.

BIGER-REQUEST-HASH formula
```
Base64Encode(RSAEncrypt(myPrivateKey, SHA256(utf8ToBytes(“someKey=someValue&anotherKey=anotherValueGET999999999999999”)))
RSAEncrypt 要利用 RSA/ECB/PKCS1Padding
```

ECB - Electronic Codebook Mode, as defined in FIPS PUB 81, http://csrc.nist.gov/publications/fips/index.html.
PKCS1Padding - The padding scheme described in PKCS #1, http://www.rsa.com/rsalabs/node.asp?id=2125.


## REST API to query Market Data and Digit Currency History Data


### Digit Currency 24 Hours Price History Query API
This API is used for querying the last 24 hours currency price history:

URL Path: https://pub-api.biger.pro/exchange/coins/query/all
HTTP Method: GET
Authentication Header: No need

#### HTTP Request URL
```
https://pub-api.biger.pro/exchange/coins/query/all

```

#### HTTP Response
```
{
    "result": "Success",
    "code": 200,
    "msg": "Success",
    "data": [
        {
            "coinCode": 102,
            "coinName": "BCH",
            "fullName": "BCH",
            "scale": 8,
            "iconUrl": "/xxxx.png",
            "status": 1,
            "coinType": 0
        },
    ...
]}
```

#### Sample
```
Request: https://pub-api.biger.pro/exchange/coins/query/all
Response: 
{
    "result": "Success",
    "code": 200,
    "msg": "Success",
    "data": [
        {
            "coinCode": 102,
            "coinName": "BCH",
            "fullName": "BCH",
            "scale": 8,
            "iconUrl": "xxx.png",
            "status": 1,
            "coinType": 0
        },
    ...
]}
```


### Exchange Market 24 Hours Price History Query API
This API is used for querying exchange market history.  

URL Path: /exchange/markets/query/all
HTTP Method: GET
Authentication Header: No



#### HTTP Request URL
```
URL Require IP Access: https://pub-api.biger.pro/exchange/markets/query/all
```

#### HTTP Response
```
{"result":"Success","code":200,"msg":"Success","data":[{"symbol":"AEUSDT","symbolDisplayName":"AE/USDT","baseCurrencyCode":212,"baseCurrencyName":"AE","quoteCurrencyCode":106,"quoteCurrencyName":"USDT","amountDivisibilityUnit":"0.001","priceDivisibilityUnit":"0.0001","last":"0.3880","rate24h":"-0.0358","open24h":"0.4024","close24h":"0.3880","low24h":"0.3857","high24h":"0.4534","volume24h":"85841.449","rate7d":"-0.0214","low7d":"0.3779","high7d":"0.4534","open7d":"0.3965","close7d":"0.3880","volume7d":"559853.902","maxPriceScale":4,"maxQuantityScale":3,"maxTotalPriceScale":7,"ticker":null},
    ...
]}
```

#### Sample
```
Request: https://pub-api.biger.pro/exchange/markets/query/all
Response: 
{"result":"Success","code":200,"msg":"Success","data":[{"symbol":"AEUSDT","symbolDisplayName":"AE/USDT","baseCurrencyCode":212,"baseCurrencyName":"AE","quoteCurrencyCode":106,"quoteCurrencyName":"USDT","amountDivisibilityUnit":"0.001","priceDivisibilityUnit":"0.0001","last":"0.3880","rate24h":"-0.0358","open24h":"0.4024","close24h":"0.3880","low24h":"0.3857","high24h":"0.4534","volume24h":"85841.449","rate7d":"-0.0214","low7d":"0.3779","high7d":"0.4534","open7d":"0.3965","close7d":"0.3880","volume7d":"559853.902","maxPriceScale":4,"maxQuantityScale":3,"maxTotalPriceScale":7,"ticker":null},
    ...
]}
```



### Market Data K-line query  API

 REST API https://pub-api.biger.pro/md/kline is dedicated to K-line history query. Please use WebSocket API for real-time K-Line subscription/query.

URL Path: /md/kline
HTTP Method: GET

Request Parameters:

Parameter | Required | Type | Description  
------ | ------ | ------ | ------------------------------------------------------
symbol | Yes | String | coin pair symbol, eg. BTCUSDT
period / interval | Yes | String | K-line timeframe. Possible values：1min，5min，15min，30min，60min，1day，1mon，1week，60，300，900，1800，3600，86400，604800, 2592000
start_time  | 	No | Integer | time in seconds since epoch. eg. 1543274801. The default value is the start time of last 200 K-lines，
end_time | No | Integer | time in seconds since epoch. eg. 1543274801. The default value is current time.

##### HTTP request URL syntax
```
https://pub-api.biger.pro/md/kline?id=0&symbol=<symbol>&start_time=<timestamp>&end_time=<timestamp>&period=<period>

```

##### HTTP response syntax
```
{“error":null,"id":0,"result":[
    [
        1492358400,   <= Time
        "7000.00",    <= Open price
        "8000.0",     <= Last price
        "8100.00",    <= High price
        "6800.00",    <= Low price
        "1000.00"     <= Volume
        "123456.00"   <= Trade value
        "BTCUSDT"     <= Symbol
    ]
    ...
]}
```

##### Sample
```
Request: 
https://pub-api.biger.pro/md/kline?id=0&symbol=BTCUSDT&start_time=1543274801&end_time=1543374801&period=1day
Response: 
{“error":null,"id":0,"result":[
  [1543190400,”4394","3863.05","4394","3701.72","1809.258054","7117136.76413459","BTCUSDT"],
  [1543276800,”3862.7","3875.11","3939.02","3686.59","1597.117575","6097170.88594629","BTCUSDT"],
  [1543363200,”3909.69","4262.39","4389.04","3887.99","1734.877599","7166445.63528313","BTCUSDT"]
]}
```


## Account And Order Restful API List 

Account and Order API are protected API, so every method call should be signed with its key.

#### Query Account Balance
URL Path: /exchange/accounts/list/accounts
HTTP Method: GET
Sample: 
```
URL Path: 	/exchange/accounts/list/accounts
Return: 	{
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

#### Order query
URL Path: 		/exchange/orders/get/orderId/{orderId}
HTTP Method: 		GET
Sample: 
URL Path: 	/exchange/orders/get/orderId/43960eab-d040-4eca-a4cd-bb20473e9960
Return: 	

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

Response parameters description

Parameters | Description | Value
----------- | --------------------------------------------------------- | ---------------
orderId | System Generated Identity | GUID
clientOrderId | Client Side Order Id | 64bit Integer 
orderState | Order Status | PENDING: System received order, waiting for processing ; NEW: Order has been processed ; PARTIALLY_FILLED ; FILLED: fulfilled ; PENDING_CANCEL: waiting for cancel ; CANCELED: Already canceled (might already have some filled qty) ; REJECTED: Engine Reject order
filledQty | FILLED Qty | total filled qty
totalPrice | Filled total volume | total filled volume,  filled volume = Sum(filledQty * filled price)
dealPrice | average price | totalPrice / filledQty
completeTime | completed time | Unix epoch milliseconds. only have this time when Order is fulfilled 
updateTime | updated time | Unix epoch milliseconds
rejectReason | reason |  String 
side | buy or sell | BUY, SELL
symbol | market symbol | Please refer Appendix A
orderType | order type | LIMIT，currently only support limit order, market order is coming soon
orderQty |  qty |  string 
Price |  price | String 



#### Query all orders
URL Path: 	/exchange/orders/current?symbol={ symbol }&side={ side }&offset={ offset=}&limit={ limit }
HTTP Method:	GET
Request Parameters: 

`Parameters` | `Required?` | `Default` | `Description` | `Value`
-------| ----- | ------ | ------- | -------------
symbol | Yes | NA | Market Symbol | Refer Appendix A
side | Yes | NA | Buy or Sell | BUY, SELL
offset | No | 0 | result start point  | used for paging
limit | No | 20 | result set item limit | Max is 100

Sample: 
URL Path: 	/exchange/orders/current?symbol=LTCUSDT&side=BUY&offset=0&limit=50
Return: 	

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

#### Place Order
Note that when creating an order, the order quantity and order price scale(number of decimal points) will be truncated if they were higher than our accepted values.
For example, eg LTC/USDT orders will have the order price truncated to 2 decimal places while order quantity will be truncated to 5 decimal places. See appendix B

URL Path: 	/exchange/orders/create
HTTP Method: 	POST
Request Body:

Parameters | Required? | Default | Description | Value
-------| ----- | ------ | ------------- | -------------
symbol | Yes |  | Market symbol | 
side | Yes |  | Buy or Sell | BUY,  SELL
price | Yes |  | Pirce | unit price - String
orderQty | Yes |  | Quantity | String to avoid rounding issues
orderType | Yes |  | Order Type | Currently, only support LIMIT order, will support Market Order

Sample: 
URL Path: 	/exchange/orders/create
Request Body:	

```
{
	"symbol" 	:  "BCHUSDT",
	"side"   		:  "BUY",
	"price"		:  "451.29",
	"orderQty" 	:  "0.14536",
	"orderType"	:  "LIMIT"
}
```

Return: 	

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

2.2.2 Cancel Order
URL Path: 		/exchange/orders/cancel/{orderId}
HTTP Method: 	PUT
Sample: 
URL Path: 	/exchange/orders/cancel/725c0119-114c-471c-a2c5-d5c51d0210dd
Return: 	

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

Common Error:	
* order.not.exist –  The Order isn't existed 
* order.update.error.user.mismatch – Doesn't belong to current user
* order.update.error.cancelled – Order has been canceled
* order.cancel.failed.wrong.state – The order cannot be caneled, its not NEW or PARTIALLY_FILLED
* ORDER CANCEL FAILURE PENDING ENGINE – The order isn't in order book, might be already filled
* The system is busy, please try again later – 


# Websocket API
 Websocket API URL为 wss://www.biger.pro/ws , market data service is provided via the API。

##### Temporary token exchange
Some websocket APIs require you to authenticate using a temporary API token. To retrieve this temporary API token, you need to request
via a HTTP POST to https://pub-api.biger.pro/tokens/exchange
with HTTP header BIGER-ACCESS-TOKEN-FOR-EXCHANGE where the value is your access token.

You should get a HTTP 200 response (if not then check your access token or contact us) that looks like the below-
```json
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

##### Syntax
```
{
  "method" : "server.ping",
  "params" : [],
  "id"     : <id>
}

```

##### Sample
```json
Request: 
    {"method": "server.ping", "params": [], "id": 1516681178}
Response: 
    {"result": "pong", "error": null, "id": 1516681178}
```

### Server time query
Get the current time of Biger MD. It is in seconds since epoch. It is suggested that client use the value to keep in sync with Biger MD.
Please note all time related parameters in websocket API request and responses are all in seconds since epoch.

##### Syntax
```
{
  "method" : "server.time",
  "params" : [],
  "id"     : <id>
}
```

##### Sample
```json
Request: 
    {"method": "server.time", "params": [], "id": 1516681178}
Response: 
    {"result": 1520437025, "error": null, "id": 1516681178}
```


### K-line APIs
K-line timeframes：60（1 minute），300（5 minutes）， 600（10 minutes），900（15 minutess），1800（30 minutes），3600（1 hour），14400（4 hours），86400（1 day），604800（1 week）， 2592000（1 month）。

#### K-line query
At most 2500 k-line entries is allowed to be requested in one request. Otherwise an argument error shall be replied.

##### Request syntax
```
{
  "method" : "kline.query",
  "params" : ["<symbol>", <start_time>, <end_time>, <interval>],
  "id"     : <id>
}
```

Parameter | Required | Type    | Description
-------   | -------  | ------- | ---------
symbol    | Yes      | String  | trade symbol
start_time| Yes      | Integer | start time in seconds since epoch
end_time  | Yes      | Integer | end time in seconds since epoch
interval  | Yes      | Integer | time frame

##### Response syntax:
```
"result": [
    [
        1492358400,   <= Time
        "7000.00",    <= Open price
        "8000.0",     <= Last price
        "8100.00",    <= High price
        "6800.00",    <= Low price
        "1000.00"     <= Volume
        "123456.00"   <= Trade value
        "BTCUSDT"     <= Symbol
    ]
    ...
]
```

##### Sample
```json
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
    "error" : null,
    "id"    : 1516681178
}
```

####  K-line subscribe
If subscribe succeeds, Biger MD will publish the lastest 2 k-lines on kline changes.

##### Request syntax
```
{
  "method" : "kline.subscribe",
  "params" : ["<symbol>", <interval>],
  "id"     : <id>
}
```

Parameter | Required | Type     | Description
-------   | -------  | -------- | --------
symbol    | Yes      | String   | trade symbol
interval  | Yes      | Integer  | K-line timeframe

##### Sample
```json
> {"method": "kline.subscribe", "params": ["BTCUSDT", 900], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
< {"method": "kline.update", "id": null, "params": [[1520436600, "8040", "8040", "8040", "8040", "9", "72360", "BTCUSDT"]]}

```

####  Kline unsubscribe
##### Syntax
```
{
  "method" : "kline.unsubscribe",
  "params" : ["<symbol>"],
  "id"     : id
}
```

Parameter | Required | Type     | Description
-------   | -------  | -------- | -------
symbol    | No       | String   | trading symbol. Unsubscribe all kline subscriptions if no symbol is provided.

##### Sample
```json
> {"method": "kline.unsubscribe", "params": [], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
```

### Price APIs
#### Price query
##### Syntax
```
{
  "method" : "price.query",
  "params" : ["<symbol>"],
  "id"     : id
}
```

Parameter | Required | Type    | Description
-------   | -------  | ------- | --------
symbol    | Yes      | String  | trade symbol

##### Sample
```json
> {"method": "price.query", "params": ["BTCUSDT"], "id": 1516681178}
< {
    "result" : "8074.00000000",
    "error"  : null,
    "id"     : 1516681178
  }
```

#### Price subscribe
BigerMD will publish the lastest price on changes.

##### Syntax
```
{
  "method" : "price.subscribe",
  "params" : ["<symbol>"],
  "id"     :  id
}
```

Parameter | Required | Type    | Description
-------   | -------  | ------- | --------
symbol    | Yes      | String  | trade symbol

##### Sample
```json
> {"method": "price.subscribe", "params": ["BTCUSDT"], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
< {"method": "price.update", "id": null, "params": ["BTCUSDT", "8050"]}
```

#### Price unsubscribe

##### Syntax
```
{
  "method" : "price.unsubscribe",
  "params" : ["<symbol>"],
  "id"     : <id>
}
```

Parameter | Required | Type     | Description
-------   | -------  | -------- | --------
symbol    | No      | String    | trading symbol. Unsubscribe all price subscriptions if no symbol is provided.

##### Sample
```json
> {"method": "price.unsubscribe", "params": [], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
```

### Deals APIs
#### Deal history query

Biger MD only allows to query the latest 100 trades.

##### Syntax
```
{
  "method" : "deals.query",
  "params" : ["<symbol>", "<limit>", "<last_id>"],
  "id"     : <id>
}
```

Parameter | Required | Type    | Description
-------   | -------  | ------- | --------
symbol    | Yes      | String  | trade symbol
limit     | Yes      | Integer | the limit of deal count in response.
last_id   | Yes      | Integer | the start id


##### Sample
```json
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

#### Deal subscribe

BigerMD will publish the lastest deals on trades.

##### Syntax
```
{
  "method" : "deals.subscribe",
  "params" : ["<symbol>"],
  "id"     : <id>
}
```

Parameter | Required | Type    | Description
-------   | -------  | ------- | --------
symbol    | Yes      | String  | trade symbol

##### Sample

```json
> {"method": "deals.subscribe", "params": ["BTCUSDT"], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}
< {"method": "deals.update", "id": null, "params": ["BTCUSDT", [{"price": "8044", "type": "buy", "time": 1520438400.361028, "amount": "2", "id": 1762}, {"price": "8078", "type": "buy", "time": 1520438300.341769, "amount": "9", "id": 1761}, {"price": "8076", "type": "buy", "time": 1520438200.324909, "amount": "10", "id": 1760}, {"price": "8056", "type": "buy", "time": 1520438100.3066709, "amount": "3", "id": 1759}, {"price": "8007", "type": "buy", "time": 1520438000.2892129, "amount": "9", "id": 1758}, {"price": "8050", "type": "buy", "time": 1520437900.2736571, "amount": "6", "id": 1757}, {"price": "8074", "type": "buy", "time": 1520437800.257802, "amount": "1", "id": 1756}, {"price": "8014", "type": "buy", "time": 1520437700.239372, "amount": "4", "id": 1755}, {"price": "8054", "type": "buy", "time": 1520437600.223423, "amount": "9", "id": 1754}, {"price": "8049", "type": "buy", "time": 1520437500.2082629, "amount": "8", "id": 1753}, {"price": "8002", "type": "buy", "time": 1520437400.1939909, "amount": "2", "id": 1752}, {"price": "8000", "type": "buy", "time": 1520437300.1761429, "amount": "5", "id": 1751}, {"price": "8002", "type": "buy", "time": 1520437200.1584849, "amount": "10", "id": 1750}, {"price": "8065", "type": "buy", "time": 1520437100.142282, "amount": "5", "id": 1749}, {"price": "8099", "type": "buy", "time": 1520437000.1258199, "amount": "5", "id": 1748}, {"price": "8009", "type": "buy", "time": 1520436900.1072299, "amount": "6", "id": 1747}, {"price": "8066", "type": "buy", "time": 1520436800.0908389, "amount": "10", "id": 1746},
...
```


#### Deal unsubscribe

##### Syntax
```
{
  "method" : "deals.unsubscribe",
  "params" : ["<symbol>"],
  "id"     : <id>
}
```

Parameter | Required | Type     | Description
-------   | -------  | -------- | --------
symbol	  | No       | String   | trading symbol. Unsubscribe all deals subscriptions if no symbol is provided.


##### Sample
```json
> {"method": "deals.unsubscribe", "params": [], "id": 1516681178}
< {"result": {"status": "success"}, "error": null, "id": 1516681178}

```

### Depth APIs
#### Depth query

##### Syntax
```
{
  "method" : "depth.query",
  "params" : ["<symbol>", <limit>, "<interval>"],
  "id"     : <id>
}
```

Parameter | Required | Type    | Description
-------   | -------  | ------- | --------
symbol    | Yes      | String  | trade symbol
limit     | Yes      | Integer | depth limit
interval  | Yes      | String  | depth price precision. Use “0” for maxmium precision. Possible values: “0”，”0.1", “0.01", “0.001", “0.0001", “0.00001", “0.000001", “0.0000001", "0.00000001"


##### Sample
```json
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

#### Depth subscribe
Biger MD will publish depth data on changes. The data can be either a difference or snapshot. It is a snapshot if the bool indicator in response is true, otherwise it is a difference data. The difference data has 3 types: add, modify or delete.

In the difference data, it is add/modify if quantity of the price level is non-zero, and it is a delete if quantity is zero.

Biger MD will publish a snapshot with 60-second interval.

##### Syntax
```
{
  "method" : "depth.subscribe",
  "params" : ["<symbol>", <limit>, "<interval>" ],
  "id"     : <id>
}
```

Parameter | Required | Type    | Description
-------   | -------  | --------| --------
symbol    | Yes      | String  | trade symbol
limit     | Yes      | Integer | depth limit
interval  | Yes      | String  | depth price precision. Use “0” for maxmium precision. Possible values: “0”，”0.1", “0.01", “0.001", “0.0001", “0.00001", “0.000001", “0.0000001", "0.00000001"


##### Sample
```json
> {"method": "depth.subscribe", "params": ["BTCUSDT", 10, "0"], "id": 1516681178}
< {"error": null, "result": {"status": "success"}, "id": 1516681178}
< {"method": "depth.update", "params": [true, {"asks": [], "bids": []}, "BTCUSDT"], "id": null}
```

#### Depth unsubscribe

##### Syntax
```
{
  "method" : "depth.unsubscribe",
  "params" : ["<symbol>"],
  "id"     : <id>
}
```

Parameter | Required | Type    | Description
-------   | -------- | ------- | --------
symbol    | No       |	String | trading symbol. Unsubscribe all deals subscriptions if no symbol is provided.

##### Sample
```json
> {"method": "depth.unsubscribe", "params": [], "id": 1516681178}
< {"error": null, "result": {"status": "success"}, "id": 1516681178}
```


#### Error Handling
Biger MD will reply error messages on failures.

##### Syntax
```
{
  “error" : {
    “code"    : <code>,
    “message" : "<message>"
  	},
  "id"     : 1516681178,
  "result" : null
}
```

`Parameter` | `Type`  | `Description`
-------     | ------- | --------
code        | Integer | error code. see below for details.
message     | String  | error message. see below for details.

##### Error Code Explanation

`Error Code` | `Description`
------- | -----------------------------------------------------------------------
6001 | Invalid argument
6005 | System timeout, it is usually a internal error causing failure to handle a request in 5s.
6012 | User authentication failure
6013 | Service busy. Usually happens on busy/slow network and it causing Biger MD has to drop messages.
6014 | Throttle limit exceeded. 
6015 | Subscription limit exceeded.

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
Once given your access token, you can use our open-source testing utility to check the validity and usability of your access token.
Check it out at [BigerFX](https://github.com/biger-exchange/biger-fx)

## Appendix D - golang signature generation example
We do not yet provide a golang client, but have sample code to illustrate how to generate the signature hash using golang. Please see [golang signature example](golang-example/main.go)
