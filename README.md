# graph

## Quick start

#### &nbsp;&nbsp;&nbsp;Run by docker

---

Build image

```sh
docker build . -t tindychu-test
```

Interactive with the image

```sh
docker run -it tindychu-test sh
```

Get all the possible paths between A - H (could test with any points between A - H)

```sh
app possible A H
```

Get shortest path) between A - H (could test with any points between A - H)

```sh
app shortest A H
```

Remove the image

```sh
docker rmi -f tindychu-test
```
