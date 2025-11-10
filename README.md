## Sistema de Navegação com Dijkstra

Como usar:

Rode o comando `go run Dijkstra.go`
Faça uma requisição POST na rota localhost:8080 com o body
```json
{
  "origem": "Rua Consolação, 200",
  "destino": "Parque Ibirapuera - Portão 2",
  "criterio": "distancia"
}
```

Curl para importar rota: 
```
curl --request POST \
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/11.6.2' \
  --data '{
  "origem": "Rua Consolação, 200",
  "destino": "Parque Ibirapuera - Portão 2",
  "criterio": "distancia"
}'
```

Possíveis entradas para origem e destino:
```
Av. Paulista, 1000
Rua Augusta, 500
Rua Consolação, 200
Rua Oscar Freire, 100
Alameda Santos, 300
Praça da República, 50
Av. Brigadeiro, 800
Shopping Iguatemi
Av. Rebouças, 1500
Av. Ibirapuera, 2000
Parque Ibirapuera - Portão 2
Av. São João, 1200
Terminal Parque Dom Pedro
Marginal Pinheiros, Km 5
```
