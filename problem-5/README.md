# Consensus-Breaking Change: User Data Structure Update

## a. Explain what does it mean by breaking consensus. 
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

## b. Explain why your change would break the consensus
Old nodes do not recognize the address field.
The creator field moved, causing stored data to be misinterpreted.
Requires all nodes to upgrade to remain compatible.