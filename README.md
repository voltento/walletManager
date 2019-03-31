docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres

psql -u postgres

create database wallets;

CREATE TABLE account (
 id text PRIMARY KEY,
 currency text not null,
 amount int  not null
);

curl -XPUT -d'{"id":"hello, world", "currency": "USD"}' localhost:8080/account_managing/add/

curl -XPUT -d'{"id":"James", "change_amount": 90}' localhost:8080/payment/change_balance

curl -XPUT -d'{"id":"jemy", "amount": 10}' localhost:8080/payment/change_balance