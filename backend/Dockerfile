ARG VERSION=0.1.0-SNAPSHOT

FROM maven:3.8.2-openjdk-16-slim AS builder
ARG VERSION
WORKDIR /usr/src/app

COPY pom.xml .
RUN mvn dependency:go-offline -B

COPY ./src ./src

RUN mvn clean package -DskipTests -DVERSION=${VERSION} \
    && cp target/sausage-store-${VERSION}.jar target/sausage-store.jar

FROM amazoncorretto:17-alpine
WORKDIR /app
RUN addgroup --system jaruser \
    && adduser -S -s /bin/false -G jaruser jaruser -D -H \
    && apk add --no-cache dumb-init curl \
    && mkdir -p /logs/reports \
    && chown -R jaruser:jaruser /logs
COPY --chown=jaruser:jaruser --from=builder /usr/src/app/target/sausage-store.jar ./sausage-store.jar
EXPOSE 8080
USER jaruser
ENTRYPOINT ["dumb-init", "java", "-jar", "-Dmyserver.bindPort=8080", "./sausage-store.jar"]