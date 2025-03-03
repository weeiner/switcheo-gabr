# Consensus-Breaking Change: User Data Structure Update

## What Does Breaking Consensus Mean? 
Breaking consensus refers to making changes to a blockchain that cause nodes running different versions of the software to disagree on the state of the blockchain. When consensus is broken, nodes cannot validate transactions or blocks correctly, leading to chain splits or forks.

## **What Changed?**

The `User` message structure was modified to **add an `address` field** and **change the `creator` field's position**.

### **New `User` Structure**

```proto
message User {
  uint64 id = 1;
  string name = 2;
  string email = 3;
  string address = 4;  // new field
  string gender = 5;
  uint64 age = 6;
  string creator = 7; // changed position
}
```

## Why this breaks consensus 
The field numbers changed (creator moved from 6 → 7) → Old nodes cannot read existing stored data correctly.
New field address added → Old versions don’t recognize it, leading to data inconsistencies.
A hard fork is required → All nodes must upgrade before this change can be deployed.