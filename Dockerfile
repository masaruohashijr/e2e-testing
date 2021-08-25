FROM centos
WORKDIR /app
COPY ./config ./config/
COPY ./features ./features/
COPY ./log ./log/
COPY ./temp ./temp/
COPY ./test ./test/
COPY ./xml ./xml/
COPY ./zarbat_tester ./zarbat_tester
COPY ./entrypoint.sh ./entrypoint.sh
RUN chmod 755 /app/entrypoint.sh
RUN chmod 755 /app/zarbat_tester
EXPOSE 5000
ENTRYPOINT ["/app/entrypoint.sh"]