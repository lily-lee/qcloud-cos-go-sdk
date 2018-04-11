package cos

var (
	Acl_Header_Keys    = []string{"x-cos-acl", "x-cos-grant-read", "x-cos-grant-write", "x-cos-grant-full-control"}
	Uploads_Param_Keys = []string{"delimiter", "encoding-type", "prefix", "max-uploads", "key-marker", "upload-id-marker"}
)

const (
	// get service
	GET_SERVICE_HOST = "service.cos.myqcloud.com"
	STS_HOST         = "sts.api.qcloud.com"
	STS_PATH         = "/v2/index.php"
	STS_URL          = "https://" + STS_HOST + STS_PATH
	STS_ACTION       = "GetFederationToken"
)

// x-cos-acl values
const (
	COS_ACL_PRIVATE           = "private"
	COS_ACL_PUBLIC_READ       = "public-read"
	COS_ACL_PUBLIC_READ_WRITE = "public-read-write"
)

const (

	/*
	 * Standard HTTP Headers
	 */

	HOST                = "Host"
	CACHE_CONTROL       = "Cache-Control"
	CONTENT_DISPOSITION = "Content-Disposition"
	CONTENT_ENCODING    = "Content-Encoding"
	CONTENT_LENGTH      = "Content-Length"
	CONTENT_MD5         = "Content-MD5"
	CONTENT_TYPE        = "Content-Type"
	CONTENT_LANGUAGE    = "Content-Language"
	DATE                = "Date"
	ETAG                = "ETag"
	LAST_MODIFIED       = "Last-Modified"
	SERVER              = "Server"
	USER_AGENT          = "User-Agent"

	/*
	 * Cos HTTP Headers
	 */

	/** Prefix for general COS headers: x-cos- */
	COS_PREFIX = "x-cos-"

	/** COS's canned ACL header: x-cos-acl */
	COS_CANNED_ACL = "x-cos-acl"

	/** Cos's alternative date header: x-cos-date */
	COS_ALTERNATE_DATE = "x-cos-date"

	/** Prefix for COS user metadata: x-cos-meta- */
	COS_USER_METADATA_PREFIX = "x-cos-meta-"

	/** COS's version ID header */
	COS_VERSION_ID = "x-cos-version-id"

	/** COS's Multi-Factor Authentication header */
	COS_AUTHORIZATION = "Authorization"

	/** COS response header for a request's cos request ID */
	REQUEST_ID = "x-cos-request-id"

	/** COS response header for TRACE ID */
	TRACE_ID = "x-cos-trace-id"

	/** COS request header indicating how to handle metadata when copying an object */
	METADATA_DIRECTIVE = "x-cos-metadata-directive"

	/** DevPay token header */
	SECURITY_TOKEN = "x-cos-security-token"

	/** Header describing what class of storage a user wants */
	STORAGE_CLASS = "x-cos-storage-class"

	/** Header for optional server-side encryption algorithm */
	SERVER_SIDE_ENCRYPTION = "x-cos-server-side-encryption"

	/** Header for the encryption algorithm used when encrypting the object with customer-provided keys */
	SERVER_SIDE_ENCRYPTION_CUSTOMER_ALGORITHM = "x-cos-server-side-encryption-customer-algorithm"

	/** Header for the customer-provided key for server-side encryption */
	SERVER_SIDE_ENCRYPTION_CUSTOMER_KEY = "x-cos-server-side-encryption-customer-key"

	/** Header for the MD5 digest of the customer-provided key for server-side encryption */
	SERVER_SIDE_ENCRYPTION_CUSTOMER_KEY_MD5 = "x-cos-server-side-encryption-customer-key-MD5"

	/** Header for optional object expiration */
	EXPIRATION = "x-cos-expiration"

	/** Header for optional object expiration */
	EXPIRES = "Expires"

	// X_COS_COPY_SOURCE
	X_COS_COPY_SOURCE = "x-cos-copy-source"

	/** ETag matching constraint header for the copy object request */
	COPY_SOURCE_IF_MATCH = "x-cos-copy-source-if-match"

	/** ETag non-matching constraint header for the copy object request */
	COPY_SOURCE_IF_NO_MATCH = "x-cos-copy-source-if-none-match"

	/** Unmodified since constraint header for the copy object request */
	COPY_SOURCE_IF_UNMODIFIED_SINCE = "x-cos-copy-source-if-unmodified-since"

	/** Modified since constraint header for the copy object request */
	COPY_SOURCE_IF_MODIFIED_SINCE = "x-cos-copy-source-if-modified-since"

	/** Range header for the get object request */
	RANGE = "Range"

	/**Range header for the copy part request */
	COPY_PART_RANGE = "x-cos-copy-source-range"

	/** Modified since constraint header for the get object request */
	GET_OBJECT_IF_MODIFIED_SINCE = "If-Modified-Since"

	/** Unmodified since constraint header for the get object request */
	GET_OBJECT_IF_UNMODIFIED_SINCE = "If-Unmodified-Since"

	/** ETag matching constraint header for the get object request */
	GET_OBJECT_IF_MATCH = "If-Match"

	/** ETag non-matching constraint header for the get object request */
	GET_OBJECT_IF_NONE_MATCH = "If-None-Match"

	/** Encrypted symmetric key header that is used in the envelope encryption mechanism */
	CRYPTO_KEY = "x-cos-key"

	/** Initialization vector (IV) header that is used in the symmetric and envelope encryption mechanisms */
	CRYPTO_IV = "x-cos-iv"

	/** JSON-encoded description of encryption materials used during encryption */
	MATERIALS_DESCRIPTION = "x-cos-matdesc"

	/** Instruction file header to be placed in the metadata of instruction files */
	CRYPTO_INSTRUCTION_FILE = "x-cos-crypto-instr-file"

	/** Header for the original, unencrypted size of an encrypted object */
	UNENCRYPTED_CONTENT_LENGTH = "x-cos-unencrypted-content-length"

	/** Header for the optional original unencrypted Content MD5 of an encrypted object */
	UNENCRYPTED_CONTENT_MD5 = "x-cos-unencrypted-content-md5"

	/** Header for optional redirect location of an object */
	REDIRECT_LOCATION = "x-cos-website-redirect-location"

	/** Header for the optional restore information of an object */
	RESTORE = "x-cos-restore"

	/** Header for the optional delete marker information of an object */
	DELETE_MARKER = "x-cos-delete-marker"

	/**
	 * Key wrapping algorithm such as "AESWrap" and "RSA/ECB/OAEPWithSHA-256AndMGF1Padding".
	 */
	CRYPTO_KEYWRAP_ALGORITHM = "x-cos-wrap-alg"
	/**
	 * Content encryption algorithm, such as "AES/GCM/NoPadding".
	 */
	CRYPTO_CEK_ALGORITHM = "x-cos-cek-alg"
	/**
	 * Tag length applicable to authenticated encrypt/decryption.
	 */
	CRYPTO_TAG_LENGTH = "x-cos-tag-len"

	/** Region where the bucket is located. This header is returned only in HEAD bucket and ListObjects response. */
	COS_BUCKET_REGION = "x-cos-bucket-region"
)

const (
	Q_SIGN_ALGORITHM_KEY   = "q-sign-algorithm"
	Q_SIGN_ALGORITHM_VALUE = "sha1"
	Q_AK                   = "q-ak"
	Q_SIGN_TIME            = "q-sign-time"
	Q_KEY_TIME             = "q-key-time"
	Q_HEADER_LIST          = "q-header-list"
	Q_URL_PARAM_LIST       = "q-url-param-list"
	Q_SIGNATURE            = "q-signature"
	SIGN_EXPIRED_TIME      = 3600
)
