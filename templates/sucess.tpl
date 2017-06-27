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
        <a href="https://github.com/fronbasal/open-redirect" class="brand-logo center" target="_blank">Open Redirect</a>
    </nav>
    <div class="container">
        <div class="card-panel grey lighten-4 hoverable">
            <h4 class="center">Sucess!</h4>
            <p class="flow-text">Sucessfully saved your domain!</p>
            <p>How to configure your DNS for redirection:</p>
            <table class="table responsive-table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Type</th>
                        <th>Target</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <th>{{ .domain }}</th>
                        <th>CNAME</th>
                        <th>{{ .host }}</th>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.99.0/js/materialize.min.js"></script>
    <script src='https://www.google.com/recaptcha/api.js'></script>
</body>

</html>
{{ end }}