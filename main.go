package main

import (
	"fmt"
	"net/http"
	pessoaProto "../pessoa/pb/pessoa"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//encode(r)
		decode(r)

	})

	http.ListenAndServe(":3000", nil)
}

func encode(r *http.Request) {

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// pessoa := &pb.Pessoa{}

	// data, err := proto.Marshal(pessoa)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

func decode(r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	bytes := make([]byte, len(data))

	p := &pb.pessoaProto{}

	
			if err := proto.Unmarshal(bytes, pessoa); err != nil {
				fmt.Println(err)
			}

		print(p.GetNome())

}
