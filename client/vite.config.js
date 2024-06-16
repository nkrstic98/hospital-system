import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    server: {
        port: 3000,
        proxy: {
            // Proxying requests from /api to your backend server
            '/api': {
                target: "http://localhost:8080", // Backend server address
                changeOrigin: true,
                rewrite: function (path) { return path.replace(/^\/api/, ''); }, // Remove /api from the path
            },
        },
    },
    preview: {
        host: true,
        port: 3000,
        proxy: {
            // Proxying requests from /api to your backend server
            '/api': {
                target: "http://localhost:8080", // Backend server address
                changeOrigin: true,
                rewrite: function (path) { return path.replace(/^\/api/, ''); }, // Remove /api from the path
            },
        },
    }
});
