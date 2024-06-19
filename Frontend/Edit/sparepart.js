const url ="http://127.0.0.1:8080";
const body = document.body
let Opsi = document.getElementById('id-sparepart')

const requestdefault = {
    method:"GET",
    redirect: "follow",
    mode: "cors"
}

fetch(url + "/sparepart", requestdefault)
  .then(response => response.json())
  .then(result => buatOpsi(result))
  .catch(error => console.error(error));


function buatOpsi(data){
    console.log(data)
    var pilihan = `<option value="">Pilih Sparepart</option>`
    for (var i = 0; i < data.length;i++){
        pilihan += `<option value="${data[i].Id_sparepart}">${data[i].Id_sparepart} ${data[i].Nama_sparepart} ${data[i].Jenis_motor}</option>`
    }
    console.log(pilihan)
    Opsi.innerHTML = pilihan
    body.replace(Opsi)
}
  

function editsparepart(){
    var id = Opsi.value
    var nama = document.getElementById('nama-sparepart') .value
    var motor = document.getElementById('jenis-motor').value
    var harga = document.getElementById('harga').value
    const send = JSON.stringify({
        "Id_sparepart": id,
        "Nama_sparepart": nama,
        "Jenis_motor": motor,
        "Harga_sparepart":harga 
      })
      const request ={
        method: "POST",
        body: send,
        redirect: "follow",
        mode: "cors"
      }
      fetch(url + "/edit/sparepart", request)
}