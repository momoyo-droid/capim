# capim
## Justificativa Técnica
1. Quais foram as 3 principais decisões técnicas mais importantes que você tomou e porquê?
    1) Framework HTTP: Para o desenvolvimento da API, optei por utilizar o framework Gin devido a sua simplicidade, ampla adoção no ecossistema Go e curva de aprendizado baixa, o que acelerou o desenvolvimento da API.
    2) Integração com Banco de Dados: Por identificar relacionamento entre as entidades (Clinica, Admin e Conta Bancária) optei por utilizar um banco relacional (PostgreSQL). Para acesso ao banco, utilizei GORM (ORM) que abstrai as operações SQL e realiza mapeamento das estruturas para as tabelas do baco de dados.
    3) Qualidade de Código: Utilização de Makefile para execução de linters e checagem de CVEs (govulncheck) a fim de garantir a qualidade de código. Além disso, utilizei logging para observabilidade, facilitando debug e monitoramento.

2. O que você faria diferente se tivesse mais tempo?
    1) Testes automatizados no CI/CD removendo a necessidade de executar `make audit` localmente e a cada alteração
    2) Melhorar cobertura de testes unitários

3. Se usou IA, como ela ajudou e onde você optou por fazer diferente do que ela sugeriu?
    A IA me ajudou a acelerar o desenvolvimento da API, oferecendo suporte sobre decisões técnicas de desenvolvimento, como qual framework Go utilizar.
    Dado a simplicidade da API, a IA havia sugerido de persistir os dados em memória, porém como possuo experiência com o uso de GORM com integração PostgreSQL, optei por seguir nesta direção.
 