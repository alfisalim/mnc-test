# MNC TEST

****

## Description
This repository contain my result for mnc test-case golang developer.
little bit about this project, i made this project using echo framework and
clean code technique. In project structure, there is 3 layers, 
namely service layer, interactor layer and adapter layer. 

Service layer used for catch each request client
and then forward the request to the next layer.
this layer also used as receiver return from interactor layers.

Interactor layer handling all about logic on this service,
and then forward ther clearly request to adapter layers.
this layer also used as receiver return from adapter layers.

Adapter layer used for build response from request, actually this layer
used for handling all about client, either third party, database, or etc.

## Note
By default, when we hit some service that unregistered in our application, echo framework
smartly will return "message not allowed", so this automatically answaring the question number 3 point 3.
but, i keep adding middleware for validate the method.

## How to Use
1. clone this repository: 
   ```https://gitlab.com/alfisalim/mnc-test.git```
2. change directory to this project : ```cd mnc-test/```
3. open terminal, cmd, or etc and type this command 
   - ```go mod download``` for install all library needed
   - ```go run main.go``` for running service
4. we can test every service on this project using tools rest client like postman or etc.

## Documentation
| Service  | Method | Url | Endpoint | Header | Request | Response | Description |
| -------- | ------ | --- | -------- | ------ | ------- | -------- | ----------- |
| GetLanguages | GET | localhost:7001 | /languages | - | - | [response](###responseGetLanguages) | service for get all data language |
| GetLanguage | GET | localhost:7001 | /language/:id | - | - | [response](###responseGetLanguage) | service for get specific data language |
| InsertLanguage | POST | localhost:7001 | /language | - | [request](###requestInsertLanguage) | [response](###responseInsertLanguage)  | service for insert data language |
| UpdateLanguage | PUT | localhost:7001 | /language/:id | - | [request](###requestInsertLanguage) | [response](###responseUpdateLanguage) | service for update data language |
| DeleteLanguage | DELETE | localhost:7001 | /language/:id | - | - | [response](###responseGetLanguages) | service for delete data language |
| PalindromeValidation | GET | localhost:7001 | /palindrome?text= | - | - | [response](###responsePalindromeValidation) | service for check isn't text contain palindrome or not |


****
****

### responseGetLanguages
``` 
{
    "status-code": 200,
    "message": "Successfully get all data",
    "data": [
        {
            "language": "Go",
            "appeared": 2009,
            "created": [
                "Robert Griesemer",
                "Rob Pike",
                "Ken Thompson"
            ],
            "functional": true,
            "object-oriented": false,
            "relation": {
                "influenced-by": [
                    "C",
                    "Python",
                    "Pascal",
                    "Smalltalk",
                    "Modula"
                ],
                "influences": [
                    "Odin",
                    "Crystal",
                    "Zig"
                ]
            }
        }
    ]
}
```

### responseGetLanguage
``` 
{
    "status-code": 200,
    "message": "Successfully get specific data",
    "data": {
        "language": "Go",
        "appeared": 2009,
        "created": [
            "Robert Griesemer",
            "Rob Pike",
            "Ken Thompson"
        ],
        "functional": true,
        "object-oriented": false,
        "relation": {
            "influenced-by": [
                "C",
                "Python",
                "Pascal",
                "Smalltalk",
                "Modula"
            ],
            "influences": [
                "Odin",
                "Crystal",
                "Zig"
            ]
        }
    }
}
```

### requestInsertLanguage
``` 
{
    "language": "Go",
    "appeared": 2009,
    "created": [
       "Robert Griesemer",
       "Rob Pike",
       "Ken Thompson"
    ],
    "functional": true,
    "object-oriented": false,
    "relation": {
       "influenced-by": [
          "C",
          "Python",
          "Pascal",
          "Smalltalk",
          "Modula"
       ],
       "influences": [
          "Odin",
          "Crystal",
          "Zig"
       ]
    }
}
```


### responseInsertLanguage
``` 
{
    "status-code": 201,
    "message": "Successfully insert data",
    "data": [
        {
            "language": "Go",
            "appeared": 2009,
            "created": [
                "Robert Griesemer",
                "Rob Pike",
                "Ken Thompson"
            ],
            "functional": true,
            "object-oriented": false,
            "relation": {
                "influenced-by": [
                    "C",
                    "Python",
                    "Pascal",
                    "Smalltalk",
                    "Modula"
                ],
                "influences": [
                    "Odin",
                    "Crystal",
                    "Zig"
                ]
            }
        }
    ]
}
```

### responseUpdateLanguage
``` 
{
    "status-code": 200,
    "message": "Successfully update data",
    "data": {
        "language": "Go",
        "appeared": 2009,
        "created": [
            "Robert Griesemer",
            "Rob Pike",
            "Ken Thompson"
        ],
        "functional": true,
        "object-oriented": false,
        "relation": {
            "influenced-by": [
                "C",
                "Python",
                "Pascal",
                "Smalltalk",
                "Modula"
            ],
            "influences": [
                "Odin",
                "Crystal",
                "Zig"
            ]
        }
    }
}
```

### responsePalindromeValidation
``` 
{
    "status-code": 200,
    "message": "Palindrome",
    "data": "kakak"
}
```


## Author
**ALFI SALIM** | Programmer | **alfisalim.12@gmail.com**
