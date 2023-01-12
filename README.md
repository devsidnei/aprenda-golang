# Fundamentos da Linguagem

### Introdução

Exibir versão do go

```go
go **version** 
// go version go1.19.5 linux/amd64
```

Exibir variaveis ambiente

```go
go **env**
/*
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/sidnei/.cache/go-build"
GOENV="/home/sidnei/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/sidnei/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/sidnei/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.19.5"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/var/www/html/cursos/aprenda-golang/1-Pacotes/go.mod"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build2834399322=/tmp/go-build -gno-record-gcc-switches"
*/
```

Iniciar modulo em go

```go
go **mod** init modulo
// criará um arquivo go.mod no local

/*
module modulo

go 1.19
*/
```

Executar arquivos:

```go
go **run** main.go
```

Criar build

```go
go **build**
// criará arquivo executavel go
```

Pacotes

Baixar pacotes externos:

```go
go **get** github.com/badoux/checkmail
// Pacotes externos são anexados ao arquivo **go.mod**
...
require github.com/badoux/checkmail v1.2.1 // indirect
...
```

Funções em modulos criadas com letra **maiscula** são utilizaveis em qualquer modulo, bastante ser importadas, ex:

```go

package auxiliar

import (
	"fmt"
)

func **Escrever**() {
	fmt.Println("Escrevendo do pacote auxiliar")
}

// pode-se usar na main
package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.**Escrever**()

	erro := checkmail.ValidateFormat("dev.sidnei@gmail.com")
	fmt.Println(erro)
}
```

Funções escritas com letra **minúscula** só são visiveis dentro do modulo;

### Variaveis

Go é **fortemente** tipada e de tipagem **implicíta! Inferencia de tipos.**

```go
var variavel1 string = "Variável 1 - Tipo explícito"
	variavel2 := "Variável 2 - Tipo implícito"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3 - Conjunto explícito"
		variavel4 string = "Variável 4 - Conjunto explícito"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel 5 - Conjunto implicíto", 100.00

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante string = "Constante 1"

	fmt.Println(constante)
```
