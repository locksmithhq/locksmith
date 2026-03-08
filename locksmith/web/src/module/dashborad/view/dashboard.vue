<template>
  <v-container fluid class="pa-4 fill-height align-start page-bg">
    <v-row>
      <!-- Header -->
      <v-col cols="12">
        <breadcrumb class="mb-3" />
        <h1 class="text-h5 font-weight-bold page-title">
          {{ $t('dashboard.title') }}
        </h1>
        <p class="text-body-2 page-subtitle mt-1">
          {{ $t('dashboard.description') }}
        </p>
      </v-col>

      <!-- KPI Cards -->
      <v-col cols="12" sm="6" lg="3">
        <v-card :to="{ name: 'projects' }" elevation="0" rounded="lg" class="stat-card">
          <v-card-text class="pa-5">
            <div class="d-flex align-center justify-space-between mb-4">
              <div class="stat-icon-wrap" style="background: rgba(var(--v-theme-primary), 0.1)">
                <v-icon color="primary" size="20">mdi-folder-outline</v-icon>
              </div>
            </div>
            <div class="text-h4 font-weight-bold stat-value mb-1">
              {{ stats.total_projects ?? '-' }}
            </div>
            <div class="text-body-2 stat-label">{{ $t('dashboard.totalProjects') }}</div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" sm="6" lg="3">
        <v-card elevation="0" rounded="lg" class="stat-card">
          <v-card-text class="pa-5">
            <div class="d-flex align-center justify-space-between mb-4">
              <div class="stat-icon-wrap" style="background: rgba(var(--v-theme-success), 0.1)">
                <v-icon color="success" size="20">mdi-account-group-outline</v-icon>
              </div>
            </div>
            <div class="text-h4 font-weight-bold stat-value mb-1">
              {{ stats.total_users ?? '-' }}
            </div>
            <div class="text-body-2 stat-label">{{ $t('dashboard.totalUsers') }}</div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" sm="6" lg="3">
        <v-card elevation="0" rounded="lg" class="stat-card">
          <v-card-text class="pa-5">
            <div class="d-flex align-center justify-space-between mb-4">
              <div class="stat-icon-wrap" style="background: rgba(103, 80, 164, 0.1)">
                <v-icon color="deep-purple" size="20">mdi-shield-check-outline</v-icon>
              </div>
            </div>
            <div class="text-h4 font-weight-bold stat-value mb-1">
              {{ stats.active_sessions ?? '-' }}
            </div>
            <div class="text-body-2 stat-label">{{ $t('dashboard.activeSessions') }}</div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" sm="6" lg="3">
        <v-card elevation="0" rounded="lg" class="stat-card">
          <v-card-text class="pa-5">
            <div class="d-flex align-center justify-space-between mb-4">
              <div class="stat-icon-wrap" style="background: rgba(var(--v-theme-warning), 0.1)">
                <v-icon color="warning" size="20">mdi-application-outline</v-icon>
              </div>
            </div>
            <div class="text-h4 font-weight-bold stat-value mb-1">
              {{ stats.total_clients ?? '-' }}
            </div>
            <div class="text-body-2 stat-label">{{ $t('dashboard.totalClients') }}</div>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Charts -->
      <v-col cols="12" md="8">
        <v-card elevation="0" rounded="lg" class="stat-card pa-4">
          <LineChart
            :title="$t('dashboard.loginsPerDay')"
            :xAxisData="lineChartXAxis"
            :data="lineChartData"
          />
        </v-card>
      </v-col>

      <v-col cols="12" md="4">
        <v-card elevation="0" rounded="lg" class="stat-card pa-4">
          <PieChart
            :title="$t('dashboard.sessionsByDevice')"
            :data="pieChartData"
          />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import breadcrumb from '@/module/core/component/breadcrumb.vue'
import LineChart from '../component/LineChart.vue'
import PieChart from '../component/PieChart.vue'
import { axiosInstance } from '@/plugins/axios'

const stats = ref({})

const lineChartXAxis = computed(() =>
  (stats.value.logins_per_day ?? []).map((d) => d.day),
)
const lineChartData = computed(() =>
  (stats.value.logins_per_day ?? []).map((d) => d.count),
)
const pieChartData = computed(() =>
  (stats.value.sessions_by_device ?? []).map((d) => ({
    name: d.device,
    value: d.count,
  })),
)

onMounted(async () => {
  try {
    const { data } = await axiosInstance.get('/dashboard/stats')
    stats.value = data
  } catch (err) {
    console.error('Failed to load dashboard stats:', err)
  }
})
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

.stat-card {
  background: #ffffff;
  border: 1px solid #eef0f6;
}

.stat-icon-wrap {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-value {
  color: #111827;
}

.stat-label {
  color: #6b7280;
}
</style>
