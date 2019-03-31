docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres

psql -u postgres

create database wallets;

CREATE TABLE account (
 id text PRIMARY KEY,
 currency text not null,
 amount int  not null
);


**API**

Add new account 

`curl -XPUT -d'{"id":"hello, world", "currency": "USD"}' localhost:8080/account_managing/add/`

Get accounts

`curl localhost:8080/brawsing/get_accounts`

Change account balance

`curl -XPUT -d'{"id":"James", "change_amount": 90}' localhost:8080/payment/change_balance`

