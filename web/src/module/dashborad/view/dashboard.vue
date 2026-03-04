<template>
  <v-container fluid class="pa-6 bg-grey-lighten-4 fill-height align-start">
    <v-row>
      <v-col cols="12">
        <div class="d-flex align-center justify-space-between mb-6">
          <div>
            <h1 class="text-h4 font-weight-bold text-grey-darken-3">
              {{ $t('dashboard.title') }}
            </h1>
            <p class="text-subtitle-1 text-grey-darken-1">
              {{ $t('dashboard.description') }}
            </p>
            <breadcrumb />
          </div>
        </div>
      </v-col>

      <v-col cols="12" md="3">
        <v-card
          :to="{ name: 'projects' }"
          class="rounded-lg"
          elevation="2"
          color="blue-lighten-5"
        >
          <v-card-text>
            <div class="d-flex justify-space-between align-center mb-2">
              <div
                class="text-subtitle-1 text-blue-darken-2 font-weight-medium"
              >
                {{ $t('dashboard.totalProjects') }}
              </div>
              <v-icon color="blue-darken-2" size="small"
                >mdi-folder-outline</v-icon
              >
            </div>
            <div class="text-h4 font-weight-bold mb-1 text-blue-darken-3">
              1
            </div>
            <div class="text-caption text-success">
              <v-icon size="x-small" color="success">mdi-arrow-up</v-icon> 2 new
              this month
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="rounded-lg" elevation="2" color="purple-lighten-5">
          <v-card-text>
            <div class="d-flex justify-space-between align-center mb-2">
              <div
                class="text-subtitle-1 text-purple-darken-2 font-weight-medium"
              >
                {{ $t('dashboard.activeClients') }}
              </div>
              <v-icon color="purple-darken-2" size="small"
                >mdi-account-group-outline</v-icon
              >
            </div>
            <div class="text-h4 font-weight-bold mb-1 text-purple-darken-3">
              1
            </div>
            <div class="text-caption text-purple-darken-1">
              Across 5 regions
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card
          :to="{ name: 'acl' }"
          class="rounded-lg"
          elevation="2"
          color="teal-lighten-5"
        >
          <v-card-text>
            <div class="d-flex justify-space-between align-center mb-2">
              <div
                class="text-subtitle-1 text-teal-darken-2 font-weight-medium"
              >
                {{ $t('dashboard.aclRoles') }}
              </div>
              <v-icon color="teal-darken-2" size="small"
                >mdi-shield-account-outline</v-icon
              >
            </div>
            <div class="text-h4 font-weight-bold mb-1 text-teal-darken-3">
              -
            </div>
            <div class="text-caption text-teal-darken-1">
              {{ $t('acl.description') }}
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <!-- <v-col cols="12" md="3">
        <v-card class="rounded-lg" elevation="2" color="orange-lighten-5">
          <v-card-text>
            <div class="d-flex justify-space-between align-center mb-2">
              <div
                class="text-subtitle-1 text-orange-darken-2 font-weight-medium"
              >
                Pending Tasks
              </div>
              <v-icon color="orange-darken-2" size="small"
                >mdi-clipboard-list-outline</v-icon
              >
            </div>
            <div class="text-h4 font-weight-bold mb-1 text-orange-darken-3">
              8
            </div>
            <div class="text-caption text-orange-darken-1">Needs attention</div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="3">
        <v-card class="rounded-lg" elevation="2" color="green-lighten-5">
          <v-card-text>
            <div class="d-flex justify-space-between align-center mb-2">
              <div
                class="text-subtitle-1 text-green-darken-2 font-weight-medium"
              >
                Revenue
              </div>
              <v-icon color="green-darken-2" size="small"
                >mdi-currency-usd</v-icon
              >
            </div>
            <div class="text-h4 font-weight-bold mb-1 text-green-darken-3">
              $24k
            </div>
            <div class="text-caption text-success">
              <v-icon size="x-small" color="success">mdi-arrow-up</v-icon>
              +12% from last month
            </div>
          </v-card-text>
        </v-card>
      </v-col> -->

      <v-col cols="12" md="8">
        <v-card class="rounded-lg pa-4" elevation="2">
          <LineChart title="Project Activity" />
        </v-card>
      </v-col>
      <v-col cols="12" md="4">
        <v-card class="rounded-lg pa-4" elevation="2">
          <PieChart title="Project Distribution" />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import breadcrumb from '@/module/core/component/breadcrumb.vue'
import LineChart from '../component/LineChart.vue'
import PieChart from '../component/PieChart.vue'
</script>
