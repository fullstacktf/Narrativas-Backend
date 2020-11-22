# Narrativas Backend

## Building: Set .env file before building

Before building the project, you must set a .env file.

Use .env.example file as reference.

    
## Endpoints

# /auth/register 

**Method:** POST

**Receives** 

```json
{   
    "username": string,
    "password": string,
    "email": string
}
```

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }


# /auth/login 

**Method:** POST

**Receives** 

```json
{   
    "username": string,
    "password": string
}
```

**Returns**

* 200 Ok, { "token": string }
* 400 Bad Request
* 422 Unprocesable Entity

# /characters/

**Method:** GET

**Header:** token

**Returns**

# /characters/$id

**Method:** GET

**Header:** token

**Returns**

# /characters/

**Method:** POST

**Header:** token

**Returns**

# /characters/$id

**Method:** PATCH

**Header:** token

**Returns**

# /characters/$id

**Method:** DELETE

**Header:** token

**Returns**

* 200 Ok
* 400 Bad Request
* 401 Unauthorized
* 403 Forbidden

