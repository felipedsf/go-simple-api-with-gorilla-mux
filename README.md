# Project: CRM Backend
*Simple Go project with GorillaMux as a router.*

---

### How to set up the project:
**go mod tidy** - Download dependencies

**go run main.go** - Start the project on port 3000

---
### How to run project unit tests:
**go test go-simple-api-with-gorilla-mux**


---

### Endpoints and how to call them:
*Inside project has a folder named __postman-collection__ with prepared to call all routes on this api.*


- Getting a single customer through: **GET /customers/{id}** 

- **curl:** `` curl --location --request GET 'http://localhost:3000/customers/1' ``


- Getting all customers through: **GET /customers**

- **curl:** `` curl --location --request GET 'http://localhost:3000/customers' ``


- Creating a customer through: **POST /customers**
- **curl:** `` curl --location --request POST 'http://localhost:3000/customers' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "Example Name",
  "role": "Example Role",
  "email": "Example Email",
  "phone": 5550199,
  "contacted": true
  }' ``


- Updating a customer through: **PUT /customers/{id}**
- **curl:** `` curl --location --request PUT 'http://localhost:3000/customers/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "PUT Example Name",
  "role": "Example Role",
  "email": "PUT Example Email",
  "phone": 5550199,
  "contacted": true
  }' ``


- Deleting a customer through: **DELETE /customers/{id}**
- **curl:** `` curl --location --request DELETE 'http://localhost:3000/customers/1' ``

