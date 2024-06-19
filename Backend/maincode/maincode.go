package maincode

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const ListMax = 500
const SparepartMax = 500

type Tanggal struct {
	Bulan          string
	Tanggal, Tahun int
}

type Servis struct {
	Id_servis         string
	Id_pelanggan      int
	Total_Harga       int
	Tanggal_kunjungan Tanggal
}

type Sparepart struct {
	Id_sparepart    string
	Nama_sparepart  string
	Harga_sparepart int
	Jumlah_terjual  int
	Jenis_motor     string
}

type Pelanggan struct {
	Id_pelanggan   int
	Nama_pelanggan string
	Jenis_motor    string
	Nomor_plat     string
}

type List_pelanggan [ListMax]Pelanggan
type List_sparepart [SparepartMax]Sparepart
type List_servis [ListMax]Servis

func KonversiSQLKeTanggal(x string) Tanggal {
	var hasil Tanggal
	Tahun, errors := strconv.Atoi(x[:4])
	if errors != nil {
		log.Print("[KONVERSI TANGGAL SQL-STANDARD] error : ", errors)
	}
	hasil.Tahun = Tahun
	if x[4:8] == "-01-" {
		hasil.Bulan = "januari"
	} else if x[4:8] == "-02-" {
		hasil.Bulan = "februari"
	} else if x[4:8] == "-03-" {
		hasil.Bulan = "maret"
	} else if x[4:8] == "-04-" {
		hasil.Bulan = "april"
	} else if x[4:8] == "-05-" {
		hasil.Bulan = "mei"
	} else if x[4:8] == "-06-" {
		hasil.Bulan = "juni"
	} else if x[4:8] == "-07-" {
		hasil.Bulan = "juli"
	} else if x[4:8] == "-08-" {
		hasil.Bulan = "agustus"
	} else if x[4:8] == "-09-" {
		hasil.Bulan = "september"
	} else if x[4:8] == "-10-" {
		hasil.Bulan = "oktober"
	} else if x[4:8] == "-11-" {
		hasil.Bulan = "november"
	} else if x[4:8] == "-12-" {
		hasil.Bulan = "desember"
	}
	Tanggal, errors2 := strconv.Atoi(x[8:])
	if errors2 != nil {
		log.Print("[KONVERSI TANGGAL SQL-STANDARD] error : ", errors2)
	}
	hasil.Tanggal = Tanggal
	return hasil
}

func KonversiTanggalSQL(x Tanggal) string {
	var hasil string
	x.Bulan = strings.ToLower(x.Bulan)
	if x.Tahun >= 0 {
		hasil = strconv.Itoa(x.Tahun)
	}
	if x.Bulan == "januari" {
		hasil = hasil + "-01-"
	} else if x.Bulan == "februari" {
		hasil = hasil + "-02-"
	} else if x.Bulan == "maret" {
		hasil = hasil + "-03-"
	} else if x.Bulan == "april" {
		hasil = hasil + "-04-"
	} else if x.Bulan == "mei" {
		hasil = hasil + "-05-"
	} else if x.Bulan == "juni" {
		hasil = hasil + "-06-"
	} else if x.Bulan == "juli" {
		hasil = hasil + "-07-"
	} else if x.Bulan == "agustus" {
		hasil = hasil + "-08-"
	} else if x.Bulan == "september" {
		hasil = hasil + "-09-"
	} else if x.Bulan == "oktober" {
		hasil = hasil + "-10-"
	} else if x.Bulan == "november" {
		hasil = hasil + "-11-"
	} else if x.Bulan == "desember" {
		hasil = hasil + "-12-"
	}
	if x.Tanggal > 0 {
		hasil = hasil + strconv.Itoa(x.Tanggal)
	}
	log.Print("[KONVERSI TANGGAL - SQL] Hasil : ", hasil)
	return hasil
}

