FROM scratch

COPY cmd/cmd .

EXPOSE 8089

ENTRYPOINT [ "./cmd" ]