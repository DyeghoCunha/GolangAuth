# Auth Microservice (Golang Puro + PostgreSQL + Docker + Kubernetes)

Microserviço de autenticação genérico, criado para ser reutilizável em qualquer plataforma. Desenvolvido 100% em Go puro, com PostgreSQL, JWT, suporte a OAuth2, envio de e-mail, autenticação 2FA com Google Authenticator, containerizado com Docker e preparado para deploy com Kubernetes.

---

## 🚀 Funcionalidades

* Cadastro e login com email e senha
* Login com Google e GitHub (OAuth2)
* Geração e validação de JWT (access/refresh tokens)
* Recuperação de senha por e-mail
* Confirmação de conta por e-mail
* Middleware para autenticação via JWT
* Autenticação de dois fatores (2FA) opcional com Google Authenticator

---

## 🔧 Tecnologias Utilizadas

* **Golang Puro** (sem frameworks externos)
* **PostgreSQL** (armazenamento de usuários e tokens)
* **Docker** (ambiente local e build)
* **Kubernetes** (deploy escalável)
* **JWT (github.com/golang-jwt/jwt)**
* **bcrypt** para hash de senhas
* **net/smtp** para envio de e-mail
* **OTP (github.com/pquerna/otp)** para TOTP/2FA
* **Google Authenticator** (compatível)

---

## 📂 Estrutura de Pastas

```plaintext
auth/
├── cmd/
├── config/
├── controller/
├── service/
├── model/
├── repository/
├── middleware/
├── util/
├── db/
├── docker/
├── k8s/
```

---

## 📃 Variáveis de Ambiente (.env.example)

```env
PORT=8080
JWT_SECRET=secreta123
JWT_EXPIRATION_MINUTES=15
REFRESH_TOKEN_SECRET=refreshsecreta
REFRESH_TOKEN_EXPIRATION_HOURS=720

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=auth_service

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=you@gmail.com
SMTP_PASS=senha_aplicativo

FRONTEND_URL=https://seusite.com
```

---

## 🛠️ Como Rodar Localmente

```bash
# 1. Clone o projeto
$ git clone https://github.com/dyeghocunha/auth-service.git && cd auth-service

# 2. Copie o .env de exemplo
$ cp .env.example .env

# 3. Suba com Docker Compose
$ docker-compose up --build
```

---

## 🌐 Deploy com Kubernetes

* Edite os manifests em `k8s/` com suas credenciais
* Aplique com `kubectl apply -f k8s/`

---

## ✅ TODO (Próximas Etapas)

Perfeito, O Grande Autista Majestoso 🧠✨
Com base nos endpoints da função `SetupRoutes()`, aqui está a documentação objetiva e organizada **por ordem de execução lógica no fluxo da aplicação**:

---

## 📘 **Documentação de Endpoints da Autenticação**

### 🔹 1. **Registrar Usuário**

* **POST `/register`**
* Cria um novo usuário com email e senha.
* ⚠️ Senha será automaticamente convertida em hash e salva.

---

### 🔹 2. **Login (sem ou com 2FA)**

* **POST `/login`**
* Verifica email e senha:

    * Se `is_two_fa_enabled = false`: gera access e refresh token direto.
    * Se `is_two_fa_enabled = true`: retorna `partial_token` e exige verificação TOTP.

---

### 🔹 3. **Ativar 2FA (autenticado)**

* **POST `/enable-2fa`**
* Gera secret TOTP para o usuário e ativa o campo `is_two_fa_enabled`.
* ⚠️ Requer `Authorization: Bearer access_token`

---

### 🔹 4. **Gerar QR Code do 2FA**

* **GET `/generate-qr`**
* Gera a imagem do QR Code para o usuário escanear no Google Authenticator.
* ⚠️ Internamente usa o segredo do 2FA já salvo.

---

### 🔹 5. **Verificar código do 2FA (com partial token)**

* **POST `/verify-2fa`**
* Recebe código TOTP + partial\_token.
* Se o código estiver correto, gera access e refresh tokens completos.

---

### 🔹 6. **Consultar dados do usuário logado**

* **GET `/profile`**
* Retorna dados básicos do usuário autenticado.
* ⚠️ Requer access token.

---

### 🔹 7. **Rota normal protegida**

* **GET `/normal-route`**
* Exemplo de rota que exige token JWT válido.

---

### 🔹 8. **Rota sensível (exige 2FA)**

* **GET `/sensitive-data`**
* Só acessível por usuários com 2FA verificado via middleware.

---

### 🔹 9. **Buscar todos os usuários**

* **GET `/user`**
* Lista de usuários cadastrados (pode ser restrita no futuro).

---

### 🔹 10. **Health Check**

* **GET `/health`**
* Apenas para checar se a API está online.

---

Se quiser, posso formatar isso num `README.md` pronto para commit também.
Posso seguir com a documentação dos controllers e models se desejar 🧩


---

## ⚖️ Licença

MIT — use à vontade e contribua se quiser ✨

---

## 💡 Observações

Este serviço é desacoplado e pode ser integrado a qualquer plataforma frontend via REST. Ideal para marketplaces, sistemas de membros, dashboards administrativos e apps SaaS.

perfeito. 
