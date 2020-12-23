# Narrativas Backend

## Building: Set .env file before building

Before building the project, you must set a .env file.

Use .env.example file as reference.

# Endpoints

| Entidad | Tipo | Url | Explicación|
---------|------|-----|------------|
|Characters         |  GET    |  /characters/   | Returns all the characters         |            
|Characters         |  GET    | /characters/:id     |  Return one character by id        |            
|Characters         |  POST    | /characters/    |     Save a character in the db     |            
|Characters         |  DELETE    | /characters/:id    |   Delete the character by id in the db       |            
|Characters         |  PUT    | /characters/    |  Modify a character        |            
|Characters         |  POST    | /characters/:id/sections    |   Add a new section of a character       |            
|Stories         |  GET    |  /stories/   |  Return all the stories        |            
|Stories        |  GET    |  /stories/:id   |   Return one story by id       |            
|Stories         |  POST    |  /stories/   |  Save a story in the db         |            
|Stories        |  POST    | /stories/:id/events/    |   Save a event of a given history       |            
|Stories         |  POST    |  /stories/:id/events/relations   |   Save a relation of two event of a given history        |            
|Stories        |   DELETE   | /stories/:id    |  Delete a story by id in the db        |            
|Authentification         |   POST   | /auth/register   |   Save a new user in the db       |            
|Authentification        |  POST    | /auth/login    |  checks if the user exists in the db       |            
|CDN        |   PUT   |  /upload/images/character   |  Hosts images of the characters        |            
|CDN        |   PUT   | /upload/images/story    |    Hosts images of the stories      |            
|CDN        |  STATIC    | /static    |   Root path to static files    |   



## /auth/register 

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


## /auth/login 

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

## /characters/

**Method:** GET

**Header:** token

**Returns**

  

**Method:** POST

**Header:** token

**Returns**

  

**Method:** PUT

**Header:** token

**Returns**


## /characters/$id

**Method:** GET

**Header:** token

**Returns**

  

**Method:** DELETE

**Header:** token

**Returns**

* 200 Ok
* 400 Bad Request
* 401 Unauthorized
* 403 Forbidden

