## How to setup

This repo only setup using docker. To prepare the container run

```
docker-compose build
```

To install all the dependency of golang library please run

```
docker-compose run --rm app /app/bin/dep ensure
```

After the installation finish, you can run this

```
docker-compose up -d
```

## How to check if it's run?

Currently the container expose port 7500. So just access on your browser `localhost:7500`