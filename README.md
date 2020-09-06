# gospec

Aplicação Golang para realizar testes em APIs. Essa aplicação irá fazer requisições de acordo com o que foi configurado nos arquivos **spec**, um modelo deste arquivo é descrito logo abaixo. Após as requisições efetuadas a aplicação irá comparar as respostas que obteve com o que é esperado, de acordo com o arquivo **spec**.

### Executando a aplicação

Tenha instalado em sua máquina:

- GoLang v1.13+

1. Crie seus arquivos spec e coloque-os em um diretório
2. Execute o comando `go run gospec.go --test-files {diretorio_dos_arquivos}`
3. A aplicação irá efetuar as chamadas e asserçoes para cada arquivo

### Example of a APIs test file

```yaml
version: '1.0'

hosts:
  - description: Testing anyhost.com
    port: 80
    host: anyhost.com
    protocol: http
    endpointsPrefix: /api/v1
    endpoints:
      - path: /endpoint
        description: Testing anyhost/api/v1/endpoit
        request:
          headers:
            - 'Authorization:bláblá'
            - 'Content-Type:application/json'
          queryParams:
            - 'query:params'
          body:
            json:
              name: 'Eclésio Melo Júnior'
              address:
                streetName: 'Some address here'
        expected:
          headers:
            - 'Content-Type:application/json'
          status: 201
          body:
            json:
              message: 'User created'
```
