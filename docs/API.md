**Show Accounts**
----
  Returns json data about all accounts.

* **URL**

  account_managing/add/

* **Method:**

  `GET`
  
*  **URL Params**

   None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[
    {"id":"hello, world","currency":"USD","amount":114},
    {"id":"f1","currency":"USD","amount":69},
    {"id":"f2","currency":"USD","amount":21}
    ]`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{"error": <msg>}`

* **Sample Call:**

  ```curl localhost:8080/browsing/accounts```
  
  
**Show Payments**
----
  Returns json data about all payments.

* **URL**

  browsing/payments

* **Method:**

  `GET`
  
*  **URL Params**

   None

* **Data Params**

   None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `[
    {"Id":1,"from_account_id":"f1","amount":2,"to_account_id":"f2"},
    {"Id":2,"from_account_id":"f1","amount":4,"to_account_id":"f2"}
    ]`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ error : <msg> }`
    
  * **Code:** 200 <br />
      **Content:** `{"error": "No data for `payments`"}`

* **Sample Call:**

  ```curl localhost:8080/browsing/payments```
  

**Add account**
----
  Add new account.

* **URL**

  /account_managing/add/

* **Method:**

  `PUT`
  
*  **URL Params**

   None

* **Data Params**

   {"id":<acc_id>, "currency": <string>, "amount": <float non negative>}

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"response":"Success"}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{ error : <msg> }`
    
  * **Code:** 400 <br />
      **Content:** `{"error": "Account id already exists"}`
      
  * **Code:** 400 <br />
      **Content:** `{"error": "got empty value for mandatory field <filed_name>"}`

* **Sample Call:**

  ```curl -XPUT -d'{"id":"f100", "currency": "USD", "amount": 100}' localhost:8080/account_managing/add/```