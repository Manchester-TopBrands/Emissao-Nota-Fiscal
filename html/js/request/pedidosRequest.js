
function headers() {
  var h = new Headers()
  h.set("Token", localStorage.getItem("Token"))
  return h
}
function rest2(body) {
    var myInit = {
      method: "POST",
      headers: headers(),
      mode: "cors",
      cache: "default",
      body: body,
    };
    fetch(
      `${window.location.protocol}//${window.location.host}/addExcel`,
      myInit
    ) //fetch('http://localhost:8080/add', myInit)
      .then((response) => {
        response.blob().then((b) => {
          b.text().then((data) => {
            let cliente = JSON.parse(data);

            const trhead = document.getElementById("thead");
            trhead.innerHTML = "";
            var a = [
              "#",
              "Produto",
              "Estoque Inicial",
              ...cliente.clientes,
              "Estoque Final",
            ];
            a.forEach((col, index) => {
              var th = document.createElement("th");
              th.setAttribute("scope", "col");
              th.style.fontColor = "black";
              if (index > 2 && index < a.length - 1) {
                //th.style.backgroundColor=""
              }
              th.innerHTML = col;
              trhead.appendChild(th);
            });
            const trbody = document.getElementById("tbody");
            trbody.innerHTML = "";
            var produtosNegativo = false;
            var button = document.getElementById("buttonconfirmed");
            var notbutton = document.getElementById("button!confirmed");
            cliente.items.forEach((item, index) => {
              var tr = document.createElement("tr");
              tr.setAttribute("scope", "row");
              var ef = item.qtds[0];
              for (var i = 1; i < item.qtds.length; i++) {
                ef -= item.qtds[i];
              }
              var tds = [index + 1, item["cod-Desc"], ...item.qtds];
              tds.forEach((col, index2) => {
                var td = document.createElement("td");
                td.innerHTML = col;
                tr.appendChild(td);
              });
              var td = document.createElement("td");
              if (ef > -0.0000001 && ef < 0) {
                ef = 0
              }
              td.innerHTML = ef;
              if (ef < 0) {
                td.style.color = "red";
                produtosNegativo = true;
              }
              tr.appendChild(td);
              trbody.appendChild(tr);
            });
            if (produtosNegativo) {
              notbutton.setAttribute("class", "card mb-4");
              button.setAttribute("class", "container my-auto d-none");
            } else {
              id = cliente.id
              button.setAttribute("class", "container my-auto");
              notbutton.setAttribute("class", "card mb-4 d-none");
            }
          });
        });
      });
}

// envio de pedidos
function enviarExcel() {
    var input = document.getElementById("file");
    console.log(input.files.length);
    if (input.files.length > 0) {
      let formData = new FormData();
      formData.append("file", input.files[0], input.files[0].name);
      //console.log("file", input.files[0], input.files[0].name);
      rest2(formData);
    }
    input.value = "";
}