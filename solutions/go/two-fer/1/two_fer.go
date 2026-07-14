package twofer

func ShareWith(name string) string {
	length:=len(name)
    if length == 0 {
        return "One for you, one for me."
    }else{
     	return "One for "+name+", one for me."   
    }
}
