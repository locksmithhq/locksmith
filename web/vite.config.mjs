// Plugins
import terser from '@rollup/plugin-terser'
import Components from 'unplugin-vue-components/vite'
import Vue from '@vitejs/plugin-vue'
import Vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'
import Fonts from 'unplugin-fonts/vite'
import { VitePWA } from 'vite-plugin-pwa'

// Utilities
import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'

function filterFontPreloads() {
  return {
    name: 'filter-font-preloads',
    transformIndexHtml(html) {
      return html.replace(
        /<link rel="preload" as="font" type="font\/(?!woff2)[^"]*"[^>]*>\n?/g,
        ''
      )
    },
  }
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    Vue({
      template: { transformAssetUrls },
    }),
    // https://github.com/vuetifyjs/vuetify-loader/tree/master/packages/vite-plugin#readme
    Vuetify(),
    Components(),
    Fonts({
      fontsource: {
        families: [
          {
            name: 'Roboto',
            weights: [100, 300, 400, 500, 700, 900],
            styles: ['normal', 'italic'],
          },
        ],
      },
    }),
    filterFontPreloads(),
    VitePWA({
      registerType: 'autoUpdate',
      includeAssets: ['favicon.ico', 'looksmith.png'],
      manifest: {
        name: 'Locksmith Security',
        short_name: 'Locksmith',
        description: 'Self-hosted IAM — Identity and Access Management',
        theme_color: '#1a1a2e',
        background_color: '#1a1a2e',
        display: 'standalone',
        start_url: '/',
        icons: [
          {
            src: 'looksmith.png',
            sizes: '192x192',
            type: 'image/png',
          },
          {
            src: 'looksmith.png',
            sizes: '512x512',
            type: 'image/png',
          },
        ],
      },
      workbox: {
        globPatterns: ['**/*.{js,css,html,ico,png,svg,woff2}'],
        navigateFallback: '/index.html',
        navigateFallbackDenylist: [/^\/api\//],
        maximumFileSizeToCacheInBytes: 5 * 1024 * 1024,
        runtimeCaching: [
          {
            urlPattern: /^\/api\//,
            handler: 'NetworkFirst',
            options: {
              cacheName: 'api-cache',
              expiration: {
                maxEntries: 50,
                maxAgeSeconds: 60 * 5,
              },
            },
          },
        ],
      },
    }),
  ],
  build: {
    minify: 'terser',
    terserOptions: {
      compress: {
        drop_console: true,
        drop_debugger: true,
        pure_funcs: ['console.log', 'console.info', 'console.debug', 'console.warn'],
        passes: 2,
      },
      mangle: {
        toplevel: true,
      },
      format: {
        comments: false,
      },
    },
    rollupOptions: {
      plugins: [terser()],
      output: {
        manualChunks(id) {
          if (id.includes('node_modules/vuetify')) return 'vuetify'
          if (id.includes('node_modules/vue') || id.includes('node_modules/vue-router')) return 'vue'
          if (id.includes('node_modules/echarts')) return 'echarts'
        },
      },
    },
  },
  optimizeDeps: {
    exclude: ['vuetify'],
  },
  define: { 'process.env': {} },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('src', import.meta.url)),
    },
    extensions: ['.js', '.json', '.jsx', '.mjs', '.ts', '.tsx', '.vue'],
  },
  server: {
    port: 3000,
    hmr: {
      clientPort: 4000,
    },
    host: '0.0.0.0',
    allowedHosts: ['auth.example.local'],
    watch: {
      usePolling: true,
    },
  },
})
