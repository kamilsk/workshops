> # Go tools, как доставлять бинарные зависимости

## Сценарий первый

Внесение правок, которые потребуют перегенерации существующего кода инструментами, принятыми в команде.

- [gomock](https://github.com/golang/mock) (BuyerX, Verticals)
- [mockery](https://github.com/vektra/mockery) (Infomodel, Core Services)

Ситуация: кто-то из команды (BuyerX | Verticals) решил исправить проблему, которую он обнаружил в
сервисе (Infomodel | Core Services).

### Подготовка

```bash
$ activate
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
```

### Занавес

```bash
$ clean
$ deactivate
```

## Сценарий второй

Внесение правок с помощью своих инструментов. Примеры:

- [golint](https://github.com/golang/lint) (Infomodel, Recommendations)
- [codecoroner](https://github.com/3rf/codecoroner)<sup id="info-1">[1](#unsupported)</sup> (Recommendations)
- [golangci-lint](https://github.com/golangci/golangci-lint) (BuyerX, Verticals)
- [goimports](https://github.com/kamilsk/go-tools/releases/tag/goimports) (Avito, but...)

## Сценарий третий

Обзор доступных подходов.

- [Откуда растут ноги](https://github.com/golang/go/issues/25922)
- [Автор gobin](https://github.com/myitcv/gobin)
- [golangci-lint](https://github.com/golangci/golangci-lint)

<sup id="unsupported">1</sup> <small>Больше не поддерживается.</small> [↩](#info-1)
