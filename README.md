# gospec

Aplicação Golang para realizar testes em APIs. Essa aplicação irá fazer requisições de acordo com o que foi configurado nos arquivos **spec**, um modelo deste arquivo é descrito logo abaixo. Após as requisições efetuadas a aplicação irá comparar as respostas que obteve com o que é esperado, de acordo com o arquivo **spec**.

### Executando a aplicação

Tenha instalado em sua máquina:

- GoLang v1.13+

1. Crie seus arquivos spec e coloque-os em um diretório
2. Execute o comando `go run gospec.go -test {diretorio_dos_arquivos}`
3. A aplicação irá efetuar as chamadas e asserçoes para cada arquivo

### Examplo de um arquivo de testes

```yaml
version: 1

testing:
  - port: 3000
    host: localhost
    protocol: http
    endpoints:
      - path: /posts
        method: GET
        expect:
          status: 200
          body:
            array: true
            json:
              - id: 1.0
                title: "json-server"
                author: "typicode"

      - path: /comments/1
        method: GET
        expect:
          status: 200
          body:
            json:
              id: 1.0
              body: "some comment"
              postId: 1.0

```
