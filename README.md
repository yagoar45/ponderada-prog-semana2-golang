# ponderada-prog-semana2-golang

Este projeto utiliza Go, Gin e Gorm, seguindo práticas de Test-Driven Development (TDD) para a implementação e testes de uma API.

## Execução dos Testes

### Database

O projeto usa PostgreSQL. A conexão é configurada via variável de ambiente `DATABASE_URL` no arquivo `.env`:
	
```env
DATABASE_URL="postgres://postgres:password@localhost:5432/database_prog_semana2?sslmode=disable"
```

## Documentação da execução dos testes 

### Database

#### Etapa Red: Teste sem ter o código 
![alt text](image.png)

#### Etapa Green: Teste com o mínimo para passar

![alt text](image-1.png)

#### Etapa Refactor: Teste refatorado para rodar tudo

![alt text](image-2.png)

### Handlers 

#### Etapa Red: Teste sem ter o código 

![alt text](image-5.png)

### Etapa Green: Teste com o mínimo para passar

![alt text](image-3.png)


### Etapa Refactor: Teste refatorado para rodar tudo

![alt text](image-4.png)


