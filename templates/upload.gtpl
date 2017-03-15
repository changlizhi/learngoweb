<html>
<head>
  <title>upload</title>
</head>
<body>

<div>
  <form enctype="multipart/form-data" action="http://localhost:9090/upload" method="post">
    <div>
      <input type="file" name="uploadfile"/>
      <input type="hidden" name="token" value="{{.}}"/>
      <input type="submit" value="upload"/>
    </div>

  </form>
</div>
</body>
</html>