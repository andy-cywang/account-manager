# account-manager
A web app backend manages merchant accounts

# Tech Stack

* Database: MongoDB
* ServeMux: Gorilla
* Container: Docker

# API Documentaton

**1. Login**
----
  Operations can only be performed after login.

* **URL**

  /login

* **Method:**

  `POST`
  
*  **URL Params**

   None
 
* **Data Params**

  **Required:**
  `{ username : "admin", password : "admin" }`
   
**2 Create Merchant
----
  Create an merchant

* **URL**

  /merchant/create

* **Method:**

  `POST`
  
*  **URL Params**

   None
 
* **Data Params**

   ```
      { 
        logo: string,
        members: map {
          "key": {
              "email": string,
              "name": string
        }
      }
   ```
    
**3. Get All Merchants**
----
  Get all merchants from database.

* **URL**

  /merchant/all

* **Method:**

  `GET`
  
*  **URL Params**

   None
 
* **Data Params**

   None


**4. Add a Memeber to Merchant**
----
  Add a memeber to merchant specified by merchant ID

* **URL**

  /merchant/member/add

* **Method:**

  `POST`
  
*  **URL Params**

   None
 
* **Data Params**

   ```
      { 
        "merchantID": string,
        "members": map {
            "key": {
                "email": string,
                "name": string
            }
        }
      }
   ```
    
**5. Update Merchant Member**
----
  Update a member from a merchant specified by merchant ID

* **URL**

  /merchant/member/update

* **Method:**

  `PUT`
  
*  **URL Params**

   None
 
* **Data Params**

   ```
      { 
        "merchantID": string,
        "members": map {
            "key": {
                "email": string,
                "name": string
            }
        }
      }
   ```
   
**6. Delete Merchant Member**
----
  Delete a member from a merchant specified by merchant ID

* **URL**

  /merchant/member/delete

* **Method:**

  `DELETE`
  
*  **URL Params**

   None
 
* **Data Params**

   ```
      { 
        "merchantID": string,
        "members": map {
            "key": {
                "email": string,
                "name": string
            }
        }
      }
   ```

**8. Update Logo**
----
  Update a merchant's logo

* **URL**

  /merchant/logo

* **Method:**

  `PUT`
  
*  **URL Params**

   None
 
* **Data Params**

   None
 
**7. Get Merchant Members**
----
  Get all member from a merchant specified by merchant ID

* **URL**

  /merchant/member

* **Method:**

  `GET`
  
*  **URL Params**

   merchantID
 
* **Data Params**

   ```
      { 
        "merchantID": string,
        "logo": string
      }
   ```
   
# How to Run
1. Run `docker-compose up`. This will bring up the server together with mongo DB.

2. Login at `host:8080/login` by given the admin credential `username:password - "admin":"admin"`.

3. All set, you are now free to make API calls.
