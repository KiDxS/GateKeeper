# GateKeeper


## Features:
- SSH key management
    - View
    - Create
    - Delete
- User authentication
    - Login
    - Change password
    - Logout

## Workflow
### SSH Management
To manage the SSH pairs we will be using IDs to access or modify these pairs.
#### Generating a SSH Key
1. User clicks the create button.
2. A form appears with the following fields:
    - label
    - password
3. Validate the fields

    **label**
    - it must not be empty

4. Generate SSH key pair based on the inputs
    If password field is empty, leave the password as blank.

5. Redirect the user to the created keypair.

#### Viewing the list of generated SSH keypairs
1. User logins to the application.
2. The dashboard will be populated with information. e.g the labels of the SSH pairs.

#### Viewing a SSH Key Pair
1. User clicks to a label of a SSH key pair.
2. Opens a new page where the user can see the key pair.

#### Deleting a SSH Key Pair
1. User views a key.
2. A delete option will appear.

### User authentication
#### Login
1. The user opens up the login page
2. Two fields will appear
    - username
    - password
3. Validate the fields

    **username**
    - it cannot be empty

    **password**
    - it cannot be empty

4. The user submits the request

    If the credentials are right:

        return a token

    if the credentials are wrong:

        return a error that their username or password is wrong.



#### Change password
1. The user clicks the change password label in the navbar which will lead to an another page.
2. There would be 3 fields that will be present
    - Current password
    - New password
    - Confirm new password
3. Validate the fields

    **Current password**
    - it cannot be empty

    **New Password & Confirm new password**
    - it cannot be empty
    - both fields should match
    - it cannot be less than 12 characters

4. The user submits the request
    If the current password is wrong:
        return a message that their current password is wrong.


## REST API Design:

**Root URL: /api/v1**

### Response example:

Content type: JSON

**Success:**
```
{
    success: true/false
    message: string
    data: {
        obj
    }
}
```

**Failed:**
```
{
   success: false
   message: string
   data: null

}
```

### Resource objects:

```
user - user-related (authentication & password)
key - SSH key pair related (view, create, delete)
```
### Endpoints:

```
Version 1:
// User object
POST /api/v1/user/login
GET /api/v1/user/logout
POST /api/v1/user/changepw

// SSH key pair object
GET /api/v1/key - Retrieve a list of SSH keypairs.
GET /api/v1/key/:id - To retrieve a single SSH pair
DELETE /api/v1/key/:id - To delete
POST /api/v1/key - To Create
```
