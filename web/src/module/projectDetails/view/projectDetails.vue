<template>
  <v-container fluid class="pa-4 fill-height align-start page-bg">
    <v-row>
      <v-col cols="12">
        <breadcrumb class="mb-3" />
        <h1 class="text-h5 font-weight-bold page-title">
          {{ $t('projectDetails.title') }}
        </h1>
        <p class="text-body-2 page-subtitle mt-1">
          {{ $t('projectDetails.description') }}
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
              {{ $t('projectDetails.tabs.config') }}
            </v-tab>
            <v-tab value="roles" class="text-capitalize font-weight-medium">
              <v-icon start size="18">mdi-shield-key-outline</v-icon>
              {{ $t('projectDetails.tabs.roles') }}
            </v-tab>
            <v-tab value="oauth" class="text-capitalize font-weight-medium">
              <v-icon start size="18">mdi-key-variant</v-icon>
              {{ $t('projectDetails.tabs.oauthClients') }}
            </v-tab>
            <v-tab value="accounts" class="text-capitalize font-weight-medium">
              <v-icon start size="18">mdi-account-group-outline</v-icon>
              {{ $t('projectDetails.tabs.users') }}
            </v-tab>
            <v-tab value="logs" class="text-capitalize font-weight-medium">
              <v-icon start size="18">mdi-history</v-icon>
              {{ $t('projectDetails.tabs.logs') }}
            </v-tab>
          </v-tabs>

          <v-divider />

          <v-card-text class="pa-4">
            <v-window v-model="controller.activeTab">
              <v-window-item value="config">
                <config-tab :controller="controller" />
              </v-window-item>
              <v-window-item value="roles">
                <roles-tab :controller="controller" />
              </v-window-item>
              <v-window-item value="oauth">
                <o-auth-clients-tab :controller="controller" />
              </v-window-item>
              <v-window-item value="accounts">
                <accounts-tab :controller="controller" />
              </v-window-item>
              <v-window-item value="logs">
                <logs-tab :controller="controller" />
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
import { projectDetailControllerImpl } from '../di/di'
import ConfigTab from '../components/ConfigTab.vue'
import OAuthClientsTab from '../components/OAuthClientsTab.vue'
import AccountsTab from '../components/AccountsTab.vue'
import LogsTab from '../components/LogsTab.vue'
import RolesTab from '../components/RolesTab.vue'

const controller = projectDetailControllerImpl()
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
