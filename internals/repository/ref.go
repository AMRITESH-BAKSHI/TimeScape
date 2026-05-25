package repository
import (
	
	"os"
	"strings"
)


func CurrentCommit()string{

	refPath := ".TimeScape/refs/main";
	
	data,err :=os.ReadFile(refPath);
	
	if(err!=nil){
		return "";
	}

	return strings.TrimSpace(string(data));
}

func UpdateCurrentCommit(hash string) error{
	refPath := ".TimeScape/refs/main";
	return os.WriteFile(refPath,[]byte(hash),0644);
}


