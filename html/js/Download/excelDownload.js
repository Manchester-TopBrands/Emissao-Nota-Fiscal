
function headers() {
  var h = new Headers()
  h.set("Token", localStorage.getItem("Token"))
  return h
}
async function downloadStructFile() {
    var myInit = {
      method: "GET",
      headers: headers(),
      mode: "cors",
      cache: "default",
    };
      const response = await fetch( `${window.location.protocol}//${window.location.host}/download`, myInit);
      var myBlob = await response.blob();
      var a = document.createElement("a");
      var url = window.URL.createObjectURL(myBlob);
      a.href = url;
      a.download = "pedidos.xlsx" 
      a.click();
      a.remove();
      window.URL.revokeObjectURL(url);
}