# jwt-auth

JWT allows distributed stateless services to trust each other by exchanging identity details via an encrypted (JWE) or signed (JWS) token. This example demos a JWE token.

JWT tokens can also carry user context, allowing the upstream services to be aware of the originating user call. 

A JWE token contains a payload encrypted with a key that is also encrypted and included in the token. Used when the claims and user context are sensitive.

Token signature can be symetric (HMAC) or assymetric (RSA).

[OWASP JWT cheatsheet](https://cheatsheetseries.owasp.org/cheatsheets/JSON_Web_Token_Cheat_Sheet_for_Java.html#introduction)

[JWS vs JWE and JWT Compact vs JSON serialization](https://medium.facilelogin.com/jwt-jws-and-jwe-for-not-so-dummies-b63310d201a3)

[Critical vulnerabilities in JSON Web Token libraries](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/)

[JSON Web Encryption is a Foot-Gun](https://paragonie.com/blog/2017/03/jwt-json-web-tokens-is-bad-standard-that-everyone-should-avoid)
