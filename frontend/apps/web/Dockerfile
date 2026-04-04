FROM oven/bun AS build
WORKDIR /app
# Copy root workspace files
COPY frontend/package.json frontend/bun.lock ./
# Copy packages directory (includes api-client)
COPY frontend/packages ./packages
# Copy web app
COPY frontend/apps/web ./apps/web
# Copy scripts
COPY scripts ./scripts
# Install all workspace dependencies
RUN bun install
# Build the web app
WORKDIR /app/apps/web
<<<<<<< HEAD
#RUN bun run build
=======
# RUN bun run build
>>>>>>> c18248eb13dc30aef2033d2c0f99f8ab05893c4b
RUN bun x tsc -b && bun x vite build
# Production stage
FROM nginx:alpine
COPY --from=build /app/apps/web/dist /usr/share/nginx/html
COPY frontend/apps/web/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]