<template>
  <v-container fluid class="pa-0">
    <!-- Section Header -->
    <div class="d-flex align-center justify-space-between mb-3">
      <div>
        <p class="text-subtitle-2 font-weight-bold section-title mb-0">
          {{ $t('logs.title') }}
        </p>
        <p class="text-caption section-subtitle mt-1">
          {{ controller.sessions.length }} {{ $t('logs.sessionsLoaded') }}
        </p>
      </div>
      <div class="d-flex" style="gap: 8px">
        <v-btn
          icon="mdi-refresh"
          variant="outlined"
          color="grey-darken-1"
          size="small"
          rounded="lg"
          @click="controller.fetchSessionsByProjectID"
        />
      </div>
    </div>

    <!-- Filters -->
    <v-row dense class="mb-4">
      <v-col cols="12" sm="6" md="4">
        <v-text-field
          v-model="controller.session.filter.search"
          :label="$t('logs.searchUser')"
          variant="outlined"
          density="compact"
          rounded="lg"
          hide-details
          prepend-inner-icon="mdi-magnify"
          bg-color="white"
          @update:modelValue="onSearchChange"
        />
      </v-col>
      <v-col cols="12" sm="6" md="3">
        <v-select
          v-model="statusFilter"
          :label="$t('logs.filterType')"
          :items="statusOptions"
          item-title="label"
          item-value="value"
          variant="outlined"
          density="compact"
          rounded="lg"
          hide-details
          bg-color="white"
          prepend-inner-icon="mdi-filter-outline"
        />
      </v-col>
    </v-row>

    <!-- Table Card -->
    <v-card elevation="0" rounded="lg" class="logs-card">
      <div class="pa-3 px-4 d-flex align-center justify-space-between logs-card-header">
        <div class="d-flex align-center" style="gap: 8px">
          <v-icon color="grey-darken-1" size="18">mdi-table-large</v-icon>
          <span class="text-caption font-weight-bold text-uppercase section-subtitle" style="letter-spacing: 0.06em">
            {{ $t('logs.sessionRegistry') }}
          </span>
        </div>
        <v-chip color="primary" variant="tonal" class="font-weight-bold">
          {{ filteredSessions.length }}
        </v-chip>
      </div>

      <v-divider />

      <v-card-text class="pa-0">
        <v-data-table-virtual
          :headers="headers"
          :items="filteredSessions"
          class="logs-table"
          :no-data-text="$t('logs.noLogs')"
          height="500"
          fixed-header
          item-value="id"
        >
          <!-- Timestamp -->
          <template #item.created_at="{ item }">
            <span class="text-caption" style="color: #374151; white-space: nowrap">
              {{ formatDate(item.created_at) }}
            </span>
          </template>

          <!-- User -->
          <template #item.account_name="{ item }">
            <div class="d-flex align-center" style="gap: 8px">
              <div class="mini-avatar">
                <span class="text-caption font-weight-bold" style="color: var(--v-theme-primary); font-size: 0.6rem">
                  {{ initials(item.account_name) }}
                </span>
              </div>
              <div>
                <p class="text-body-2 font-weight-medium mb-0" style="color: #111827; line-height: 1.2">
                  {{ item.account_name }}
                </p>
                <p class="text-caption mb-0 section-subtitle">{{ item.account_email }}</p>
              </div>
            </div>
          </template>

          <!-- Device/Browser -->
          <template #item.browser="{ item }">
            <div>
              <div class="d-flex align-center" style="gap: 4px">
                <v-icon size="13" color="grey">{{ deviceIcon(item.device_type) }}</v-icon>
                <span class="text-body-2" style="color: #374151">
                  {{ item.browser || '—' }}
                  <span v-if="item.browser_version" class="text-caption section-subtitle">
                    {{ item.browser_version }}
                  </span>
                </span>
              </div>
              <p class="text-caption section-subtitle mb-0 mt-0">
                {{ item.os || '—' }}
                <span v-if="item.os_version">{{ item.os_version }}</span>
              </p>
            </div>
          </template>

          <!-- IP / Location -->
          <template #item.ip_address="{ item }">
            <div>
              <p class="text-body-2 font-weight-medium mb-0" style="color: #374151; font-family: 'Fira Code', monospace; font-size: 0.78rem">
                {{ item.ip_address || '—' }}
              </p>
              <p class="text-caption section-subtitle mb-0">
                <span v-if="item.location_city || item.location_country">
                  {{ [item.location_city, item.location_country].filter(Boolean).join(', ') }}
                </span>
                <span v-else>{{ $t('logs.unknownLocation') }}</span>
              </p>
            </div>
          </template>

          <!-- Client -->
          <template #item.client_name="{ item }">
            <v-chip variant="tonal" color="primary" class="font-weight-bold">
              {{ item.client_name }}
            </v-chip>
          </template>

          <!-- Status -->
          <template #item.revoked="{ item }">
            <v-chip
              :color="item.revoked ? 'error' : 'success'"
              variant="tonal"
              class="font-weight-bold"
            >
              <v-icon start size="10">
                {{ item.revoked ? 'mdi-close-circle-outline' : 'mdi-check-circle-outline' }}
              </v-icon>
              {{ item.revoked ? $t('logs.revoked') : $t('logs.active') }}
            </v-chip>
            <p v-if="item.revoked && item.revoked_reason" class="text-caption section-subtitle mb-0 mt-1">
              {{ item.revoked_reason }}
            </p>
          </template>

          <!-- Last Activity -->
          <template #item.last_activity="{ item }">
            <span class="text-caption" style="color: #374151; white-space: nowrap">
              {{ formatDate(item.last_activity) }}
            </span>
          </template>

          <!-- No Data -->
          <template #no-data>
            <div class="d-flex flex-column align-center justify-center py-16 text-center">
              <div class="empty-icon-wrap mb-4">
                <v-icon size="28" color="grey-lighten-1">mdi-history</v-icon>
              </div>
              <p class="text-subtitle-2 font-weight-semibold text-grey-darken-2">{{ $t('logs.noSessions') }}</p>
              <p class="text-body-2 text-grey mt-1" style="max-width: 380px">
                {{ $t('logs.noSessionsDesc') }}
              </p>
            </div>
          </template>
        </v-data-table-virtual>
      </v-card-text>

    </v-card>

    <!-- Pagination -->
    <div class="d-flex justify-center mt-4" v-if="controller.session.filter.totalPages > 1">
      <v-pagination
        v-model="controller.session.filter.page"
        :length="controller.session.filter.totalPages"
        density="comfortable"
        rounded="lg"
        active-color="primary"
        variant="flat"
        @update:model-value="controller.fetchSessionsByProjectID"
      />
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})

