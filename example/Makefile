name?=example-server
version?=0.0.1
license?=false
package=server

# 检查操作系统
ifeq ($(OS),Windows_NT)
    package := $(package).exe
endif

clean:
	@if [ -d "./$(name)" ]; then rm -rf "./$(name)"; fi

build:clean
	@echo "Start build ${package}..."
	@mkdir -p "$(name)/resource"
	@cp -r "./resource" "$(name)"
	@go build -ldflags "-X 'github.com/lrayt/small-sparrow/core.AppName=$(name)' -X 'github.com/lrayt/small-sparrow/core.Version=$(version)' -X 'github.com/lrayt/small-sparrow/core.VerifyLicense=$(license)'" -a -v -o "$(name)/${package}" "./cmd/server"
	@tar zcvf "$(name)-v$(version).tar.gz" "$(name)"
	@rm -rf "./$(name)"
	@echo "Packaging successful!"