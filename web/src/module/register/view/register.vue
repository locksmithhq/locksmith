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
      <div class="language-selector-wrapper">
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
              <v-icon start size="18" color="grey-darken-1"
                >mdi-translate</v-icon
              >
              {{ currentLanguageName }}
              <v-icon end size="18" color="grey-darken-1"
                >mdi-chevron-down</v-icon
              >
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
              <v-list-item-title class="font-weight-medium">{{
                lang.name
              }}</v-list-item-title>
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
        <component is="style" v-if="controller.registerConfig.custom_css">
          {{ controller.registerConfig.custom_css }}
        </component>
        <div
          ref="customHtmlContainer"
          v-html="controller.registerConfig.custom_html"
          class="fill-height pa-0"
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
            class="bg-gradient-register d-flex flex-column justify-center align-center pa-12 w-100 fill-height"
            :style="{
              position: 'relative',
              overflow: 'hidden',
              ...backgroundStyle,
            }"
          >
            <div class="decorative-circle circle-1"></div>
            <div class="decorative-circle circle-2"></div>

            <div
              style="position: relative; z-index: 1"
              class="text-center text-white"
            >
              <img
                v-if="logoUrl"
                :src="logoUrl"
                alt="Logo"
                style="max-height: 80px; max-width: 240px; object-fit: contain"
                class="mb-6"
              />
              <v-icon
                v-else
                size="80"
                color="white"
                class="mb-4"
                style="opacity: 0.9"
              >
                mdi-account-plus-outline
              </v-icon>
              <h1 class="text-h3 font-weight-bold mb-2">Join Us</h1>
              <p class="text-h6 font-weight-regular" style="opacity: 0.85">
                Create your account to continue
              </p>
            </div>
          </div>
        </v-col>

        <!-- Form Side -->
        <v-col
          cols="12"
          :md="controller.registerConfig.layout === 'centered' ? 5 : 6"
          class="bg-white d-flex flex-column justify-center align-center pa-8 pa-md-16 fill-height"
          :order="controller.registerConfig.layout === 'split-left' ? 1 : 0"
        >
          <div class="w-100" style="max-width: 440px">
            <!-- Centered Logo -->
            <div
              class="text-center mb-8"
              v-if="controller.registerConfig.layout === 'centered'"
            >
              <img
                v-if="logoUrl"
                :src="logoUrl"
                alt="Logo"
                style="max-height: 50px; max-width: 200px; object-fit: contain"
              />
              <v-icon v-else size="48" :color="primaryColor">
                mdi-account-plus-outline
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
                type="password"
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
                class="text-none font-weight-bold mb-6"
                @click="controller.register"
                :loading="controller.loading"
              >
                {{ $t('auth.signUp') }}
              </v-btn>
            </v-form>

            <!-- Social Registration -->
            <template v-if="controller.registerConfig.show_social">
              <div class="d-flex align-center my-8">
                <v-divider></v-divider>
                <span
                  class="text-caption text-grey-darken-1 px-4 text-uppercase font-weight-bold"
                  style="letter-spacing: 1px"
                >
                  {{ $t('clientDetails.orSignUpWith') }}
                </span>
                <v-divider></v-divider>
              </div>

              <v-row dense>
                <v-col cols="6">
                  <v-btn
                    variant="outlined"
                    block
                    rounded="lg"
                    class="text-none"
                    size="large"
                    border
                  >
                    <v-icon start size="20" color="red">mdi-google</v-icon>
                    <span class="font-weight-bold">Google</span>
                  </v-btn>
                </v-col>
                <v-col cols="6">
                  <v-btn
                    variant="outlined"
                    block
                    rounded="lg"
                    class="text-none"
                    size="large"
                    border
                  >
                    <v-icon start size="20" color="black">mdi-github</v-icon>
                    <span class="font-weight-bold">GitHub</span>
                  </v-btn>
                </v-col>
              </v-row>
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
import { computed, watch, useTemplateRef } from 'vue'
import { useI18n } from 'vue-i18n'
import { registerControllerImpl } from '../di/di'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const { locale } = useI18n()
const controller = registerControllerImpl()
const customHtmlContainer = useTemplateRef('customHtmlContainer')

// Logic to handle events in custom HTML
watch(
  () => controller.registerConfig.custom_html,
  () => {
    if (controller.registerConfig.use_custom_html) {
      setTimeout(() => {
        if (customHtmlContainer.value) {
          const registerBtn =
            customHtmlContainer.value.querySelector('button[type="submit"]') ||
            customHtmlContainer.value.querySelector('#register-button') ||
            customHtmlContainer.value.querySelector('button')

          if (registerBtn) {
            registerBtn.addEventListener('click', (e) => {
              e.preventDefault()
              controller.register()
            })
          }

          const form = customHtmlContainer.value.querySelector('form')
          if (form) {
            form.addEventListener('submit', (e) => {
              e.preventDefault()
              controller.register()
            })
          }
        }
      }, 100)
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
  const config = controller.registerConfig || {}
  
  const { background_type, background_color, background_image } = config

  if (background_type === 'image' && background_image) {
    return {
      backgroundImage: `url(${background_image})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
    }
  }

  if (background_color) {
    return {
      background: background_color,
    }
  }

  const loginConfig = controller.client?.login || {}
  if (loginConfig.background_type === 'image' && loginConfig.background_image) {
    return {
      backgroundImage: `url(${loginConfig.background_image})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
    }
  }

  if (loginConfig.background_color) {
    return {
      background: loginConfig.background_color,
    }
  }

  return {}
})

const primaryColor = computed(() => {
  return controller.registerConfig?.primary_color || controller.client?.login?.primary_color || 'primary'
})

const primaryColorHex = computed(() => {
  const color = primaryColor.value
  return color === 'primary' ? 'rgb(var(--v-theme-primary))' : color
})

const logoUrl = computed(() => {
  return controller.registerConfig?.logo_url || controller.client?.login?.logo_url || ''
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

.bg-gradient-register {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
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

.opacity-70 {
  opacity: 0.7;
}

@media (max-width: 960px) {
  .language-selector-wrapper {
    top: 1rem;
    right: 1rem;
  }
}
</style>
