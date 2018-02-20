FROM scratch

COPY quoteserver config.json /app/
WORKDIR "/app"
EXPOSE 44445
CMD ["./quoteserver"]
