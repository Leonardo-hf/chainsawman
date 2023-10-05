FROM bitnami/spark:3.3.2

USER root
ENV LIVY_HOME /opt/bitnami/livy
WORKDIR /opt/bitnami/

#ADD apache-livy-0.7.1-incubating-bin.zip .

RUN apt-get update && apt-get install -y curl && apt-get install -y unzip

RUN curl -OL http://archive.apache.org/dist/incubator/livy/0.7.1-incubating/apache-livy-0.7.1-incubating-bin.zip

RUN unzip "apache-livy-0.7.1-incubating-bin" \
    && rm -rf "apache-livy-0.7.1-incubating-bin.zip" \
    && mv "apache-livy-0.7.1-incubating-bin" $LIVY_HOME \
    && mkdir $LIVY_HOME/logs \
    && chown -R 1001:1001 $LIVY_HOME

USER 1001

CMD ["sh", "-c", "/opt/bitnami/livy/bin/livy-server"]