<template>
  <v-row no-gutters class="fill-height">
    <!-- Left Side: Form -->
    <v-col
      cols="12"
      md="6"
      class="bg-white d-flex flex-column justify-center align-center pa-8 pa-md-16"
    >
      <div class="w-100" style="max-width: 440px">
        <!-- Logo -->
        <Logo />

        <!-- Header -->
        <div class="mb-8">
          <h1 class="text-h4 font-weight-bold mb-2">
            {{ $t('login.welcome_back') }}
          </h1>
          <p class="text-body-1 text-medium-emphasis">
            {{ $t('login.welcome_back_subtitle') }}
          </p>
        </div>

        <!-- Toggle -->
        <!-- <v-sheet class="bg-grey-lighten-4 rounded-lg pa-1 mb-8" elevation="0">
          <v-item-group
            mandatory
            :disabled="controller.loading"
            v-model="controller.loginOption"
            selected-class="bg-primary"
          >
            <v-row no-gutters>
              <v-col cols="6">
                <v-item v-slot="{ isSelected, toggle }">
                  <v-btn
                    block
                    flat
                    class="text-medium-emphasis text-capitalize"
                    height="44"
                    :elevation="isSelected ? 2 : 0"
                    :variant="isSelected ? 'flat' : 'text'"
                    @click="toggle"
                    >Sign In</v-btn
                  >
                </v-item>
              </v-col>
              <v-col cols="6">
                <v-item v-slot="{ isSelected, toggle }">
                  <v-btn
                    block
                    flat
                    :variant="isSelected ? 'flat' : 'text'"
                    class="text-medium-emphasis text-capitalize"
                    :elevation="isSelected ? 2 : 0"
                    height="44"
                    @click="toggle"
                    >Signup</v-btn
                  >
                </v-item>
              </v-col>
            </v-row>
          </v-item-group>
        </v-sheet> -->

        <!-- Form -->
        <v-form
          @submit.prevent="controller.login"
          :disabled="controller.loading"
          ref="meuPingulin"
        >
          <v-text-field
            v-model="controller.email"
            :label="$t('login.email_label')"
            variant="outlined"
            prepend-inner-icon="mdi-email-outline"
            color="primary"
            class="mb-4"
            :rules="controller.emailRules"
            rounded="lg"
            hide-details="auto"
            @keyup.enter="controller.login"
          />

          <v-text-field
            v-model="controller.password"
            :label="$t('login.password_label')"
            type="password"
            variant="outlined"
            prepend-inner-icon="mdi-lock-outline"
            color="primary"
            class="mb-8"
            :rules="controller.passwordRules"
            rounded="lg"
            @keyup.enter="controller.login"
            hide-details="auto"
          />

          <v-btn
            block
            color="#2563eb"
            height="56"
            class="text-capitalize text-body-1 font-weight-bold rounded-lg mb-8"
            flat
            @click="controller.login"
            :loading="controller.loading"
            :disabled="controller.loading"
          >
            {{ $t('login.continue') }}
          </v-btn>
        </v-form>

        <!-- Divider -->
        <!-- <div class="d-flex align-center mb-8">
                    <v-divider></v-divider>
                    <span class="mx-8 text-caption text-medium-emphasis text-uppercase text-no-wrap">Or Continue
                        With</span>
                    <v-divider></v-divider>
                </div> -->

        <!-- Socials -->
        <!-- <div class="d-flex justify-center gap-4 mb-12">
                    <v-btn icon="mdi-google" variant="outlined" color="grey-lighten-1"
                        class="rounded-circle mx-2"></v-btn>
                    <v-btn icon="mdi-apple" variant="flat" color="black" class="rounded-circle mx-2"></v-btn>
                    <v-btn icon="mdi-facebook" variant="flat" color="#1877F2" class="rounded-circle mx-2"></v-btn>
                </div> -->

        <!-- Footer -->
        <p
          class="text-caption text-medium-emphasis px-4"
          style="line-height: 1.6"
        >
          {{ $t('login.footer_text') }}
        </p>
      </div>
    </v-col>
    <Keys />
  </v-row>
</template>

<script setup>
import Keys from '@/module/core/component/keys.vue'
import Logo from '@/module/core/component/logo.vue'
import { loginControllerImpl } from '../di/di'

const controller = loginControllerImpl()
</script>
