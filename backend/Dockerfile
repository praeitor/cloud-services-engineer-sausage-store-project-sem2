#Это так называемый мультистейдж docker-файл. Сначала делаем образ, который соберёт нам java-приложение,
#а затем из него берём собранный артефакт и кладём его в новый образ, чтобы оставить в нём меньше лишнего

# build
FROM maven:3.9.6-eclipse-temurin-17 AS builder
# задаём переменную VERSION, которая будет использоваться для сборки проекта
### ARG VERSION=${VERSION} # временно убираем, в pom уже поправил имя будет пока статичное
WORKDIR /usr/src/app
COPY ./src ./src
### COPY *.xml ./ # пока убраем
COPY pom.xml .
COPY checkstyle.xml .

###RUN mvn package -Dversion.application=${VERSION}
RUN mvn clean package -DskipTests

# release
FROM eclipse-temurin:17-jre
### ARG VERSION=${VERSION}
WORKDIR /app
# создаём пользователя jaruser
RUN apt-get update && apt-get install -y curl dumb-init \
  && rm -rf /var/lib/apt/lists/*

RUN groupadd -r jaruser && useradd -r -g jaruser jaruser \
  && mkdir -p /logs/reports \
  && chown -R jaruser:jaruser /logs
  
COPY --chown=jaruser:jaruser \
     --from=builder /usr/src/app/target/sausage-store.jar ./sausage-store.jar
###     --from=builder /usr/src/app/target/sausage-store-${VERSION}.jar ./sausage-store.jar
EXPOSE 8080
# приложение будет запускаться под пользователем jaruser
USER jaruser
ENTRYPOINT ["dumb-init", "java", "-jar", "-Dmyserver.bindPort=8080", "./sausage-store.jar"]