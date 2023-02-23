from web3 import Web3, HTTPProvider


def send_data(address_1, address_2, data):
    rpc = 'http://127.0.0.1:8545'
    web3 = Web3(HTTPProvider(rpc))
    web3.eth.sendTransaction({
        'to': address_1,
        'from': address_2,
        'value': web3.toWei('2', 'milli'),
        'gas': 6600000,
        'gasPrice': web3.toWei(2, 'gwei'),
        'data': data.encode('utf-8').hex()
        })


if __name__ == 'main':
    address_1 = '0xBb952208532B149fb6FCd11816293BD6628A50b5'
    address_2 = '0xD5546d740581d336F1AA16Dc5FCe415F789638BD'
    send_data(address_1, address_2, 'hello')
