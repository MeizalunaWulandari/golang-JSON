# Golang JSON
JSON (JavaScript Object Notation), merupakan struktur format data yang bentuknya seperti Object di JavaScript<br>
`JSON` merupakan struktur format data paling banyak digunakan saat kita membuat `RESTful API`<br>
Golang sudah menyediakan `package json`, dimana kita bisa menggunakannya untuk melakukan konversi data `JSON` `(encode)`atau sebaliknya `(decode)`

## JSON Encode
Goalng Telah menyediakan function untuk melakukan konversi data ke `JSON`, yaitu menggunakan function `json.Marshal(interface{})`<br>
Karena parameternya `interface{}`, maka kita bisa memasukkan tipe data apapun kedalam function `Marshal`