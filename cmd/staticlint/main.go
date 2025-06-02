// пакет для запуска анализаторов
// чтобы запустить локально, нужно в среде разработки создать деплоймент, запускающий ./cmd/staticlint/main.go
// нажать edit configuration, внутри окна конфигации найти строку program arguments и вставить в поле ввода ./...
// так как multichecker требует входные аргументы, без них он не будет запущен
// для расширенной конфигурации можно подсмотреть help выполнив консольную команду staticcheck -help
// для примера можно добавить в конец какого-нибудь main.go строку os.Exit(0)
package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"

	"github.com/sprint1/cmd/staticlint/analyzers"
)

func main() {
	myChecks := []*analysis.Analyzer{
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		inspect.Analyzer,
		analyzers.OsExitAnalyzer,
	}

	// добавляем анализаторы из staticcheck, которые указаны в файле конфигурации
	for _, v := range staticcheck.Analyzers {
		myChecks = append(myChecks, v.Analyzer)
	}

	multichecker.Main(
		myChecks...,
	)
}
