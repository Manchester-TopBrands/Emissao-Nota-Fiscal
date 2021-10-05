function headers() {
  var h = new Headers()
  h.set("Token", localStorage.getItem("Token"))
  return h
}
function enviarXml() {
    var input = getElement("fileXml");

    if (input.files.length > 0) {
      let formData = new FormData();
      formData.append("file", input.files[0], input.files[0].name);
      rest(formData);
    }
    input.value = "";
}

async function rest(body) {
    var myInit = criarInit(body,'POST')
    const response = await fetch(
      `${window.location.protocol}//${window.location.host}/addEntradas`,
      myInit
    );
    if (response.status == "200") {
      const cardOk = getElement("card!confirmed");
      cardOk.setAttribute("class", "card mb-4 d-none");
      alert("Enviado com Sucesso!!!");
      return;
    }
    const data = JSON.parse(await (await response.blob()).text());
    //"card-!confirmed" ,"data-desc" ,"data-title"
    const cardOk = getElement("card!confirmed");
    var text = [data.status, data.error];
    var title = getElement("datatitle");
    var desc = getElement("datadesc");
    title.innerHTML = text[0];
    desc.innerHTML = text[1];
    cardOk.setAttribute("class", "card mb-4");
    return;
}