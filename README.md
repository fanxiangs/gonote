# gonote

## git book
### build book
```shell
#  Generate a `SUMMARY.md`
docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 gitbook book sm
# install plugin
docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 gitbook gitbook install
# run serve or build
docker run --rm -v "$PWD:/srv/gitbook" -p 4000:4000 gitbook gitbook serve
```
### build gitbook image
```shell
# build
docker build -t gitbook .
# pull
docker pull fanxiangs/gitbook
```

