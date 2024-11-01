<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Create</title>
</head>
<body>
<form action="/v1/article" method="post">
    {{.xsrf}}
    <label>Title
        <input type="text" name="title">
    </label>
    <label>Content
        <input type="text" name="content">
    </label>
    <input type="submit" value="submit">
</form>
</body>
</html>