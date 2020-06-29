# PLM-File-Preview

文件预览服务

1、对Word\Excel\PPT\PDF 在线预览  
2、实现CAD标准dwg\dxf 文件预览  
3、压缩文件预览：tar.gz\tar.bzip2\tar.xz\zip\rar\tar\brotli\bzip2\flate\gzip\lz4\snappy\xz\zstandard  
4、支持txt,java,php,py,md,js,css等所有纯文本  
5、支持jpg，jpeg，png，gif等图片预览

效果图

![avatar](/tmp/bg.jpg)

构建镜像步骤  
1、首先构建基础环境镜像  
`docker build -t plm-files-preview-services -f Dockerfile .`  

2、然后执行启动服务指令  
`docker-compose up -d`
