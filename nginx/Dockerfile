# Use the official Nginx 1.21.6 image as the base image
FROM nginx:1.21.6

# Set the working directory to /websites
WORKDIR /websites

# Copy the startup script to the working directory and make it executable
COPY ./nginx-startup.sh /websites/nginx-startup.sh
RUN chmod +x /websites/nginx-startup.sh

# Install Vim for debugging purposes
RUN apt-get update && apt-get install -y vim

# Create the sites-available and sites-enabled directories
RUN mkdir -p /etc/nginx/sites-available
RUN mkdir -p /etc/nginx/sites-enabled

# Disable the default Nginx configuration by renaming it
RUN mv /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.disabled

# Create a default virtual host in the sites-available directory by copying the disabled configuration
RUN cp /etc/nginx/conf.d/default.conf.disabled /etc/nginx/sites-available/default

# Create directories for a sample website
RUN mkdir -p /websites/site-one

# Update the main Nginx configuration file to include virtual hosts from the sites-enabled directory
RUN sed 's@include /etc/nginx/conf.d/\*.conf;@include /etc/nginx/conf.d/*.conf;\n    include /etc/nginx/sites-enabled/*;@g' /etc/nginx/nginx.conf >> ./new_nginx.conf && mv new_nginx.conf /etc/nginx/nginx.conf

# Expose port 80 for HTTP traffic
EXPOSE 80
