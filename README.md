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
### Struct
JSON Object di golang direpresentasikan dengan type data `struct` <br>
Dimana tiap atribut di `JSON Object` merupakan atribut di `struct`

## JSON Decode
Untuk melakukan konversi dari JSON ke tipe data di golang `(decode)`, kita bisa menggunakan function `json.Unmarshal(byte[], interface{})`<br>
Diaman, `byte[]` adalah data `JSON`nya, sedangkan `interface{}` adalah tempat menyimpan hasil konversi, biasanya berupa `pointer`

## JSON Array
Selain tipe data dalam bentuk object, biasanya dalam JSON, kadang kita menggunakan tipe data `Array`<br>
`Array` di `JSON` mirip dengan `array di JavaScript`, dia bisa berisikan tipe `data primitif`, atau `tipe data kompleks`(`Object` atau `Array`)<br>
Di golang, JSON Array direpesentasikan dalam bentuk slice<br>
Konversi `dari JSON` atau `ke JSON` dilakukan secara otomatis oleh `package json` menggunakan `tipe data slice`
### Decode JSON Array
Selain menggunakan `Array` pada atribut di `Object`<br>
Kita juga bisa melakukan `Encode` atau `Decode` langsung `JSON Array`nya<br>
`Encode` dan `Decode` JSON Array bisa menggunakan tipe data `Slice`

## JSON Tag
Secara default atrubut yang terdapat di `struct` dan `JSON` akan di mapping sesuai dengan nama atribut yang sama (Case Sensitive)<br>
Kadang style yang berbeda antara penamaan atribut di Struct dan JSON, misalnya  di JSON kita ingin menggunakan `snake_case`, tapi di `Struct` kita ingin menggunakan `PascalCase`<br>
Untungnya `package json` mendukung `Tag Reflection`<br>
Kita bisa menambahkan tag reflection dengan nama `JSON`, lalu diikuti dengan atribut yang kita inginkan ketika konversi `dari atau ke JSON`

## JSON Map
Saat menggunakan `JSON`, kadang mungkin kita menemukan kasus data `JSON`nya `dinamis`<br>
Artinya atributnya tidak menentu, bisa bertambah dan bisa berkurang, dan tidak tetap<br>
Pada kasus seperti ini, menggunakan `struct` akan menyulitkan, karena pada `struct`, kita harus menentukan semua `atribut`nya <br>
Untuk kasus seperti ini, kita bisa menggunakan tipe data `map[string]interface{}` <br>
secara otomatis, atribut akan menjadi `key` di `map` dan `value` menjadi `value` di `map`<br>
Namun karena value berupa `interface{}`, maka kita harus melakukan konversi secara manual jika ingin mengambil `value`nya<br>
Dan type data `map` tidak mendukung `JSON Tag` lagi

## Streaming Decoder
Kadang data `JSON` berasal dari input berupa `io.Reader` `(File, Network, Request Body)`<br>
Kita bisa saja membaca semua datanya lalu menyimpannya kedalam variabel, baru lalukan konversi dari `JSON` , namun hal seperi ini sebenarnya tidak perlu dilakukan karena `package JSON` memiliki fitur untuk membaca dari `stream`