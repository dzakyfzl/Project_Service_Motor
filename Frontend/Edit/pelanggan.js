const url ="http://127.0.0.1:8080";
const body = document.body
let Opsi = document.getElementById('id-pelanggan')

const requestdefault = {
    method:"GET",
    redirect: "follow",
    mode: "cors"
}

fetch(url + "/pelanggan", requestdefault)
  .then(response => response.json())
  .then(result => buatOpsi(result))
  .catch(error => console.error(error));


function buatOpsi(data){
    console.log(data)
    var pilihan = `<option value="">Pilih Pelanggan</option>`
    for (var i = 0; i < data.length;i++){
        pilihan += `<option value="${data[i].Id_pelanggan}">${data[i].Id_pelanggan} ${data[i].Nama_pelanggan} ${data[i].Nomor_plat}</option>`
    }
    console.log(pilihan)
    Opsi.innerHTML = pilihan
    body.replace(Opsi)
}
  

function editpelanggan(){
    var id = Opsi.value
    var nama = document.getElementById('nama-pelanggan').value
    var motor = document.getElementById('jenis-motor').value
    var plat = document.getElementById('plat-nomor').value
    const send = JSON.stringify({
        "Id_pelanggan": parseInt(id),
        "Nama_pelanggan": nama,
        "Jenis_motor": motor,
        "Nomor_plat": plat 
      })
      const request ={
        method: "POST",
        body: send,
        redirect: "follow",
        mode: "cors"
      }
      fetch(url + "/edit/pelanggan", request)
}