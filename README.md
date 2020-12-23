# Narrativas Backend

## Building: Set .env file before building

Before building the project, you must set a .env file.

Use .env.example file as reference.

# Endpoints

| Id | Entidad | Tipo | Url | Explicación|
|----|---------|------|-----|------------|
1|Characters         |  GET    |  /characters/   | Returns all the characters         |            
2|Characters         |  GET    | /characters/:id     |  Return one character by id        |            
3|Characters         |  POST    | /characters/    |     Save a character in the db     |            
4|Characters         |  DELETE    | /characters/:id    |   Delete the character by id in the db       |            
5|Characters         |  PUT    | /characters/    |  Modify a character        |            
6|Characters         |  POST    | /characters/:id/sections    |   Add a new section of a character       |            
7|Stories         |  GET    |  /stories/   |  Return all the stories        |            
8|Stories        |  GET    |  /stories/:id   |   Return one story by id       |            
9|Stories         |  POST    |  /stories/   |  Save a story in the db         |            
10|Stories        |  POST    | /stories/:id/events/    |   Save a event of a given history       |            
11|Stories         |  POST    |  /stories/:id/events/relations   |   Save a relation of two event of a given history        |            
12|Stories        |   DELETE   | /stories/:id    |  Delete a story by id in the db        |            
13|Authentification         |   POST   | /auth/register   |   Save a new user in the db       |            
14|Authentification        |  POST    | /auth/login    |  checks if the user exists in the db       |            
15|CDN        |   PUT   |  /upload/images/character   |  Hosts images of the characters        |            
16|CDN        |   PUT   | /upload/images/story    |    Hosts images of the stories      |            
17|CDN        |  STATIC    | /static    |   Root path to static files    |   



## 1. /characters/

**Method:** GET

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 2. /characters/:id

**Method:** GET

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 3. /characters/

**Method:** POST

**Receives** 

```json
{
    "name": "Character",
    "biography": "Biography",
    "image": "image678.png",
    "sections": [
        {
            "title": "Section 1",
            "fields": [
                {
                    "name": "Field 1",
                    "value": "255",
                    "description": "test"
                },
                {
                    "name": "Field 2",
                    "value": "333",
                    "description": "test"
                }
            ]
        },
        {
            "title": "Field 3",
            "fields": [
                {
                    "name": "Test",
                    "value": "255",
                    "description": "test"
                }
            ]
        }
    ]
}
```

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 4. /characters/:id

**Method:** DELETE

**Returns**

* 200 Ok
* 400 Bad Request, { "error": msg }
* 401 Unauthorized
* 403 Forbidden

## 5. /characters/

**Method:** PUT

**Receives** 

```json
{
    "id": 4,
    "userid": 2,
    "name": "ESTOHASIDOCAMBIADO",
    "biography": "123",
    "image": "ESTOHASIDOCAMBIADO.png",
    "sections": [
        {
            "id": 7,
            "character_id": 6,
            "title": "prueba",
            "fields": [
                {
                    "id": 10,
                    "section_id": 7,
                    "name": "Titulo 1",
                    "value": "255",
                    "description": "prueba"
                },
                {
                    "id": 11,
                    "section_id": 7,
                    "name": "Titulo 2",
                    "value": "333",
                    "description": "prueba"
                },
                {
                    "id": 8,
                    "section_id": 5,
                    "name": "Titulo 2",
                    "value": "333",
                    "description": "prueba"
                },
                {
                    "id": 8,
                    "section_id": 5,
                    "name": "Titulo 2",
                    "value": "333",
                    "description": "prueba"
                }
            ]
        },
        {
            "id": 8,
            "character_id": 6,
            "title": "Titulo 3",
            "fields": [
                {
                    "id": 12,
                    "section_id": 8,
                    "name": "Titulo 1",
                    "value": "255",
                    "description": "prueba"
                }
            ]
        }
    ]
}
```

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 6. /characters/:id/sections

**Method:** POST

**Receives** 

```json
"title": "Section 1",
            "fields": [
                {
                    "name": "Field 1",
                    "value": "255",
                    "description": "test"
                },
                {
                    "name": "Field 2",
                    "value": "333",
                    "description": "test"
                }
            ]
```

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 7. /stories/

**Method:** GET

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 8. /stories/:id

**Method:** GET

**Receives** 

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 9. /stories/

**Method:** POST

**Receives** 

```json
{
    "initial_event_id": null,
    "image": "image2.jpg",
    "title": "title",
    "description": "blabla"
}
```

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 10. /stories/:id/events/

**Method:** POST

**Receives** 

```json
{
    "event_title": "prueba222",
    "event_description": "evento"
}
```

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }

## 11. /stories/:id/events/relations

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

## 12. /stories/:id

**Method:** DELETE

**Returns**

* 200 Ok
* 400 Bad Request, { "error": msg }
* 401 Unauthorized
* 403 Forbidden


## 13. /auth/register 

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


## 14. /auth/login 

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

## 15. /upload/images/character

**Method:** PUT

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

## 16. /upload/images/story

**Method:** PUT

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

## 17. /static

**Method:** STATIC

**Receives** 

Imagen binary

**Returns**

* 201 Created
* 400 Bad Request, { "error": msg }
* 422 Unprocesable Entity, { "error": msg }