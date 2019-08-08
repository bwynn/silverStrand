# silverStrand

## Docker Configuration

See [docs](https://docs.docker.com/engine/examples/postgresql_service/) for full
configuration details.

Run the PostgreSQL server container via
```sh
docker run --rm -P --name silverStrand silver_strand
```

Run `psql` in the command line
```sh
psql -h localhost -p 32782 -d docker -U docker --password
```
