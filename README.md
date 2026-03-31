# capim

## Execução
- Pré requisitos: Docker e Postman instalados.
1. `cp .env.example .env` -> Copiar variaveis de ambiente
2. `docker compose up --build` -> Sobe containers
3. API está disponível em `http://localhost:3000`
OBS: Caso a API não conecte ao banco na primeira execução, aguarde alguns segundos e reinicie o container da API:
`docker compose restart api`
- Exemplo de POST request:
``curl -X POST http://localhost:3000/sellers \
  -H "Content-Type: application/json" \
  -d '{
    "document": "123",
    "legal_name": "Teste LTDA",
    "business_name": "Teste",
    "bank_account": {
      "bank_code": "001",
      "agency_number": "123",
      "account_number": "456"
    },
    "owner": [
      {
        "name": "Ana",
        "phone": "119999999",
        "email": "ana@email.com"
      }
    ]
  }'```
  
- Qualidade de código:
1. `make audit` -> Execução para monitoramento da qualidade de código
2. `make test` -> Execução de testes e visualização da % de cobertura

## Justificativa Técnica
1. Quais foram as 3 principais decisões técnicas mais importantes que você tomou e porquê?
    1) Framework HTTP: Para o desenvolvimento da API, optei por utilizar o framework Gin devido a sua simplicidade, ampla adoção no ecossistema Go e curva de aprendizado baixa, o que acelerou o desenvolvimento da API.
    2) Integração com Banco de Dados: Por identificar relacionamento entre as entidades (Clinica, Admin e Conta Bancária) optei por utilizar um banco relacional (PostgreSQL). Para acesso ao banco, utilizei GORM (ORM) que abstrai as operações SQL e realiza mapeamento das estruturas para as tabelas do baco de dados.
    3) Qualidade de Código: Utilização de Makefile para execução de linters e checagem de CVEs - Common Vulnerabilities and Exposures - (`govulncheck`) a fim de garantir a qualidade de código. Além disso, utilizei logging para observabilidade, facilitando debug e monitoramento.

2. O que você faria diferente se tivesse mais tempo?
    1) Testes automatizados no CI/CD removendo a necessidade de executar `make audit` localmente e a cada alteração
    2) Melhorar cobertura de testes unitários
    3) Melhorar lógica para atualização dos dados de Owners (Administradores)

3. Se usou IA, como ela ajudou e onde você optou por fazer diferente do que ela sugeriu? <br>
    A IA me ajudou a acelerar o desenvolvimento da API, oferecendo suporte sobre decisões técnicas de desenvolvimento, como qual framework Go utilizar.
    Dado a simplicidade da API, a IA havia sugerido de persistir os dados em memória, porém como possuo experiência com o uso de GORM com integração PostgreSQL, optei por seguir nesta direção.
 
