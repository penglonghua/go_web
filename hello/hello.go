/**
* 最简单的例子 hello world
*/

package main

import (
	"net/http"
	"fmt"
)

func main() {



	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello , you've requested: %s\n", r.URL.Path)
	})

	// addr 也可以写完 localhost:9090
	http.ListenAndServe(":9090", nil)     // FIXME: 这个地方即使 端口冲突了，也没有报错，代码上有问题.


}
