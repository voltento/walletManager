    docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres
    
    psql -u postgres
    
    create database wallets;

    CREATE TABLE account (
     id text PRIMARY KEY,
     currency text not null,
     amount double precision not null,
     constraint account_nonnegative check (amount >= 0)
    );


    CREATE TABLE payment (
     id serial PRIMARY KEY,
     from_account text not null REFERENCES account(id),
     to_account text not null REFERENCES account(id),
     amount double precision not null
    );


**API**

Add new account 

`curl -XPUT -d'{"id":"hello, world", "currency": "USD"}' localhost:8080/account_managing/add/`

Get accounts

`curl localhost:8080/browsing/accounts`

Change account balance

`curl -XPUT -d'{"id":"James", "change_amount": 90}' localhost:8080/payment/change_balance`

Transfer money

`curl -XPUT -d'{"from_account":"James", "to_account": "foo", "change_amount": 90}' localhost:8080/payment/send_money`
