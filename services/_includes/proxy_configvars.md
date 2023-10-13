## Environment Variables

| Name | Type | Default Value | Description |
|------|------|---------------|-------------|
| OCIS_TRACING_ENABLED<br/>PROXY_TRACING_ENABLED | bool | false | Activates tracing.|
| OCIS_TRACING_TYPE<br/>PROXY_TRACING_TYPE | string |  | The type of tracing. Defaults to '', which is the same as 'jaeger'. Allowed tracing types are 'jaeger' and '' as of now.|
| OCIS_TRACING_ENDPOINT<br/>PROXY_TRACING_ENDPOINT | string |  | The endpoint of the tracing agent.|
| OCIS_TRACING_COLLECTOR<br/>PROXY_TRACING_COLLECTOR | string |  | The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces. Only used if the tracing endpoint is unset.|
| OCIS_LOG_LEVEL<br/>PROXY_LOG_LEVEL | string |  | The log level. Valid values are: 'panic', 'fatal', 'error', 'warn', 'info', 'debug', 'trace'.|
| OCIS_LOG_PRETTY<br/>PROXY_LOG_PRETTY | bool | false | Activates pretty log output.|
| OCIS_LOG_COLOR<br/>PROXY_LOG_COLOR | bool | false | Activates colorized log output.|
| OCIS_LOG_FILE<br/>PROXY_LOG_FILE | string |  | The path to the log file. Activates logging to this file if set.|
| PROXY_DEBUG_ADDR | string | 127.0.0.1:9205 | Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed.|
| PROXY_DEBUG_TOKEN | string |  | Token to secure the metrics endpoint.|
| PROXY_DEBUG_PPROF | bool | false | Enables pprof, which can be used for profiling.|
| PROXY_DEBUG_ZPAGES | bool | false | Enables zpages, which can be used for collecting and viewing in-memory traces.|
| PROXY_HTTP_ADDR | string | 0.0.0.0:9200 | The bind address of the HTTP service.|
| PROXY_HTTP_ROOT | string | / | Subdirectory that serves as the root for this HTTP service.|
| PROXY_TRANSPORT_TLS_CERT | string | ~/.ocis/proxy/server.crt | Path/File name of the TLS server certificate (in PEM format) for the external http services. If not defined, the root directory derives from $OCIS_BASE_DATA_PATH:/proxy.|
| PROXY_TRANSPORT_TLS_KEY | string | ~/.ocis/proxy/server.key | Path/File name for the TLS certificate key (in PEM format) for the server certificate to use for the external http services. If not defined, the root directory derives from $OCIS_BASE_DATA_PATH:/proxy.|
| PROXY_TLS | bool | true | Enable/Disable HTTPS for external HTTP services. Must be set to 'true' if the built-in IDP service an no reverse proxy is used. See the text description for details.|
| OCIS_REVA_GATEWAY | string | com.owncloud.api.gateway | The CS3 gateway endpoint.|
| OCIS_GRPC_CLIENT_TLS_MODE | string |  | TLS mode for grpc connection to the go-micro based grpc services. Possible values are 'off', 'insecure' and 'on'. 'off': disables transport security for the clients. 'insecure' allows using transport security, but disables certificate verification (to be used with the autogenerated self-signed certificates). 'on' enables transport security, including server certificate verification.|
| OCIS_GRPC_CLIENT_TLS_CACERT | string |  | Path/File name for the root CA certificate (in PEM format) used to validate TLS server certificates of the go-micro based grpc services.|
| OCIS_URL<br/>OCIS_OIDC_ISSUER<br/>PROXY_OIDC_ISSUER | string | https://localhost:9200 | URL of the OIDC issuer. It defaults to URL of the builtin IDP.|
| OCIS_INSECURE<br/>PROXY_OIDC_INSECURE | bool | false | Disable TLS certificate validation for connections to the IDP. Note that this is not recommended for production environments.|
| PROXY_OIDC_ACCESS_TOKEN_VERIFY_METHOD | string | jwt | Sets how OIDC access tokens should be verified. Possible values are 'none' and 'jwt'. When using 'none', no special validation apart from using it for accessing the IPD's userinfo endpoint will be done. When using 'jwt', it tries to parse the access token as a jwt token and verifies the signature using the keys published on the IDP's 'jwks_uri'.|
| PROXY_OIDC_SKIP_USER_INFO | bool | false | Do not look up user claims at the userinfo endpoint and directly read them from the access token. Incompatible with 'PROXY_OIDC_ACCESS_TOKEN_VERIFY_METHOD=none'.|
| OCIS_CACHE_STORE<br/>PROXY_OIDC_USERINFO_CACHE_STORE | string | memory | The type of the cache store. Supported values are: 'memory', 'ocmem', 'etcd', 'redis', 'redis-sentinel', 'nats-js', 'noop'. See the text description for details.|
| OCIS_CACHE_STORE_NODES<br/>PROXY_OIDC_USERINFO_CACHE_STORE_NODES | []string | [] | A comma separated list of nodes to access the configured store. This has no effect when 'memory' or 'ocmem' stores are configured. Note that the behaviour how nodes are used is dependent on the library of the configured store.|
| OCIS_CACHE_DATABASE | string | ocis | The database name the configured store should use.|
| PROXY_OIDC_USERINFO_CACHE_TABLE | string | userinfo | The database table the store should use.|
| OCIS_CACHE_TTL<br/>PROXY_OIDC_USERINFO_CACHE_TTL | Duration | 10s | Default time to live for user info in the user info cache. Only applied when access tokens has no expiration. The duration can be set as number followed by a unit identifier like s, m or h. Defaults to '10s' (10 seconds).|
| OCIS_CACHE_SIZE<br/>PROXY_OIDC_USERINFO_CACHE_SIZE | int | 0 | The maximum quantity of items in the user info cache. Only applies when store type 'ocmem' is configured. Defaults to 512.|
| PROXY_OIDC_JWKS_REFRESH_INTERVAL | uint64 | 60 | The interval for refreshing the JWKS (JSON Web Key Set) in minutes in the background via a new HTTP request to the IDP.|
| PROXY_OIDC_JWKS_REFRESH_TIMEOUT | uint64 | 10 | The timeout in seconds for an outgoing JWKS request.|
| PROXY_OIDC_JWKS_REFRESH_RATE_LIMIT | uint64 | 60 | Limits the rate in seconds at which refresh requests are performed for unknown keys. This is used to prevent malicious clients from imposing high network load on the IDP via ocis.|
| PROXY_OIDC_JWKS_REFRESH_UNKNOWN_KID | bool | true | If set to 'true', the JWKS refresh request will occur every time an unknown KEY ID (KID) is seen. Always set a 'refresh_limit' when enabling this.|
| PROXY_OIDC_REWRITE_WELLKNOWN | bool | false | Enables rewriting the /.well-known/openid-configuration to the configured OIDC issuer. Needed by the Desktop Client, Android Client and iOS Client to discover the OIDC provider.|
| OCIS_JWT_SECRET<br/>PROXY_JWT_SECRET | string |  | The secret to mint and validate JWT tokens.|
| PROXY_ROLE_ASSIGNMENT_DRIVER | string | default | The mechanism that should be used to assign roles to user upon login. Supported values: 'default' or 'oidc'. 'default' will assign the role 'user' to users which don't have a role assigned at the time they login. 'oidc' will assign the role based on the value of a claim (configured via PROXY_ROLE_ASSIGNMENT_OIDC_CLAIM) from the users OIDC claims.|
| PROXY_ROLE_ASSIGNMENT_OIDC_CLAIM | string | roles | The OIDC claim used to create the users role assignment.|
| PROXY_ENABLE_PRESIGNEDURLS | bool | true | Allow OCS to get a signing key to sign requests.|
| PROXY_ACCOUNT_BACKEND_TYPE | string | cs3 | Account backend the PROXY service should use. Currently only 'cs3' is possible here.|
| PROXY_USER_OIDC_CLAIM | string | preferred_username | The name of an OpenID Connect claim that is used for resolving users with the account backend. The value of the claim must hold a per user unique, stable and non re-assignable identifier. The availability of claims depends on your Identity Provider. There are common claims available for most Identity providers like 'email' or 'preferred_username' but you can also add your own claim.|
| PROXY_USER_CS3_CLAIM | string | username | The name of a CS3 user attribute (claim) that should be mapped to the 'user_oidc_claim'. Supported values are 'username', 'mail' and 'userid'.|
| OCIS_MACHINE_AUTH_API_KEY<br/>PROXY_MACHINE_AUTH_API_KEY | string |  | Machine auth API key used to validate internal requests necessary to access resources from other services.|
| PROXY_AUTOPROVISION_ACCOUNTS | bool | false | Set this to 'true' to automatically provision users that do not yet exist in the users service on-demand upon first sign-in. To use this a write-enabled libregraph user backend needs to be setup an running.|
| PROXY_ENABLE_BASIC_AUTH | bool | false | Set this to true to enable 'basic authentication' (username/password).|
| PROXY_INSECURE_BACKENDS | bool | false | Disable TLS certificate validation for all HTTP backend connections.|
| PROXY_HTTPS_CACERT | string |  | Path/File for the root CA certificate used to validate the server’s TLS certificate for https enabled backend services.|
| PROXY_POLICIES_QUERY | string |  | Defines the 'Complete Rules' variable defined in the rego rule set this step uses for its evaluation. Rules default to deny if the variable was not found.|