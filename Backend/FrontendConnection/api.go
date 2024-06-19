package api

import (
	"encoding/json"
	"log"
	"net/http"
	"pkg/maincode"

	"github.com/gorilla/mux"
)

type Memesan struct {
	Id_servis    string
	Id_sparepart string
}

type ServisDanPelanggan struct {
	Id_servis         string
	Id_pelanggan      int
	Nama_pelanggan    string
	Jenis_motor       string
	Nomor_plat        string
	Total_Harga       int
	Tanggal_kunjungan maincode.Tanggal
}

type List_servisdanpelanggan [250]ServisDanPelanggan

func RunServer() {
	comm := mux.NewRouter()
	comm.HandleFunc("/tambah/pelanggan", tPelanggan).Methods("POST")
	comm.HandleFunc("/tambah/sparepart", tSparepart).Methods("POST")
	comm.HandleFunc("/tambah/servis", tServis).Methods("POST")
	comm.HandleFunc("/tambah/pesanan", tMemesan).Methods("POST")
	comm.HandleFunc("/cari/pelanggan/nama", cPelangganNama).Methods("POST")
	comm.HandleFunc("/cari/pelanggan/pembelian-sparepart", cPelangganSparepart).Methods("POST")
	comm.HandleFunc("/cari/sparepart/nama", cSparepart).Methods("POST")
	comm.HandleFunc("/cari/servis/nama-pelanggan", cServisNamaPelanggan).Methods("POST")
	comm.HandleFunc("/cari/servis/tanggal", cServisTanggal).Methods("POST")
	comm.HandleFunc("/edit/pelanggan", ePelanggan).Methods("POST")
	comm.HandleFunc("/edit/sparepart", eSparepart).Methods("POST")
	comm.HandleFunc("/hapus/pelanggan", hPelanggan).Methods("POST")
	comm.HandleFunc("/hapus/sparepart", hSparepart).Methods("POST")
	comm.HandleFunc("/hapus/servis", hServis).Methods("POST")
	comm.HandleFunc("/hapus/pesanan", hPesananDariServis).Methods("POST")
	comm.HandleFunc("/pesanan", TampilkanPesanan).Methods("POST")
	comm.HandleFunc("/pelanggan", TampilkanPelanggan).Methods("GET")
	comm.HandleFunc("/sparepart", TampilkanSparepart).Methods("GET")
	comm.HandleFunc("/servis", TampilkanServis).Methods("GET")

	err := http.ListenAndServe(":8080", comm)
	if err != nil {
		log.Fatalln("There's an error with the server, ", err)
	}

}

func tPelanggan(x http.ResponseWriter, r *http.Request) {
	var pelanggan maincode.Pelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&pelanggan)
	log.Print("[TAMBAH PELANGGAN] Nama : ", pelanggan.Nama_pelanggan)
	maincode.TambahPelanggan(pelanggan)
}

func tSparepart(x http.ResponseWriter, r *http.Request) {
	var sparepart maincode.Sparepart
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&sparepart)
	log.Print("[TAMBAH SPAREPART] Nama : ", sparepart.Nama_sparepart)
	maincode.TambahSparepart(sparepart)
}

func tServis(x http.ResponseWriter, r *http.Request) {
	var servis maincode.Servis
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&servis)
	log.Print("[PESAN SERVIS] ID Pelanggan : ", servis.Id_pelanggan)
	maincode.DaftarServis(servis.Id_pelanggan, servis.Tanggal_kunjungan)
}

func tMemesan(x http.ResponseWriter, r *http.Request) {
	var order Memesan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&order)
	log.Print("[TAMBAH PESANAN] ID Sparepart : ", order.Id_sparepart, ", ID Servis : ", order.Id_servis)
	maincode.PesanSparepart(order.Id_servis, order.Id_sparepart)
}

func cPelangganNama(x http.ResponseWriter, r *http.Request) {
	var pelanggan maincode.Pelanggan
	var n int
	var ArrPelanggan maincode.List_pelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&pelanggan)
	log.Print("[CARI PELANGGAN] Nama : ", pelanggan.Nama_pelanggan)
	maincode.CariPelanggan(pelanggan.Nama_pelanggan, &ArrPelanggan, &n)
	json.NewEncoder(x).Encode(ArrPelanggan[0:n])
}

func cPelangganSparepart(x http.ResponseWriter, r *http.Request) { //MASUKKAN ID SPAREPART
	var sparepart maincode.Sparepart
	var n int
	var ArrPelanggan maincode.List_pelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&sparepart)
	log.Print("[CARI PELANGGAN] Sparepart : ", sparepart.Id_sparepart)
	maincode.CariPelangganSparepart(sparepart.Id_sparepart, &ArrPelanggan, &n)
	json.NewEncoder(x).Encode(ArrPelanggan[0:n])
}

func cSparepart(x http.ResponseWriter, r *http.Request) {
	var sparepart maincode.Sparepart
	var n int
	var ArrSparepart maincode.List_sparepart
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&sparepart)
	log.Print("[CARI SPAREPART] Nama : ", sparepart.Nama_sparepart)
	maincode.CariSparepart(sparepart.Nama_sparepart, &ArrSparepart, &n)
	json.NewEncoder(x).Encode(ArrSparepart[0:n])
}

func cServisNamaPelanggan(x http.ResponseWriter, r *http.Request) { //INPUT ID PELANGGAN
	var pelanggan maincode.Pelanggan
	var n int
	var ArrServis maincode.List_servis
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&pelanggan)
	log.Print("[CARI SERVIS] Id Pelanggan : ", pelanggan.Id_pelanggan)
	maincode.CariServisNamaPelanggan(pelanggan.Id_pelanggan, &ArrServis, &n)
	json.NewEncoder(x).Encode(ArrServis[0:n])
}

