const url ="http://127.0.0.1:8080";
const body = document.body
let tableservis = document.getElementById('tables-servis')



const requestdefaultservis = {
  method:"GET",
  redirect: "follow",
  mode: "cors"
}

fetch(url + "/servis", requestdefaultservis)
  .then(response => response.json())
  .then(result => buatTabelServis(result))
  .catch(error => console.error(error));


function buatTabelServis(data){
  var row = `<thead>
                    <tr>
                        <th>Id</th>
                        <th>Nama Pelanggan</th>
                        <th>Plat Nomor</th>
                        <th>Jenis Motor</th>
                        <th>Tanggal Servis</th>
                        <th>Biaya</th>
                        <th>Opsi</th>
                    </tr>
                </thead>`
  row += `<tbody>`
  for (var i = 0; i < data.length;i++){
   row +=             `<tr>
                          <td>
                            ${data[i].Id_servis}
                          </td>
                          <td>
                            ${data[i].Nama_pelanggan}
                          </td>
                          <td>
                            ${data[i].Nomor_plat}
                          </td>
                          <td>
                            ${data[i].Jenis_motor}
                          </td>
                          <td>
                            ${data[i].Tanggal_kunjungan.Tanggal} ${data[i].Tanggal_kunjungan.Bulan} ${data[i].Tanggal_kunjungan.Tahun}
                          </td>
                          <td>
                            Rp.${data[i].Total_Harga}
                          </td>
                          <td>
                            <button id="tombol-edit" onclick="tampilkanpesanan('${data[i].Id_servis}')" type="button" class="btn btn-primary btn-sm">Pesanan</button>
                            <button id="tombol-hapus" onclick="hapusservis('${data[i].Id_servis}')" type="button" class="btn btn-danger btn-sm">Hapus</button>
                          </td>
                      </tr>`
  row += `</tbody>`
  }
  tableservis.innerHTML = row
  body.replace(tableservis)
};

//<=========================== pesanan ===================================>

function tampilkanpesanan(data){
    console.log(data)
    const atas = document.getElementById('kotak-pencarian')
    var penambahan = `<button id="tambah-pesanan" onclick="tambahpesanan('${data}')">+ Tambah</button>`
    var row = `<thead>
                    <tr>
                        <th>Id</th>
                        <th>Nama Sparepart</th>
                        <th>Jenis Motor</th>
                        <th>Harga</th>
                        <th>Hapus</th>
                    </tr>
                </thead>`
    row += `<tbody>`
    
    const send = JSON.stringify({
      "Id_servis": data
    })
    console.log(send)
    const request ={
      method: "POST",
      body: send,
      redirect: "follow",
      mode: "cors"
    }
    fetch(url + "/pesanan", request)
        .then(response => response.json())
        .then(result => buatTabelPesanan(result,data))
        .catch(error => console.error(error));
    atas.innerHTML = penambahan
    tableservis.innerHTML = row
    body.replace(tableservis)
    body.replace(atas)
};

