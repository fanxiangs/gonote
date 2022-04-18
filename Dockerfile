FROM fellah/gitbook
ENV BOOKDIR /srv/gitbook
WORKDIR $BOOKDIR
VOLUME $BOOKDIR
RUN sed -i 's@confirm: true@confirm: false@g' /root/.gitbook/versions/3.2.0/lib/output/website/copyPluginAssets.js \
    && npm install -g gitbook-summary
EXPOSE 4000
CMD ["gitbook", "--help"]

