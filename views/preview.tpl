<html>

<head>
    <title>{{.plm}}</title>
    <style>
        .pdf {
            height: 30rem;
            border: 1rem solid rgba(0, 0, 0, .1);
        }
    </style>
</head>

<body>
    <div id="pdf"></div>
</body>
<script src="https://cdn.bootcdn.net/ajax/libs/pdfobject/2.1.1/pdfobject.min.js"></script>
<script>PDFObject.embed("{{.url}}", "#pdf");</script>

</html>