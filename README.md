# 📘 Textbook Supply Chain Management - Hyperledger Fabric

This project implements a blockchain-based **Textbook Supply Chain Management System** using **Hyperledger Fabric**. It enables transparent and secure tracking of textbooks through three main roles:

- **Author**: Full control — can create, read, update, and delete textbooks.
- **Dealer**: Can read textbooks and update only the **price**.
- **Institution**: Read-only access to all textbook records.

## 🛠 Technologies Used

- 🔗 Hyperledger Fabric v2.3
- ⚙️ Minifabric (for easy network setup)
- 🧠 Chaincode in Go
- 🐳 Docker

---

## 📦 Chaincode Functionality

| Function           | Description                              | Access Roles           |
|--------------------|------------------------------------------|------------------------|
| `CreateTextbook`   | Adds a new textbook to the ledger        | Author only            |
| `ReadTextbook`     | Retrieves textbook details by ID         | All (Author, Dealer, Institution) |
| `UpdateTextbook`   | Updates textbook info                    | Author (all fields), Dealer (price only) |
| `DeleteTextbook`   | Removes a textbook from the ledger       | Author only            |

### 📚 Textbook Data Model

```go
type Textbook struct {
  ID     string `json:"id"`
  Title  string `json:"title"`
  Author string `json:"author"`
  Year   string `json:"year"`
  Price  string `json:"price"`
}
```
###  Start network using script file

```
./startNetwork.sh
```
### Minifab commands to deploy and invoke chaincode

```
sudo chmod -R 777 vars/
```
```
sudo chmod -R 777 vars/
```
```
cp -r ../Chaincode/* vars/chaincode/CHF_PROJECT/go/
```
```
minifab ccup -n CHF_PROJECT -l go -v 1.0 -d false -r false
```
```
minifab invoke -n CHF_PROJECT -p '"CreateTextbook","book001","Blockchain Basics","Daniel Drescher","2017","45"'
```
```
minifab query -n CHF_PROJECT -p '"ReadTextbook","book001"'

```
```
minifab invoke -n CHF_PROJECT -p '"DeleteTextbook","book001"'
```
```
minifab invoke -n CHF_PROJECT -p '"UpdateTextbook","book001","Advanced Blockchain","Daniel Drescher","2019","55"'
```
