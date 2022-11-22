package main

//Press Command + Shift + P  (Generate Interface Stubs)

// type Reader interface{
// 	Read(p []byte) (n int, err error)
// }

type MyReader struct {
	Name string `json:"name,omitempty"`
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

func main(){
	
	

}


