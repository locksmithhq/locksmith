<template>
  <v-container fluid class="pa-0 fill-height">
    <!-- Loading State -->
    <v-row
      v-if="!controller.client?.client_id"
      class="fill-height align-center justify-center bg-grey-lighten-4"
    >
      <v-progress-circular indeterminate color="primary"></v-progress-circular>
    </v-row>

    <!-- 404 State -->
    <v-row
      v-else-if="!controller.client?.signup || !controller.client?.signup?.enabled"
      class="fill-height align-center justify-center bg-grey-lighten-4"
    >
      <v-col cols="12" sm="6" md="4" class="text-center">
        <div class="text-h1 font-weight-bold text-grey-lighten-1 mb-4">404</div>
        <div class="text-h5 font-weight-bold mb-2">{{ $t('notFound.title') }}</div>
        <div class="text-body-2 text-grey-darken-1 mb-6">{{ $t('notFound.description') }}</div>
      </v-col>
    </v-row>

    <v-row
      v-else
      no-gutters
      class="fill-height"
      :class="{
        'justify-center align-center':
          controller.registerConfig.layout === 'centered',
      }"
      :style="
        controller.registerConfig.layout === 'centered' ? backgroundStyle : {}
      "
    >
      <!-- Language Selector (Global Overlay) -->
      <div
        class="language-selector-wrapper"
        v-if="!controller.registerConfig.use_custom_html"
      >
        <v-menu transition="scale-transition" location="bottom end">
          <template v-slot:activator="{ props }">
            <v-btn
              variant="flat"
              v-bind="props"
              class="text-none font-weight-bold"
              rounded="lg"
              color="white"
              elevation="1"
              border
            >
              <v-icon start size="18" color="grey-darken-1">mdi-translate</v-icon>
              {{ currentLanguageName }}
              <v-icon end size="18" color="grey-darken-1">mdi-chevron-down</v-icon>
            </v-btn>
          </template>

          <v-list density="compact" rounded="lg" class="mt-2" border>
            <v-list-item
              v-for="lang in availableLanguages"
              :key="lang.code"
              :value="lang.code"
              @click="changeLanguage(lang.code)"
              :active="currentLocale === lang.code"
              color="primary"
            >
              <template v-slot:prepend>
                <span class="mr-3">{{ lang.flag }}</span>
              </template>
              <v-list-item-title class="font-weight-medium">{{ lang.name }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>

      <!-- Custom HTML Content -->
      <v-col
        v-if="controller.registerConfig.use_custom_html"
        cols="12"
        class="pa-0 fill-height"
      >
        <div
          ref="customHtmlContainer"
          class="fill-height pa-0 d-flex flex-column justify-center align-center"
        ></div>
      </v-col>

      <!-- Default Layout -->
      <template v-else>
        <!-- Branding Side -->
        <v-col
          v-if="controller.registerConfig.layout !== 'centered'"
          cols="12"
          md="6"
          class="d-none d-md-flex pa-0 fill-height"
          :order="controller.registerConfig.layout === 'split-left' ? 0 : 1"
        >
          <div
            class="bg-gradient-auth d-flex flex-column justify-center align-center pa-12 w-100 fill-height"
            :style="{
              position: 'relative',
              overflow: 'hidden',
              ...backgroundStyle,
            }"
          >
            <div class="decorative-circle circle-1"></div>
            <div class="decorative-circle circle-2"></div>

            <div
              style="position: relative; z-index: 1; width: 100%"
              class="text-center text-white d-flex flex-column align-center"
            >
              <v-img
                v-if="logoUrl"
                :src="logoUrl"
                contain
                max-height="160"
                max-width="280"
                style="width: 100%"
              />
              <v-icon
                v-else
                size="100"
                color="white"
                style="opacity: 0.9"
              >
                mdi-shield-lock-outline
              </v-icon>
            </div>
          </div>
        </v-col>

        <!-- Form Side -->
        <v-col
          cols="12"
          :md="controller.registerConfig.layout === 'centered' ? 5 : 6"
          class="d-flex flex-column justify-center align-center pa-8 pa-md-16 fill-height"
          :class="{
            'bg-white': controller.registerConfig.layout !== 'centered',
          }"
          :order="controller.registerConfig.layout === 'split-left' ? 1 : 0"
        >
          <div class="w-100" style="max-width: 440px">
            <!-- Centered Logo -->
            <div
              class="text-center mb-8"
              v-if="controller.registerConfig.layout === 'centered'"
            >
              <v-img
                v-if="logoUrl"
                :src="logoUrl"
                contain
                max-height="100"
                max-width="200"
                style="width: 100%; margin: 0 auto"
              />
              <v-icon v-else size="48" :color="primaryColor">
                mdi-shield-lock-outline
              </v-icon>
            </div>

            <!-- Error State -->
            <v-alert
              v-if="controller.error"
              type="error"
              variant="tonal"
              class="mb-6 rounded-lg text-caption"
              closable
              @click:close="controller.error = null"
            >
              {{ controller.error }}
            </v-alert>

            <v-form
              @submit.prevent="controller.register"
              :disabled="controller.loading"
            >
              <v-text-field
                v-model="controller.name"
                :label="$t('clientDetails.fullName')"
                :variant="controller.registerConfig.input_variant || 'outlined'"
                density="comfortable"
                prepend-inner-icon="mdi-account-outline"
                hide-details="auto"
                class="mb-4"
                :rules="controller.nameRules"
                placeholder="John Doe"
                :color="primaryColor"
              ></v-text-field>

              <v-text-field
                v-model="controller.email"
                :label="$t('auth.email')"
                :variant="controller.registerConfig.input_variant || 'outlined'"
                density="comfortable"
                prepend-inner-icon="mdi-email-outline"
                hide-details="auto"
                class="mb-4"
                :rules="controller.emailRules"
                placeholder="email@example.com"
                :color="primaryColor"
              ></v-text-field>

              <v-text-field
                v-model="controller.password"
                :label="$t('auth.password')"
                :variant="controller.registerConfig.input_variant || 'outlined'"
                density="comfortable"
                prepend-inner-icon="mdi-lock-outline"
                :append-inner-icon="showPassword ? 'mdi-eye-off-outline' : 'mdi-eye-outline'"
                @click:append-inner="showPassword = !showPassword"
                :type="showPassword ? 'text' : 'password'"
                hide-details="auto"
                class="mb-8"
                :rules="controller.passwordRules"
                placeholder="••••••••"
                :color="primaryColor"
              ></v-text-field>

              <v-btn
                :color="primaryColor"
                size="x-large"
                block
                elevation="0"
                rounded="lg"
                class="text-none font-weight-bold mb-6"
                @click="controller.register"
                :loading="controller.loading"
              >
                {{ $t('auth.signUp') }}
              </v-btn>
            </v-form>

            <!-- Social Registration -->
            <template v-if="controller.registerConfig.show_social && controller.client?.social_providers?.length">
              <div class="d-flex align-center my-6">
                <v-divider></v-divider>
                <span
                  class="text-caption text-grey px-3 text-uppercase font-weight-medium"
                  style="letter-spacing: 1px; white-space: nowrap"
                >
                  {{ $t('clientDetails.orSignUpWith') }}
                </span>
                <v-divider></v-divider>
              </div>

              <v-btn
                v-if="controller.client?.social_providers?.includes('google')"
                variant="outlined"
                block
                rounded="lg"
                class="text-none social-btn"
                size="large"
                :loading="socialLoading === 'google'"
                @click="startSocialLogin('google')"
              >
                <template v-slot:prepend>
                  <svg width="18" height="18" viewBox="0 0 24 24" style="flex-shrink: 0; display: block">
                    <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                    <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                    <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                    <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
                  </svg>
                </template>
                <span class="font-weight-medium" style="color: #3c4043">Continue with Google</span>
              </v-btn>
            </template>

            <!-- Footer -->
            <div class="text-center mt-12">
              <span class="text-body-2 text-grey-darken-1">
                {{ $t('clientDetails.alreadyHaveAccount') }}
              </span>
              <router-link
                :to="{ name: 'auth', query: $route.query }"
                class="text-decoration-none text-body-2 font-weight-bold ml-1"
                :style="{ color: primaryColorHex }"
              >
                {{ $t('auth.signIn') }}
              </router-link>
            </div>
          </div>
        </v-col>
      </template>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, computed, watch, useTemplateRef } from 'vue'
