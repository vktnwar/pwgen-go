# PWGen-Go

Uma ferramenta de linha de comando simples e eficiente para geração de senhas seguras, escrita em Go.

## ✨ Características

- 🔐 Geração de senhas seguras e aleatórias
- 📏 Comprimento customizável (padrão: 12 caracteres)
- 🔤 Controle sobre tipos de caracteres (maiúsculas, minúsculas, números, símbolos)
- 🔢 Geração de múltiplas senhas de uma vez
- ⚡ Rápido e leve
- 🛠️ Interface de linha de comando intuitiva

## 🚀 Instalação

### Opção 1: Build local
```bash
git clone https://github.com/seu-usuario/pwgen-go
cd pwgen-go
go build -o pwgen
```

### Opção 2: Executar diretamente
```bash
git clone https://github.com/seu-usuario/pwgen-go
cd pwgen-go
go run main.go [opções]
```

### Opção 3: Instalar globalmente
```bash
go install github.com/seu-usuario/pwgen-go@latest
```

## 📖 Uso

```bash
./pwgen [opções]
```

## ⚙️ Opções

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `-l` | int | `12` | Comprimento da senha |
| `-upper` | bool | `true` | Incluir letras maiúsculas |
| `-lower` | bool | `true` | Incluir letras minúsculas |
| `-n` | bool | `true` | Incluir números |
| `-sy` | bool | `false` | Incluir símbolos |
| `-count` | int | `1` | Número de senhas a gerar |

## 💡 Exemplos

### Senha básica (padrão)
```bash
./pwgen
# Saída: Senha com 12 caracteres, letras maiúsculas, minúsculas e números
```

### Senha longa com símbolos
```bash
./pwgen -l 20 -sy
# Saída: Senha com 20 caracteres incluindo símbolos
```

### Múltiplas senhas
```bash
./pwgen -l 16 -count 5
# Saída: 5 senhas com 16 caracteres cada
```

### Senha apenas numérica (PIN)
```bash
./pwgen -upper=false -lower=false -l 6 -count 3
# Saída: 3 PINs de 6 dígitos
```

### Senha sem símbolos
```bash
./pwgen -l 15 -sy=false -count 2
# Saída: 2 senhas de 15 caracteres sem símbolos especiais
```

### Senha apenas com letras
```bash
./pwgen -n=false -sy=false -l 10
# Saída: Senha de 10 caracteres apenas com letras
```

## 🔧 Build

### Build simples
```bash
go build -o pwgen
```

### Build para múltiplas plataformas
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o pwgen.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o pwgen

# MacOS
GOOS=darwin GOARCH=amd64 go build -o pwgen
```

### Build otimizado (tamanho reduzido)
```bash
go build -ldflags="-s -w" -o pwgen
```

## 🧪 Teste

Execute alguns comandos para testar:

```bash
# Teste básico
./pwgen

# Teste avançado
./pwgen -l 25 -sy -count 3

# Teste de validação (deve dar erro)
./pwgen -upper=false -lower=false -n=false -sy=false
```

## 📋 Requisitos

- Go 1.16 ou superior

## 🛠️ Desenvolvimento

```bash
# Clonar repositório
git clone https://github.com/seu-usuario/pwgen-go
cd pwgen-go

# Executar em modo desenvolvimento
go run main.go [opções]

# Executar testes (quando implementados)
go test ./...
```

## 📄 Licença

MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🤝 Contribuição

Contribuições são bem-vindas! Por favor, abra uma issue ou envie um pull request.

## 🚧 Próximas Funcionalidades

- [ ] Validação de força da senha
- [ ] Exportar senhas para arquivo
- [ ] Modo interativo
- [ ] Exclusão de caracteres ambíguos
- [ ] Templates de senha personalizados