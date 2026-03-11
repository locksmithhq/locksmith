<template>
  <v-row no-gutters class="fill-height">
    <!-- Left Side: Form -->
    <v-col
      cols="12"
      md="6"
      class="d-flex flex-column justify-center align-center pa-8 pa-md-12"
      style="background: #f7f8fc"
    >
      <div class="w-100" style="max-width: 420px">
        <!-- Logo -->
        <div class="mb-10">
          <Logo />
        </div>

        <!-- Error -->
        <v-alert
          v-if="controller.error"
          type="error"
          variant="tonal"
          rounded="lg"
          density="compact"
          class="mb-5 text-body-2"
          closable
          @click:close="controller.error = null"
        >
          {{ controller.error }}
        </v-alert>

        <!-- Form -->
        <v-form
          @submit.prevent="controller.login"
          :disabled="controller.loading"
          ref="meuPingulin"
          class="mb-5"
        >
          <div class="field-label">{{ $t('login.email_label') }}</div>
          <v-text-field
            v-model="controller.form.email"
            variant="outlined"
            prepend-inner-icon="mdi-email-outline"
            color="primary"
            class="mb-5"
            :error-messages="controller.errors.email"
            rounded="lg"
            hide-details="auto"
            density="comfortable"
            bg-color="white"
            placeholder="email@example.com"
            @keyup.enter="controller.login"
          />

          <div class="field-label">{{ $t('login.password_label') }}</div>
          <v-text-field
            v-model="controller.form.password"
            :type="showPassword ? 'text' : 'password'"
            variant="outlined"
            prepend-inner-icon="mdi-lock-outline"
            :append-inner-icon="showPassword ? 'mdi-eye-off-outline' : 'mdi-eye-outline'"
            @click:append-inner="showPassword = !showPassword"
            color="primary"
            class="mb-7"
            :error-messages="controller.errors.password"
            rounded="lg"
            hide-details="auto"
            density="comfortable"
            bg-color="white"
            placeholder="••••••••"
            @keyup.enter="controller.login"
          />

          <v-btn
            block
            color="primary"
            height="50"
            class="text-none font-weight-bold text-body-2"
            rounded="lg"
            elevation="0"
            @click="controller.login"
            :loading="controller.loading"
            :disabled="controller.loading"
          >
            {{ $t('login.continue') }}
          </v-btn>
        </v-form>

        <!-- Footer -->
        <p class="login-footer text-center">
          {{ $t('login.footer_text') }}
        </p>
      </div>
    </v-col>

    <Keys />
  </v-row>
</template>

<script setup>
import { ref } from 'vue'
import Keys from '@/module/core/component/keys.vue'
import Logo from '@/module/core/component/logo.vue'
import { loginControllerImpl } from '../di/di'

const controller = loginControllerImpl()
const showPassword = ref(false)
</script>

<style scoped>
.fill-height {
  height: 100vh !important;
}

.field-label {
  font-size: 12px;
  font-weight: 600;
  color: #374151;
  margin-bottom: 6px;
  letter-spacing: 0.01em;
}

.login-footer {
  font-size: 11px;
  color: #9ca3af;
  line-height: 1.6;
}
</style>
