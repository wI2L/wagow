FROM       scratch
MAINTAINER William Poussier <william.poussier@gmail.com>
ADD        wagow wagow
ENV        PORT 8080
EXPOSE     8080
ENTRYPOINT ["/wagow"]
