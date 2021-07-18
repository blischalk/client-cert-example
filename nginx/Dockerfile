FROM nginx:latest
RUN apt-get update\
 && apt-get install -y wget
# Install mkcert tool for installing a local CA and generating testing certificates
# NOTE: Storing secrets in a Docker image is bad practice.
# This is being done simply for expedient demo purposes
RUN wget https://github.com/FiloSottile/mkcert/releases/download/v1.4.3/mkcert-v1.4.3-linux-amd64 -O mkcert\
  && chmod 755 mkcert\
  && ./mkcert -install\
  # Generate an ssl cert for the server
  # Will make yourcertname+4-key.pem and yourcertname+4.pem
  && ./mkcert yourcertname testssl.local localhost 127.0.0.1 ::1\
  # Generate an client side ssl cert
  # Will make client-adhoc-client-key.pem and client-adhoc-client.pem
  && ./mkcert -client client-adhoc
EXPOSE 443
COPY nginx.conf /etc/nginx/nginx.conf
