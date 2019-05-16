FROM centos:centos7
ADD . /opt/
RUN chgrp -R 0 /opt \
    && chmod -R g+rwX /opt
CMD [ "/opt/run.sh" ]
