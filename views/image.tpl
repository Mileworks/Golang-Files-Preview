<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <title>{{.plm}}</title>
    <link rel="stylesheet" href="https://cdn.bootcdn.net/ajax/libs/viewerjs/1.6.0/viewer.min.css">
    <style>
        * {
            margin: 0;
            padding: 0;
        }

        #dowebok {
            width: 800px;
            margin: 0 auto;
            font-size: 0;
        }

        #dowebok li {
            display: inline-block;
            width: 50px;
            height: 50px;
            margin-left: 1%;
            padding-top: 1%;
        }
    </style>
</head>

<body>
<div id="dowebok">
    <img src="{{.url}}">
</div>
<script src="http://code.jquery.com/jquery-3.5.1.min.js"
        integrity="sha256-9/aliU8dGd2tb6OSsuzixeV4y/faTqgFtohetphbbj0="
        crossorigin="anonymous"></script>
<script src="https://cdn.bootcdn.net/ajax/libs/viewerjs/1.6.0/viewer.min.js"></script>
<script>
    var viewer = new Viewer(document.getElementById('dowebok'));
</script>
</body>

</html>