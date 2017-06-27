{{ define "index.tpl" }}
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
            <h4 class="center">Welcome to Open Redirect!</h4>
            <p class="flow-text">This is a free and open source redirection service to redirect your domain somewhere without the need of a server!</p>
        </div>
        <div class="card-panel grey lighten-4">
            <form action="/add" method="POST">
                <h4 class="center">Add your domain</h4>
                <p>Your domain</p>
                <input name="source" type="text" placeholder="example.com" required>
                <p>Target url</p>
                <input name="target" type="url" placeholder="https://example.com" required>
                <div class="center">
                    <div class="g-recaptcha" data-sitekey="{{ .siteKey }}"></div>
                    <button type="submit" class="btn btn-large waves-effect waves-light">Submit
                            <i class="material-icons right">send</i>
                    </button>
                </div>
            </form>
        </div>
    </div>
    <script type="text/javascript" src="https://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.99.0/js/materialize.min.js"></script>
    <script src='https://www.google.com/recaptcha/api.js'></script>
</body>

</html>
{{ end }}