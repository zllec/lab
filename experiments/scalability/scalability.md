# Wed Nov 26 07:54:39 PM PST 2025
- Just got an idea how can i demonstrate ways of scaling a simple application where a participant can signup as proof of attendance.
- Rough idea would be 3 phases: 1000 users, 100k users, and 1M users.
- Technology of choice would be: Go, PostgreSQL, Docker
- For the phase 1: Monolith, go server -> postgresql
- I will have to find a way to load test this, make sure to set constraints on the resources for both db and server

```yaml
db:
  image: postgres:16-alpine
  volumes:
    - postgres_data:/var/lib/postgresql/data
volumes:
  postgres_data:
```

- `postgres_data` means a named volume. This lives in the host `/var/lib/docker/volumes/project-name_volume-name`
- the volume name `project-name_volume-name` is composed by Docker itself. If not specified, Docker will use the project directory name as prefix and second half is the volume name from the docker-compose.yml

- Done setting up server + postgresq. Server is able to write to postgresql
- To do:
  - [ ] Limit server and db size - CPU, Memory
  - [ ] Add monitoring - Prometheus and grafana
  - [ ] Perf test 
