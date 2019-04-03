CREATE database wallets;

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


INSERT INTO account(id, currency, amount) VALUES('user1', 'USD', 10.0);
INSERT INTO account(id, currency, amount) VALUES('user2', 'USD', 10.0);
INSERT INTO account(id, currency, amount) VALUES('user3', 'RUB', 10.0);