import { useI18n } from 'vue-i18n'
import { registerControllerImpl } from '../di/di'
import { useRouter, useRoute } from 'vue-router'
import { axiosInstance } from '@/plugins/axios'

const router = useRouter()
const route = useRoute()
const { locale } = useI18n()
const controller = registerControllerImpl()
const customHtmlContainer = useTemplateRef('customHtmlContainer')
const showPassword = ref(false)
const socialLoading = ref(null)

const startSocialLogin = async (provider) => {
  socialLoading.value = provider
  try {
    const res = await axiosInstance.get(`/oauth2/social/${provider}/begin`, {
      params: {
        client_id: route.query.client_id,
        redirect_uri: route.query.redirect_uri,
        state: route.query.state || '',
        code_challenge: route.query.code_challenge || '',
        code_challenge_method: route.query.code_challenge_method || '',
      },
    })
    window.location.href = res.data.auth_url
  } catch (err) {
    controller.error = err?.response?.data || `Failed to start ${provider} login`
    socialLoading.value = null
  }
}

// Shadow DOM isolation for custom HTML — same approach as auth.vue
watch(
  () => [
    controller.registerConfig?.custom_html,
    controller.registerConfig?.custom_css,
    controller.registerConfig?.use_custom_html,
  ],
  () => {
    if (controller.registerConfig?.use_custom_html) {
      setTimeout(() => {
        if (customHtmlContainer.value) {
          if (!customHtmlContainer.value.shadowRoot) {
            customHtmlContainer.value.attachShadow({ mode: 'open' })
          }

          const shadow = customHtmlContainer.value.shadowRoot
          const html = controller.registerConfig?.custom_html || ''
          const css = controller.registerConfig?.custom_css || ''

          shadow.innerHTML = `
            <style>
              :host { display: block; height: 100%; width: 100%; overflow: auto; }
              ${css}
            </style>
            ${html}
          `

          const syncInputs = () => {
            const nameInput =
              shadow.querySelector('#name') ||
              shadow.querySelector('input[name="name"]')
            const emailInput =
              shadow.querySelector('#email') ||
              shadow.querySelector('input[name="email"]') ||
              shadow.querySelector('input[type="email"]')
            const passwordInput =
              shadow.querySelector('#password') ||
              shadow.querySelector('input[name="password"]') ||
              shadow.querySelector('input[type="password"]')

            if (nameInput) controller.name = nameInput.value
            if (emailInput) controller.email = emailInput.value
            if (passwordInput) controller.password = passwordInput.value
          }

          const inputs = shadow.querySelectorAll('input')
          inputs.forEach((input) => {
            input.addEventListener('input', syncInputs)
            input.addEventListener('change', syncInputs)
          })

          const registerBtn =
            shadow.querySelector('button[type="submit"]') ||
            shadow.querySelector('#register-button') ||
            shadow.querySelector('button')

          if (registerBtn) {
            registerBtn.addEventListener('click', (e) => {
              if (registerBtn.type !== 'submit') {
                e.preventDefault()
                syncInputs()
                controller.register()
              }
            })
          }

          const form = shadow.querySelector('form')
          if (form) {
            form.addEventListener('submit', (e) => {
              e.preventDefault()
              syncInputs()
              controller.register()
            })
          }
        }
      }, 50)
    }
  },
  { immediate: true },
)