func KonversiId(id, x string) string {
	var angka string
	angka = id[1:]
	bil, err := strconv.Atoi(angka)
	bil += 1
	if err != nil {
		log.Print(err)
	}
	temp := strconv.Itoa(bil)
	if len(temp) == 2 {
		angka = x + "0" + temp
	} else if len(temp) == 1 {
		angka = x + "00" + temp
	} else {
		angka = x + temp
	}
	return angka
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/servicemotor")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TambahPelanggan(A Pelanggan) {
	input := "INSERT INTO pelanggan(Nama_pelanggan, Jenis_motor, Nomor_plat) VALUES ('" + A.Nama_pelanggan + "','" + A.Jenis_motor + "','" + A.Nomor_plat + "');"
	db, err := Connect()
	if err != nil {
		log.Print("[TAMBAH PELANGGAN] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print(errors)
	}
	defer log.Print("[TAMBAH PELANGGAN] Success")
	defer rows.Close()
}

func TambahSparepart(A Sparepart) { //HANYA INPUT NAMA, HARGA, JENIS MOTOR
	var id string
	input1 := "SELECT Id_sparepart FROM sparepart ORDER BY Id_sparepart DESC;"
	db, err := Connect()
	if err != nil {
		log.Print("[TAMBAH SPAREPART] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	rows1, errors1 := db.QueryContext(ctx, input1)
	if errors1 != nil {
		log.Print("[TAMBAH SPAREPART] Error : ", errors1)
	}
	if !rows1.Next() {
		id = "P000"
	} else {
		rows1.Scan(&id)
	}
	input2 := "INSERT INTO sparepart(Id_sparepart, Nama_sparepart, Harga_sparepart, Jumlah_terjual, Jenis_motor) VALUES ('" + KonversiId(id, "P") + "','" + A.Nama_sparepart + "'," + strconv.Itoa(A.Harga_sparepart) + "," + strconv.Itoa(0) + ",'" + A.Jenis_motor + "');"
	rows2, errors2 := db.QueryContext(ctx, input2)
	if errors2 != nil {
		log.Print("[TAMBAH SPAREPART] Error : ", errors2)
	}
	defer log.Print("[TAMBAH SPAREPART] Success ")
	defer rows1.Close()
	defer rows2.Close()
}

func DaftarServis(idPelanggan int, date Tanggal) { //MASUKKAN ID
	var id string
	TanggalTerkonversi := KonversiTanggalSQL(date)
	input1 := "SELECT Id_servis FROM servis ORDER BY Id_servis DESC;"
	db, err := Connect()
	if err != nil {
		log.Print("[DAFTAR SERVIS] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	rows1, errors1 := db.QueryContext(ctx, input1)
	if errors1 != nil {
		log.Print("[DAFTAR SERVIS] Error : ", errors1)
	}
	if !rows1.Next() {
		id = "S000"
	} else {
		rows1.Scan(&id)
	}
	input2 := "INSERT INTO servis(Id_servis, Id_pelanggan, Total_harga, Tanggal_kunjungan) VALUES ('" + KonversiId(id, "S") + "','" + strconv.Itoa(idPelanggan) + "','" + "0" + "','" + TanggalTerkonversi + "');"
	rows2, errors2 := db.QueryContext(ctx, input2)
	if errors2 != nil {
		log.Print("[DAFTAR SERVIS] Error : ", errors2)
	}
	defer log.Print("[DAFTAR SERVIS] Success ")
	defer rows1.Close()
	defer rows2.Close()
}

func PesanSparepart(idServis, idSparepart string) {
	input := "INSERT INTO memesan(Id_servis, Id_sparepart) VALUES ('" + idServis + "','" + idSparepart + "');"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[PESAN PESANAN] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[PESAN PESANAN] Error : ", errors)
	}
	defer rows.Close()
	input2 := "UPDATE sparepart SET Jumlah_terjual = Jumlah_terjual+1 WHERE Id_sparepart = '" + idSparepart + "';"
	log.Print("[PESAN PESANAN] Sql : ", input2)
	rows2, errors2 := db.QueryContext(ctx, input2)
	if errors2 != nil {
		log.Print("[PESAN PESANAN] Error : ", errors2)
	}
	defer log.Print("[TAMBAH PESANAN] Success ")
	defer rows2.Close()
	HitungTotalHargaServis(idServis, 0)
}

func HitungTotalHargaServis(idServis string, BiayaServis int) { //CARI BERDASARKAN ID SERVIS DAN MENGEMBALIKAN VARIABEL SERVIS
	var idSparepart [SparepartMax]string
	var idServisTemp [ListMax]string
	var idSparepartTemp [SparepartMax]string
	var i, i2 int
	var t Servis
	var input2 string
	input := "SELECT Id_sparepart, Id_servis FROM memesan;"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[HITUNG TOTAL SERVIS] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[HITUNG TOTAL SERVIS] Error1 : ", errors)
	}
	for rows.Next() {
		rows.Scan(&idSparepartTemp[i], &idServisTemp[i])
		i++
	}
	i2 = 0
	for y := 0; y < i; y++ { //SEQUENTIAL SEARCH
		if idServisTemp[y] == idServis {
			idSparepart[i2] = idSparepartTemp[y]
			i2++
		}
	}
	i = i2
	defer rows.Close()
	for x := 0; x < i; x++ {
		if x == i-1 {
			input2 = input2 + "'" + idSparepart[x] + "'"
		} else {
			input2 = input2 + "'" + idSparepart[x] + "',"
		}
	}
	input2 = "SELECT SUM(Harga_sparepart) AS total FROM sparepart WHERE Id_sparepart IN(" + input2 + ");"
	log.Print("[HITUNG TOTAL SERVIS] Id Sparepart : ", input2)
	rows2, errors2 := db.QueryContext(ctx, input2)
	if errors2 != nil {
		log.Print("[HITUNG TOTAL SERVIS] Error2 : ", errors2)
	}
	for rows2.Next() {
		rows2.Scan(&t.Total_Harga)
	}
	t.Total_Harga += BiayaServis
	input3 := "UPDATE servis SET Total_harga= " + strconv.Itoa(t.Total_Harga) + " WHERE Id_servis = '" + idServis + "';"
	rows3, errors3 := db.QueryContext(ctx, input3)
	if errors3 != nil {
		log.Print("[HITUNG TOTAL SERVIS] Error : ", errors3)
	}
	defer log.Print("[HITUNG TOTAL SERVIS] Success")
	defer rows3.Close()
	defer rows2.Close()
}

func CariPelanggan(nPelanggan string, y *List_pelanggan, n *int) { //MASUKKAN NAMA PELANGGAN
	var i int = 0
	var x int = 0
	var output List_pelanggan
	input := "SELECT * FROM pelanggan;"
	db, err := Connect()
	if err != nil {
		log.Print("[CARI PELANGGAN] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[CARI PELANGGAN] Error : ", errors)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&output[i].Id_pelanggan, &output[i].Nama_pelanggan, &output[i].Jenis_motor, &output[i].Nomor_plat)
		i++
	}
	*n = i
	for i = 0; i < *n; i++ { //SEQUENTIAL SEARCH
		if strings.EqualFold(strings.ToLower(output[i].Nama_pelanggan), strings.ToLower(nPelanggan)) {
			y[x] = output[i]
			x++
		}
	}
	*n = x
	defer log.Print("[CARI PELANGGAN] Success ")
}

func CariPelangganSparepart(idSparepart string, out *List_pelanggan, n *int) {
	var i, x int
	var idServis [ListMax]string
	var idServisTemp [ListMax]string
	var idSparepartTemp [ListMax]string
	input := "SELECT Id_servis, Id_sparepart FROM memesan;"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[CARI PELANGGAN] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[CARI PELANGGAN] Error : ", errors)
	}
	for rows.Next() {
		rows.Scan(&idServisTemp[i], &idSparepartTemp[i])
		i++
	}
	defer rows.Close()
	*n = i
	x = 0
	for i = 0; i < *n; i++ { //SEQUENTIAL SEARCH
		if idSparepartTemp[i] == idSparepart {
			idServis[x] = idServisTemp[i]
			x++
		}
	}
	*n = x + 1
	input2 := ""
	for i = 0; i < *n; i++ {
		if i == *n-1 {
			input2 = input2 + "'" + idServis[i] + "');"
		} else {
			input2 = input2 + "'" + idServis[i] + "',"
		}
	}
	input2 = "SELECT pelanggan.Id_pelanggan, pelanggan.Nama_pelanggan, pelanggan.Jenis_motor, pelanggan.Nomor_plat  FROM servis INNER JOIN pelanggan ON servis.Id_pelanggan = pelanggan.Id_pelanggan WHERE servis.Id_servis IN(" + input2
	i = 0
	rows2, errors2 := db.QueryContext(ctx, input2)
	if errors2 != nil {
		log.Print("[CARI PELANGGAN] Error : ", errors2)
	}
	defer rows2.Close()
	for rows2.Next() {
		rows2.Scan(&out[i].Id_pelanggan, &out[i].Nama_pelanggan, &out[i].Jenis_motor, &out[i].Nomor_plat)
		i++
	}
	defer log.Print("[CARI PELANGGAN] Success ")
	*n = i
}

func CariSparepart(nSparepart string, y *List_sparepart, n *int) { //NAMA SPAREPART
	var i int = 0
	var x int = 0
	var output List_sparepart
	input := "SELECT * FROM sparepart;"
	db, err := Connect()
	if err != nil {
		log.Print("[CARI SPAREPART] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[CARI SPAREPART] Error : ", errors)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&output[i].Id_sparepart, &output[i].Nama_sparepart, &output[i].Harga_sparepart, &output[i].Jumlah_terjual, &output[i].Jenis_motor)
		i++
	}
	*n = i
	for i = 0; i < *n; i++ { //SEQUENTIAL SEARCH
		if strings.EqualFold(strings.ToLower(output[i].Nama_sparepart), strings.ToLower(nSparepart)) {
			y[x] = output[i]
			x++
		}
	}
	defer log.Print("[CARI SPAREPART] Success ")
	*n = x
}

