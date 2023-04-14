## **Company microservice**
**Requirements**: docker, docker-compose
#### **Stack**:
- Go with chi router
- Mongo db


#### Run the app
- In order to run the app:
1. 	`$ docker-compose build`
2. `$ docker-compose run`

## Api Documention

#### Routes:
- `/register` POST - Route for creating a user
- `/login` POST - Logins a user

##### Company Routes:
- `/api/v1/company/{name}` GET - Gets the company with `{name}` in uri.
- `/api/v1/company` 		POST - Creates the company with provided json payload.
- `/api/v1/company/{name}` DELETE- Deleted the company with provided name of company `{name}` uri param.
- `/api/v1/company/{name}` PATCH - Updates the company with provided `{name}` uri param and `application`
