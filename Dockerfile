FROM iron/base
WORKDIR /app
# copy binary into image
COPY web /app/
COPY index.html /app/
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["./web"]
