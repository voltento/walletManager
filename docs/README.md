**Description**<br />
This is wallet manager. It provide basic functional for storage wallets with different currency and operations with them <br />
This software solves the problem of storing wallets users wallets and transfering money between them

**Functional**<br />
- Account creating
- Account balance changing
- Transfer money between accounts with the same currency

**Restrictions**<br />
- Account balance can't be negative
- Account currency isn't able to be changed

**API**<br />
Check the API.md file

**Deploy**

Run the wallet mgr and the database in the one command with docker-compose

`docker-compose -f .deploy/docker-compose.yml up --force-recreate --build`

After that command wallet will be available by address `localhost:5432`

Check API.md for to find out usage examples 

If you wanna use PostgreSQL without provided docker file, you have to pre initialize the database with the sql queries
with are described in the file `.deploy/db/init.sql`

**Testing**

Run tests in two simple steps

- Run wallet manager instance `docker-compose -f .deploy/docker-compose.yml up --force-recreate --build`
- Run tests `go test ./test`


**Dependencies**
- github.com/go-kit - programming toolkit for building microservices
- github.com/gorilla - http/ip network library
- github.com/go-pg - PostgreSQL client
 
