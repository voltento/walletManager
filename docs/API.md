**Show User**
----
  Returns json data about a single user.

* **URL**

  /account_managing/add/

* **Method:**

  `PUT`
  
*  **URL Params**

   None

* **Data Params**

  **Required:**
  
  `{"id": string, "currency": string, "amount: float}'

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"response":"Success"}`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/users/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```