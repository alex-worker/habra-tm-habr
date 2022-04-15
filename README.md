# habra-tm-habr
Make some ™ on html page habrahabr

## Docker papers

```bash
  docker build ./docker
```

```bash
  docker-compose --file ./docker/docker-compose.yml build --force
```

```bash
  docker ps -s  
```

## Git papers

```shell
  git config user.email "alex[ ]worker.org"
```

```shell
  git config user.name "alex-worker"
```

```shell
  git remote set-url origin "https://alex-worker@github.com/alex-worker/habra-tm-habr.git"
```

## "net/http/pprof"

```
    r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    r.HandleFunc("/debug/pprof/profile", pprof.Profile)
    r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    r.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

```bash
  go tool pprof main http://127.0.0.1:9090/debug/pprof/profile
```
