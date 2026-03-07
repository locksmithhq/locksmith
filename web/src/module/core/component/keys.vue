<template>
  <!-- Right Side: Visual -->
  <v-col
    cols="12"
    md="6"
    class="d-none d-md-flex flex-column align-center justify-center position-relative overflow-hidden"
    style="background: linear-gradient(145deg, #0f172a 0%, #1e3a5f 55%, #1d4ed8 100%)"
  >
    <!-- Decorative circles -->
    <div class="deco-circle deco-1"></div>
    <div class="deco-circle deco-2"></div>
    <div class="deco-circle deco-3"></div>

    <!-- Keys illustration -->
    <div class="key-container mb-8">
      <svg
        width="480"
        height="480"
        viewBox="0 -100 500 500"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        style="filter: drop-shadow(0 32px 64px rgba(0, 0, 0, 0.4))"
      >
        <defs>
          <linearGradient id="gold-gradient" x1="0%" y1="0%" x2="100%" y2="100%">
            <stop offset="0%" stop-color="#FCD34D" />
            <stop offset="50%" stop-color="#F59E0B" />
            <stop offset="100%" stop-color="#D97706" />
          </linearGradient>
          <linearGradient id="gold-light" x1="0%" y1="0%" x2="100%" y2="100%">
            <stop offset="0%" stop-color="#FEF3C7" />
            <stop offset="100%" stop-color="#FCD34D" />
          </linearGradient>
          <linearGradient id="gold-dark" x1="0%" y1="0%" x2="100%" y2="100%">
            <stop offset="0%" stop-color="#D97706" />
            <stop offset="100%" stop-color="#92400E" />
          </linearGradient>
        </defs>

        <!-- Keyring -->
        <circle cx="250" cy="100" r="80" stroke="url(#gold-gradient)" stroke-width="10" fill="none" />

        <!-- 1. Golden Car Key Fob -->
        <g class="key-sway-1" transform-origin="250 170">
          <g transform="translate(195, 145) rotate(20)">
            <circle cx="35" cy="-10" r="10" stroke="url(#gold-gradient)" stroke-width="4" fill="none" />
            <rect x="0" y="0" width="70" height="120" rx="25" fill="url(#gold-dark)" />
            <rect x="5" y="5" width="60" height="110" rx="20" fill="url(#gold-gradient)" />
            <circle cx="35" cy="35" r="14" fill="url(#gold-light)" opacity="0.8" />
            <path d="M29 35 L41 35 M35 29 L35 41" stroke="#92400E" stroke-width="2" stroke-linecap="round" />
            <circle cx="35" cy="70" r="14" fill="url(#gold-light)" opacity="0.8" />
            <path d="M28 70 L42 70" stroke="#92400E" stroke-width="2" stroke-linecap="round" />
          </g>
        </g>

        <!-- 2. Antique Skeleton Key -->
        <g class="key-sway-2" transform-origin="250 170">
          <g transform="translate(255, 165) rotate(-10)">
            <circle cx="20" cy="-10" r="8" stroke="url(#gold-gradient)" stroke-width="3" fill="none" />
            <path d="M20 0 C 5 0 0 15 0 25 C 0 35 10 40 20 40 C 30 40 40 35 40 25 C 40 15 35 0 20 0 Z M 20 10 C 25 10 28 15 28 25 C 28 32 25 35 20 35 C 15 35 12 32 12 25 C 12 15 15 10 20 10 Z" fill="url(#gold-gradient)" />
            <rect x="16" y="40" width="8" height="100" fill="url(#gold-gradient)" />
            <path d="M16 110 L 35 110 L 35 125 L 28 125 L 28 130 L 35 130 L 35 135 L 16 135 Z" fill="url(#gold-gradient)" />
          </g>
        </g>

        <!-- 3. Standard House Key -->
        <g class="key-sway-3" transform-origin="250 170">
          <g transform="translate(290, 140) rotate(-25)">
            <circle cx="20" cy="-10" r="8" stroke="url(#gold-gradient)" stroke-width="3" fill="none" />
            <circle cx="20" cy="20" r="22" fill="url(#gold-gradient)" />
            <circle cx="20" cy="20" r="8" fill="rgba(255,255,255,0.25)" />
            <rect x="12" y="35" width="16" height="70" fill="url(#gold-gradient)" />
            <path d="M28 60 L38 60 L38 70 L28 70 L28 80 L38 80 L38 90 L28 90" fill="url(#gold-gradient)" />
          </g>
        </g>

        <!-- 4. Small Padlock Key -->
        <g class="key-sway-4" transform-origin="250 170">
          <g transform="translate(155, 135) rotate(35)">
            <circle cx="15" cy="-8" r="6" stroke="url(#gold-gradient)" stroke-width="3" fill="none" />
            <circle cx="15" cy="15" r="15" fill="url(#gold-gradient)" />
            <rect x="10" y="25" width="10" height="40" fill="url(#gold-gradient)" />
            <rect x="20" y="50" width="8" height="10" fill="url(#gold-gradient)" />
          </g>
        </g>
      </svg>
    </div>

    <!-- Brand text -->
    <div class="brand-block text-center">
      <p class="brand-eyebrow mb-2">LOCKSMITH</p>
      <h2 class="brand-heading mb-3">Your Identity Platform</h2>
      <p class="brand-description mb-8">
        Self-hosted IAM with OAuth 2.0,<br>multi-tenant projects and RBAC.
      </p>

      <div class="feature-list">
        <div v-for="f in features" :key="f.text" class="feature-pill">
          <v-icon :icon="f.icon" size="14" style="color: #fbbf24" class="mr-2" />
          <span>{{ f.text }}</span>
        </div>
      </div>
    </div>
  </v-col>
