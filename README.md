# Crud Gorilla Mux

Tentativa de um crud simples usando o [gorilla mux](https://github.com/gorilla/mux)

## Requisitos

Antes de começar, você precisará ter o Go instalado no seu sistema. Para instalar o Go, visite [golang.org](https://golang.org/doc/install).

## Instalação

1. Clone o repositório para o seu sistema:

```bash
    git clone https://github.com/KayoRonald/crud-mux-http
```
2. Acesse o diretório do projeto:

```bash
    cd crud-mux-http
```

3. Instalar as dependências

```bash
    go mod tidy
```
4. Inicie o servidor:

```bash
    go run main.go
```

## Endpoints

A API possui os seguintes endpoints:

| Método | Endpoint | Descrição |
| ------ | -------- | --------- |
| GET | /tasks/ | Retorna todos os taskss |
| GET | /tasks/:id | Retorna um tasks pelo ID |
| POST | /tasks | Cria um novo tasks |
| PUT | /tasks/:id | Atualiza um tasks existente |
| DELETE | /tasks/:id | Deleta um tasks pelo ID |

Os dados são retornados em formato JSON.

## Contribuição

Se você encontrar algum bug ou tiver alguma sugestão de melhoria, sinta-se à vontade para abrir uma issue ou um pull request.

<div align="center">
  <img src="https://www.pngkit.com/png/full/380-3801403_go-programming-language-logo-golang-logo-png.png" width="110" title="Golang"/>
  <img src="https://miro.medium.com/v2/resize:fit:400/1*5QBUnkCjT_m0amIHeweqGg.png" width="100" alt="Mux" title="Mux" />
</div>

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para obter mais informações.
