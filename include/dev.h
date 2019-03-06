#define WASM_EXPORT extern "C"
#ifdef __cplusplus
extern "C" {
#endif
//
extern char *http_get(char *url);
extern char *http_post(char *url, char *contentType, char *body);

extern char *md5(char *);
extern char *sha1(char *);
extern char *sha512(char *);
extern char *sha256(char *);
extern char *base64_encode(char *);
extern char *base64_decode(char *);
//
#ifdef __cplusplus
}
#endif