# Problem 5 - Crude Chain

## _Set Up:_

### Pre-requisites:
1. This project is built and run with:
   - **Ignite CLI @v28.7.0**
   - **Cosmos SDK @v0.50.12**

### Start the Blockchain

1. Navigate to the `crude-chain` folder.
2. Run the following command to start the blockchain:
   ```sh
   ignite chain build
   ignite chain serve
   ```
## _Resources:_  
This project is built using **Ignite** and **Go**, and the blockchain keeps track of **user details**.

### **User Data Structure**
Each user contains the following fields:
- **id**: `uint64` (unique user ID)
- **name**: `string`
- **email**: `string`
- **age**: `uint64`
- **gender**: `string`
- **creator**: `string` (address of the user creator)

## _Available Commands_

### **Create a User**
To create a new user:
```sh
crude-chaind tx crudechain create-user "Alice" "alice@example.com" "female" 25 --from alice --chain-id crudechain --yes
```
Other Examples
```sh
crude-chaind tx crudechain create-user "Gabriel" "gabriel@example.com" "male" 25 --from alice --chain-id crudechain --yes
```
```sh
crude-chaind tx crudechain create-user "John" "johnl@example.com" "male" 30 --from alice --chain-id crudechain --yes
```

### **Get All Users**
To fetch the list of all users:
```sh
crude-chaind query crudechain list-user --chain-id crudechain
```

### **Get User by ID**
Fetch a specific user using their ID (Alice):
```sh
crude-chaind query crudechain show-user 0 --chain-id crudechain
```

### **Get User by Age**
Fetch a specific user using their Age (Alice, Gabriel):
```sh
crude-chaind query crudechain users-by-age --age 25 --chain-id crudechain 
```

### **Get Users by Gender**
Filter users by gender (Gabriel, John):
```sh
crude-chaind query crudechain users-by-gender --gender male --chain-id crudechain
```

### **Update User Details**
To update an existing user's details:
```sh
crude-chaind tx crudechain update-user 0 "Alice Updated" "alice_new@example.com" "Female" 26 --from alice --chain-id crudechain --yes
```
Verify the update:
```sh
crude-chaind query crudechain show-user 0 --chain-id crudechain
```

### **Delete Specific User**
To delete a user (Alice):
```sh
crude-chaind tx crudechain delete-user 0 --from alice --chain-id crudechain --yes
```
Verify the update:
```sh
crude-chaind query crudechain list-user --chain-id crudechain  
```