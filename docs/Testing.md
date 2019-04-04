**Testing**

Run tests in two simple steps

- Run wallet manager instance `docker-compose -f .deploy/docker-compose.yml up --force-recreate --build`
- Run tests `go test test`