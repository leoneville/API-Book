# CRUD de livros desenvolvido em Go!

Um CRUD desenvolvido em Golang que **Cria**, **Busca**, **Atualiza** e **Deleta** o cadastro de livros no banco de dados MySQL.

## Rodando o Projeto

Para conseguir rodar o projeto localmente, basta você rodar no banco de dados MySQL o script que se encontra na pasta `src/sql/sql.sql` para criar o banco e a tabela do CRUD.

Crie também um arquivo `.env` na raiz do projeto baseado nos campos do arquivo `.env.example`.

Após ter realizado os passos acima, rode o comando `go mod tidy` para baixar os packages externos utilizados no código.

Enfim, rode o comando `go run main.go` na raiz do projeto para subir a aplicação localmente na porta que você definiu no `.env`.
