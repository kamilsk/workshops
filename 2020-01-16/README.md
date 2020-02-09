> # 👨‍🏫 Go tools, как доставлять бинарные зависимости
>
> If you:
>
> - want to use a go-based tool (e.g. `stringer`) while working on a module, and
> - want to ensure that everyone is using the same version of that tool while tracking the tool's version in your module's `go.mod` file
>
> then one currently recommended approach is to add a `tools.go` file to your module that includes import statements for the tools of interest
> (such as `import _ "golang.org/x/tools/cmd/stringer"`), along with a `// +build tools` build constraint. The import statements allow
> the go command to precisely record the version information for your tools in your module's `go.mod`, while the `// +build tools`
> build constraint prevents your normal builds from actually importing your tools.

## 🏆 Мотивация

В [Avito](https://tech.avito.ru) много команд, пишущих на Go свои сервисы, и каждая такая команда, как правило, использует наиболее удобный для себя
набор инструментов при решении ежедневных задач (кодогенерация, форматирование, линтинг). При отсутствии единого подхода очевидны две проблемы,
которые могут возникнуть при разработке:

- кросс-командная разработка (сценарий [первый](#-сценарий-первый) и [второй](#-сценарий-второй))
- консистентность инструментария внутри команды (версии, форки, etc)

В сообществе Go уже есть подход, который позволяет решать данные проблемы. Его я и хочу внедрить на уровне компании.

## 👨‍💻 Сценарий первый

Внесение правок, которые потребуют перегенерации существующего кода инструментами, принятыми в команде.

- [gomock](https://github.com/golang/mock) (BuyerX, Verticals)
- [mockery](https://github.com/vektra/mockery) (Infomodel, Core Services)

Ситуация: кто-то из команды (Core Services | Infomodel) решил исправить проблему,
которую он обнаружил в сервисе (BuyerX | Verticals).

### 👷‍♂️ Подготовка

```bash
$ source bin/activate
$ build
$ case-1
```

### 🤦‍♂️ Проблема

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

### 🤷‍♂️ Занавес

```bash
$ clean
$ deactivate
```

## 👨‍💻 Сценарий второй

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

### 🤦‍♂️ Проблема

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
root@hash:/app# git apply /git/retag.patch
root@hash:/app# go mod tidy
root@hash:/app# go test -race ./...
root@hash:/app# go build -o "${GOPATH}"/bin/linter -v github.com/golangci/golangci-lint/cmd/golangci-lint
root@hash:/app# ls -la $(which linter)
```

### 🤷‍♂️ Занавес

```makefile
lint:
	@go build -o "$(GOPATH)"/bin/linter -v github.com/golangci/golangci-lint/cmd/golangci-lint
	@golangci-lint run ./...
```

## 👨‍💻 Сценарий третий

```bash
$ case-3
```

### 🤵 Идеальный мир

```bash
# архитектура делает заготовку частью go.avito.ru/gl/app-boilerplate
root@hash:/app# git apply /git/boilerplate.patch

# команда прозрачно интегрирует свой инструментарий
root@hash:/app# git apply /git/team.patch

# любой разработчик в Avito понимает, как ему интегрироваться
root@hash:/app# git apply /git/contributor.patch

# точка старта
root@hash:/app# make deps

# использование
root@hash:/app# make go-generate format test lint # make generate если есть avito cli
# или
root@hash:/app# source bin/activate
root@hash:/app# which goimports golangci-lint mockgen
```

### 👨‍🎓 Примеры

- [Dapr](https://github.com/dapr/dapr)
- [GolangCI-Lint](https://github.com/golangci/golangci-lint)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)<sup id="info-3">[3](#confused)</sup>

## 🕵️‍♂️ Обзор

- [Paul Jolly](https://github.com/myitcv)
- [И его issue](https://github.com/golang/go/issues/25922)
- [Официальная позиция](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)
- [И небольшой ресёрч](https://github.com/under-the-hood/egg)

---

- <sup id="unsupported">1</sup> <small>Больше не поддерживается.</small> [↩](#info-1)
- <sup id="exception">2</sup> <small>Редко кто запускает `avito service fmt`.</small> [↩](#info-2)
- <sup id="confused">3</sup> <small>Спорное позиционирование.</small> [↩](#info-3)

made with ❤️ for everyone
