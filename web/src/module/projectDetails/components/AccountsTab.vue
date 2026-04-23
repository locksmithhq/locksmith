<template>
  <v-container fluid class="pa-0">
    <!-- Section Header -->
    <div class="d-flex align-center justify-space-between mb-3">
      <div>
        <p class="text-subtitle-2 font-weight-bold section-title mb-0">
          {{ $t('projectDetails.users') || 'User Directory' }}
        </p>
        <p class="text-caption section-subtitle mt-1">
          {{ controller.account.filter.totalItems }} items registered
        </p>
      </div>
      <v-btn
        color="primary"
        prepend-icon="mdi-account-plus-outline"
        height="32"
        rounded="lg"
        elevation="0"
        class="text-capitalize font-weight-medium"
        @click="controller.account.openDialog"
      >
        {{ $t('projectDetails.addUser') }}
      </v-btn>
    </div>

    <!-- Search -->
    <v-text-field
      v-model="controller.account.filter.search"
      density="compact"
      variant="outlined"
      :placeholder="$t('projectDetails.search')"
      prepend-inner-icon="mdi-magnify"
      hide-details
      rounded="lg"
      bg-color="white"
      style="max-width: 360px"
      class="mb-4"
      @update:modelValue="controller.searchAccounts"
    />

    <!-- Account Table -->
    <v-data-table-server
      v-model:page="controller.account.filter.page"
      v-model:items-per-page="controller.account.filter.limit"
      :headers="headers"
      :items="controller.accounts"
      :items-length="controller.account.filter.totalItems"
      :loading="controller.account.loading"
      density="comfortable"
      rounded="lg"
      class="account-table-card"
      @update:options="controller.loadAccountOptions"
    >
      <template #item.name="{ item }">
        <div class="d-flex align-center" style="gap: 10px; min-width: 0">
          <div class="item-avatar flex-shrink-0">
            <v-icon color="primary" size="16">mdi-account-outline</v-icon>
          </div>
          <div style="min-width: 0">
            <div class="d-flex align-center" style="gap: 6px">
              <p class="text-body-2 font-weight-bold section-title mb-0 text-truncate">{{ item.name }}</p>
              <v-tooltip v-for="provider in (item.social_providers || [])" :key="provider" :text="providerLabel(provider)" location="top">
                <template v-slot:activator="{ props }">
                  <span v-bind="props" class="provider-badge">
                    <svg v-if="provider === 'google'" width="12" height="12" viewBox="0 0 24 24">
                      <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                      <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                      <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                      <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
                    </svg>
                    <v-icon v-else size="12" color="grey-darken-2">mdi-shield-account-outline</v-icon>
                  </span>
                </template>
              </v-tooltip>
            </div>
            <p class="text-caption section-subtitle mb-0">@{{ item.username }} · {{ item.email }}</p>
          </div>
        </div>
      </template>

      <template #item.role_name="{ item }">
        <v-chip color="primary" variant="tonal" class="font-weight-bold">
          {{ item.role_name || $t('users.noRole') }}
        </v-chip>
      </template>

      <template #item.actions="{ item }">
        <div class="d-flex align-center justify-end" style="gap: 4px">
          <v-btn
            icon="mdi-devices"
            variant="text"
            size="x-small"
            color="secondary"
            :title="$t('users.viewDevices')"
            :to="`/${$route.params.locale}/projects/${$route.params.id}/users/${item.id}`"
          />
          <v-btn
            icon="mdi-pencil-outline"
            variant="text"
            size="x-small"
            color="primary"
            @click="controller.account.openEditDialog(item)"
          />
          <v-btn icon="mdi-delete-outline" variant="text" size="x-small" color="error" />
        </div>
      </template>

      <template #no-data>
        <div class="d-flex flex-column align-center justify-center py-16 text-center">
          <div class="empty-icon-wrap mb-4">
            <v-icon size="28" color="grey-lighten-1">mdi-account-off-outline</v-icon>
          </div>
          <p class="text-subtitle-2 font-weight-semibold text-grey-darken-2">{{ $t('users.noUsersTitle') }}</p>
          <p class="text-body-2 text-grey mt-1">{{ $t('users.noUsersDesc') }}</p>
        </div>
      </template>
    </v-data-table-server>

    <!-- Device Sessions Dialog -->
    <v-dialog v-model="controller.deviceDialog.open" max-width="780px" scrollable>
      <v-card rounded="xl" elevation="0" border>
        <v-card-item class="pa-6 pb-3">
          <div class="d-flex align-start justify-space-between">
            <div>
              <v-card-title class="text-h6 font-weight-bold pa-0 section-title">
                {{ $t('users.devicesTitle') }}
              </v-card-title>
              <v-card-subtitle class="text-body-2 pa-0 mt-1">
                {{ controller.deviceDialog.accountName }}
              </v-card-subtitle>
            </div>
            <v-btn
              icon="mdi-close"
              variant="text"
              size="small"
              color="grey"
              @click="controller.deviceDialog.close"
            />
          </div>
        </v-card-item>

        <v-divider />

        <v-card-text class="pa-4">
          <div v-if="controller.deviceDialog.sessions.length === 0" class="d-flex flex-column align-center justify-center py-10 text-center">
            <v-icon size="32" color="grey-lighten-1" class="mb-3">mdi-devices</v-icon>
            <p class="text-subtitle-2 text-grey-darken-2">{{ $t('users.noDevices') }}</p>
          </div>

          <div v-else>
            <div
              v-for="(session, index) in controller.deviceDialog.sessions"
              :key="session.id"
              :class="{ 'mb-3': index < controller.deviceDialog.sessions.length - 1 }"
            >
              <v-card rounded="lg" elevation="0" class="session-card" :class="{ 'session-card--revoked': session.revoked }">
                <!-- Card Header -->
                <div class="d-flex align-center justify-space-between px-4 pt-3 pb-2">
                  <div class="d-flex align-center" style="gap: 10px">
                    <div class="device-icon-wrap flex-shrink-0">
                      <v-icon size="18" color="primary">{{ deviceIcon(session.device_type) }}</v-icon>
                    </div>
                    <div>
                      <div class="d-flex align-center" style="gap: 6px">
                        <span class="text-body-2 font-weight-bold section-title">
                          {{ session.browser || '—' }}
                          <span v-if="session.browser_version" class="font-weight-regular text-caption section-subtitle"> {{ session.browser_version }}</span>
                        </span>
                        <v-chip
                          :color="session.revoked ? 'error' : 'success'"
                          variant="tonal"
                          size="x-small"
                          class="font-weight-bold"
                        >
                          {{ session.revoked ? $t('logs.revoked') : $t('logs.active') }}
                        </v-chip>
                      </div>
                      <p class="text-caption section-subtitle mb-0">
                        {{ session.client_name }}
                      </p>
                    </div>
                  </div>
                  <v-btn
                    v-if="!session.revoked"
                    variant="tonal"
                    color="error"
                    size="small"
                    rounded="lg"
                    :loading="session._revoking"
                    class="flex-shrink-0"
                    @click="handleRevoke(session)"
                  >
                    {{ $t('users.revokeAccess') }}
                  </v-btn>
                </div>

                <v-divider />

                <!-- Detail Grid -->
                <div class="px-4 py-3 session-detail-grid">
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.deviceType') }}</span>
                    <span class="detail-value">{{ session.device_type || '—' }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.deviceName') }}</span>
                    <span class="detail-value">{{ session.device_name || '—' }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.os') }}</span>
                    <span class="detail-value">{{ session.os ? session.os + (session.os_version ? ' ' + session.os_version : '') : '—' }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.ipAddress') }}</span>
                    <span class="detail-value">{{ session.ip_address || '—' }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.location') }}</span>
                    <span class="detail-value">{{ [session.location_city, session.location_region, session.location_country].filter(Boolean).join(', ') || '—' }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.lastActivity') }}</span>
                    <span class="detail-value">{{ formatDate(session.last_activity) }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.createdAt') }}</span>
                    <span class="detail-value">{{ formatDate(session.created_at) }}</span>
                  </div>
                  <div class="detail-item">
                    <span class="detail-label">{{ $t('users.expiresAt') }}</span>
                    <span class="detail-value">{{ formatDate(session.expires_at) }}</span>
                  </div>
                  <div v-if="session.revoked_reason" class="detail-item detail-item--full">
                    <span class="detail-label">{{ $t('users.revokedReason') }}</span>
                    <span class="detail-value">{{ session.revoked_reason }}</span>
                  </div>
                </div>
              </v-card>
            </div>
          </div>
        </v-card-text>

        <v-card-actions v-if="controller.deviceDialog.filter.totalPages > 1" class="pa-4 pt-0">
          <v-pagination
            v-model="controller.deviceDialog.filter.page"
            :length="controller.deviceDialog.filter.totalPages"
            density="comfortable"
            rounded="lg"
            active-color="primary"
            variant="flat"
            @update:model-value="controller.deviceDialog.fetch"
          />
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Account Dialog -->
    <v-dialog v-model="controller.account.dialog" max-width="500px" persistent scrollable>
      <v-card rounded="xl" elevation="0" border>
        <v-card-item class="pa-6 pb-3">
          <div class="d-flex align-start justify-space-between">
            <div>
              <v-card-title class="text-h6 font-weight-bold pa-0 section-title">
                {{ controller.account.form.id ? $t('users.editUser') : $t('users.addUser') }}
              </v-card-title>
              <v-card-subtitle class="text-body-2 pa-0 mt-1">
                {{ controller.account.form.id ? $t('users.editSubtitle') : $t('users.createSubtitle') }}
              </v-card-subtitle>
            </div>
            <v-btn
              icon="mdi-close"
              variant="text"
              size="small"
              color="grey"
              @click="controller.account.cancelDialog"
            />
          </div>
        </v-card-item>

        <v-divider />

        <v-card-text class="pa-6 pb-2">
          <v-row dense>
            <v-col cols="12">
              <v-text-field
                v-model="controller.account.form.name"
                variant="outlined"
                density="compact"
                :label="$t('users.name')"
                rounded="lg"
                hide-details="auto"
                class="mb-3"
                :error-messages="controller.account.errors.name"
              />
            </v-col>
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="controller.account.form.username"
                variant="outlined"
                density="compact"
                rounded="lg"
                :label="$t('users.username')"
                hide-details="auto"
                class="mb-3"
                :error-messages="controller.account.errors.username"
              />
            </v-col>
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="controller.account.form.email"
                variant="outlined"
                density="compact"
                rounded="lg"
                :label="$t('users.email')"
                hide-details="auto"
                class="mb-3"
                :error-messages="controller.account.errors.email"
              />
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="controller.account.form.password"
                type="password"
                variant="outlined"
                density="compact"
                rounded="lg"
                :label="$t('users.password')"
                hide-details="auto"
                class="mb-3"
                :error-messages="controller.account.errors.password"
              />
            </v-col>
            <v-col cols="12">
              <v-select
                v-model="controller.account.form.role_name"
                :items="controller.permission.permissions"
                item-title="title"
                item-value="title"
                variant="outlined"
                :label="$t('users.roles')"
                density="compact"
                rounded="lg"
                hide-details="auto"
                class="mb-3"
                :error-messages="controller.account.errors.role_name"
              />
              <v-switch
                inset
                color="primary"
                v-model="controller.account.form.must_change_password"
                :label="$t('users.mustChangePassword')"
                hide-details="auto"
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
            @click="controller.account.cancelDialog"
          >
            Cancel
          </v-btn>
          <v-btn
            color="primary"
            variant="flat"
            rounded="lg"
            class="text-capitalize"
            @click="controller.account.save"
          >
            {{ $t('users.saveAccount') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})

const headers = computed(() => [
  { title: t('users.name'), key: 'name', sortable: false },
  { title: t('users.roles'), key: 'role_name', sortable: false, width: '160px' },
  { title: '', key: 'actions', sortable: false, align: 'end', width: '120px' },
])

function providerLabel(provider) {
  const labels = { google: 'Google' }
  return labels[provider] || provider
}

function deviceIcon(deviceType) {
  const icons = {
    mobile: 'mdi-cellphone',
    desktop: 'mdi-monitor',
    tablet: 'mdi-tablet',
    tv: 'mdi-television',
    watch: 'mdi-watch',
  }
  return icons[deviceType] || 'mdi-devices'
}

function formatDate(val) {
  if (!val) return '—'
  const normalized = val
    .replace(' ', 'T')
    .replace(/(\.\d{3})\d+/, '$1')
    .replace(/\+00:00$/, 'Z')
    .replace(/\+00$/, 'Z')
  const d = new Date(normalized)
  if (isNaN(d.getTime())) return val
  return d.toLocaleString(undefined, {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  })
}

async function handleRevoke(session) {
  session._revoking = true
  try {
    await props.controller.deviceDialog.revoke(session.id)
  } finally {
    session._revoking = false
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
.account-table-card {
  border: 1px solid #eef0f6;
}
.provider-badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 18px;
  height: 18px;
  border-radius: 4px;
  background: #f1f3f4;
  border: 1px solid #e0e0e0;
  flex-shrink: 0;
}

.item-avatar {
  width: 30px;
  height: 30px;
  border-radius: 6px;
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
.empty-icon-wrap {
  width: 60px;
  height: 60px;
  border-radius: 14px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
}
.device-icon-wrap {
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
}
.session-card {
  background: #f9fafb;
  border: 1px solid #eef0f6;
}
.session-card--revoked {
  background: #fff5f5;
  border-color: #fee2e2;
}
.session-detail-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px 24px;
}
.detail-item {
  display: flex;
  flex-direction: column;
  gap: 2px;
}
.detail-item--full {
  grid-column: 1 / -1;
}
.detail-label {
  font-size: 11px;
  font-weight: 600;
  color: #9ca3af;
  text-transform: uppercase;
  letter-spacing: 0.04em;
}
.detail-value {
  font-size: 13px;
  color: #111827;
  font-weight: 500;
}
</style>
