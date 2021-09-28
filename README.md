# graph

<img width="180px" src="https://github.com/tindyChu/graph/blob/main/graph.png">

- a. Write a function that returns all the possible paths between A­-H.
- b. Write a function that returns the least number of hops (shortest path) between A­-H.

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
