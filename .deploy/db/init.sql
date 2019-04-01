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