const availableLanguages = [
  { name: 'English', code: 'en', flag: '🇺🇸' },
  { name: 'Português', code: 'pt-br', flag: '🇧🇷' },
]

const currentLocale = computed(() => locale.value)
const currentLanguageName = computed(() => {
  const lang = availableLanguages.find((l) => l.code === locale.value)
  return lang ? lang.name : 'Language'
})

const changeLanguage = (code) => {
  router.push({
    name: route.name,
    params: { ...route.params, locale: code },
    query: route.query,
    replace: true,
  })
  locale.value = code
}

const backgroundStyle = computed(() => {
  const { background_type, background_color, background_image } = controller.registerConfig || {}

  if (background_type === 'image' && background_image) {
    return {
      backgroundImage: `url(${background_image})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
    }
  }

  if (background_color) {
    return { background: background_color }
  }

  return {}
})

const primaryColor = computed(() => {
  return controller.registerConfig?.primary_color || 'primary'
})

const primaryColorHex = computed(() => {
  const color = primaryColor.value
  return color === 'primary' ? 'rgb(var(--v-theme-primary))' : color
})

const logoUrl = computed(() => {
  return controller.registerConfig?.logo_url || ''
})
</script>

<style scoped>
.fill-height {
  height: 100vh !important;
}

.language-selector-wrapper {
  position: fixed;
  top: 1.5rem;
  right: 1.5rem;
  z-index: 100;
}

.bg-gradient-auth {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.decorative-circle {
  position: absolute;
  border-radius: 50%;
}

.circle-1 {
  top: -50px;
  right: -50px;
  width: 340px;
  height: 340px;
  background: rgba(255, 255, 255, 0.12);
}

.circle-2 {
  bottom: -100px;
  left: -100px;
  width: 440px;
  height: 440px;
  background: rgba(255, 255, 255, 0.08);
}

:deep(.v-btn.text-none) {
  letter-spacing: normal;
}

.social-btn {
  background: #fff !important;
  border-color: #dadce0 !important;
  color: #3c4043 !important;
  box-shadow: 0 1px 2px rgba(0,0,0,0.08) !important;
  transition: box-shadow 0.2s, background 0.2s !important;
}

.social-btn:hover {
  background: #f8f9fa !important;
  box-shadow: 0 2px 6px rgba(0,0,0,0.12) !important;
}

@media (max-width: 960px) {
  .language-selector-wrapper {
    top: 1rem;
    right: 1rem;
  }
}
</style>
