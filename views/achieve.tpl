<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>{{.plm}}</title>

    <link rel="stylesheet" href="https://static.runoob.com/assets/js/jquery-treeview/jquery.treeview.css"/>
    <link rel="stylesheet" href="https://static.runoob.com/assets/js/jquery-treeview/screen.css"/>

    <script src="https://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
    <script src="https://static.runoob.com/assets/js/jquery-treeview/jquery.cookie.js"></script>
    <script src="https://static.runoob.com/assets/js/jquery-treeview/jquery.treeview.js"
            type="text/javascript"></script>

    <script type="text/javascript">
        $(document).ready(function () {
            $("#browser").treeview({
                toggle: function () {
                    console.log("%s was toggled.", $(this).find(">span").text());
                }
            });
        });
    </script>
</head>
<body>

<h1 id="banner">{{.base}}</h1>
<div id="main">
    <ul id="browser" class="filetree treeview-famfamfam">
		<li><span class="folder">{{.top}}</span>
			<ul id="{{.top}}">
                {{range $key,$value := .treeData }}
                    <li><a href="{{$value}}" target="_blank"><span class="file">{{$key}}</span></a></li>
                {{end}}
			</ul>
		</li>
    </ul>



</div>

</body>
</html>
