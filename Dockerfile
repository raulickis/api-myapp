FROM alpine
COPY main /
COPY .env /
EXPOSE 9990
CMD cd / ; /main
