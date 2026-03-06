<template>
  <v-container fluid class="pa-4 fill-height align-start page-bg">
    <v-row>
      <!-- Page Header -->
      <v-col cols="12">
        <breadcrumb class="mb-3" />
        <div v-if="controller.account" class="d-flex align-center" style="gap: 14px">
          <div class="account-avatar">
            <v-icon size="22" color="primary">mdi-account-outline</v-icon>
          </div>
          <div>
            <h1 class="text-h5 font-weight-bold page-title mb-0">{{ controller.account.name }}</h1>
            <p class="text-body-2 page-subtitle mb-0">
              @{{ controller.account.username }} · {{ controller.account.email }}
              <v-chip v-if="controller.account.role_name" color="primary" variant="tonal" size="x-small" class="ml-2 font-weight-bold">
                {{ controller.account.role_name }}
              </v-chip>
            </p>
          </div>
        </div>
        <div v-else class="d-flex align-center" style="gap: 14px">
          <v-skeleton-loader type="avatar" width="44" height="44" />
          <v-skeleton-loader type="heading" width="220" />
        </div>
      </v-col>

      <!-- Split Layout -->
      <v-col cols="12">
        <div class="split-layout">

          <!-- Left: Device List -->
          <div class="panel-left">
            <div class="panel-header px-4 py-3">
              <p class="text-subtitle-2 font-weight-bold section-title mb-0">Devices</p>
              <p class="text-caption section-subtitle mb-0">{{ controller.sessions.items.length }} session(s)</p>
            </div>
            <v-divider />

            <div v-if="controller.sessions.loading" class="d-flex justify-center py-8">
              <v-progress-circular indeterminate color="primary" size="24" />
            </div>

            <div v-else-if="controller.sessions.items.length === 0" class="d-flex flex-column align-center justify-center py-10 text-center px-4">
              <v-icon size="32" color="grey-lighten-1" class="mb-2">mdi-devices</v-icon>
              <p class="text-caption text-grey-darken-1">No sessions found</p>
            </div>

            <div v-else class="session-list">
              <div
                v-for="session in controller.sessions.items"
                :key="session.id"
                class="session-row px-4 py-3"
                :class="{ 'session-row--active': controller.selectedSession?.id === session.id, 'session-row--revoked': session.revoked }"
                @click="controller.sessions.select(session)"
              >
                <div class="d-flex align-center" style="gap: 10px">
                  <div class="device-icon-wrap flex-shrink-0" :class="{ 'device-icon-wrap--active': controller.selectedSession?.id === session.id }">
                    <v-icon size="16" :color="controller.selectedSession?.id === session.id ? 'white' : 'primary'">
                      {{ deviceIcon(session.device_type) }}
                    </v-icon>
                  </div>
                  <div style="min-width: 0">
                    <p class="text-body-2 font-weight-bold section-title mb-0 text-truncate">
                      {{ session.browser || 'Unknown Browser' }}
                      <span v-if="session.browser_version" class="font-weight-regular text-caption section-subtitle"> {{ session.browser_version }}</span>
                    </p>
                    <p class="text-caption section-subtitle mb-0 text-truncate">
                      {{ session.os || '—' }}<span v-if="session.os_version"> {{ session.os_version }}</span>
                    </p>
                    <div class="d-flex align-center mt-1" style="gap: 4px">
                      <v-chip :color="session.revoked ? 'error' : 'success'" variant="tonal" size="x-small" class="font-weight-bold">
                        {{ session.revoked ? 'Revoked' : 'Active' }}
                      </v-chip>
                      <span class="text-caption section-subtitle">{{ formatDateShort(session.last_activity) }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div v-if="controller.sessions.filter.totalPages > 1" class="px-2 py-2">
              <v-pagination
                v-model="controller.sessions.filter.page"
                :length="controller.sessions.filter.totalPages"
                density="compact"
                rounded="lg"
                active-color="primary"
                variant="flat"
                size="small"
                @update:model-value="controller.sessions.fetch"
              />
            </div>
          </div>

          <!-- Right: Session Detail + Tokens -->
          <div class="panel-right">
            <!-- No session selected -->
            <div v-if="!controller.selectedSession" class="d-flex flex-column align-center justify-center h-100 text-center py-16">
              <v-icon size="48" color="grey-lighten-1" class="mb-3">mdi-cursor-pointer</v-icon>
              <p class="text-subtitle-2 text-grey-darken-2">Select a device to see details</p>
            </div>

            <template v-else>
              <!-- Session Details Card -->
              <div class="pa-4">
                <div class="d-flex align-center justify-space-between mb-3">
                  <div class="d-flex align-center" style="gap: 10px">
                    <div class="device-icon-wrap">
                      <v-icon size="18" color="primary">{{ deviceIcon(controller.selectedSession.device_type) }}</v-icon>
                    </div>
                    <div>
                      <p class="text-body-1 font-weight-bold section-title mb-0">
                        {{ controller.selectedSession.browser || 'Unknown Browser' }}
                        <span v-if="controller.selectedSession.browser_version" class="font-weight-regular text-caption section-subtitle"> {{ controller.selectedSession.browser_version }}</span>
                      </p>
                      <p class="text-caption section-subtitle mb-0">{{ controller.selectedSession.client_name }}</p>
                    </div>
                    <v-chip
                      :color="controller.selectedSession.revoked ? 'error' : 'success'"
                      variant="tonal"
                      size="x-small"
                      class="font-weight-bold"
                    >
                      {{ controller.selectedSession.revoked ? 'Revoked' : 'Active' }}
                    </v-chip>
                  </div>
                  <v-btn
                    v-if="!controller.selectedSession.revoked"
                    variant="tonal"
                    color="error"
                    size="small"
                    rounded="lg"
                    :loading="controller.selectedSession._revoking"
                    @click="controller.sessions.revoke(controller.selectedSession)"
                  >
                    Revoke Session
                  </v-btn>
                </div>

                <v-card elevation="0" rounded="lg" class="detail-card">
                  <div class="detail-grid pa-4">
                    <div class="detail-item">
                      <span class="detail-label">Device Type</span>
                      <span class="detail-value">{{ controller.selectedSession.device_type || '—' }}</span>
                    </div>
                    <div class="detail-item">
                      <span class="detail-label">Operating System</span>
                      <span class="detail-value">
                        {{ controller.selectedSession.os || '—' }}
                        <span v-if="controller.selectedSession.os_version"> {{ controller.selectedSession.os_version }}</span>
                      </span>
                    </div>
                    <div class="detail-item">
                      <span class="detail-label">IP Address</span>
                      <span class="detail-value">{{ controller.selectedSession.ip_address || '—' }}</span>
                    </div>
                    <div class="detail-item">
                      <span class="detail-label">Location</span>
                      <span class="detail-value">
                        {{ [controller.selectedSession.location_city, controller.selectedSession.location_region, controller.selectedSession.location_country].filter(Boolean).join(', ') || '—' }}
                      </span>
                    </div>
                    <div class="detail-item">
                      <span class="detail-label">Last Activity</span>
                      <span class="detail-value">{{ formatDate(controller.selectedSession.last_activity) }}</span>
                    </div>
                    <div class="detail-item">
                      <span class="detail-label">First Login</span>
                      <span class="detail-value">{{ formatDate(controller.selectedSession.created_at) }}</span>
                    </div>
                    <div class="detail-item">
                      <span class="detail-label">Expires At</span>
                      <span class="detail-value">{{ formatDate(controller.selectedSession.expires_at) }}</span>
                    </div>
                    <div class="detail-item">
                      <span class="detail-label">Client</span>
                      <span class="detail-value">{{ controller.selectedSession.client_name }}</span>
                    </div>
                    <div v-if="controller.selectedSession.revoked_reason" class="detail-item detail-item--full">
                      <span class="detail-label">Revoke Reason</span>
                      <span class="detail-value">{{ controller.selectedSession.revoked_reason }}</span>
                    </div>
                  </div>
                </v-card>
              </div>

              <v-divider />

              <!-- Token Rotations -->
              <div class="pa-4">
                <div class="d-flex align-center justify-space-between mb-3">
                  <div>
                    <p class="text-subtitle-2 font-weight-bold section-title mb-0">Token Rotations</p>
                    <p class="text-caption section-subtitle mb-0">All refresh tokens linked to this session, ordered by date</p>
                  </div>
                  <v-progress-circular v-if="controller.tokens.loading" indeterminate color="primary" size="18" width="2" />
                </div>

                <div v-if="!controller.tokens.loading && controller.tokens.items.length === 0" class="d-flex flex-column align-center justify-center py-8 text-center">
                  <v-icon size="32" color="grey-lighten-1" class="mb-2">mdi-key-chain-variant</v-icon>
                  <p class="text-caption text-grey-darken-1">No token rotations for this session</p>
                </div>

                <div v-else-if="controller.tokens.items.length > 0">
                  <!-- Token Chain Trace -->
                  <div class="token-chain">
                    <template v-for="(token, index) in chainedTokens" :key="token.id">
                      <div
                        class="token-card pa-3"
                        :class="token.revoked ? 'token-card--revoked' : 'token-card--active'"
                        style="cursor: pointer"
                        @click="openTokenDialog(token)"
                      >
                        <!-- Card header -->
                        <div class="d-flex align-center justify-space-between mb-2">
                          <div class="d-flex align-center" style="gap: 6px">
                            <span class="chain-gen-badge">#{{ index + 1 }}</span>
                            <v-chip :color="token.revoked ? 'error' : 'success'" variant="tonal" size="x-small" class="font-weight-bold">
                              {{ token.revoked ? 'Revoked' : 'Active' }}
                            </v-chip>
                            <v-chip color="primary" variant="tonal" size="x-small" class="font-weight-bold">
                              {{ token.rotation_count }}x rotated
                            </v-chip>
                          </div>
                          <div class="d-flex flex-column align-end" style="gap: 1px">
                            <span class="detail-value detail-value--mono" style="font-size: 11px; line-height: 1.3">{{ truncateId(token.id) }}</span>
                            <span class="text-caption section-subtitle">{{ formatDate(token.created_at) }}</span>
                          </div>
                        </div>

                        <v-divider class="mb-2" />

                        <!-- Detail grid -->
                        <div class="token-detail-grid">
                          <div class="detail-item">
                            <span class="detail-label">Client</span>
                            <span class="detail-value">{{ token.client_name || '—' }}</span>
                          </div>
                          <div class="detail-item">
                            <span class="detail-label">Expires At</span>
                            <span class="detail-value">{{ formatDate(token.expires_at) }}</span>
                          </div>
                          <div class="detail-item">
                            <span class="detail-label">Last Used</span>
                            <span class="detail-value">{{ formatLastUsed(token) }}</span>
                          </div>
                          <div v-if="token.parent_token_id" class="detail-item">
                            <span class="detail-label">Parent Token</span>
                            <span class="detail-value detail-value--mono">{{ truncateId(token.parent_token_id) }}</span>
                          </div>
                          <div v-if="token.revoked_at" class="detail-item">
                            <span class="detail-label">Revoked At</span>
                            <span class="detail-value">{{ formatDate(token.revoked_at) }}</span>
                          </div>
                          <div v-if="token.revoked_reason" class="detail-item token-detail-grid--full">
                            <span class="detail-label">Revoke Reason</span>
                            <span class="detail-value">{{ token.revoked_reason }}</span>
                          </div>
                        </div>
                      </div>

                      <!-- Chain connector (between cards, not after last) -->
                      <div v-if="index < chainedTokens.length - 1" class="chain-connector">
                        <div class="chain-connector-line" />
                        <div class="chain-connector-badge">
                          <v-icon size="11">mdi-rotate-right</v-icon>
                          rotated
                        </div>
                        <div class="chain-connector-line" />
                      </div>
                    </template>
                  </div>

                  <div v-if="controller.tokens.filter.totalPages > 1" class="d-flex justify-center mt-3">
                    <v-pagination
                      v-model="controller.tokens.filter.page"
                      :length="controller.tokens.filter.totalPages"
                      density="comfortable"
                      rounded="lg"
                      active-color="primary"
                      variant="flat"
                      @update:model-value="controller.tokens.fetch"
                    />
                  </div>
                </div>
              </div>
            </template>
          </div>

        </div>
      </v-col>
    </v-row>

    <!-- Token Detail Dialog -->
    <v-dialog v-model="tokenDialog" max-width="520" rounded="lg">
      <v-card v-if="selectedToken" rounded="lg" elevation="0">
        <div class="d-flex align-center justify-space-between px-5 pt-5 pb-3">
          <div class="d-flex align-center" style="gap: 10px">
            <div class="token-dialog-icon">
              <v-icon size="18" color="primary">mdi-key-variant</v-icon>
            </div>
            <div>
              <p class="text-subtitle-2 font-weight-bold section-title mb-0">Token Details</p>
              <p class="text-caption section-subtitle mb-0">Rotation #{{ selectedTokenIndex + 1 }}</p>
            </div>
          </div>
          <v-btn icon="mdi-close" variant="text" size="small" @click="tokenDialog = false" />
        </div>

        <v-divider />

        <v-card-text class="px-5 pt-4 pb-5">
          <!-- Access Token Section -->
          <p class="text-caption font-weight-bold section-subtitle mb-2" style="text-transform: uppercase; letter-spacing: 0.06em">Access Token (derived)</p>
          <v-card elevation="0" rounded="lg" class="detail-card mb-4">
            <div class="detail-grid pa-4">
              <div class="detail-item">
                <span class="detail-label">Issued At</span>
                <span class="detail-value">{{ formatDate(selectedToken.created_at) }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Expires At</span>
                <span class="detail-value">{{ formatDate(accessTokenExpiresAt(selectedToken)) }}</span>
              </div>
              <div class="detail-item detail-item--full">
                <span class="detail-label">Status</span>
                <v-chip :color="isAccessTokenActive(selectedToken) ? 'success' : 'warning'" variant="tonal" size="x-small" class="font-weight-bold">
                  {{ isAccessTokenActive(selectedToken) ? 'Active' : 'Expired' }}
                </v-chip>
              </div>
            </div>
          </v-card>

          <!-- Refresh Token Section -->
          <p class="text-caption font-weight-bold section-subtitle mb-2" style="text-transform: uppercase; letter-spacing: 0.06em">Refresh Token</p>
          <v-card elevation="0" rounded="lg" :class="selectedToken.revoked ? 'detail-card--revoked' : 'detail-card'">
            <div class="detail-grid pa-4">
              <div class="detail-item detail-item--full">
                <span class="detail-label">Token ID</span>
                <span class="detail-value detail-value--mono">{{ selectedToken.id }}</span>
              </div>
              <div class="detail-item detail-item--full">
                <span class="detail-label">Session ID</span>
                <span class="detail-value detail-value--mono">{{ selectedToken.session_id }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Client</span>
                <span class="detail-value">{{ selectedToken.client_name || '—' }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Client ID</span>
                <span class="detail-value detail-value--mono">{{ truncateId(selectedToken.client_id) }}</span>
              </div>
              <div v-if="selectedToken.parent_token_id" class="detail-item detail-item--full">
                <span class="detail-label">Parent Token ID</span>
                <span class="detail-value detail-value--mono">{{ selectedToken.parent_token_id }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Created At</span>
                <span class="detail-value">{{ formatDate(selectedToken.created_at) }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Expires At</span>
                <span class="detail-value">{{ formatDate(selectedToken.expires_at) }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Last Used</span>
                <span class="detail-value">{{ formatLastUsed(selectedToken) }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Rotations</span>
                <span class="detail-value">{{ selectedToken.rotation_count }}x</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">Status</span>
                <v-chip :color="selectedToken.revoked ? 'error' : 'success'" variant="tonal" size="x-small" class="font-weight-bold">
                  {{ selectedToken.revoked ? 'Revoked' : 'Active' }}
                </v-chip>
              </div>
              <div v-if="selectedToken.revoked_at" class="detail-item">
                <span class="detail-label">Revoked At</span>
                <span class="detail-value">{{ formatDate(selectedToken.revoked_at) }}</span>
              </div>
              <div v-if="selectedToken.revoked_reason" class="detail-item detail-item--full">
                <span class="detail-label">Revoke Reason</span>
                <span class="detail-value">{{ selectedToken.revoked_reason }}</span>
              </div>
            </div>
          </v-card>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, computed } from 'vue'
import Breadcrumb from '@/module/core/component/breadcrumb.vue'
import { userDetailControllerImpl } from '../di/di'

const controller = userDetailControllerImpl()

const tokenDialog = ref(false)
const selectedToken = ref(null)
const selectedTokenIndex = ref(0)

// Build token chain ordered by parent_token_id linkage (oldest → newest)
const chainedTokens = computed(() => {
  const items = controller.tokens.items
  if (!items.length) return []
  const byId = new Map(items.map(t => [t.id, t]))
  const childOf = {}
  items.forEach(t => { if (t.parent_token_id) childOf[t.parent_token_id] = t })
  const root = items.find(t => !t.parent_token_id || !byId.has(t.parent_token_id))
  if (!root) return items
  const chain = []
  const visited = new Set()
  let current = root
  while (current && !visited.has(current.id)) {
    chain.push(current)
    visited.add(current.id)
    current = childOf[current.id]
  }
  items.forEach(t => { if (!visited.has(t.id)) chain.push(t) })
  return chain
})

function openTokenDialog(token) {
  selectedToken.value = token
  selectedTokenIndex.value = chainedTokens.value.indexOf(token)
  tokenDialog.value = true
}

function accessTokenExpiresAt(token) {
  const d = parseDate(token.created_at)
  if (!d) return null
  return new Date(d.getTime() + 5 * 60 * 1000).toISOString()
}

function isAccessTokenActive(token) {
  const d = parseDate(token.created_at)
  if (!d) return false
  return Date.now() < d.getTime() + 5 * 60 * 1000
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

function parseDate(val) {
  if (!val) return null
  const normalized = val
    .replace(' ', 'T')                 // space → T
    .replace(/(\.\d{3})\d+/, '$1')     // truncate microseconds to milliseconds
    .replace(/\+00:00$/, 'Z')
    .replace(/\+00$/, 'Z')
  const d = new Date(normalized)
  return isNaN(d.getTime()) ? null : d
}

function truncateId(id) {
  if (!id) return '—'
  return id.length > 13 ? id.slice(0, 8) + '…' + id.slice(-4) : id
}

function formatLastUsed(token) {
  if (!token.last_used_at) return token.revoked ? '—' : 'Not used yet'
  return formatDate(token.last_used_at)
}

function formatDate(val) {
  const d = parseDate(val)
  if (!d) return '—'
  return d.toLocaleString(undefined, {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  })
}

function formatDateShort(val) {
  const d = parseDate(val)
  if (!d) return '—'
  return d.toLocaleString(undefined, {
    month: 'short',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  })
}
</script>

<style scoped>
.page-bg {
  background: #f7f8fc;
}
.page-title { color: #111827; }
.page-subtitle { color: #6b7280; }
.section-title { color: #111827; }
.section-subtitle { color: #6b7280; }

.account-avatar {
  width: 44px;
  height: 44px;
  border-radius: 10px;
  background: linear-gradient(135deg, rgba(var(--v-theme-primary), 0.14), rgba(var(--v-theme-primary), 0.07));
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

/* Split layout */
.split-layout {
  display: flex;
  gap: 0;
  background: white;
  border: 1px solid #eef0f6;
  border-radius: 12px;
  overflow: hidden;
  min-height: 600px;
}

.panel-left {
  width: 300px;
  flex-shrink: 0;
  border-right: 1px solid #eef0f6;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.panel-right {
  flex: 1;
  overflow-y: auto;
  min-width: 0;
}

.panel-header {
  background: #f9fafb;
}

.session-list {
  flex: 1;
  overflow-y: auto;
}

.session-row {
  cursor: pointer;
  border-bottom: 1px solid #f3f4f6;
  transition: background 0.15s ease;
}
.session-row:last-child { border-bottom: none; }
.session-row:hover { background: #f3f4f6; }
.session-row--active { background: rgba(var(--v-theme-primary), 0.07); border-left: 3px solid rgb(var(--v-theme-primary)); }
.session-row--active:hover { background: rgba(var(--v-theme-primary), 0.1); }
.session-row--revoked { opacity: 0.65; }

.device-icon-wrap {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  background: linear-gradient(135deg, rgba(var(--v-theme-primary), 0.12), rgba(var(--v-theme-primary), 0.06));
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.device-icon-wrap--active {
  background: rgb(var(--v-theme-primary));
}

.token-dialog-icon {
  width: 36px;
  height: 36px;
  border-radius: 9px;
  background: linear-gradient(135deg, rgba(var(--v-theme-primary), 0.12), rgba(var(--v-theme-primary), 0.06));
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

/* Detail card */
.detail-card {
  background: #f9fafb;
  border: 1px solid #eef0f6;
}
.detail-card--revoked {
  background: #fff5f5;
  border: 1px solid #fee2e2;
}
.detail-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px 24px;
}
.detail-item {
  display: flex;
  flex-direction: column;
  gap: 2px;
}
.detail-item--full { grid-column: 1 / -1; }
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

/* Token Chain Trace */
.token-chain {
  display: flex;
  flex-direction: column;
  align-items: stretch;
}

.token-card {
  background: #f9fafb;
  border: 1px solid #eef0f6;
  border-radius: 10px;
  transition: box-shadow 0.15s ease;
}
.token-card:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}
.token-card--revoked {
  background: #fff5f5;
  border-color: #fee2e2;
}
.token-card--active {
  border-color: #bbf7d0;
  background: #f0fdf4;
}

.chain-gen-badge {
  font-size: 11px;
  font-weight: 700;
  color: #6b7280;
  background: #f3f4f6;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  padding: 1px 6px;
  line-height: 1.6;
}

.chain-connector {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0;
}
.chain-connector-line {
  width: 2px;
  height: 8px;
  background: #d1d5db;
}
.chain-connector-badge {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 10px;
  font-weight: 600;
  color: #6b7280;
  background: #f3f4f6;
  border: 1px solid #e5e7eb;
  border-radius: 20px;
  padding: 2px 8px;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}
.token-detail-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px 20px;
}
.token-detail-grid--full { grid-column: 1 / -1; }
.detail-value--mono {
  font-family: 'Courier New', Courier, monospace;
  font-size: 12px;
  word-break: break-all;
}
</style>
