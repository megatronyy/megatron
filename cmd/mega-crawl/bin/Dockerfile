FROM centos

MAINTAINER quwenbin

COPY mega-crawl /opt/crawl/

ENV HOME /opt/crawl
WORKDIR ${HOME}
RUN chmod +x mega-crawl

ENTRYPOINT[ "mega-crawl" ]
