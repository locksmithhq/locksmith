<template>
  <v-container fluid class="pa-6 bg-grey-lighten-4 fill-height align-start">
    <v-row>
      <v-col cols="12">
        <div class="d-flex align-center justify-space-between mb-6">
          <div>
            <h1 class="text-h4 font-weight-bold text-grey-darken-3">
              {{ $t('acl.title') }}
            </h1>
            <p class="text-subtitle-1 text-grey-darken-1">
              {{ $t('acl.description') }}
            </p>
            <breadcrumb />
          </div>
          <v-menu
            v-model="controller.showFilter"
            :close-on-content-click="false"
            location="bottom end"
            offset="8"
            transition="slide-y-transition"
          >
            <template v-slot:activator="{ props }">
              <v-btn
                v-bind="props"
                :color="controller.showFilter ? 'primary' : 'grey-darken-3'"
                variant="flat"
                prepend-icon="mdi-filter-variant"
                rounded="lg"
                class="text-none"
              >
                {{ $t('acl.filters') }}
              </v-btn>
            </template>

            <FilterBuilder
              v-model:rules="controller.filter.rules"
              v-model:logic="controller.filter.logic"
              v-model:autoSort="controller.filter.autoSort"
              :columns="filterColumns"
              @apply="controller.applyFilter"
              @clear="controller.clearAllFilters"
            />
          </v-menu>
        </div>

        <v-card elevation="0" rounded="lg" border>
          <v-divider></v-divider>

          <v-data-table-server
            :headers="headers"
            :options="controller.options"
            :items="controller.acls"
            :items-length="controller.totalItems"
            :loading="controller.loading"
            density="compact"
            class="elevation-0"
            fixed-header
            height="calc(100vh - 290px)"
            @update:options="controller.loadItems"
            multi-sort
          >
            <template v-slot:item.p_type="{ item }">
              <v-chip
                size="small"
                :color="item.p_type === 'p' ? 'primary' : 'secondary'"
                label
              >
                {{ item.p_type }}
              </v-chip>
            </template>
          </v-data-table-server>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import breadcrumb from '@/module/core/component/breadcrumb.vue'
import FilterBuilder from '../component/FilterBuilder.vue'
import { aclControllerImpl } from '@/module/acl/di/di'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const controller = aclControllerImpl()

const headers = [
  {
    title: t('acl.table.ptype'),
    key: 'p_type',
    align: 'start',
    sortable: false,
  },
  { title: t('acl.table.v0'), key: 'v0', sortable: true },
  { title: t('acl.table.v1'), key: 'v1', sortable: true },
  { title: t('acl.table.v2'), key: 'v2', sortable: true },
  { title: t('acl.table.v3'), key: 'v3', sortable: true },
]

const filterColumns = [
  { title: t('acl.table.v0'), value: 'v0' },
  { title: t('acl.table.v1'), value: 'v1' },
  { title: t('acl.table.v2'), value: 'v2' },
  { title: t('acl.table.v3'), value: 'v3' },
]
</script>

<style scoped>
.search-field {
  max-width: 400px;
}
</style>
