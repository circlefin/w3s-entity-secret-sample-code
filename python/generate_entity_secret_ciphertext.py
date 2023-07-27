# Copyright (c) 2023, Circle Technologies, LLC. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import base64
import codecs
# Installed by `pip install pycryptodome`
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
from Crypto.Hash import SHA256

# Paste your entity public key here.
public_key_string = 'PASTE_YOUR_PUBLIC_KEY_HERE'

# If you already have a hex encoded entity secret, you can paste it here. the length of the hex string should be 64.
hex_encoded_entity_secret = 'PASTE_YOUR_HEX_ENCODED_ENTITY_SECRET_KEY_HERE'

# The following sample codes generate a distinct entity secret ciphertext with each execution.
if __name__ == '__main__':
    entity_secret = bytes.fromhex(hex_encoded_entity_secret)

    if len(entity_secret) != 32:
        print("invalid entity secret")
        exit(1)

    public_key = RSA.importKey(public_key_string)

    # encrypt data by the public key
    cipher_rsa = PKCS1_OAEP.new(key=public_key, hashAlgo=SHA256)
    encrypted_data = cipher_rsa.encrypt(entity_secret)

    # encode to base64
    encrypted_data_base64 = base64.b64encode(encrypted_data)

    print("Hex encoded entity secret:", codecs.encode(entity_secret, 'hex').decode())
    print("Entity secret ciphertext:", encrypted_data_base64.decode())
