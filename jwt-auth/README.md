# jwt-auth

JWT allows distributed stateless services to trust each other by exchanging identity details via an encrypted (JWE) or signed (JWS) token. This example demos the usage of a JWT signed token with symmetric encryption over a secure channel (TLS).

JWT tokens can also carry user context, allowing the upstream services to be aware of the originating user call.

Token encryption (signature) can be symetric (HMAC) or assymetric (RSA).

[OWASP JWT cheatsheet](https://cheatsheetseries.owasp.org/cheatsheets/JSON_Web_Token_Cheat_Sheet_for_Java.html#introduction)

[JWS vs JWE and JWT Compact vs JSON serialization](https://medium.facilelogin.com/jwt-jws-and-jwe-for-not-so-dummies-b63310d201a3)