import binascii
hex_str = '53616d706c6520537472696e67'
hex_bytes = hex_str.encode('utf-8')
print(hex_bytes)
str_bin = binascii.unhexlify(hex_bytes)
print(str_bin.decode('utf-8'))
