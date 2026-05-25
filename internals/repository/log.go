package repository


import (
	"fmt"
	"TimeScape/internals/objects"
	"TimeScape/internals/storage"
)




func Loghistory()error{

	currentHash := CurrentCommit();
	
	for currentHash!=""{
		data,err :=storage.ReadObject(currentHash);

		if(err!=nil){
			return err;
		}

		commit :=objects.ParseCommitData(data);

		fmt.Println("commit: ",currentHash);
		fmt.Println("message: ",commit.MSG);
		fmt.Println("Tree: ",commit.TreeHash);
		fmt.Println()
		
		currentHash=commit.ParentHash;
	}
return nil;
}