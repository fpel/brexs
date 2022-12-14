# BREXS-TEST
App final do teste Brexs


## Estrutura do código.

- `cmd`: onde fica o chamador principal da aplicação
- `domain`: onde fica as estruturas utilizadas
- `services`: onde ficam os serviços e utilitários
- `http`: onde ficam as configurações do servidor HTTP


## Arquitetura
### Packages escolhidos 
- **mux**: usado para registrar as rotas do server REST
- **logrus**: usado para logs

## Executando os testes
Na raiz do projeto, executar o comando: `go test ./...`

Só foi desenvolvido os testes sobre as rotas, pois elas próprias já vão validar as regras de negócio da aplicação.

Se eu "mockasse" as funções da pasta `services` nos testes dos serviços, aí sim eu precisaria implementar testes unitários sobre as regras de negócio.

## Executando

**Modo console**
```bash
export ASSUME_NO_MOVING_GC_UNSAFE_RISK_IT_WITH=go1.19
go run cmd/main.go console input-routes.csv
```

**Modo server**
```bash
export ASSUME_NO_MOVING_GC_UNSAFE_RISK_IT_WITH=go1.19
go run cmd/main.go server input-routes.csv
```

Para cancelar, pressione `CTRL+C`.


## API Rest

### Consultando a melhor rota

**Endpoint**
```
http://localhost:4000
```
**Método**
`GET`

**Enviar no Body**

```json
{
    "origin": "GRU",
    "destiny": "ORL"
}
```


### Salvando uma nova rota

**Endpoint**
```
http://localhost:4000
```
**Método**
`POST`

**Enviar no Body**

```json
{
    "origin": "BRL",
    "destiny": "CCG",
    "cost": 120
}
```









