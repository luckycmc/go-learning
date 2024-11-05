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
<form action="/v1/article/update/{{.article.Id}}" method="post">
    {{.xsrf}}
    <label>Title
        <input type="text" name="title" value="{{.article.Title}}">
    </label>
    <label>Content
        <input type="text" name="content" value="{{.article.Content}}">
    </label>
    <input type="submit" value="submit">
</form>
</body>
</html>