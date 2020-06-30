<html>

<head>
    <title>{{.plm}}</title>
    <link href="https://unpkg.com/video.js/dist/video-js.min.css" rel="stylesheet">
    <script src="https://unpkg.com/video.js/dist/video.min.js"></script>
</head>

<body>
    <video id="xttblog" class="video-js vjs-default-skin" controls preload="none" width="950" height="664"
           data-setup="{}">
        <source src="{{.url}}}" type='video/mp4' />
    </video>

    <script type="text/javascript">
        var myPlayer = videojs('xttblog');
        videojs("example_video1").ready(function(){
            var myPlayer = this;
            myPlayer.play();
        });
    </script>
</body>
</html>