# 🚀 Consensus-Breaking Change: User Data Structure Update

## **1️⃣ What Changed?**
The `User` message structure was modified to **add an `address` field** and **change the `creator` field's position**.

### **Old `User` Structure**
```proto
message User {
  uint64 id = 1;
  string name = 2;
  string email = 3;
  string gender = 4;
  uint64 age = 5;
  string creator = 6;
}
