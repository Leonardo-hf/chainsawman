FROM bitnami/spark:3.3.2

ADD ./jar/common-latest-jar-with-dependencies.jar jars/
RUN rm -rf jars/okhttp-3.12.12.jar jars/okio-1.14.0.jar
USER root

ENV HADOOP_HOME="/opt/hadoop"
ENV HADOOP_CONF_DIR="$HADOOP_HOME/etc/hadoop"
ENV HADOOP_LOG_DIR="/var/log/hadoop"
ENV PATH="$HADOOP_HOME/hadoop/sbin:$HADOOP_HOME/bin:$PATH"

WORKDIR /opt

RUN apt-get update && apt-get install -y openssh-server && apt-get install -y curl

RUN ssh-keygen -t rsa -f /root/.ssh/id_rsa -P '' && \
    cat /root/.ssh/id_rsa.pub >> /root/.ssh/authorized_keys

#COPY ./hadoop/hadoop-3.3.2.tar.gz /opt/
RUN curl -OL https://archive.apache.org/dist/hadoop/common/hadoop-3.3.2/hadoop-3.3.2.tar.gz
RUN tar -xzvf /opt/hadoop-3.3.2.tar.gz && \
  mv hadoop-3.3.2 hadoop && \
  rm -rf /opt/hadoop-3.3.2.tar.gz && \
  mkdir /var/log/hadoop

ADD ./hadoop/ssh/ssh_config /root/.ssh/config
COPY ./hadoop/conf/ $HADOOP_CONF_DIR