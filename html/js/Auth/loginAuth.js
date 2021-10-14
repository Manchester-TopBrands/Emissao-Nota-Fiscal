var token;
var username;

// REQUEST API
async function restlogin(body) {
  var myInit = criarInit(body,'POST')
  const response = await fetch(
    `${window.location.protocol}//${window.location.host}/addLogin`,
    myInit
  );
  const data = JSON.parse(await (await response.blob()).text());
  if (response.status == "200") {
    token = response.headers.get("Token");
    username = response.headers.get("Name");
    document.cookie = ("Token=" + token + "; path=/");
    document.cookie = ("Name=" + username + "; path=/");
    location.assign(`${window.location.protocol}//${window.location.host}/html/index2.html`)
  }
  alert(data.error)
 
}

// BOTÃO DE ENVIO 
function enviarLogin() {
  var username = getElement("userEmail").value
  username = username.toLowerCase()
  var userpassword = getElement("userPassword").value
  let body = criarObjeto();
  body.username = username
  body.userpassword = userpassword
  data = JSON.stringify(body)
  restlogin(data);
}
// SET TOKEN ON LOCALSTORAGE VS COOKIE
function headers() {
  var h = new Headers()
  return h
}
