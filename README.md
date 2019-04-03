Check the `docs` folder
**Start**

`cd .deploy &&  docker-compose build && docker-compose up`

**API**

Add new account 

`curl -XPUT -d'{"id":"f1", "currency": "USD"}' localhost:8080/account_managing/add/`

Get payments

``

Change account balance

`curl -XPUT -d'{"id":"f1", "change_amount": 90}' localhost:8080/payment/change_balance`

Transfer money

`curl -XPUT -d'{"from_account":"f1", "to_account": "f2", "change_amount": 1}' localhost:8080/payment/send_money`
