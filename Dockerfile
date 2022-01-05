FROM fellah/gitbook
ENV BOOKDIR /srv/gitbook
WORKDIR $BOOKDIR
VOLUME $BOOKDIR
EXPOSE 4000
CMD ["gitbook", "--help"]

