package main

import "fmt"

/*
	map
	map[key_type]value_type
	キーの型には==で比較可能な型がすべて使える=>文字列型、数値型、struct
*/

// Goに入りては・・・ 構造体をマップのキーに使う（キーの利用箇所を限定し安全性を確保）
func main() {
	tactics()
}

/*
	空の構造体をmapのキーに活用する
	構造体を利用することでキーに対するアクセス制御を行うことが可能
*/
type key_A struct{}
type key_B struct{}

func tactics() {
	// キーにも値にもinterface{}型を用いてどんな型でも利用可能にすると便利
	m := make(map[interface{}]interface{})
	m[key_A{}] = "value_A"
	m[key_B{}] = "value_B"

	v := m[key_A{}]

	fmt.Printf("%v\n", v)
}
