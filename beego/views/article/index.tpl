<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>List</title>
</head>
<body>
<div>
    <ul>
        {{range .articles}}
        <li>
            {{.Title}}
            {{.Content}}
        </li>
        {{end}}
    </ul>
</div>
</body>
</html>