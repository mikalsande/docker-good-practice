# Docker good practices for golang

At every stage you can check the size of the container by running `make size`.

And you can check out the contents of the container by running `make console`.


## Checkout the initial docker build
```
cat Dockerfile1
make run1
```

## Add a non-root user
```
make diff2
cat Dockerfile2
make run2
```

## Use dumb-init as entrypoint
```
make diff3
cat Dockerfile3
make run3
```

## Avoid temporary files
```
make diff4
cat Dockerfile4
make run4
```

## Avoid temporary files
```
make diff5
cat Dockerfile5
make run5
```

