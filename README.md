# Eth collision match address

It's a high-performance program that generates an Ethereum private key and checks it is on the white list.

## Usage

Notice: This program should use combined with the model database [eth-collision / eth-address-all](https://github.com/eth-collision/eth-address-all). The model is generated by the bloom filter and it contains 180,000,000 Ethereum addresses that do have not zero amounts. 

```
git clone https://github.com/eth-collision/eth-address-all

git clone https://github.com/eth-collision/eth-collision-match-address
cd eth-collision-match-address
go run main.go
```
