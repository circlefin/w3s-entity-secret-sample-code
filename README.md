# README: Entity Secret Generation and Encryption

## Getting Started

To generate an entity secret and encrypt with the entity public key, and register the entity secret ciphertext follow the steps below:

1. Choose a programming language: Select the programming language you are using for your application. We provide sample code snippets for Python and Golang.

2. Follow the sample code to generate an entity secret.

3. Acquire the entity public key: Use the provided API endpoint `GET /config/entity/publicKey` to obtain the entity public key securely. This public key is required for the encryption process.

4. Input the acquired entity public key in the sample code, and run sample code to encrypt the entity secret and encode in base64, you will get **entity secret ciphertext** accordingly.

5. Register the **entity secret ciphertext** in the Configurator Page in the developer dashboard and click Register. The entity secret only needs to be generated, encrypted, encoded and registered once, unless you need to rotate the entity secret.

6. Now you can append an **entity secret ciphertext** in the API requests. Note that the encryption and encoding of entity secret should be executed every time you append in an API request to prevent replay attack. There is no need to register an updated entity secret ciphertext unless you need to rotate the entity secret. Here’s the sample API request for reference:

```bash
curl --location  --request POST 'https://api.circle.com/v1/w3s/developer/walletSets' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer [TEST_API_KEY]' \
--data '{ \
    "idempotencyKey": "b1433df1-8676-4610-b8c9-ef8b5de3c79d", \
    "name": "Entity WalletSet A", \
    "entitySecretCiphertext": "[ENTITY_SECRET_CIPHERTEXT]" \
}'
```

**Note**: Make sure to install related libraries for encryption before using the sample code. For Python sample code please first `pip install pycryptodome`.

**Note**: By running the sample code you will get a hex-encoded entity secret, please store the information carefully by yourself, as it is required for critical API requests and Circle does not store the information.

**Note**: Please refrain from directly embedding the hex-encoded entity secret within the code.
