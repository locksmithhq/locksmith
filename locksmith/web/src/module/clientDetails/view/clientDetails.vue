<template>
  <v-container fluid class="pa-4 fill-height align-start page-bg">
    <v-row>
      <v-col cols="12">
        <breadcrumb class="mb-3" />
        <h1 class="text-h5 font-weight-bold page-title">
          {{ controller.client.name || $t('clientDetails.title') }}
        </h1>
        <p class="text-body-2 page-subtitle mt-1">
          {{ $t('clientDetails.description') }}
        </p>
      </v-col>

      <v-col cols="12">
        <v-card elevation="0" rounded="lg" class="details-card">
          <v-tabs
            v-model="controller.activeTab"
            color="primary"
            align-tabs="start"
          >
            <v-tab value="config" class="text-capitalize font-weight-medium">
              <v-icon start size="18">mdi-cog-outline</v-icon>
              {{ $t('clientDetails.tabs.config') }}
            </v-tab>
            <v-tab value="login" class="text-capitalize font-weight-medium">
              <v-icon start size="18">mdi-login-variant</v-icon>
              {{ $t('clientDetails.tabs.login') }}
            </v-tab>
            <v-tab value="register" class="text-capitalize font-weight-medium">
              <v-icon start size="18">mdi-account-plus-outline</v-icon>
              {{ $t('clientDetails.tabs.register') }}
            </v-tab>
          </v-tabs>

          <v-divider />

          <v-card-text class="pa-4">
            <v-window v-model="controller.activeTab">
              <v-window-item value="config">
                <client-config-tab :controller="controller" />
              </v-window-item>
              <v-window-item value="login">
                <login-tab :controller="controller" />
              </v-window-item>
              <v-window-item value="register">
                <register-tab :controller="controller" />
              </v-window-item>
            </v-window>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import Breadcrumb from '@/module/core/component/breadcrumb.vue'
import { clientDetailControllerImpl } from '../di/di'
import ClientConfigTab from '../components/ClientConfigTab.vue'
import LoginTab from '../components/LoginTab.vue'
import RegisterTab from '../components/RegisterTab.vue'

const controller = clientDetailControllerImpl()
</script>

<style scoped>
.page-bg {
  background: #f7f8fc;
}
.page-title {
  color: #111827;
}
.page-subtitle {
  color: #6b7280;
}
.details-card {
  background: white;
  border: 1px solid #eef0f6;
}
:deep(.v-window-item) {
  min-height: 400px;
}
</style>