func cServisTanggal(x http.ResponseWriter, r *http.Request) { //INPUT TANGGAL
	var tanggal maincode.Tanggal
	var n int
	var ArrServis maincode.List_servis
	var ArrPelanggan maincode.List_pelanggan
	var ArrGab List_servisdanpelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&tanggal)
	log.Print("[CARI SERVIS] Tanggal : ", tanggal.Tanggal, "-", tanggal.Bulan, "-", tanggal.Tahun)
	maincode.CariServisTanggal(tanggal, &ArrServis, &ArrPelanggan, &n)
	for i := 0; i < n; i++ {
		ArrGab[i].Id_pelanggan = ArrPelanggan[i].Id_pelanggan
		ArrGab[i].Nama_pelanggan = ArrPelanggan[i].Nama_pelanggan
		ArrGab[i].Jenis_motor = ArrPelanggan[i].Jenis_motor
		ArrGab[i].Nomor_plat = ArrPelanggan[i].Nomor_plat
		ArrGab[i].Total_Harga = ArrServis[i].Total_Harga
		ArrGab[i].Tanggal_kunjungan = ArrServis[i].Tanggal_kunjungan
		ArrGab[i].Id_servis = ArrServis[i].Id_servis
	}
	json.NewEncoder(x).Encode(ArrGab[0:n])
}

func ePelanggan(x http.ResponseWriter, r *http.Request) {
	var pelanggan maincode.Pelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&pelanggan)
	log.Print("[EDIT PELANGGAN] ID Pelanggan : ", pelanggan.Id_pelanggan)
	maincode.EditPelanggan(pelanggan.Id_pelanggan, pelanggan)
}

func eSparepart(x http.ResponseWriter, r *http.Request) {
	var sparepart maincode.Sparepart
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&sparepart)
	log.Print("[EDIT SPAREPART] ID Sparepart : ", sparepart.Id_sparepart)
	maincode.EditSparepart(sparepart.Id_sparepart, sparepart)
}

func hPelanggan(x http.ResponseWriter, r *http.Request) {
	var pelanggan maincode.Pelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&pelanggan)
	log.Print("[HAPUS PELANGGAN] ID Pelanggan : ", pelanggan.Id_pelanggan)
	maincode.HapusPelanggan(pelanggan.Id_pelanggan)
}

func hSparepart(x http.ResponseWriter, r *http.Request) {
	var sparepart maincode.Sparepart
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&sparepart)
	log.Print("[HAPUS SPAREPART] ID Sparepart : ", sparepart.Id_sparepart)
	maincode.HapusSparepart(sparepart.Id_sparepart)
}

func hServis(x http.ResponseWriter, r *http.Request) {
	var servis maincode.Servis
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&servis)
	log.Print("[HAPUS SERVIS] ID Servis : ", servis.Id_servis)
	maincode.HapusServis(servis.Id_servis)
}

func hPesananDariServis(x http.ResponseWriter, r *http.Request) {
	var pesanan Memesan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&pesanan)
	log.Print("[HAPUS PESANAN] ID Servis : ", pesanan.Id_servis)
	maincode.HapusPesananDariServis(pesanan.Id_sparepart, pesanan.Id_servis)
}

func TampilkanPelanggan(x http.ResponseWriter, r *http.Request) {
	var n int
	var ArrPelanggan maincode.List_pelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	log.Print("[TAMPILKAN TABEL PELANGGAN]")
	maincode.TampilkanPelanggan(&ArrPelanggan, &n)
	json.NewEncoder(x).Encode(ArrPelanggan[0:n])
}

func TampilkanSparepart(x http.ResponseWriter, r *http.Request) {
	var n int
	var ArrSparepart maincode.List_sparepart
	x.Header().Set("Access-Control-Allow-Origin", "*")
	log.Print("[TAMPILKAN TABEL SPAREPART]")
	maincode.TampilkanSparepart(&ArrSparepart, &n)
	json.NewEncoder(x).Encode(ArrSparepart[0:n])
}

func TampilkanServis(x http.ResponseWriter, r *http.Request) {
	var n int
	var ArrServis maincode.List_servis
	var ArrPelanggan maincode.List_pelanggan
	var ArrGab List_servisdanpelanggan
	x.Header().Set("Access-Control-Allow-Origin", "*")
	log.Print("[TAMPILKAN TABEL SERVIS]")
	maincode.TampilkanServis(&ArrServis, &ArrPelanggan, &n)
	for i := 0; i < n; i++ {
		ArrGab[i].Id_pelanggan = ArrPelanggan[i].Id_pelanggan
		ArrGab[i].Nama_pelanggan = ArrPelanggan[i].Nama_pelanggan
		ArrGab[i].Jenis_motor = ArrPelanggan[i].Jenis_motor
		ArrGab[i].Nomor_plat = ArrPelanggan[i].Nomor_plat
		ArrGab[i].Total_Harga = ArrServis[i].Total_Harga
		ArrGab[i].Tanggal_kunjungan = ArrServis[i].Tanggal_kunjungan
		ArrGab[i].Id_servis = ArrServis[i].Id_servis
	}
	json.NewEncoder(x).Encode(ArrGab[0:n])
}

func TampilkanPesanan(x http.ResponseWriter, r *http.Request) {
	var Servis maincode.Servis
	var ArrPesanan maincode.List_sparepart
	var n int
	x.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&Servis)
	log.Print("[TAMPILKAN PESANAN] ID Servis : ", Servis.Id_servis)
	maincode.TampilkanPesanan(Servis.Id_servis, &ArrPesanan, &n)
	json.NewEncoder(x).Encode(ArrPesanan[0:n])
}
