package storage

import ("os" ; "path/filepath")


func StoreObject(hash string, data []byte) error{

	folder := hash[0:2];
	filename :=hash[2:];

	objectDir := filepath.Join(".TimeScape","objects",folder);

	err := os.MkdirAll(objectDir,0755);

	if err!=nil {
		return err;
	}
	objectPath := filepath.Join(objectDir,filename)

	_,err =os.Stat(objectPath);
	if err==nil {
		return nil;
	}
	return os.WriteFile(objectPath, data, 0644)

}