# 跨包接口兼容性

1. [When an Interface Depends on Another Interface in Go](https://medium.com/swlh/when-an-interface-depends-on-another-interface-in-go-a32d988cd21e)
1. [Example of cross-package interfaces in golang](https://gist.github.com/deinspanjer/14b34f4c2e05a9be7c5c5ce941c34ddc)

Without `interface{Worker}`:

```bash
🕙[2020-09-02 09:47:42.838] ❯ go build
# interfacedep
./main.go:13:12: cannot use asyncRunner (type async.Runner) as type runner.Runner in argument to runner.Run:
        async.Runner does not implement runner.Runner (wrong type for Run method)
                have Run(async.Worker)
                want Run(runner.Worker)
```
