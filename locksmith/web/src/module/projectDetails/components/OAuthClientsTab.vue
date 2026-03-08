<template>
  <v-container fluid class="pa-0">
    <!-- Section Header -->
    <div class="d-flex align-center justify-space-between mb-3">
      <div>
        <p class="text-subtitle-2 font-weight-bold section-title mb-0">
          {{ $t('projectDetails.oauthClients') || 'OAuth Clients' }}
        </p>
        <p class="text-caption section-subtitle mt-1">
          {{ controller.clients.length }} {{ $t('projectDetails.configuredClients') }}
        </p>
      </div>
      <v-btn
        color="primary"
        prepend-icon="mdi-plus"
        height="32"
        rounded="lg"
        elevation="0"
        class="text-capitalize font-weight-medium"
        @click="controller.oauthClient.openDialog()"
      >
        {{ $t('projectDetails.addClient') }}
      </v-btn>
    </div>

    <!-- Search -->
    <v-text-field
      v-model="searchQuery"
      density="compact"
      variant="outlined"
      :placeholder="$t('projectDetails.search')"
      prepend-inner-icon="mdi-magnify"
      hide-details
      rounded="lg"
      bg-color="white"
      style="max-width: 360px"
      class="mb-4"
    />

    <!-- Client Cards -->
    <div v-if="filteredClients.length > 0" class="d-flex flex-column" style="gap: 12px">
      <v-card
        v-for="(client, index) in filteredClients"
        :key="index"
        elevation="0"
        rounded="lg"
        class="client-card"
      >
        <!-- Header -->
        <div class="pa-3 d-flex align-center justify-space-between">
          <div class="d-flex align-center" style="gap: 12px">
            <div class="item-avatar">
              <v-icon color="primary" size="18">mdi-key-chain-variant</v-icon>
            </div>
            <div>
              <p class="text-subtitle-2 font-weight-bold section-title mb-0">{{ client.name }}</p>
              <div class="d-flex align-center" style="gap: 2px">
                <span class="text-caption section-subtitle">UUID: {{ client.id }}</span>

                <v-btn
                  icon="mdi-content-copy"
                  variant="text"
                  size="x-small"
                  color="grey"
                  @click="copyToClipboard(client.id, 'UUID')"
                />
              </div>
            </div>
          </div>
          <div class="d-flex" style="gap: 4px">
            <v-btn
              variant="tonal"
              color="primary"
              rounded="lg"
              size="small"
              class="text-capitalize font-weight-medium"
              :to="`/${$i18n.locale}/projects/${controller.route.params.id}/clients/${client.id}`"
            >
              <v-icon start size="16">mdi-eye-outline</v-icon> {{ $t('projectDetails.clientDetails') }}
            </v-btn>
            <v-btn
              icon="mdi-pencil-outline"
              variant="text"
              size="small"
              color="grey-darken-1"
              @click="controller.oauthClient.openEditDialog(client)"
            />
            <v-btn icon="mdi-delete-outline" variant="text" size="small" color="error" />
          </div>
        </div>

        <v-divider />

        <v-card-text class="pa-3">
          <v-row dense>
            <v-col cols="12" md="6">
              <p class="field-label mb-1">{{ $t('projectDetails.clientId') }}</p>
              <div class="field-box d-flex align-center">
                <code class="text-primary font-weight-bold text-body-2 flex-grow-1">{{ client.client_id }}</code>
                <v-btn
                  icon="mdi-content-copy"
                  variant="text"
                  size="x-small"
                  color="primary"
                  @click="copyToClipboard(client.client_id, 'Client ID')"
                />
              </div>
            </v-col>

            <v-col cols="12" md="6">
              <p class="field-label mb-1">{{ $t('projectDetails.clientSecret') }}</p>
              <div class="field-box d-flex align-center">
                <code class="text-body-2 flex-grow-1" style="color: #374151">
                  {{ client.client_secret
                    ? visibleSecrets[client.id] ? client.client_secret : '••••••••••••••••'
                    : 'N/A' }}
                </code>
                <v-btn
                  v-if="client.client_secret"
                  :icon="visibleSecrets[client.id] ? 'mdi-eye-off-outline' : 'mdi-eye-outline'"
                  variant="text"
                  size="x-small"
                  color="grey"
                  @click="visibleSecrets[client.id] = !visibleSecrets[client.id]"
                />
                <v-btn
                  v-if="client.client_secret"
                  icon="mdi-content-copy"
                  variant="text"
                  size="x-small"
                  color="primary"
                  @click="copyToClipboard(client.client_secret, 'Client Secret')"
                />
              </div>
            </v-col>

            <v-col cols="12" md="6" class="mt-2">
              <p class="field-label mb-2">{{ $t('projectDetails.grantTypes') }}</p>
              <div class="d-flex flex-wrap" style="gap: 6px">
                <v-chip
                  v-for="(grant, gIndex) in (client.grant_types || 'standard').split(' ')"
                  :key="gIndex"
                  size="x-small"
                  variant="tonal"
                  color="primary"
                  class="font-weight-bold"
                >
                  {{ grant }}
                </v-chip>
              </div>
            </v-col>

            <v-col cols="12" md="6" class="mt-2">
              <p class="field-label mb-2">{{ $t('projectDetails.redirectUris') }}</p>
              <div class="d-flex flex-wrap" style="gap: 6px">
                <div
                  v-for="(uri, uriIndex) in client.redirect_uris?.split(' ') || []"
                  :key="uriIndex"
                  class="uri-chip d-flex align-center"
                  style="gap: 4px"
                >
                  <v-icon size="10" color="success">mdi-check-circle-outline</v-icon>
                  <span class="text-caption font-weight-medium" style="color: #374151">{{ uri }}</span>
                </div>
                <span v-if="!client.redirect_uris" class="text-caption section-subtitle" style="font-style: italic">
                  {{ $t('projectDetails.noUrisDefined') }}
                </span>
              </div>
            </v-col>
          </v-row>
        </v-card-text>

        <v-divider />

        <div class="pa-2 px-3 d-flex flex-wrap footer-bar" style="gap: 12px">
          <div class="d-flex align-center" style="gap: 6px">
            <v-icon size="13" color="grey">mdi-clock-outline</v-icon>
            <span class="text-caption section-subtitle">
              {{ $t('projectDetails.createdAt') }}: <strong style="color: #374151">{{ client.created_at }}</strong>
            </span>
          </div>
          <div class="d-flex align-center" style="gap: 6px">
            <v-icon size="13" color="grey">mdi-update</v-icon>
            <span class="text-caption section-subtitle">
              {{ $t('projectDetails.updatedAt') }}: <strong style="color: #374151">{{ client.updated_at }}</strong>
            </span>
          </div>
        </div>
      </v-card>
    </div>

    <!-- Empty State -->
    <div v-else class="d-flex flex-column align-center justify-center py-16 text-center">
      <div class="empty-icon-wrap mb-4">
        <v-icon size="28" color="grey-lighten-1">mdi-key-variant</v-icon>
      </div>
      <p class="text-subtitle-2 font-weight-semibold text-grey-darken-2">{{ $t('projectDetails.noClients') }}</p>
      <p class="text-body-2 text-grey mt-1">{{ $t('projectDetails.noClientsDesc') }}</p>
    </div>

    <!-- Dialog -->
    <v-dialog v-model="controller.oauthClient.dialog" max-width="520px" persistent>
      <v-card rounded="xl" elevation="0" border>
        <v-card-item class="pa-6 pb-3">
          <div class="d-flex align-start justify-space-between">
            <div>
              <v-card-title class="text-h6 font-weight-bold pa-0 section-title">
                {{ controller.oauthClient.form.id ? $t('projectDetails.editClient') : $t('projectDetails.addClient') }}
              </v-card-title>
              <v-card-subtitle class="text-body-2 pa-0 mt-1">
                {{ controller.oauthClient.form.id ? $t('projectDetails.editClient') : $t('projectDetails.addClient') }}
              </v-card-subtitle>
            </div>
            <v-btn
              icon="mdi-close"
              variant="text"
              size="small"
              color="grey"
              @click="controller.oauthClient.dialog = false"
            />
          </div>
        </v-card-item>

        <v-divider />

        <v-card-text class="pa-6 pb-2">
          <v-row dense>
            <v-col cols="12">
              <v-text-field
                v-model="controller.oauthClient.form.name"
                :error-messages="controller.oauthClient.errors.name"
                :label="$t('projectDetails.clientName')"
                variant="outlined"
                density="compact"
                rounded="lg"
                hide-details="auto"
                class="mb-3"
              />
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="controller.oauthClient.form.client_id"
                :error-messages="controller.oauthClient.errors.client_id"
                :label="$t('projectDetails.clientId')"
                variant="outlined"
                density="compact"
                rounded="lg"
                hide-details="auto"
                class="mb-3"
                :readonly="!!controller.oauthClient.form.id"
              />
            </v-col>
            <v-col cols="12">
              <v-textarea
                v-model="controller.oauthClient.form.redirect_uris"
                :error-messages="controller.oauthClient.errors.redirect_uris"
                :label="$t('projectDetails.redirectUris')"
                variant="outlined"
                density="compact"
                rounded="lg"
                rows="3"
                hide-details="auto"
                class="mb-3"
              />
            </v-col>
            <v-col cols="12">
              <v-select
                v-model="controller.oauthClient.form.grant_types"
                :items="grantTypes"
                :label="$t('projectDetails.grantTypes')"
                variant="outlined"
                density="compact"
                rounded="lg"
                hide-details="auto"
                multiple
                chips
              />
            </v-col>
          </v-row>
        </v-card-text>

        <v-card-actions class="pa-6 pt-3" style="gap: 8px">
          <v-spacer />
          <v-btn
            color="grey"
            variant="tonal"
            rounded="lg"
            class="text-capitalize"
            @click="controller.oauthClient.cancelDialog()"
          >
            {{ $t('projectDetails.cancel') }}
          </v-btn>
          <v-btn
            color="primary"
            variant="flat"
            rounded="lg"
            class="text-capitalize"
            @click="controller.oauthClient.save()"
          >
            {{ $t('projectDetails.save') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Snackbar -->
    <v-snackbar v-model="snackbar" :timeout="2000" color="success" location="top" rounded="lg" elevation="4">
      <div class="d-flex align-center" style="gap: 8px">
        <v-icon icon="mdi-check-circle" color="white" />
        <span class="font-weight-medium">{{ snackbarText }}</span>
      </div>
    </v-snackbar>
  </v-container>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})

