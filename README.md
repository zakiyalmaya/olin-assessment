# Olin Assessment
This is the answer to the test conducted by Olin (PT DPAI)

## Postgresql Test

The answer to the postgresql test are located in the `postgresql_assessment` folder, which contains the query answer in `.sql` file and the simulation result in `.png` file.  

## Golang Test

The answer to the Golang test is implemented in the form of a REST API using clean architecture. Below is the API contract

1. **Question 1**

    - Request

    `GET /question1?nums={nums}&target={target}`

        curl --location 'http://localhost:8080/question1?nums=2%2C7%2C6%2C1%2C4&target=9'

    *Query Param*

    | param | description | example |
    | :---: | :---: | :---: |
    | nums | array of number | 2,7,6,5,8 |
    | target | target of sum 2 number | 9 |

    - Response

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "success",
            "answer": [
                0,
                1
            ]
        }
        ```

2. **Question 2**

    - Request

    `GET /question2?nums={nums}`

        curl --location 'http://localhost:8080/question2?nums=-1%2C1%2C0%2C-1%2C2'

    *Query Param*

    | param | description | example |
    | :---: | :---: | :---: |       
    | nums | array of number | -1,1,0,-1,2 |

    - Response

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "success",
            "answer": []
        }
        ```

3. **Question 3**

    - Request
    
    `GET /question3?sentence={sentence}&words={words}`

        curl --location 'http://localhost:8080/question3?sentence=wordgoodgoodgoodbestword&words=word%2Cgood%2Cbest%2Cword'

    *Query Param*

    | param | description | example |
    | :---: | :---: | :---: |
    | sentence | sentence in string | wordgoodgoodgoodbestword |
    | words | array of word | word,good,best,word |

    - Response

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "success",
            "answer": []
        }
        ```