func CariServisNamaPelanggan(idPelanggan int, out *List_servis, n *int) { // MASUKKAN ID PELANGGAN
	var i int
	var Tanggal string
	var arrTemp List_servis
	db, err := Connect()
	if err != nil {
		log.Print("[CARI SERVIS] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	input := "SELECT Id_servis, Id_pelanggan, Tanggal_kunjungan, Total_harga FROM servis;"
	rows2, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[CARI SERVIS] Error : ", errors)
	} else {
		for rows2.Next() {
			rows2.Scan(&arrTemp[i].Id_servis, &arrTemp[i].Id_pelanggan, &Tanggal, &arrTemp[i].Total_Harga)
			arrTemp[i].Tanggal_kunjungan = KonversiSQLKeTanggal(Tanggal)
			log.Print("[CARI SERVIS] Data :", arrTemp[i])
			i++
		}
		*n = i
		x := 0
		for i = 0; i < *n; i++ { //SEQUENTIAL SEARCH
			if arrTemp[i].Id_pelanggan == idPelanggan {
				out[x] = arrTemp[i]
				x++
			}
		}
		defer log.Print("[CARI SERVIS] Success ")
		*n = x
	}
}

func CariServisTanggal(in Tanggal, out *List_servis, out2 *List_pelanggan, n *int) {
	var i int = 0
	var Tanggal string
	input := "SELECT servis.Id_servis, servis.Id_pelanggan, servis.Total_harga, servis.Tanggal_kunjungan, pelanggan.Id_pelanggan, pelanggan.Nama_pelanggan, pelanggan.Jenis_motor, pelanggan.Nomor_plat FROM servis INNER JOIN pelanggan ON pelanggan.Id_pelanggan = servis.Id_pelanggan WHERE servis.Tanggal_kunjungan LIKE '%" + KonversiTanggalSQL(in) + "%';"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[CARI SERVIS] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[CARI SERVIS] Error : ", errors)
	}
	for rows.Next() {
		rows.Scan(&out[i].Id_servis, &out[i].Id_pelanggan, &out[i].Total_Harga, &Tanggal, &out2[i].Id_pelanggan, &out2[i].Nama_pelanggan, &out2[i].Jenis_motor, &out2[i].Nomor_plat)
		out[i].Tanggal_kunjungan = KonversiSQLKeTanggal(Tanggal)
		i++
	}
	defer rows.Close()
	defer log.Print("[CARI SERVIS] Success ")
	*n = i
}

