## How to build

```bash
pushd libcxx
mkdir build
pushd build
cmake ..
make
popd
popd
```

## Proof of concept

```bash
go test -race ./...
```
