package repository

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepository() error{


	repoPath :=".TimeScape"
	_,err :=os.Stat(repoPath);
if(err==nil){
	fmt.Println("file already exists")
	return nil
}


directories := []string{filepath.Join(repoPath,"objects"),filepath.Join(repoPath,"refs"),}


for _,dir :=range directories{
	err:=os.MkdirAll(dir,0755);
	if(err!=nil){
		return err;
	}
}


headPath := filepath.Join(repoPath,"HEAD");


err= os.WriteFile(headPath,[]byte("ref: refs/main"),0644);


if(err!=nil){
	return err;
}


indexPath := filepath.Join(repoPath,"index");

_,err=os.Create(indexPath);

if(err!=nil){
	return nil;
}


fmt.Println("initialised the empty TimeScape repository");


return nil;

}