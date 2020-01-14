> # Go tools, как доставлять бинарные зависимости

## Сценарий первый

Внесение правок, которые потребуют перегенерации существующего кода инструментами, принятыми в команде.

- [gomock](https://github.com/golang/mock) (BuyerX, Verticals)
- [mockery](https://github.com/vektra/mockery) (Infomodel, Core Services)

Ситуация: кто-то из команды (BuyerX | Verticals) решил исправить проблему, которую он обнаружил в
сервисе (Infomodel | Core Services).

### Подготовка

```bash
$ source bin/activate
$ build
$ case-1
```

### Проблема

```bash
root@hash:/app# tree
root@hash:/app# server &
root@hash:/app# curl -H 'Content-Type: application-json' \
                     -H 'X-Timeout: 500ms' \
                     -d '{"hello":"world"}' \
                     http://localhost/echo?delay=1s
root@hash:/app# curl -H 'Content-Type: application-json' \
                     -H 'X-Timeout: 500ms' \
                     -d '{"hello":"world"}' \
                     http://localhost/?uri=http://localhost/echo?delay=1s
root@hash:/app# curl -H 'Content-Type: application-json' \
                     -H 'X-Timeout: 500ms' \
                     -d '{"hello":"world"}' \
                     http://localhost/?uri=http://localhost/echo?delay=1m
root@hash:/app# curl -H 'Content-Type: application-json' \
                     -H 'X-Timeout: 500ms' \
                     -d '{"hello":"world"}' \
                     http://localhost/?uri=http://localhost/echo?delay=1h
```

Из последнего примера видно, что процесс может надолго зависнуть, если
целевой сервер будет медленно отвечать. Поняв в чём проблема идём исправлять данный сценарий.

```bash
# патчим сервис и готовимся тестировать
root@hash:/app# git apply /git/context.patch

# делаем smoke-тестирование
root@hash:/app# go build -o /go/bin/server ./cmd/server.go
root@hash:/app# server &
root@hash:/app# curl -H 'Content-Type: application-json' \
                     -H 'X-Timeout: 500ms' \
                     -d '{"hello":"world"}' \
                     http://localhost/?uri=http://localhost/echo?delay=1h

# пытаемся запустить unit-тесты
root@hash:/app# go test -race ./...
root@hash:/app# go generate ./...

# google:mockgen, установка
root@hash:/app# go get github.com/golang/mock/mockgen
root@hash:/app# which mockgen
root@hash:/app# go generate ./...
root@hash:/app# go test -race ./...

# чиним тесты
root@hash:/app# git apply /git/tests.patch
root@hash:/app# go test -race ./...

# что будет если изучить go.mod и go.sum
root@hash:/app# go mod why golang.org/x/tools
root@hash:/app# go mod tidy # https://github.com/golang/mock/blob/master/go.mod#L3
root@hash:/app# go get github.com/vektra/mockery/cmd/mockery
root@hash:/app# go mod why github.com/vektra/mockery
root@hash:/app# go mod tidy # https://github.com/vektra/mockery/blob/master/go.mod#L4
```

### Занавес

```bash
$ clean
$ deactivate
```

## Сценарий второй

Внесение правок с помощью своих инструментов. Примеры:

- [golint](https://github.com/golang/lint) (Infomodel, Recommendations)
- [golangci-lint](https://github.com/golangci/golangci-lint) (BuyerX, Verticals)
- [codecoroner](https://github.com/3rf/codecoroner)<sup id="info-1">[1](#unsupported)</sup> (Recommendations)
- [goimports](https://github.com/kamilsk/go-tools/releases/tag/goimports) (Avito, but...<sup id="info-2">[2](#exception)</sup>)

Ситуация: в предыдущей серии разработчик обнаружил несколько ошибок, которые ему подсветили линтеры.
Инженерная культура не позволяет пройти мимо и не нанести пользу.

```bash
$ source bin/activate
$ case-2 golangci-lint run ./...
$ case-2
```

### Проблема

```bash
# исправляем ошибки
root@hash:/app# git apply /git/linter.patch
root@hash:/app# golangci-lint run ./...
root@hash:/app# go test -race ./...

# наносим пользу
root@hash:/app# git apply /git/good.patch
root@hash:/app# go test -race ./...

# прячем golangci-lint
root@hash:/app# git apply /git/ignore.patch
root@hash:/app# go test -race ./...
root@hash:/app# go mod tidy

# забыли про https://golang.org/pkg/go/build/#hdr-Build_Constraints
root@hash:/app# git apply /git/rewrite.patch
root@hash:/app# go mod tidy
root@hash:/app# go test -race ./...
root@hash:/app# go build -o "${GOPATH}"/bin/linter -v github.com/golangci/golangci-lint/cmd/golangci-lint
root@hash:/app# ls -la $(which linter)
```

### Занавес

```makefile
lint:
	@go build -o "$(GOPATH)"/bin/linter -v github.com/golangci/golangci-lint/cmd/golangci-lint
	@golangci-lint run ./...
```

## Сценарий третий

Обзор доступных подходов.

- [Откуда растут ноги](https://github.com/golang/go/issues/25922)
- [Автор gobin](https://github.com/myitcv/gobin)

### Примеры

- [Dapr](https://github.com/dapr/dapr)
- [GolangCI-Lint](https://github.com/golangci/golangci-lint)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)<sup id="info-3">[3](#confused)</sup>

---

- <sup id="unsupported">1</sup> <small>Больше не поддерживается.</small> [↩](#info-1)
- <sup id="exception">2</sup> <small>Редко кто запускает `avito service fmt`.</small> [↩](#info-2)
- <sup id="confused">3</sup> <small>Спорное позиционирование.</small> [↩](#info-3)
