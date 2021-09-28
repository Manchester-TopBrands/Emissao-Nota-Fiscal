
function headers() {
  var h = new Headers()
  h.set("Token", localStorage.getItem("Token"))
  return h
}
// resposta URL para Download automático 
async function downloadFile(name, data) {
    var myInit = {
      method: "GET",
      headers: headers(),
      mode: "cors",
      cache: "default",
    };
    var values = [[data.xml, "xml"], [data.pdf, "pdf"]]
    
    for (var i =0; i<values.length; i++) {
      const response = await fetch(values[i][0], myInit);
      var myBlob = await response.blob();
      var a = document.createElement("a");
      var url = window.URL.createObjectURL(myBlob);
      a.href = url;
      a.download = name +"."+ values[i][1]
      a.click();
      a.remove();
      window.URL.revokeObjectURL(url);

      var a = document.createElement("a");
          a.setAttribute("class", "badge badge-success");
          a.innerHTML = "Sucesso";
      var td2 = document.getElementById(name+ values[i][1])
      while (td2.firstChild) {
        td2.firstChild.remove();
      }

      td2.appendChild(a);
     // var column1 = document.getElementById("th1").value
      

    }
}

  