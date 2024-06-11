
const url ="http://127.0.0.1:8080";
const body = document.body
let table = document.getElementById('tables')

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

function kirimidsparepart(data){
  console.log(data)
}

