FROM quay.io/keycloak/keycloak:21.0 as builder

# Enable health and metrics support
ENV KC_HEALTH_ENABLED=true
ENV KC_METRICS_ENABLED=true

# Configure a database vendor
ENV KC_DB=postgres

WORKDIR /opt/keycloak
# for demonstration purposes only, please make sure to use proper certificates in production instead
# RUN keytool -genkeypair -storepass password -storetype PKCS12 -keyalg RSA -keysize 2048 -dname "CN=server" -alias server -ext "SAN:c=DNS:localhost,IP:127.0.0.1" -keystore conf/server.keystore

# import the theme from the keycloakify build file
# ADD keycloakify-starter/build_keycloak/target/keycloakify-starter-keycloak-theme-4.1.3.jar providers/keycloakify-theme.jar 
RUN /opt/keycloak/bin/kc.sh build

# run the optimized build
FROM quay.io/keycloak/keycloak:21.0.2
COPY --from=builder /opt/keycloak/ /opt/keycloak/
ENTRYPOINT ["/opt/keycloak/bin/kc.sh"]