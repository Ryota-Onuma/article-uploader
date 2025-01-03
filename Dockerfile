FROM golang:1.23.4-bullseye

RUN useradd -m -s /bin/bash app && \
    apt update && \
    mv /usr/local/go /home/app/go

ENV PATH=$PATH:/home/app/go/bin

RUN chown -R -v app:app /home/app

USER app

WORKDIR /home/app/workspace

COPY --chown=app:app . .

EXPOSE 3000

CMD ["go", "run", "cmd/main.go"]