</template>

<script setup>
const features = [
  { icon: 'mdi-shield-lock-outline', text: 'OAuth 2.0 & PKCE' },
  { icon: 'mdi-account-multiple-outline', text: 'Multi-tenant management' },
  { icon: 'mdi-key-variant', text: 'Role-based access control' },
]
</script>

<style scoped>
.deco-circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.05);
  pointer-events: none;
}
.deco-1 { width: 500px; height: 500px; top: -160px; right: -160px; }
.deco-2 { width: 340px; height: 340px; bottom: -100px; left: -100px; }
.deco-3 { width: 180px; height: 180px; top: 42%; left: 6%; background: rgba(255,255,255,0.03); }

.key-container {
  position: relative;
  z-index: 1;
  animation: float 8s ease-in-out infinite;
}

.brand-block {
  position: relative;
  z-index: 1;
}

.brand-eyebrow {
  font-size: 11px;
  font-weight: 800;
  letter-spacing: 0.22em;
  color: rgba(255, 255, 255, 0.45);
  margin: 0;
}

.brand-heading {
  font-size: 1.5rem;
  font-weight: 700;
  color: #ffffff;
  line-height: 1.3;
}

.brand-description {
  font-size: 0.875rem;
  color: rgba(255, 255, 255, 0.55);
  line-height: 1.7;
  margin: 0;
}

.feature-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}

.feature-pill {
  display: inline-flex;
  align-items: center;
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 999px;
  padding: 6px 16px;
  font-size: 12px;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(4px);
}

.key-sway-1 { animation: sway-1 5s ease-in-out infinite; }
.key-sway-2 { animation: sway-2 6s ease-in-out infinite; }
.key-sway-3 { animation: sway-3 7s ease-in-out infinite; }
.key-sway-4 { animation: sway-4 4s ease-in-out infinite; }

@keyframes float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-28px); }
}
@keyframes sway-1 {
  0%, 100% { transform: rotate(0deg); }
  50% { transform: rotate(8deg); }
}
@keyframes sway-2 {
  0%, 100% { transform: rotate(0deg); }
  50% { transform: rotate(-5deg); }
}
@keyframes sway-3 {
  0%, 100% { transform: rotate(0deg); }
  50% { transform: rotate(10deg); }
}
@keyframes sway-4 {
  0%, 100% { transform: rotate(0deg); }
  50% { transform: rotate(-8deg); }
}
</style>
