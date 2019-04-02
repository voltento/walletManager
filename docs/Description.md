**Deploy**

`cd .deploy &&  docker-compose build && docker-compose up`

**API**

Add new account 

`curl -XPUT -d'{"id":"hello, world", "currency": "USD"}' localhost:8080/account_managing/add/`

Get accounts

`curl localhost:8080/browsing/accounts`

Change account balance

`curl -XPUT -d'{"id":"James", "change_amount": 90}' localhost:8080/payment/change_balance`

Transfer money

`curl -XPUT -d'{"from_account":"James", "to_account": "foo", "change_amount": 90}' localhost:8080/payment/send_money`