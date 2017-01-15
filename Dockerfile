FROM busybox
COPY gotpl /gotpl
ENV GOTPL_DATA /input/data
ENV GOTPL_OUTPUT ""
ENV GOTPL_TEMPLATE /input/template
ENTRYPOINT /gotpl
CMD -data ${GOTPL_DATA} -output ${GOTPL_OUTPUT} -template ${GOTPL_TEMPLATE}
