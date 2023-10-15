# Primera etapa: Instala Node.js
FROM golang AS nodejs

RUN apt-get update && apt-get install -y curl
RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource.gpg.key | apt-key add -
RUN curl -sL https://deb.nodesource.com/setup_19.x | bash -
RUN apt-get install -y nodejs

# Segunda etapa: Construye tu aplicación Go
FROM golang

# Establece el directorio de trabajo en /go/src
WORKDIR /go/src

# Copia el contenido actual del directorio del proyecto al contenedor en /go/src
COPY . .

# Copia Node.js desde la primera etapa a la segunda etapa
COPY --from=nodejs /usr/bin/node /usr/bin/
COPY --from=nodejs /usr/bin/npm /usr/bin/

# Ejecuta el comando "tail -f /dev/null" para mantener el contenedor en ejecución
CMD ["tail", "-f", "/dev/null"]
