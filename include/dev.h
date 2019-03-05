#define WASM_EXPORT extern "C"
extern "C"
{
#endif
//
    extern char *http_get(char *url);
    extern char *http_post(char *url, char *contentType, char *body);

    char * md5(char *);
    char * sha1(char *);
    char * sha512(char *);
    char * sha256(char *);
    char * base64_encode(char *);
    char * base64_decode(char *)
//
#ifdef __cplusplus
}
#endif