const statusFilter = ref('all')

const statusOptions = computed(() => [
  { label: t('logs.allSessions'), value: 'all' },
  { label: t('logs.active'), value: 'active' },
  { label: t('logs.revoked'), value: 'revoked' },
])

const headers = computed(() => [
  { title: t('logs.columns.timestamp'), key: 'created_at', width: '160px' },
  { title: t('logs.columns.user'), key: 'account_name', minWidth: '180px' },
  { title: t('logs.columns.client'), key: 'client_name', width: '140px' },
  { title: t('logs.columns.deviceBrowser'), key: 'browser', minWidth: '160px' },
  { title: t('logs.columns.ipLocation'), key: 'ip_address', minWidth: '150px' },
  { title: t('logs.columns.lastActivity'), key: 'last_activity', width: '160px' },
  { title: t('logs.columns.status'), key: 'revoked', width: '120px' },
])

const filteredSessions = computed(() => {
  if (statusFilter.value === 'active') {
    return props.controller.sessions.filter((s) => !s.revoked)
  }
  if (statusFilter.value === 'revoked') {
    return props.controller.sessions.filter((s) => s.revoked)
  }
  return props.controller.sessions
})

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

function initials(name) {
  if (!name) return '?'
  return name
    .split(' ')
    .slice(0, 2)
    .map((w) => w[0])
    .join('')
    .toUpperCase()
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

function onSearchChange() {
  props.controller.session.filter.page = 1
  props.controller.fetchSessionsByProjectID()
}
</script>

<style scoped>
.section-title {
  color: #111827;
}
.section-subtitle {
  color: #6b7280;
}
.logs-card {
  background: white;
  border: 1px solid #eef0f6;
}
.logs-card-header {
  background: white;
}
.mini-avatar {
  width: 28px;
  height: 28px;
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
.logs-table :deep(thead th) {
  background-color: white !important;
  font-weight: 700 !important;
  color: #9ca3af !important;
  height: 40px !important;
  border-bottom: 1px solid #f3f4f6 !important;
  text-transform: uppercase !important;
  font-size: 0.65rem !important;
  letter-spacing: 0.06em !important;
}
.logs-table :deep(tbody tr:hover) {
  background-color: #fafafa !important;
}
.logs-table :deep(tbody td) {
  padding-top: 10px !important;
  padding-bottom: 10px !important;
}
</style>
