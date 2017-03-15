<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>use ws socket</title>
</head>
<body>
<script type="text/javascript">
  var sock = null;
  var wsuri = "ws://127.0.0.1:1234/ws";
  window.onload = function() {
    console.log("onload")
    sock = new WebSocket(wsuri)
    sock.onopen = function() {
      console.log("connected to " + wsuri)
    }
    sock.onclose = function(e) {
      console.log("connected closed (" + e.code + ")")
    }
    sock.onmessage = function(e) {
      console.log("message received: " + e.data)
    }
  };
  function send() {
    var msg = document.getElementById("message").value;
    sock.send(msg)
  }
</script>
<h1>websocket echo test</h1>
<form>
  <p>Message:<input id="message" type="text" value="Hello,cls"/></p>
</form>
<button onclick="send();"> send</button>
</body>
</html>