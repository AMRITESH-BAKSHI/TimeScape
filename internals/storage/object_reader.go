package storage

import(
	 
"os" 
"path/filepath")



func ReadObject(hash string) ([]byte,error){

	folder := hash[0:2];
	filename :=hash[2:];

	objectLocation := filepath.Join(".TimeScape","objects",folder,filename)


	_,err := os.Stat(objectLocation);

	if(err!=nil){
		return nil,err;
	}

	data,err :=os.ReadFile(objectLocation);

	return data,err;


}