function tambahpesanan(servis){
  console.log(servis)
  const atas = document.getElementById('kotak-pencarian')
  var penambahan = `
            <input type="text" id="input" placeholder="Nama Sparepart">
            <button id="search" onclick="sparepartcaridarinama()">Search</button>`
  
  const request = {
    method:"GET",
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/sparepart", request)
  .then(response => response.json())
  .then(result => buatTabelSparepart(result,servis))
  .catch(error => console.error(error));
  atas.innerHTML = penambahan
  body.replace(atas)
}

function buatTabelSparepart(data,servis){
  const tampilan = document.getElementById('tables-servis')
  var row = ``
  for (var i = 0; i < data.length;i++){
    console.log(data[i])
    console.log(servis)
   row +=             `<tr>
                          <td>
                            ${data[i].Id_sparepart}
                          </td>
                          <td>
                            ${data[i].Nama_sparepart}
                          </td>
                          <td>
                            ${data[i].Jenis_motor}
                          </td>
                          <td>
                            ${data[i].Jumlah_terjual}
                          </td>
                          <td>
                            <button id="tombol-pilih" onclick="pilihsparepart('${servis}','${data[i].Id_sparepart}')" type="button" class="btn btn-primary btn-sm">Pilih</button>
                          </td>
                      </tr>`
  row += `</tbody>`
  }
  tampilan.innerHTML = row
  body.append(tampilan)
};

function pilihsparepart(servis,sparepart){
  console.log(servis)
  console.log(sparepart)
  const send = JSON.stringify({
    "Id_sparepart": sparepart,
    "Id_servis" : servis
  })
  const request = {
    method:"POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/tambah/pesanan", request)
}

function sparepartcaridarinama(){
  var input = document.getElementById('input').value
  const send = JSON.stringify({
    "Nama_sparepart": input
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/cari/sparepart/nama", request)
  .then(response => response.json())
  .then(result => tambahpesanan(result))
  .catch(error => console.error(error));
};


function buatTabelPesanan(data,servis){
    console.log(servis)
    const atas = document.getElementById('kotak-pencarian')
    var penambahan = `<button id="tambah-pesanan" onclick="tambahpesanan('${servis}')">+ Tambah</button>`
    var row = `<thead>
                    <tr>
                        <th>Id</th>
                        <th>Nama Sparepart</th>
                        <th>Jenis Motor</th>
                        <th>Harga</th>
                        <th>Hapus</th>
                    </tr>
                </thead>`
  row += `<tbody>`
  for (var i = 0; i < data.length;i++){
   row +=             `<tr>
                          <td>
                            ${data[i].Id_sparepart}
                          </td>
                          <td>
                            ${data[i].Nama_sparepart}
                          </td>
                          <td>
                            ${data[i].Jenis_motor}
                          </td>
                          <td>
                            ${data[i].Harga_sparepart}
                          </td>
                          <td>
                            <button id="tombol-hapus" onclick="hapuspesanan('${data[i].Id_sparepart}','${servis}')" type="button" class="btn btn-danger btn-sm">Hapus</button>
                          </td>
                      </tr>`
  row += `</tbody>`
  }
  atas.innerHTML = penambahan
  tableservis.innerHTML = row
  body.replace(tableservis)
  body.replace(atas)
}

function hapuspesanan(data,servis){
  const send = JSON.stringify({
    "Id_sparepart": data,
    "Id_servis" : servis
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/hapus/pesanan", request)
}

//<==================BATAS AKHIR PESANAN=========================>


function serviscaridaritanggal(){
  var tanggal = document.getElementById('tanggal').value
  var bulan = document.getElementById('bulan').value
  var tahun = document.getElementById('tahun').value
  const send = JSON.stringify({
    "Tanggal": parseInt(tanggal),
    "Bulan":bulan,
    "Tahun":parseInt(tahun)
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/cari/servis/tanggal", request)
  .then(response => response.json())
  .then(result => buatTabelServis(result))
  .catch(error => console.error(error));
};

function hapusservis(data){
  console.log(data)
  const send = JSON.stringify({
    "Id_servis": data
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/hapus/servis", request)
}

//<================ PENAMBAHAN SERVIS ================>

var opsipelanggan = document.getElementById('form-nama-pelanggan')

const requestdefaultpelanggan = {
    method:"GET",
    redirect: "follow",
    mode: "cors"
}

fetch(url + "/pelanggan", requestdefaultpelanggan)
  .then(response => response.json())
  .then(result => buatOpsi(result))
  .catch(error => console.error(error));

function buatOpsi(data){
    var hasil = `<option value="">Nama Pelanggan</option>`
    for (var i = 0; i < data.length;i++){
        hasil += `<option value="${data[i].Id_pelanggan}">${data[i].Nama_pelanggan}</option>`
    }
    opsipelanggan.innerHTML = hasil
    body.replace(opsipelanggan)
}

function tambahservis(){
    var Id = opsipelanggan.value
    var tanggal = document.getElementById('form-tanggal').value
    var bulan = document.getElementById('form-bulan').value
    var tahun = document.getElementById('form-tahun').value
    const send = JSON.stringify({
        "Id_pelanggan": parseInt(Id),
        "Tanggal_kunjungan":
        {"Tanggal": parseInt(tanggal),
        "Bulan":bulan,
        "Tahun":parseInt(tahun)}
      })
      const request ={
        method: "POST",
        body: send,
        redirect: "follow",
        mode: "cors"
      }
      fetch(url + "/tambah/servis", request)
}