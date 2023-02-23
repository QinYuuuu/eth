
from web3 import Web3, HTTPProvider

address = '0xa2103D4f9FA6A257B932fed53D39280242C30d26'
rpc = 'http://127.0.0.1:8545'
abi = [
        {
            "inputs": [],
            "name": "sayHello",
            "outputs": [
                {
                    "internalType": "string",
                    "name": "",
                    "type": "string"
                }
                ],
            "stateMutability": "pure",
            "type": "function"
            }
            ]
web3 = Web3(HTTPProvider(rpc))
c = web3.eth.contract(address, abi=abi)
print(c.caller.sayHello())
