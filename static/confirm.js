
function confirmForm() {
  const file = document.getElementById("upload").value;
  if (file == undefined || file == null || file == "") {
    window.alert("You need to select a file!");
    return false;
  }
  return true;
}