func EditPelanggan(idPelanggan int, x Pelanggan) { //MEMEASUKKAN ID PELANGGAN DAN STRUKTUR PELANGGAN
	input := "UPDATE pelanggan SET Nama_pelanggan = '" + x.Nama_pelanggan + "', Jenis_motor = '" + x.Jenis_motor + "', Nomor_plat= '" + x.Nomor_plat + "' WHERE Id_pelanggan = " + strconv.Itoa(idPelanggan) + ";"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[EDIT PELANGGAN] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[EDIT PELANGGAN] Error : ", errors)
	}
	defer log.Print("[EDIT PELANGGAN] Success ")
	defer rows.Close()
}

func EditSparepart(idSparepart string, x Sparepart) { //MEMEASUKKAN ID SPAPREPART DAN STRUKTUR SPAREPART
	input := "UPDATE sparepart SET Nama_sparepart = '" + x.Nama_sparepart + "', Jenis_motor = '" + x.Jenis_motor + "', Harga_sparepart = " + strconv.Itoa(x.Harga_sparepart) + " WHERE Id_Sparepart = '" + idSparepart + "';"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[EDIT SPAREPART] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[EDIT SPAREPART] Error : ", errors)
	}
	defer log.Print("[EDIT SPAREPART] Success ")
	defer rows.Close()
}

