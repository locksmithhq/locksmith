<template>
  <!-- Right Side: Visual -->
  <v-col
    cols="12"
    md="6"
    class="d-none d-md-flex align-center justify-center position-relative overflow-hidden"
    style="background: linear-gradient(135deg, #bae6fd 0%, #60a5fa 100%)"
  >
    <!-- Decorative Background Elements -->
    <div
      class="position-absolute"
      style="
        top: 10%;
        right: 10%;
        width: 2px;
        height: 100px;
        background: rgba(255, 255, 255, 0.3);
      "
    ></div>
    <div
      class="position-absolute"
      style="
        bottom: 20%;
        left: 20%;
        width: 2px;
        height: 60px;
        background: rgba(255, 255, 255, 0.3);
      "
    ></div>
    <div
      class="position-absolute rounded-circle"
      style="
        top: 20%;
        left: 10%;
        width: 20px;
        height: 20px;
        background: rgba(255, 255, 255, 0.2);
      "
    ></div>
    <div class="key-container">
      <!-- Increased size from 500 to 600 and adjusted viewBox for zoom effect -->
      <svg
        width="600"
        height="600"
        viewBox="0 -100 500 500"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        style="filter: drop-shadow(0 40px 80px rgba(0, 0, 0, 0.3))"
      >
        <defs>
          <linearGradient
            id="gold-gradient"
            x1="0%"
            y1="0%"
            x2="100%"
            y2="100%"
          >
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
        <circle
          cx="250"
          cy="100"
          r="80"
          stroke="url(#gold-gradient)"
          stroke-width="10"
          fill="none"
        />

        <!-- 1. Golden Car Key Fob (Hanging Left) -->
        <g class="key-sway-1" transform-origin="250 170">
          <!-- Adjusted translate to attach to ring -->
          <g transform="translate(195, 145) rotate(20)">
            <!-- Connector Ring -->
            <circle
              cx="35"
              cy="-10"
              r="10"
              stroke="url(#gold-gradient)"
              stroke-width="4"
              fill="none"
            />

            <!-- Fob Body -->
            <rect
              x="0"
              y="0"
              width="70"
              height="120"
              rx="25"
              fill="url(#gold-dark)"
            />
            <rect
              x="5"
              y="5"
              width="60"
              height="110"
              rx="20"
              fill="url(#gold-gradient)"
            />

            <!-- Buttons -->
            <circle
              cx="35"
              cy="35"
              r="14"
              fill="url(#gold-light)"
              opacity="0.8"
            />
            <path
              d="M29 35 L41 35 M35 29 L35 41"
              stroke="#92400E"
              stroke-width="2"
              stroke-linecap="round"
            />

            <circle
              cx="35"
              cy="70"
              r="14"
              fill="url(#gold-light)"
              opacity="0.8"
            />
            <path
              d="M28 70 L42 70"
              stroke="#92400E"
              stroke-width="2"
              stroke-linecap="round"
            />
          </g>
        </g>

        <!-- 2. Antique Skeleton Key (Hanging Center-Right) -->
        <g class="key-sway-2" transform-origin="250 170">
          <g transform="translate(255, 165) rotate(-10)">
            <!-- Connector Ring -->
            <circle
              cx="20"
              cy="-10"
              r="8"
              stroke="url(#gold-gradient)"
              stroke-width="3"
              fill="none"
            />

            <!-- Bow -->
            <path
              d="M20 0 C 5 0 0 15 0 25 C 0 35 10 40 20 40 C 30 40 40 35 40 25 C 40 15 35 0 20 0 Z M 20 10 C 25 10 28 15 28 25 C 28 32 25 35 20 35 C 15 35 12 32 12 25 C 12 15 15 10 20 10 Z"
              fill="url(#gold-gradient)"
            />

            <!-- Shaft -->
            <rect
              x="16"
              y="40"
              width="8"
              height="100"
              fill="url(#gold-gradient)"
            />

            <!-- Bit -->
            <path
              d="M16 110 L 35 110 L 35 125 L 28 125 L 28 130 L 35 130 L 35 135 L 16 135 Z"
              fill="url(#gold-gradient)"
            />
          </g>
        </g>

        <!-- 3. Standard House Key (Hanging Far Right) -->
        <g class="key-sway-3" transform-origin="250 170">
          <!-- Adjusted translate to attach to ring -->
          <g transform="translate(290, 140) rotate(-25)">
            <!-- Connector Ring -->
            <circle
              cx="20"
              cy="-10"
              r="8"
              stroke="url(#gold-gradient)"
              stroke-width="3"
              fill="none"
            />

            <circle cx="20" cy="20" r="22" fill="url(#gold-gradient)" />
            <circle cx="20" cy="20" r="8" fill="#60a5fa" />
            <rect
              x="12"
              y="35"
              width="16"
              height="70"
              fill="url(#gold-gradient)"
            />
            <path
              d="M28 60 L38 60 L38 70 L28 70 L28 80 L38 80 L38 90 L28 90"
              fill="url(#gold-gradient)"
            />
          </g>
        </g>

        <!-- 4. Small Padlock Key (Hanging Far Left) -->
        <g class="key-sway-4" transform-origin="250 170">
          <!-- Adjusted translate to attach to ring -->
          <g transform="translate(155, 135) rotate(35)">
            <!-- Connector Ring -->
            <circle
              cx="15"
              cy="-8"
              r="6"
              stroke="url(#gold-gradient)"
              stroke-width="3"
              fill="none"
            />

            <circle cx="15" cy="15" r="15" fill="url(#gold-gradient)" />
            <rect
              x="10"
              y="25"
              width="10"
              height="40"
              fill="url(#gold-gradient)"
            />
            <rect
              x="20"
              y="50"
              width="8"
              height="10"
              fill="url(#gold-gradient)"
            />
          </g>
        </g>
      </svg>
    </div>
  </v-col>
</template>

<style scoped>
.key-container {
  animation: float 8s ease-in-out infinite;
}

/* Increased rotation angles for more noticeable animation */
.key-sway-1 {
  animation: sway-1 5s ease-in-out infinite;
}

.key-sway-2 {
  animation: sway-2 6s ease-in-out infinite;
}

.key-sway-3 {
  animation: sway-3 7s ease-in-out infinite;
}

.key-sway-4 {
  animation: sway-4 4s ease-in-out infinite;
}

@keyframes float {
  0%,
  100% {
    transform: translateY(0);
  }

  50% {
    transform: translateY(-30px);
  }
}

@keyframes sway-1 {
  0%,
  100% {
    transform: rotate(0deg);
  }

  50% {
    transform: rotate(8deg);
  }
}

@keyframes sway-2 {
  0%,
  100% {
    transform: rotate(0deg);
  }

  50% {
    transform: rotate(-5deg);
  }
}

@keyframes sway-3 {
  0%,
  100% {
    transform: rotate(0deg);
  }

  50% {
    transform: rotate(10deg);
  }
}

@keyframes sway-4 {
  0%,
  100% {
    transform: rotate(0deg);
  }

  50% {
    transform: rotate(-8deg);
  }
}
</style>