const grantTypes = [
  { title: 'Authorization Code', value: 'authorization_code' },
  { title: 'Implicit', value: 'implicit' },
  { title: 'Password', value: 'password' },
  { title: 'Client Credentials', value: 'client_credentials' },
  { title: 'Refresh Token', value: 'refresh_token' },
]

const searchQuery = ref('')
const snackbar = ref(false)
const snackbarText = ref('')
const visibleSecrets = ref({})

const filteredClients = computed(() => {
  if (!searchQuery.value) return props.controller.clients
  const q = searchQuery.value.toLowerCase()
  return props.controller.clients.filter(
    (c) =>
      c.name?.toLowerCase().includes(q) ||
      c.client_id?.toLowerCase().includes(q) ||
      c.id?.toLowerCase().includes(q),
  )
})

const copyToClipboard = async (text, label) => {
  try {
    await navigator.clipboard.writeText(text)
    snackbarText.value = `${label} copied to clipboard!`
    snackbar.value = true
  } catch (err) {
    console.error('Failed to copy: ', err)
  }
}
</script>

<style scoped>
.section-title {
  color: #111827;
}
.section-subtitle {
  color: #6b7280;
}
.footer-bar {
  background: #f9fafb;
}
.client-card {
  background: white;
  border: 1px solid #eef0f6;
  transition: box-shadow 0.2s ease, border-color 0.2s ease;
}
.client-card:hover {
  box-shadow: 0 4px 16px -6px rgba(0, 0, 0, 0.08) !important;
  border-color: rgba(var(--v-theme-primary), 0.2) !important;
}
.item-avatar {
  width: 36px;
  height: 36px;
  border-radius: 8px;
  background: linear-gradient(
    135deg,
    rgba(var(--v-theme-primary), 0.12),
    rgba(var(--v-theme-primary), 0.06)
  );
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.field-label {
  font-size: 0.65rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  color: #9ca3af;
}
.field-box {
  background: #f9fafb;
  border: 1px solid #f0f0f5;
  border-radius: 8px;
  padding: 10px 12px;
}
.uri-chip {
  background: #f3f4f6;
  border-radius: 20px;
  padding: 3px 10px;
}
.empty-icon-wrap {
  width: 60px;
  height: 60px;
  border-radius: 14px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
}
code {
  font-family: 'Fira Code', monospace;
}
</style>
