# Exercicio
Mariana Santos

## Run
  O projeto baseia-se numa aplicação servidor-cliente, em que o servidor executa localmente no porto 9000 e o cliente no porto 9001
  
  Correr o servidor: 
    make run-server
  
  Correr o cliente:
    make run-client 

    #abrir browser em http://localhost:9001 
    
  Criar uma migraçao:
    .\bin\server.exe <nome da Migação>




## Nota
    É necessário ter uma instacancia da base de dados MySql, e alterar o ficheiro config nos campos <assinalados>


# Casos de uso

Caso de uso   | Requisitos    | Descricao
------------- | ------------- | -----------
login(username,password)        | Person com username e password tem de existir na dase de dados. Username e password são obrigatorios | Cliente envia uma token com o seu <usename:password> para o servidor. Aquando da validação no servidor, o cliente recebe uma token com <username:id> que deve usar para próximas comunicações. Depois é redirecionado para  para home page. Se atingir 3 tentativas em 10 segundos com credenciais erradas, tem de aguardar por 30 segundos.
addUser(name, age, username, password, family, role)   | User que adiciona tem de ter role de admin. Parametros username e password são obrigatórios. O número de pessoas que pretencem a family não pode exceder o max_persons. O username não estar já em uso. A password tem de ter entre 6 a 10 caracteres dos quais um símbolo e um número  | Um user com uma válida e com o role de admin adiciona uma nova Person com os atribuitos dados. 
deleteUser(username) | User que remove tem de ter role de admin e token válida. Tem de existir uma Person com esse username | Um user com uma token válida e com um role de admin elimina uma Person. 
updateUser(oldUsername, name, age, username,password, family, role ) | User que atualiza tem de ter role de admin e token válida. oldUsername é obrigatório. username não pode existir. password deve ter entre 6 a 10 caracteres dos quais um símbolo e um número|  Um user com uma token válida e com o role de admin atualiza os dados de outra Person. À excepção de oldeUsername, todos os campos que nao forem prenchidos permanecem como antes do update. 
createFamily(name, max_persons)| noame e max_persons são obrigatórios.  name nao pode existir.  max_persons tem de ser superior a 0  | Um user com uma token válida cria uma familia. 
readFamily(name) | name é obrigatório, e tem de existir. | Um user com uma token válida recebe uma lista com as Person que pretencem à familia.  
