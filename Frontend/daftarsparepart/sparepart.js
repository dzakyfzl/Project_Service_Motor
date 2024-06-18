
const url ="http://127.0.0.1:8080";
const body = document.body
let table = document.getElementById('tables-sparepart')

const requestdefault = {
  method:"GET",
  redirect: "follow",
  mode: "cors"
}

fetch(url + "/sparepart", requestdefault)
  .then(response => response.json())
  .then(result => buatTabel(result))
  .catch(error => console.error(error));

function buatTabel(data){
  var row = `<thead>
                <tr>
                    <th>Id</th>
                    <th>Nama</th>
                    <th>Jenis Motor</th>
                    <th>Jumlah terjual</th>
                    <th>Opsi</th>
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
                            <button id="tombol-edit" onclick="" type="button" class="btn btn-primary btn-sm">Edit</button>
                            <button id="tombol-hapus" onclick="hapussparepart('${data[i].Id_sparepart}')" type="button" class="btn btn-danger btn-sm">Hapus</button>
                          </td>
                      </tr>`
  row += `</tbody>`
  }
  table.innerHTML = row
  body.replace(table)
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
  .then(result => buatTabel(result))
  .catch(error => console.error(error));
};

function hapussparepart(data){
  console.log(data)
  const send = JSON.stringify({
    "Id_sparepart": data
  })
  const request ={
    method: "POST",
    body: send,
    redirect: "follow",
    mode: "cors"
  }
  fetch(url + "/hapus/sparepart", request)
};


//<======================TAMBAH SPAREPART=======================>

function tambahsparepart(){
    var nama = document.getElementById('nama-sparepart') .value
    var motor = document.getElementById('jenis-motor').value
    var harga = document.getElementById('harga').value
    const send = JSON.stringify({
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
      fetch(url + "/tambah/sparepart", request)
}