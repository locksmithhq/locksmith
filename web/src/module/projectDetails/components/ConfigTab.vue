<template>
  <v-row no-gutters>
    <v-col cols="12">
      <v-card elevation="0" rounded="lg" class="inner-card">
        <!-- Header -->
        <div class="pa-4 d-flex align-center justify-space-between">
          <div>
            <p class="text-subtitle-1 font-weight-bold card-title mb-0">
              {{ controller.defaultProject.name }}
            </p>
            <p class="text-caption card-subtitle mt-1">
              ID: {{ controller.defaultProject.id }}
            </p>
          </div>
          <v-btn
            v-if="!controller.isEdit"
            prepend-icon="mdi-pencil-outline"
            variant="tonal"
            color="primary"
            size="small"
            rounded="lg"
            class="text-capitalize font-weight-medium"
            @click="controller.editProject"
          >
            {{ $t('common.edit') || 'Edit' }}
          </v-btn>
        </div>

        <v-divider />

        <v-card-text class="pa-4">
          <p class="field-section-label mb-4">
            {{ $t('projectDetails.settings') || 'Project Settings' }}
          </p>

          <v-form :disabled="!controller.isEdit">
            <v-row dense>
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="controller.project.name"
                  :label="$t('projectDetails.projectName')"
                  variant="outlined"
                  density="compact"
                  rounded="lg"
                  hide-details="auto"
                />
              </v-col>
              <v-col cols="12">
                <v-textarea
                  v-model="controller.project.description"
                  :label="$t('projectDetails.projectDescription')"
                  variant="outlined"
                  density="compact"
                  rounded="lg"
                  rows="3"
                  hide-details="auto"
                  placeholder="Describe your project"
                />
              </v-col>
            </v-row>
          </v-form>

          <div class="d-flex justify-end mt-5" style="gap: 8px" v-if="controller.isEdit">
            <v-btn
              color="grey"
              variant="tonal"
              rounded="lg"
              class="text-capitalize"
              @click="controller.resetProject"
            >
              {{ $t('common.cancel') || 'Cancel' }}
            </v-btn>
            <v-btn
              color="primary"
              variant="flat"
              rounded="lg"
              class="text-capitalize"
              elevation="0"
              @click="controller.saveProject"
            >
              {{ $t('common.save') || 'Save' }}
            </v-btn>
          </div>
        </v-card-text>

        <v-divider />

        <!-- Footer -->
        <div class="pa-3 px-4 d-flex flex-wrap footer-bar" style="gap: 12px">
          <div class="d-flex align-center" style="gap: 6px">
            <v-icon size="13" color="grey">mdi-clock-outline</v-icon>
            <span class="text-caption card-subtitle">
              {{ $t('projectDetails.createdAt') || 'Created' }}:
              <strong class="text-dark">{{ controller.defaultProject.created_at }}</strong>
            </span>
          </div>
          <div class="d-flex align-center" style="gap: 6px">
            <v-icon size="13" color="grey">mdi-update</v-icon>
            <span class="text-caption card-subtitle">
              {{ $t('projectDetails.updatedAt') || 'Updated' }}:
              <strong class="text-dark">{{ controller.defaultProject.updated_at }}</strong>
            </span>
          </div>
          <div class="d-flex align-center" style="gap: 6px">
            <span class="status-dot" />
            <span class="text-caption font-weight-bold text-uppercase card-subtitle" style="letter-spacing: 0.05em">
              {{ $t('projectDetails.active') || 'Active' }}
            </span>
          </div>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>

<script setup>
const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})
</script>

<style scoped>
.inner-card {
  border: 1px solid #eef0f6;
}
.card-title {
  color: #111827;
}
.card-subtitle {
  color: #6b7280;
}
.text-dark {
  color: #374151;
}
.footer-bar {
  background: #f9fafb;
}
.field-section-label {
  font-size: 0.65rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  color: #9ca3af;
}
.status-dot {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: #22c55e;
  display: inline-block;
}
</style>
