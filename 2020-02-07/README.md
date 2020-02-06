> # 👨‍🏫 Go module, как перевести свой PaaS-сервис

## 🏆 Мотивация

> In Go 1.14, module support will be considered ready for production use, and all users will be encouraged
> to migrate to modules from other dependency management systems.

## 🤦‍♂️ Проблема

- [GOPATH](https://github.com/golang/go/wiki/GOPATH)
- [GO15VENDOREXPERIMENT](https://github.com/golang/go/wiki/PackageManagementTools#go15vendorexperiment)
- [GO111MODULE](https://github.com/golang/go/wiki/PackageManagementTools#go111module)

## 👨‍💻 Три простых шага

0. `echo 'export GOPRIVATE=go.avito.ru' | tee -a ~/.bash_profile ~/.zshrc`.
1. [go mod init](https://tip.golang.org/pkg/cmd/go/internal/modconv/?m=all#pkg-variables).
2. Удалить legacy.
3. Обновить Dockerfile.

## 🕵️‍♂️ Обзор

- [Proposal: Versioned Go Modules](https://go.googlesource.com/proposal/+/master/design/24301-versioned-go.md)
- [Go & Versioning](https://research.swtch.com/vgo)
- [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules)
- [Go modules by example is a series of work-along guides](https://github.com/go-modules-by-example)
- [Selected Go-internal packages factored out from the standard library](https://github.com/rogpeppe/go-internal)
- [Versioned Go Prototype](https://github.com/golang/vgo)

## 👨‍🔬 Перенести в research

- [A wrapper around the go tool that automatically sets GOPATH](https://github.com/myitcv/go)
- [Make temporary edits to your Go module dependencies](https://github.com/rogpeppe/gohack)

---

made with ❤️ for everyone
