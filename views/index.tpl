<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <title>引流小密圈</title>

    <!-- Bootstrap -->
    <link rel="stylesheet" href="static/css/bootstrap.min.css">

    <!-- HTML5 shim 和 Respond.js 是为了让 IE8 支持 HTML5 元素和媒体查询（media queries）功能 -->
    <!-- 警告：通过 file:// 协议（就是直接将 html 页面拖拽到浏览器中）访问页面时 Respond.js 不起作用 -->
    <!--[if lt IE 9]>
    <script src="https://cdn.jsdelivr.net/npm/html5shiv@3.7.3/dist/html5shiv.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/respond.js@1.4.2/dest/respond.min.js"></script>
    <![endif]-->
</head>
<body>
<div class="container">
    <div class="jumbotron">
        <button type="button" id="jumbotron-close" class="close" aria-label="Close"><span aria-hidden="true">×</span>
        </button>
        <h2>欢迎加入引流小密圈</h2>
        <p>本软件主要为了避免大家扎堆，对商家正常经营造成干扰，也为了达到比较好的扫码体验。<br>
            本项目提供有“引流小密圈”社群，未入群的，联系 <code>hyx2011</code> 拉你入群。
        </p>
        <ul>
            <li><h4>请将新更新的资源解压后放置到 <code>qrcodes</code> 目录下</h4></li>
            <li><h4>已扫码的文件会转移至 <code>scanned</code> 目录下，并自动解密，随时可以再扫</h4></li>
            <li><h4>为避免封控，请勿将本项目提供的资源分享给别人，大家偷偷用即可</h4></li>
            <li><h4>直接点击二维码下方的刷新按钮即可随机切换群二维码，刷新会将当前的二维码转移至 <code>scanned</code> 目录
            </h4></li>
        </ul>
    </div>
    <div class="text-center" style="margin-top: 15px">
        <a href="javascript:;" class="btn btn-primary btn-lg" data-loading-text="Loading..." id="fresh-code" role="button"
           style="width: 160px">刷新二维码</a>
        <a href="javascript:;" class="thumbnail" style="max-width: 305px;margin: 15px auto;">
            <img src="static/WX20240712-112149@2x.png" id="qrcode" alt="" style="max-width: 300px"/>
        </a>
    </div>
</div>

<script src="static/js/jquery.min.js"></script>
<script src="static/js/bootstrap.min.js"></script>

<script>
    (function () {
        $('#jumbotron-close').click(function () {
            $('.jumbotron').remove()
        })

        const $btn = $(this).button('loading');
        $.ajax({
            url: "/",
            method: "post",
            success: function (data) {
                $('#qrcode').attr("src", "/qrcode?img=" + data.img)
                $btn.button('reset')
            },
            error: function (err) {
                alert("没有找到任何资源，请确保 qrcodes 目录下有资源包")
            }
        })

        $('#fresh-code').click(function () {
            $.ajax({
                url: "/",
                method: "post",
                success: function (data) {
                    $('#qrcode').attr("src", "/qrcode?img=" + data.img)
                    $btn.button('reset')
                },
                error: function (err) {
                    alert("没有找到任何资源，请确保 qrcodes 目录下有资源包")
                }
            })
        })
    })()
</script>
</body>
</html>