<html>
<head>
  <title>gptl头</title>
</head>
<body>
<div>
  <form action="http:localhost:9090/login" method="post">
    <div>
      username:<input type="text" name="username"/>
      password:<input type="password" name="password"/>
      <input type="submit"/>
    </div>
    <div>
      <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banana">banana</option>
      </select>
    </div>
    <div>

      <input type="radio" name="gender" value="1"/>nan
      <input type="radio" name="gender" value="2"/>nv
    </div>

    <div>
      <input type="checkbox" name="interest" value="football">football
      <input type="checkbox" name="interest" value="basketball">basketball
      <input type="checkbox" name="interest" value="tennis">tennis
    </div>
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="file" name="uploadfile"/>
  </form>
</div>
</body>
</html>