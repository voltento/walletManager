**Deploy**

Run the wallet mgr and the database in the one command with docker-compose

`docker-compose -f .deploy/docker-compose.yml up --force-recreate --build`

After that command wallet will be available by address `localhost:5432`

Check API.md for to find out usage examples 
