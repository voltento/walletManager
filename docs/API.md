**Show Accounts**
----
  Returns json data about all accounts.

* **URL**

  browsing/add/

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
    **Content:** `{ "error" : <msg> }`
    
  * **Code:** 200 <br />
      **Content:** `{"error": "No data for `payments`"}`

* **Sample Call:**

  ```curl localhost:8080/browsing/payments```
  

**Add account**
----
  Add new account.

* **URL**

  /accmamaging/add/

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
    **Content:** `{ "error" : <msg> }`
    
  * **Code:** 400 <br />
      **Content:** `{"error": "Account id already exists"}`
      
  * **Code:** 400 <br />
      **Content:** `{"error": "got empty value for mandatory field <filed_name>"}`

* **Sample Call:**

  ```curl -XPUT -d'{"id":"f100", "currency": "USD", "amount": 100}' localhost:8080/accmamaging/add/```
  
  
**Transfer money**
----
  Transfer money between accounts.

* **URL**

  /payment/send_money

* **Method:**

  `PUT`
  
*  **URL Params**

   None

* **Data Params**

   {"from_account":<acc_id> "to_account": <acc_id>, "change_amount": <float>}

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"response":"Success"}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{"error": "Few balance for the operation. Account id: `f1`"}`
    
  * **Code:** 400 <br />
    **Content:** `{"error": "Can't send 0 amount"}`

* **Sample Call:**

  ```curl -XPUT -d'{"from_account":"f1", "to_account": "f2", "amount": 10.0}' localhost:8080/payment/send_money```
  

  
**Change account balance**
----
  Change account balance for amount.

* **URL**

  /payment/change_balance

* **Method:**

  `PUT`
  
*  **URL Params**

   None

* **Data Params**

   {"id":<acc_id>, "change_amount": <float positive or negative>}

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"response":"Success"}`
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** `{"error": "Few balance for the operation. Account id: `f1`"}`

* **Sample Call:**

  ```curl -XPUT -d'{"id":"f1", "change_amount": 90}' localhost:8080/payment/change_balance```