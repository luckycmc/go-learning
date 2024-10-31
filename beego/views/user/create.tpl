<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Document</title>
</head>
<body>
<form action="/user/add_user" method="post">
    {{.xsrfdata}}
    <label>
        name:
        <input type="text" name="name">
    </label>
    <input type="submit" value="submit">
</form>
</body>
</html>