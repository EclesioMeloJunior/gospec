# gospec

Aplicação Golang para realizar testes em APIs

### Example of a APIs test file

```yaml
version: '1.0'

testing:
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
