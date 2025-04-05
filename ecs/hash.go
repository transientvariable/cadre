package ecs

const (
	// ZeroLengthMD5 defines the MD5 digest that is generated from zero-length content, e.g. md5.New().Write([]byte{}).
	ZeroLengthMD5 = `d41d8cd98f00b204e9800998ecf8427e`

	// ZeroLengthSHA256 defines the MD5 digest that is generated from zero-length content, e.g. sha256.Sum256([]byte{}).
	ZeroLengthSHA256 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
)

// Hash represents a cryptographic hash digest.
type Hash struct {
	Adler32 string `json:"adler32,omitempty"`
	Md5     string `json:"md5,omitempty"`
	Sha1    string `json:"sha1,omitempty"`
	Sha256  string `json:"sha256,omitempty"`
	Sha512  string `json:"sha512,omitempty"`
	Ssdeep  string `json:"ssdeep,omitempty"`
}
