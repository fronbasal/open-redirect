{{ define "success.tpl" }}
<!DOCTYPE html>
<html>

<head>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.99.0/css/materialize.min.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Welcome to open redirect! An open source redirection service.">
    <meta name="author" content="Fronbasal">
    <title>Open Redirect</title>
</head>

<body>
    <nav class="red lighten-2">
        <a href="/" class="brand-logo center">Open Redirect</a>
    </nav>
    <div class="container">
        <div class="card-panel grey lighten-4 hoverable">
            <h4 class="center">Oooops</h4>
            <p class="flow-text">The domain you tried to access ({{ .domain }}) does not exist in our database but is pointed to our server. Maybe try contacting the domain owner for more information.</p>
        </div>
    </div>
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.99.0/js/materialize.min.js"></script>
    <script src='https://www.google.com/recaptcha/api.js'></script>
</body>

</html>
{{ end }}