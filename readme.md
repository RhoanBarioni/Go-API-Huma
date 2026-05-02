# 🌐 Guia de Operações CRUD e Verbos HTTP

## 📌 Mapeamento de Operações CRUD

| Ação | Verbo HTTP | Descrição | Idempotente |
|------|------------|-----------|-------------|
| Create | POST | Cria um novo recurso | Não |
| Read | GET | Recupera um recurso ou lista recursos | Sim |
| Update (Total) | PUT | Substitui completamente um recurso | Sim |
| Update (Parcial) | PATCH | Atualiza parcialmente um recurso | Não |
| Delete | DELETE | Remove um recurso | Sim |

---

## 🧠 `context.Context` (Go)

O `context.Context` é usado para controlar o ciclo de vida de requisições em aplicações Go.

### 📌 Principais funções:

- **Cancelamento**
  - Encerra operações quando o cliente desconecta ou ocorre timeout.

- **Deadlines (prazos)**
  - Permite definir tempo máximo de execução (ex: 2s para resposta).

- **Metadados**
  - Transporta informações adicionais como:
    - Request ID
    - Token de autenticação
    - Dados de tracing

---

## 📡 Códigos de Status HTTP

### 🟢 2xx — Sucesso

- **200 OK**
  - Requisição bem-sucedida com retorno de dados.

- **201 Created**
  - Recurso criado com sucesso (ex: POST).

- **204 No Content**
  - Sucesso sem retorno de conteúdo.

---

### 🟡 3xx — Redirecionamento

- **301 Moved Permanently**
  - URL mudou permanentemente.

- **302 Found**
  - Redirecionamento temporário.

- **304 Not Modified**
  - Conteúdo não alterado (cache).

---

### 🔴 4xx — Erros do cliente

- **400 Bad Request**
  - Requisição inválida (JSON malformado, campos faltando).

- **401 Unauthorized**
  - Requer autenticação.

- **403 Forbidden**
  - Autenticado, mas sem permissão.

- **404 Not Found**
  - Recurso não encontrado.

- **405 Method Not Allowed**
  - Método HTTP incorreto.

- **409 Conflict**
  - Conflito com estado atual (ex: duplicidade).

- **422 Unprocessable Entity**
  - Validação falhou.

---

### 🔴 5xx — Erros do servidor

- **500 Internal Server Error**
  - Erro genérico no backend.

- **502 Bad Gateway**
  - Resposta inválida de outro servidor.

- **503 Service Unavailable**
  - Serviço indisponível.

- **504 Gateway Timeout**
  - Timeout na resposta do servidor.

---

## 💡 Resumo rápido

- **2xx** → Sucesso  
- **3xx** → Redirecionamento  
- **4xx** → Erro do cliente  
- **5xx** → Erro do servidor  