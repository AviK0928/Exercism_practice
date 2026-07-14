package hamming
import "errors"

func Distance(a, b string) (int, error) {
	if len(a)!=len(b) {
        return 0, errors.New("DNA Length must be same")
    }
    var count int
    i:=0
    for i!=len(a){
        if a[i]!=b[i] {
            count++
        }
        i++
    }
    return count, nil
}