func HapusPelanggan(idPelanggan int) {
	var idServis [ListMax]string
	var i int = 0
	var n int
	var adaServis bool = false
	input1 := "SELECT Id_servis FROM servis WHERE Id_pelanggan = " + strconv.Itoa(idPelanggan) + " ;"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[HAPUS PELANGGAN] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input1)
	if errors != nil {
		log.Print("[HAPUS PELANGGAN] Error : ", errors)
	}
	for rows.Next() {
		rows.Scan(&idServis[i])
		adaServis = true
		i++
	}
	n = i + 1
	defer rows.Close()
	if adaServis {
		idServisString := ""
		input3 := ""
		for i = 0; i < n-1; i++ {
			if i == n-2 {
				idServisString = idServisString + "'" + idServis[i] + "')"
			} else {
				idServisString = idServisString + "'" + idServis[i] + "',"
			}
		}
		inputConf := "SELECT Id_Servis FROM memesan WHERE Id_servis IN(" + idServisString + ";"
		confirm, errorsConf := db.QueryContext(ctx, inputConf)
		if errorsConf != nil {
			log.Print("[HAPUS PELANGGAN] Error : ", errorsConf)
		}
		defer confirm.Close()
		if confirm.Next() {
			inputs2 := "DELETE FROM memesan WHERE Id_servis IN(" + idServisString + ";"
			rows2, errors2 := db.QueryContext(ctx, inputs2)
			if errors2 != nil {
				log.Print("[HAPUS PELANGGAN] Error : ", errors2)
			}
			defer rows2.Close()
		}
		input3 = "DELETE FROM servis WHERE Id_pelanggan =" + strconv.Itoa(idPelanggan) + ";"
		defer rows.Close()
		rows3, errors3 := db.QueryContext(ctx, input3)
		if errors3 != nil {
			log.Print("[HAPUS PELANGGAN] Error : ", errors3)
		}
		defer rows3.Close()
	}
	input4 := "DELETE FROM pelanggan WHERE Id_pelanggan = " + strconv.Itoa(idPelanggan) + ";"
	rows4, errors4 := db.QueryContext(ctx, input4)
	if errors4 != nil {
		log.Print("[HAPUS PELANGGAN] Error : ", errors4)
	}
	defer log.Print("[HAPUS PELANGGAN] Success ")
	defer rows4.Close()
}

func HapusSparepart(idSparepart string) {
	input := "DELETE FROM memesan WHERE Id_sparepart = '" + idSparepart + "';"
	input2 := "DELETE FROM sparepart WHERE Id_sparepart = '" + idSparepart + "';"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[HAPUS SPAREPART] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[HAPUS SPAREPART] Error : ", errors)
	}
	defer rows.Close()
	rows2, errors2 := db.QueryContext(ctx, input2)
	if errors2 != nil {
		log.Print("[HAPUS SPAREPART] Error : ", errors2)
	}
	defer log.Print("[HAPUS SPAREPART] Success ")
	defer rows2.Close()
}

func HapusPesananDariServis(idSparepart, idServis string) {
	input := "DELETE FROM memesan WHERE Id_sparepart = '" + idSparepart + "' AND Id_servis = '" + idServis + "';"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[HAPUS PESANAN] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[HAPUS PESANAN] Error : ", errors)
	}
	defer log.Print("[HAPUS PESANAN] Success ")
	defer rows.Close()
}

func HapusServis(idServis string) {
	input := "DELETE FROM  memesan WHERE Id_servis = '" + idServis + "';"
	input2 := "DELETE FROM servis WHERE Id_servis = '" + idServis + "';"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[HAPUS SERVIS] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[HAPUS SERVIS] Error : ", errors)
	}
	defer rows.Close()
	rows2, errors2 := db.QueryContext(ctx, input2)
	if errors2 != nil {
		log.Print("[HAPUS SERVIS] Error : ", errors2)
	}
	defer log.Print("[HAPUS SERVIS] Success ")
	defer rows2.Close()
}

