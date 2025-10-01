/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "export",
  rewrites: async () => [
    {
      source: "/api/:path*",
      destination: "http://localhost:8081/api/:path*",
    },
  ],
};
module.exports = nextConfig;
