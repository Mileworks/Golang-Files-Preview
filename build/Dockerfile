FROM sgrio/java:jre_8_ubuntu
MAINTAINER long0419 "borrip0419@gmail.com"

ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8

WORKDIR /opt/

EXPOSE 8079

COPY fonts/* /usr/share/fonts/chienes/

ADD plm-files-preview.tar.gz .

# libreoffice
RUN apt-get update  && apt-get install -y locales &&\
    apt-get install -y locales && apt-get install -y language-pack-zh-hans &&\
	localedef -i zh_CN -c -f UTF-8 -A /usr/share/locale/locale.alias zh_CN.UTF-8 && locale-gen zh_CN.UTF-8 &&\
	apt-get install -y tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
	apt-get install -y libxrender1 && apt-get install -y libxt6 && apt-get install -y libxext-dev && apt-get install -y libfreetype6-dev &&\
	apt-get install -y wget && apt-get install -y ttf-mscorefonts-installer && apt-get install -y fontconfig &&\
	cd /tmp &&\
    apt-get update && \
    apt-get -y -q install \
        libreoffice \
        libreoffice-writer \
        ure \
        libreoffice-java-common \
        libreoffice-core \
        libreoffice-common \
        openjdk-8-jre \
        fonts-opensymbol \
        hyphen-fr \
        hyphen-de \
        hyphen-en-us \
        hyphen-it \
        hyphen-ru \
        fonts-dejavu \
        fonts-dejavu-core \
        fonts-dejavu-extra \
        fonts-droid-fallback \
        fonts-dustin \
        fonts-f500 \
        fonts-fanwood \
        fonts-freefont-ttf \
        fonts-liberation \
        fonts-lmodern \
        fonts-lyx \
        fonts-sil-gentium \
        fonts-texgyre \
        fonts-tlwg-purisa && \
    apt-get -y -q remove libreoffice-gnome && \
    apt -y autoremove && \
    rm -rf /var/lib/apt/lists/* && \
    cd /usr/share/fonts/chienes &&\
    mkfontscale &&\
    mkfontdir &&\
    fc-cache -fv

RUN adduser --home=/opt/libreoffice --disabled-password --gecos "" --shell=/bin/bash libreoffice

ENTRYPOINT ["nohup","/opt/plm-files-preview","&"]
