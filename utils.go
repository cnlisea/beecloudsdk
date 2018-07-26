package beecloudsdk

import "crypto/md5"

//EncryptMD5 encrypt given []byte with MD5 algorithm
func EncryptMD5(data []byte) []byte {
	if data == nil {
		return nil
	}
	encrypter := md5.New()
	encrypter.Write(data)
	return encrypter.Sum(nil)
}
