Check the `docs` folder
**Start**

`cd .deploy &&  docker-compose build && docker-compose up`

**API**

Change account balance

`curl -XPUT -d'{"id":"f1", "change_amount": 90}' localhost:8080/payment/change_balance`
