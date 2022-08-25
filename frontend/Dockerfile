FROM nginx:1.23-alpine
RUN apk add --update nodejs npm
# RUN npm install -g @angular/cli
COPY ./ /opt/onlyone/frontend
WORKDIR /opt/onlyone/frontend
RUN npm install
RUN npm build
RUN rm -rf /usr/share/nginx/html
RUN mv dist/frontend/ /usr/share/nginx
RUN mv /usr/share/nginx/frontend /usr/share/nginx/html
WORKDIR /opt
RUN rm -rf onlyone

# COPY dist/frontend/ /usr/share/nginx/html

