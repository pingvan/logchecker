# logchecker

Линтер для Go, проверяющий вызовы логирования (`log/slog` и `go.uber.org/zap`) на соответствие best practices.

## Правила

| Правило | Описание |
|---------|----------|
| `LowercaseLetterRule` | Сообщение лога должно начинаться со строчной буквы |
| `EnglishLanguageRule` | Сообщение лога должно содержать только английские буквы, цифры и пробелы |
| `NoSpecialCharsRule` | Сообщение лога не должно содержать спецсимволы и эмодзи |
| `NoSensitiveDataRule` | Сообщение и ключи полей не должны содержать чувствительные данные (`password`, `token`, `secret` и др.) |

### Примеры

```go
// LowercaseLetterRule
slog.Info("Started server")   // плохо: начинается с заглавной
slog.Info("started server")   // хорошо

// EnglishLanguageRule
slog.Info("сервер запущен")   // плохо: не английские символы
slog.Info("server started")   // хорошо

// NoSpecialCharsRule
slog.Info("done!")            // плохо: содержит '!'
slog.Info("done")             // хорошо

// NoSensitiveDataRule
slog.Info("user password reset")                      // плохо: содержит "password"
zap.String("api_key", key)                            // плохо: ключ поля содержит "api_key"
slog.Info("user requested reset")                     // хорошо
```

## Установка

### Standalone

```bash
go install github.com/pingvan/logchecker/cmd/logchecker@latest
```

### Запуск

```bash
logchecker ./...
```

## Использование с golangci-lint (module plugin)

golangci-lint поддерживает подключение внешних линтеров через [module plugin system](https://golangci-lint.run/plugins/module-plugins/).

### 1. Создайте `.custom-gcl.yml` в корне проекта

```yaml
version: v1.64.0
plugins:
  - module: 'github.com/pingvan/logchecker'
    import: 'github.com/pingvan/logchecker/plugin'
    version: v1.0.0
```

> Замените `version` в `plugins` на нужную версию logchecker, а корневой `version` — на версию golangci-lint.

### 2. Соберите кастомный golangci-lint

```bash
golangci-lint custom
```

Эта команда создаст бинарник `custom-gcl` в текущей директории, содержащий logchecker.

### 3. Включите линтер в `.golangci.yml`

```yaml
linters:
  enable:
    - logchecker
```

### 4. Запустите

```bash
./custom-gcl run ./...
```

## Поддерживаемые логгеры

- `log/slog` — все методы логирования (`Info`, `Error`, `Warn`, `Debug` и т.д.)
- `go.uber.org/zap` — методы типа `*zap.Logger` (`Info`, `Error`, `Warn`, `Debug` и т.д.)

## Архитектура

```
cmd/logchecker/         — CLI точка входа (singlechecker)
plugin/                 — Точка входа для golangci-lint плагина
internal/
  logchecker/           — Основной анализатор
    analyzer.go         — Создание и запуск анализатора
    loggers.go          — Определение поддерживаемых логгеров
    utils.go            — Извлечение аргументов из вызовов
  rules/                — Реализации правил
    rule.go             — Интерфейс Rule
    registry.go         — Реестр всех правил (AllRules)
    *_rule.go           — Отдельные правила
```

## Разработка

### Запуск тестов

```bash
go test ./...
```

### Добавление нового правила

1. Создайте файл `internal/rules/my_rule.go`, реализующий интерфейс `Rule`:

```go
type Rule interface {
    Name() string
    CheckRule(pass *analysis.Pass, call *ast.CallExpr, msg ast.Expr, args []ast.Expr)
}
```

2. Зарегистрируйте правило в `internal/rules/registry.go`:

```go
var AllRules = []Rule{
    // ...существующие правила
    &MyRule,
}
```

3. Добавьте тестовые данные в `testdata/src/` и тест в `internal/rules/`.