func TampilkanPelanggan(output *List_pelanggan, n *int) {
	var i int = 0
	input := "SELECT * FROM pelanggan"
	db, err := Connect()
	if err != nil {
		log.Print("[TAMPILKAN TABEL PELANGGAN] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[TAMPILKAN TABEL PELANGGAN] Error : ", errors)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&output[i].Id_pelanggan, &output[i].Nama_pelanggan, &output[i].Jenis_motor, &output[i].Nomor_plat)
		i++
	}
	defer log.Print("[TAMPILKAN TABEL PELANGGAN] Success ")
	*n = i
}

func TampilkanSparepart(output *List_sparepart, n *int) {
	var i int = 0
	var temp int
	var tempbox Sparepart
	input := "SELECT Id_sparepart, Nama_sparepart, Harga_sparepart, Jumlah_terjual, Jenis_motor FROM sparepart;"
	db, err := Connect()
	if err != nil {
		log.Print("[TAMPILKAN TABEL SPAREPART] Error : ", err)
	}
	defer db.Close()
	ctx := context.Background()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[TAMPILKAN TABEL SPAREPART] Error : ", errors)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&output[i].Id_sparepart, &output[i].Nama_sparepart, &output[i].Harga_sparepart, &output[i].Jumlah_terjual, &output[i].Jenis_motor)
		i++
	}
	*n = i
	for y := 0; y < *n; y++ { //SELECTION SORT
		temp = output[y].Jumlah_terjual
		tempbox = output[y]
		for x := y + 1; x < *n; x++ {
			if temp < output[x].Jumlah_terjual {
				tempbox = output[y]
				output[y] = output[x]
				output[x] = tempbox
			}
		}
	}
	defer log.Print("[TAMPILKAN TABEL SPAREPART] Success ")
}

func TampilkanServis(out *List_servis, out2 *List_pelanggan, n *int) {
	var i int = 0
	var Tanggal string
	input := "SELECT servis.Id_servis, servis.Id_pelanggan, servis.Total_harga, servis.Tanggal_kunjungan, pelanggan.Id_pelanggan, pelanggan.Nama_pelanggan, pelanggan.Jenis_motor, pelanggan.Nomor_plat FROM servis INNER JOIN pelanggan ON pelanggan.Id_pelanggan = servis.Id_pelanggan;"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[TAMPILKAN TABEL SERVIS] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[TAMPILKAN TABEL SERVIS] Error : ", errors)
	}
	for rows.Next() {
		rows.Scan(&out[i].Id_servis, &out[i].Id_pelanggan, &out[i].Total_Harga, &Tanggal, &out2[i].Id_pelanggan, &out2[i].Nama_pelanggan, &out2[i].Jenis_motor, &out2[i].Nomor_plat)
		out[i].Tanggal_kunjungan = KonversiSQLKeTanggal(Tanggal)
		i++
	}
	defer rows.Close()
	defer log.Print("[TAMPILKAN TABEL SERVIS] Success ")
	*n = i
}

func TampilkanPesanan(idServis string, out *List_sparepart, n *int) {
	var i int = 0
	input := "SELECT sparepart.Id_sparepart, sparepart.Nama_sparepart, sparepart.Jenis_motor, sparepart.Harga_sparepart FROM memesan INNER JOIN sparepart ON sparepart.Id_sparepart = memesan.Id_sparepart WHERE Id_servis=" + "'" + idServis + "';"
	ctx := context.Background()
	db, err := Connect()
	if err != nil {
		log.Print("[TAMPILKAN TABEL PESANAN] Error : ", err)
	}
	defer db.Close()
	rows, errors := db.QueryContext(ctx, input)
	if errors != nil {
		log.Print("[TAMPILKAN TABEL PESANAN] Error : ", errors)
	}
	for rows.Next() {
		rows.Scan(&out[i].Id_sparepart, &out[i].Nama_sparepart, &out[i].Jenis_motor, &out[i].Harga_sparepart)
		i++
	}
	defer rows.Close()
	defer log.Print("[TAMPILKAN TABEL PESANAN] Success ")
	*n = i
}
