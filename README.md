# Test Matrix
<img src="https://img.shields.io/badge/golang-1.18-blue"> <img src=https://img.shields.io/badge/project-matrix-orange>
<img src="https://img.shields.io/badge/api-GIN-green">

## Scope
App builded folow the below instruction:

Nível 1:
Desenvolva uma API que esteja de acordo com os requisitos propostos acima, que seja capaz
de validar uma sequência de letras válidas.
Nível 2:
Use um banco de dados de sua preferência para armazenar as sequencias verificadas pela API.
Esse banco deve garantir a unicidade, ou seja, apenas 1 registro por sequência.
Disponibilizar um outro endpoint "/stats" que responde um HTTP GET. A resposta deve ser um
Json que retorna as estatísticas de verificações de sequências, onde deve informar a
quantidade de sequências válidas, quantidade de sequências inválidas, e a proporção de
sequências válidas em relação ao total. Segue exemplo da resposta:
{"count_valid": 40, "count_invalid": 60: "ratio": 0.4}
Nível 3:
Construir um Docker composse para executar a API, para possibilitar a execução em qualquer ambiente

## Usage
After build and run docker file

``` shell

sudo docker pull redis:alpine
docker build --pull --rm -f "dockerfile" -t matrix:latest "."
docker-compose -f matrix.yaml up

```
You can use API builded in go-gin, to test all methods on port 3001 (fixed):
POST - /sequence
Body Request sample:

```json

{
"letters": ["DUHBHB", 
"DUBUBD", 
"UBUUHU", 
"BHBDHH",
"BDDDDUB", 
"UDBDUH"]
}


```

Response Sample:
```json
{
    "is_valid": false
}
```

Get - /stats
Response Sample:

```json

{
    "count_invalid": 1,
    "count_valid": 0,
    "ration": 1
}

``` 

This two methods above have dependence to redis database.


## Author 
Fagner Ribeiro
