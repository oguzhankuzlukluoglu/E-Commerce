import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: '/api/products/:path*',
        destination: 'http://localhost:8080/products/:path*',
      },
      {
        source: '/api/login',
        destination: 'http://localhost:8081/login',
      },
      {
        source: '/api/register',
        destination: 'http://localhost:8081/register',
      },
      {
        source: '/api/auth/:path*',
        destination: 'http://localhost:8081/auth/:path*',
      },
      // Diğer mikroservisler için de benzer şekilde eklenebilir
    ]
  },
};

export default nextConfig;
