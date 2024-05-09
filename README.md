# Olin Assessment
This is the answer to the test conducted by Olin (PT DPAI)

## Postgresql Test

The answer to the postgresql test are located in the `postgresql` folder, which contains the query answer in `.sql` file and the simulation result in `.png` file.  

## Golang Test

The answer to the Golang test is implemented in the form of a REST API using clean architecture and located in the `golang` folder. To run this application, follow this step:
- Clone the repository
   
    ```sh
    https://github.com/zakiyalmaya/olin-assessment.git
    ```

- Navigate to the `golang` folder using the following command

    ```sh
    cd .\golang\
    ```

- Run the appliaction using the following command

    ```sh
    go run .\main.go
    ```

**API contract**

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

        curl --location 'http://localhost:8080/question2?nums=-1%2C1%2C0%2C-1%2C2%2C-4'

    *Query Param*

    | param | description | example |
    | :---: | :---: | :---: |       
    | nums | array of number | -1,1,0,-1,2,-4 |

    - Response

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "success",
            "answer": {
                "Result": [
                    [
                        -1,
                        -1,
                        2
                    ],
                    [
                        -1,
                        0,
                        1
                    ]
                ],
                "TimeComplexity": "The time complexity is O(n^2) because there is a double nested loop. The first loop has a complexity of O(n), iterating through n array elements, while the second loop, which checks the result array, also has a complexity of O(n)."
            }
        }
        ```

3. **Question 3**

    - Request
    
    `GET /question3?sentence={sentence}&words={words}`

        curl --location 'http://localhost:8080/question3?sentence=wordgoodwordgoodbestword&words=word%2Cgood%2Cbest%2Cword'

    *Query Param*

    | param | description | example |
    | :---: | :---: | :---: |
    | sentence | sentence in string | wordgoodwordgoodbestword |
    | words | array of word | word,good,best,word |

    - Response

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "success",
            "answer": [
                8
            ]
        }
        ```
