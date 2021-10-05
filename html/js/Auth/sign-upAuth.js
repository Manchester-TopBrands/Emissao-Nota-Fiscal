
// REQUEST API
async function enviarLogout() {
  var myInit = criarInit(null,'GET')
  const response = await fetch(
    `${window.location.protocol}//${window.location.host}/addLogout`,
    myInit
  );
    document.cookie = ("Token=" + "" + "; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT");
    location.assign(
      `${window.location.protocol}//${window.location.host}/html/login.html`
    );
}
// BOTÃO DE ENVIO
// SET TOKEN ON LOCALSTORAGE VS COOKIE
