const url ="http://127.0.0.1:8080";
const body = document.body
let tablepelanggan = document.getElementById('tables-pelanggan')



const requestdefaultpelanggan = {
  method:"GET",
  redirect: "follow",
  mode: "cors"
}

fetch(url + "/pelanggan", requestdefaultpelanggan)
  .then(response => response.json())
  .then(result => buatTabelPelanggan(result))
  .catch(error => console.error(error));


function buatTabelPelanggan(data){
  var row = `<thead>
                <tr>
                    <th>Id</th>
                    <th>Nama</th>
                    <th>Plat Nomor</th>
                    <th>Jenis Motor</th>
                    <th>Opsi</th>
                </tr>
                </thead>`
  row += `<tbody>`
  for (var i = 0; i < data.length;i++){
   row +=             `<tr>
                          <td>
                            ${data[i].Id_pelanggan}
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
                            <button id="tombol-hapus" onclick="hapuspelanggan(${data[i].Id_pelanggan})" type="button" class="btn btn-danger btn-sm">Hapus</button>
                          </td>
                      </tr>`
  row += `</tbody>`
  }
  tablepelanggan.innerHTML = row
  body.replace(tablepelanggan)
};

function pelanggancaridarinama(){
  var input = document.getElementById('input').value
  const send = JSON.stringify({
    "Nama_pelanggan": input
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/cari/pelanggan/nama", request)
  .then(response => response.json())
  .then(result => buatTabelPelanggan(result))
  .catch(error => console.error(error));
};

function hapuspelanggan(data){
  console.log(data)
  const send = JSON.stringify({
    "Id_pelanggan": data
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/hapus/pelanggan", request)
}

//<===========================SPAREPART===========================>

let tablesparepart = document.getElementById('tables-sparepart')
const requestdefaultsparepart = {
  method:"GET",
  redirect: "follow",
  mode: "cors"
}

fetch(url + "/sparepart", requestdefaultsparepart)
  .then(response => response.json())
  .then(result => buatTabelSparepart(result))
  .catch(error => console.error(error));

function buatTabelSparepart(data){
  var row = `<thead>
                <tr>
                    <th>Id</th>
                    <th>Nama</th>
                    <th>Jenis Motor</th>
                    <th>Jumlah terjual</th>
                    <th>Pilih</th>
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
                            ${data[i].Jumlah_terjual}
                          </td>
                          <td>
                            <button id="tombol-pilih" onclick="kirimidsparepart('${data[i].Id_sparepart}')" type="button" class="btn btn-primary btn-sm">Pilih</button>
                          </td>
                      </tr>`
  row += `</tbody>`
  }
  tablesparepart.innerHTML = row
  body.replace(tablesparepart)
};

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
  .then(result => buatTabelSparepart(result))
  .catch(error => console.error(error));
};

function kirimidsparepart(data){
  const send = JSON.stringify({
    "Id_sparepart": data
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/cari/pelanggan/pembelian-sparepart", request)
  .then(response => response.json())
  .then(result => hasilTabelPelanggan(result))
  .catch(error => console.error(error));
}
  
//<===================TABEL LANJUTAN===================>


function hasilTabelPelanggan(data){
  const atas = document.getElementById('kotak-pencarian')
  var tambahan = `
              <div id="penambahan">
              <a href="registpelanggan.html" id="tambah">+ Tambah</a>
              <a href="editpelanggan.html" id="tambah" style="margin-right: 10px; padding-left: 5px; padding-right: 5px;">Edit</a>
              </div>`
  
  var row = `<thead>
                <tr>
                    <th>Id</th>
                    <th>Nama</th>
                    <th>Plat Nomor</th>
                    <th>Jenis Motor</th>
                    <th>Opsi</th>
                </tr>
                </thead>`
  row += `<tbody>`
  for (var i = 0; i < data.length;i++){
   row +=             `<tr>
                          <td>
                            ${data[i].Id_pelanggan}
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
                            <button id="tombol-hapus" onclick="hapuspelanggan(${data[i].Id_pelanggan})" type="button" class="btn btn-danger btn-sm">Hapus</button>
                          </td>
                      </tr>`
  row += `</tbody>`
  }
  tablesparepart.innerHTML = row
  atas.innerHTML = tambahan
  body.replace(atas)
  body.replace(tablesparepart)
};

//<================================PENDAFTARAN PELANGGAN================================>

function tambahpelanggan(){
  var nama = document.getElementById('nama-pelanggan') .value
  var motor = document.getElementById('jenis-motor').value
  var plat = document.getElementById('plat-nomor').value
  const send = JSON.stringify({
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
    fetch(url + "/tambah/pelanggan", request)
}