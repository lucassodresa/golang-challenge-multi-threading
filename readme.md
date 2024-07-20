# CEP Search Challenge with Multithreading and APIs

## Description

In this challenge, you will utilize concepts of Multithreading and APIs to fetch the fastest result between two distinct APIs.

The two requests will be made simultaneously to the following APIs:

- [BrasilAPI](https://brasilapi.com.br/docs#tag/CEP/paths/~1cep~1v1~1{cep}/get)
- [ViaCEP](http://viacep.com.br/)

## Requirements

1. **Fastest Response:**

   - The fastest response between the two APIs should be accepted, and the slower response discarded.

2. **Display the Result:**

   - The result of the request should be displayed in the terminal with the address data, along with information on which API provided the response.

3. **Time Limit:**
   - The response time should be limited to 1 second. Otherwise, a timeout error message should be displayed.
