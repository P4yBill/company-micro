
## **Company microservice**
**Requirements**: docker, docker-compose
#### **Stack**:
- Go with chi router
- Mongo db
- JWT Authentication
	- Access token
	- Refresh token
#### Run the app
- In order to run the app:
1. Include a `.env` file (or rename `.env.example` to `.env`)
2. 	`$ docker-compose build`
3. `$ docker-compose up`

# Api Documention

### Auth
---
- **POST** `/register` - Registers user.
 	- **Content-Type:** `x-www-form-urlencoded`
	- **Params**: 
		- `email` `Required` 
				  - *Description:* Email of the user
				  - *Validation:* `Email` 
				  - *Type*: `string`

		- `password` `Required` 
				  - *Description:* password of the user
				  - *Validation:* `Len greather than: 5` 
				  - *Type*: `string`

		- `username` `Required` 
				  - *Description:* Name of the user
				  - *Type*: `string`

	- **Response Status:**
		- OK (200) - If register successfully.
		- Bad Request (400) - If params were not valid.
		- Status Conflict (409) - If user already exists.
		- Internal Server error (500) - Something went wrong with create the user.
	- **Response Content-Type:** `application/json`
	---
- **POST** `/login` - Logins user.
 	- **Content-Type:** `x-www-form-urlencoded`
	- **Params**: 
		- `email` `Required` 
				  - *Description:* Email of the user
				  - *Validation:* `Email` 
				  - *Type*: `string`

		- `password` `Required` 
				  - *Description:* Password of the user
				  - *Validation:* `Len greather than: 5` 
				  - *Type*: `string`

	- **Response Status:**
		- Bad Request (400) - If params were invalid.
		- Internal Server error (500) - Something went wrong with retrieving the user.
		- OK (200) - If login successfully.

	- **Response Content-Type:** `application/json`
---
- **POST** `/refresh` - Logins user.
 	- **Content-Type:** `x-www-form-urlencoded`
	- **Params**: 
		- `grant_type` `Required` 
				  - *Description:* Grant type
				  - *Validation:* Need to be setted to `"refresh_token"` 
				  - *Type*: `string`

		- `refresh_token` `Required` 
				  - *Description:* Refresh token
				  - *Validation:* `Token validation` 
				  - *Type*: `string`

	- **Response Status:**
		- Bad Request (400) - If params were invalid / token ivalid.
		- Internal Server error (500) - Something went wrong with creating the token.
		- OK (200) - If token refreshed successfully
	- **Response Content-Type:** `application/json`
---
### Company Routes
---
- **GET** `/api/v1/company/{name}` - Gets the company with `{name}` in uri.
	- **Params - *URI***: 
		- `{name}` `Required` 
				  - *Description:* Gets the information of a company
				  - *Validation:* `Max Characters: 15` 
				  - *Type*: `string`
	- **Response Status:**
		- Bad Request (400) - If name was not valid.
		- Not Found (404) - If company does not exists.
		- OK (200) - If response was successful.
	- **Response Content-Type:** `application/json`
---
- **POST** `/api/v1/company` **`REQUIRES AUTH`**- Creates the company with provided json payload.
	- **Content-Type:** `x-www-form-urlencoded`
	- **Params:** 
		- `name` `Required` **`Unique`**
				  - *Description:* Description of the company
				  - *Validation:* `Max Characters: 15` 
				  - *Type*: `string`

		- `description` `Optional` 
			  - *Description:* Description of the company
			  - *Validation:*  `Max Characters: 3000`
			  - *Type*: `string`
		- `employees_count` `Required` 
			  - *Description:* Amount of employeers in the company
			  - *Type*: `number`

		- `registered` `Required` 
			  - *Description:* Amount of employeers in the company
			  - *Type*: `boolean`

		- `type` `Required` 
			  - *Description:* Amount of employeers in the company
			  - *Type*: `"Corporations" | "NonProfit" | "Cooperative" | "Sole Proprietorship"`
	- **Response Status:**
		- Bad Request (400) - Invalid input given or company exists already.
		- Internal Server Error (500) - Something went wrong while creating the company.
		- Created (201) - If successfully created

	- **Response Headers (201):**
		- **`Location:`**` /api/v1/company/{name}` - Resource location for the create company with name `{name}`

	- **Response Content-Type:** `application/json`
---
- **DELETE** `/api/v1/company/{name}` **`REQUIRES AUTH`**- Delete the company with provided name of company `{name}` uri param.
	- **Params - *URI*:** 
		- `{name}` `Required` 
				  - *Description:* Gets the information of a company
				  - *Validation:* `Max Characters: 15` 
				  - *Type*: `string`
	- **Response Status:**
		 - Bad Request (400) - If company name is not valid
		 - Not Found (404) - If company does not exists
		- No Content (204) - Deleted successfully
---
- **PATCH** `/api/v1/company{name}` **`REQUIRES AUTH`** - Updates the company with `{name}` uri param and 
	- **Content-Type:** `application/merge-patch+json`
	- **Params:** 
		- `description` `Optional` 
			  - *Description:* Description of the company
			  - *Validation:*  `Max Characters: 3000`
			  - *Type*: `string`
		- `employees_count` `Optional` 
			  - *Description:* Amount of employeers in the company
			  - *Type*: `number`

		- `registered` `Optional` 
			  - *Description:* Amount of employeers in the company
			  - *Type*: `boolean`

		- `type` `Optional` 
			  - *Description:* Amount of employeers in the company
			  - *Type*: `"Corporations" | "NonProfit" | "Cooperative" | "Sole Proprietorship"`
	- **Response Status:**
		 - Bad Request (400) - If company name or parameters given are not valid
		 - Not Found (404) - If company does not exists
		- OK (200) - Update successfully
	
	- **Response Content-Type:** `application/json`
