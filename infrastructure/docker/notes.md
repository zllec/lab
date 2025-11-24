# Notes on Docker

```bash
# create a container but not run it
docker container create --name nginx-1 docker.io/library/nginx

# run the created container
docker container run nginx-1

# display container ID not truncated
docker container ps -a --no-trunc

# inspect stopped container | more explicit command
docker container inspect CONTAINER_ID|NAME

# pause a container
docker container pause nginx-1

# running in detached 
docker run -d nginx

# check container log
docker logs container_name

# attach to container
docker attach container_name

# start an interactive shell in the running container
docker exec -it container_name /bin/bash
```

- Docker use [[Graph Driver]]  (storage driver) to extract and map the container's future root file system and store it on the host as regular folder. The extracted rootfs is then uncompressed and merged. The rootfs shows where the container's files are on the host. For this example, it's in the `MergedDir`

```bash
# assuming overlay2 storage driver was used
"GraphDriver": {
            "Data": {
                "ID": "afe4e670aeae7ce1efe3cff300ec4cfc4c5c694b6b255ae17e5e417d010a5058",
                "LowerDir": "/var/lib/docker/overlay2/f95afe8bf681d06acc1fcfced621450710535d3a27ba1fad080e59044140ab6c-init/diff:/var/lib/docker/overlay2/e335cb7ecd5c0121371a01543c965c47701f4124743fee9fa1aeee3510957245/diff:/var/lib/docker/overlay2/a6faf58a5de85f9ec6a0261a68752abfe270be31439335e65b95fa0a89353be0/diff:/var/lib/docker/overlay2/2c7e1341f1a28e6e1d9a83801ffc5e5f8c5e7eba0487c02f1915058e7155376a/diff:/var/lib/docker/overlay2/048cf0af8e6229dec368e8c162cbe2fca865089174056aa9d05cf5cae959b0da/diff:/var/lib/docker/overlay2/8ace673882f25c23fbf6b7c60a8d62cb3ded2ccae48c1f4eee919d7c3caffce4/diff:/var/lib/docker/overlay2/0ecc9d096b17467d2dd0d521f3d2858c4a3247e4ff4cc6100d50c6c8804c5518/diff:/var/lib/docker/overlay2/83b7c7bd0bde188a5e0acdf15fd0b4283891ee9310eeae36b763aea5a8a1ff9f/diff",
                "MergedDir": "/var/lib/docker/overlay2/f95afe8bf681d06acc1fcfced621450710535d3a27ba1fad080e59044140ab6c/merged",
                "UpperDir": "/var/lib/docker/overlay2/f95afe8bf681d06acc1fcfced621450710535d3a27ba1fad080e59044140ab6c/diff",
                "WorkDir": "/var/lib/docker/overlay2/f95afe8bf681d06acc1fcfced621450710535d3a27ba1fad080e59044140ab6c/work"
            },
            "Name": "overlay2"
        },
```
- You can get the container's PID using `docker inspect`
- Containers can have different layers with different PID
- Think of containers as box of processes
- Paused vs Stopped
	- paused means the container's processes are just suspended but is technically "running"
	- stopped means that all container's processes are terminated but the file system is still intact and can be restarted again later
- Restart vs Start
	- you can restart a `created` container
	- you can restart an `exited` container
	- running `docker restart` on a running container is the same as running `docker stop` and `docker start`
### Signals
- sending signal to a container not only terminate or kill, but also to trigger certain behavior
- `docker kill -s SIGUSR1 my-container-name`

### Exec
- allows you to execute commands to a running container

### Mon Nov 24 06:51:05 PM PST 2025
- Some containers images include more than one executable file. Like redis, `redis-service` also comes with `redis-cli`
- use `docker run [OPTION] IMAGE [COMMAND] [ARG...]`
  - `docker run redis:latest redis-cli --version`
  - you can run redis-cli without installing it on your machine. Neat!
- upload a png file to a redis server using `redis-cli` from the docker run command
  - first, cat the file then pipe it to the docker run command
  - second, the command must be "interactive" to keep the STDIN open
  - `cat ~/avatar.png | docker run -i redis:latest redis-cli -h HOST -p PORT -x set avatar:user123`
- there are containers that the entrypoint is a binary
  - use `docker inspect alpine/httpie:latest` to see `Cmd` and `Entrypoint`. These two properties shows the default arg when you run `docker run image`
  - use `docker run -it --entrypoint sh alpine/httpie:latest` to override the default run arguments. This will open an interactive shell in the docker image
