FROM scratch

COPY quoteserver config.json /app/
WORKDIR "/app"
EXPOSE 44442
CMD ["./quoteserver"]
