package hashing

import ("fmt" ;"crypto/sha1");


func GenerateSHA1(data []byte) string {

	hash :=sha1.Sum(data);

	return fmt.Sprintf("%x",hash);

}