# Golang JSON
JSON (JavaScript Object Notation), merupakan struktur format data yang bentuknya seperti Object di JavaScript<br>
`JSON` merupakan struktur format data paling banyak digunakan saat kita membuat `RESTful API`<br>
Golang sudah menyediakan `package json`, dimana kita bisa menggunakannya untuk melakukan konversi data `JSON` `(encode)`atau sebaliknya `(decode)`

## JSON Encode
Goalng Telah menyediakan function untuk melakukan konversi data ke `JSON`, yaitu menggunakan function `json.Marshal(interface{})`<br>
Karena parameternya `interface{}`, maka kita bisa memasukkan tipe data apapun kedalam function `Marshal`

## JSON Object
Pada materi sebelumnya kita melakukan encode data seperti `string, number, boolean, dan type data primitif lainnya`<br>
Walaupun bisa dilakukan karena sesuai dengan kontrak `interface{}`, namun tidak sesuai dengan kontrak `JSON`<br>
Jika mengikuti kontrak `JSON` bentuknya adalah `Object` dan `Array`, sedangkan pada materi sebelumnya harya berupa `value` saja
