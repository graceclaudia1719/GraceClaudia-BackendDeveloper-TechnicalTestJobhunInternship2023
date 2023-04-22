# Grace Claudia - Backend Developer - Technical Test Jobhun Internship 2023

## Prerequisite
- golang
- mysql

## Library
- echo
## ADDITIONAL NOTES

1. tidak ada relasi maupun atribut yang berkaitan antar mahasiswa dan jurusan sehingga tidak dapat mendapatkan jurusan mahasiswa maupun mengupdate jurusan mahasiswa
2. jurusan dan mahasiswa merupakan relasi many to many (mempertimbangkan adanya \*double degree) sehingga ada 2 pertimbangan penyelesaian yang memungkinkan:
   - ditambahkan entitas mahasiswa_jurusan yang atributnya id mahasiswa dan id jurusan
   - ditambahkan atribut jurusan bertipe array of string pada entitas mahasiswa

3. request insert mahasiswa dengan hobi berupa array of string merupakan pertimbangan karena biasanya terjadi relation many to many antar entitas sehingga diperlukan entitas tambahan agar menjembatani many to many menjadi many to one, dalam hal ini seperti mahasiswa dan hobi sehingga diperlukan mahasiswa_hobi. Dalam implementasinya, apabila hobi sering digunakan, dapat dipertimbangkan untuk menjadikan hobi sebagai atribut tambahan array of string pada mahasiswa agar tidak perlu mengakses 3 tabel hanya untuk mendapatkan hobi.
4. diasumsikan insert selalu untuk new mahasiswa karena tidak ada atribute yang mendefine keunikan antar mahasiswa (contoh: email, username)

double degree menurut situs akupintar.id merupakan program kuliah yang memungkinkan mahasiswa memperoleh dua gelar sarjana yang berbeda. Program double degree dapat dilakukan di dua jurusan yang berbeda dalam satu kampus yang sama, bisa juga di kampus yang berbeda.

## TESTING
dokumen testing lengkap dapat diakses di [sini](https://docs.google.com/document/d/1ZQgdigxIWvTsyme-C4AedS8UtevSX4YkvxFA1V7sv3Q/edit?usp=sharing)

## AUTHOR
Grace Claudia, [linkedin](www.linkedin.com/in/graceclaudia/)

