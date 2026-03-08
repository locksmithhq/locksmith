<template>
  <v-container fluid class="pa-4 fill-height align-start page-bg">
    <v-row>
      <!-- Header -->
      <v-col cols="12">
        <breadcrumb class="mb-3" />
        <div class="d-flex align-center justify-space-between">
          <div>
            <h1 class="text-h5 font-weight-bold page-title">
              {{ $t('project.title') }}
            </h1>
            <p class="text-body-2 page-subtitle mt-1">
              {{ $t('project.description') }}
            </p>
          </div>
          <v-btn
            color="primary"
            prepend-icon="mdi-plus"
            height="34"
            class="text-capitalize font-weight-medium"
            rounded="lg"
            elevation="0"
            @click="controller.openDialog()"
          >
            {{ $t('project.new') }}
          </v-btn>
        </div>
      </v-col>

      <!-- Search -->
      <v-col cols="12">
        <v-text-field
          v-model="controller.filter.projectName"
          density="compact"
          variant="outlined"
          :placeholder="$t('project.search')"
          prepend-inner-icon="mdi-magnify"
          single-line
          hide-details
          rounded="lg"
          bg-color="white"
          @update:model-value="controller.applyFilter"
        />
      </v-col>

      <!-- Empty state -->
      <v-col v-if="controller.projects.length === 0" cols="12">
        <div class="d-flex flex-column align-center justify-center py-16 text-center">
          <div class="empty-icon-wrap mb-4">
            <v-icon size="32" color="grey-lighten-1">mdi-folder-open-outline</v-icon>
          </div>
          <p class="text-subtitle-1 font-weight-semibold text-grey-darken-2">{{ $t('project.empty_title') }}</p>
          <p class="text-body-2 text-grey mt-1">{{ $t('project.empty_description') }}</p>
          <v-btn
            color="primary"
            variant="tonal"
            rounded="lg"
            class="mt-5 text-capitalize font-weight-medium"
            prepend-icon="mdi-plus"
            @click="controller.openDialog()"
          >
            {{ $t('project.new') }}
          </v-btn>
        </div>
      </v-col>

      <!-- Project Cards -->
      <v-col
        v-for="project in controller.projects"
        :key="project.id"
        cols="12"
        md="6"
        lg="4"
      >
        <v-card
          class="project-card h-100 d-flex flex-column"
          elevation="0"
          rounded="lg"
          @click="viewProjectDetails(project.id)"
        >
          <v-card-item class="pa-5 pb-4">
            <div class="d-flex align-start" style="gap: 14px">
              <div class="project-avatar flex-shrink-0">
                {{ project.name?.slice(0, 2).toUpperCase() }}
              </div>

              <div class="flex-grow-1 min-w-0">
                <div class="d-flex align-center justify-space-between" style="gap: 8px">
                  <span class="text-subtitle-2 font-weight-bold text-truncate card-title">
                    {{ project.name }}
                  </span>
                  <v-chip
                    :color="project.status === 'active' ? 'success' : 'default'"
                    size="x-small"
                    variant="flat"
                    class="font-weight-bold text-uppercase flex-shrink-0"
                    rounded="sm"
                  >
                    {{ project.status }}
                  </v-chip>
                </div>
                <p class="text-body-2 mt-1 card-description">
                  {{ project.description || '—' }}
                </p>
              </div>
            </div>
          </v-card-item>

          <v-spacer />
          <v-divider />

          <v-card-actions class="px-4 py-2" @click.stop>
            <v-btn
              variant="text"
              size="small"
              color="primary"
              class="text-capitalize font-weight-medium"
              @click="viewProjectDetails(project.id)"
            >
              {{ $t('project.viewDetails') }}
            </v-btn>
            <v-spacer />
            <v-btn
              icon="mdi-pencil-outline"
              variant="text"
              size="small"
              color="grey-darken-1"
              @click="controller.openEditDialog(project)"
            />
            <v-btn
              icon="mdi-delete-outline"
              variant="text"
              size="small"
              color="error"
              @click="controller.openDeleteDialog(project)"
            />
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>

    <!-- Create / Edit Dialog -->
    <v-dialog v-model="controller.dialog" max-width="480px">
      <v-card rounded="xl" elevation="0" border>
        <v-card-item class="pa-6 pb-3">
          <div class="d-flex align-start justify-space-between">
            <div>
              <v-card-title class="text-h6 font-weight-bold pa-0 card-title">
                {{ controller.form.id ? $t('project.edit_project') : $t('project.new_project') }}
              </v-card-title>
              <v-card-subtitle class="text-body-2 pa-0 mt-1">
                {{ controller.form.id ? $t('project.edit_subtitle') : $t('project.create_subtitle') }}
              </v-card-subtitle>
            </div>
            <v-btn
              icon="mdi-close"
              variant="text"
              size="small"
              color="grey"
              @click="controller.dialog = false"
            />
          </div>
        </v-card-item>

        <v-divider />

        <v-card-text class="pa-6 pb-2">
          <v-text-field
            v-model="controller.form.name"
            :error-messages="controller.errors.name"
            :label="$t('project.name')"
            variant="outlined"
            density="compact"
            rounded="lg"
            class="mb-3"
          />
          <v-text-field
            v-model="controller.form.domain"
            :error-messages="controller.errors.domain"
            :label="$t('project.domain')"
            variant="outlined"
            density="compact"
            rounded="lg"
            class="mb-3"
            @blur="controller.fetchRoles"
          />
          <v-textarea
            v-model="controller.form.description"
            :error-messages="controller.errors.description"
            :label="$t('project.description_input')"
            variant="outlined"
            density="compact"
            rounded="lg"
            rows="3"
            auto-grow
          />
        </v-card-text>

        <v-card-actions class="pa-6 pt-3" style="gap: 8px">
          <v-spacer />
          <v-btn
            color="grey"
            variant="tonal"
            rounded="lg"
            class="text-capitalize"
            @click="controller.dialog = false"
          >
            {{ $t('project.cancel') }}
          </v-btn>
          <v-btn
            color="primary"
            variant="flat"
            rounded="lg"
            class="text-capitalize"
            @click="controller.save"
          >
            {{ $t('project.save') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Delete Dialog -->
    <v-dialog v-model="controller.deleteDialog" max-width="400px">
      <v-card rounded="xl" elevation="0" border>
        <v-card-item class="pa-6 pb-4">
          <div class="d-flex align-start" style="gap: 14px">
            <div class="delete-icon-wrap flex-shrink-0">
              <v-icon color="error" size="18">mdi-alert-outline</v-icon>
            </div>
            <div>
              <v-card-title class="text-subtitle-1 font-weight-bold pa-0 card-title">
                {{ $t('project.delete_title') || 'Delete Project' }}
              </v-card-title>
              <v-card-subtitle class="text-body-2 pa-0 mt-1" style="white-space: normal">
                {{ $t('project.delete_message') || 'This action cannot be undone. The project and all associated data will be permanently removed.' }}
              </v-card-subtitle>
            </div>
          </div>
        </v-card-item>

        <v-card-actions class="pa-6 pt-0" style="gap: 8px">
          <v-spacer />
          <v-btn
            color="grey"
            variant="tonal"
            rounded="lg"
            class="text-capitalize"
            @click="controller.deleteDialog = false"
          >
            {{ $t('project.cancel') }}
          </v-btn>
          <v-btn
            color="error"
            variant="flat"
            rounded="lg"
            class="text-capitalize"
            @click="controller.confirmDelete"
          >
            {{ $t('projectDetails.delete') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import breadcrumb from '@/module/core/component/breadcrumb.vue'
import { projectControllerImpl } from '@/module/project/di/di'
import { useRouter } from 'vue-router'

const controller = projectControllerImpl()
const router = useRouter()

const viewProjectDetails = (projectId) => {
  router.push({ name: 'project-details', params: { id: projectId } })
}
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

.card-title {
  color: #111827;
}

.card-description {
  color: #6b7280;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.project-card {
  background: #ffffff;
  border: 1px solid #eef0f6;
  cursor: pointer;
  transition: box-shadow 0.2s ease, border-color 0.2s ease, transform 0.15s ease;
}

.project-card:hover {
  box-shadow: 0 8px 24px -6px rgba(0, 0, 0, 0.1) !important;
  border-color: rgba(var(--v-theme-primary), 0.25) !important;
  transform: translateY(-2px);
}

.project-avatar {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  background: linear-gradient(
    135deg,
    rgba(var(--v-theme-primary), 0.14),
    rgba(var(--v-theme-primary), 0.07)
  );
  color: rgb(var(--v-theme-primary));
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 12px;
  letter-spacing: 0.5px;
}

.empty-icon-wrap {
  width: 64px;
  height: 64px;
  border-radius: 16px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
}

.delete-icon-wrap {
  width: 34px;
  height: 34px;
  border-radius: 8px;
  background: rgba(var(--v-theme-error), 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
}

.min-w-0 {
  min-width: 0;
}
</style>
