FROM ronmi/mingo

COPY kta /usr/bin/kta
WORKDIR /app
ENTRYPOINT ["/usr/bin/kta"]
