<template>
  <v-container fluid class="pa-0">
    <v-row no-gutters>
      <!-- Workspace -->
      <v-col cols="12" md="8" lg="9" class="pa-2">
        <v-card elevation="0" rounded="lg" class="workspace-card fill-height">
          <!-- Workspace Header -->
          <div class="pa-3 px-4 d-flex align-center justify-space-between workspace-header">
            <div class="d-flex align-center" style="gap: 10px">
              <div class="workspace-icon">
                <v-icon color="primary" size="16">mdi-shield-key-outline</v-icon>
              </div>
              <p class="text-subtitle-2 font-weight-bold section-title mb-0 text-uppercase" style="letter-spacing: 0.05em">
                {{ $t('projectDetails.permissionsWorkspace') || 'Permission Workspace' }}
              </p>
            </div>
            <v-btn
              color="primary"
              prepend-icon="mdi-content-save-outline"
              height="32"
              rounded="lg"
              elevation="0"
              class="text-capitalize font-weight-medium"
              @click="controller.permission.save"
            >
              {{ $t('common.save') || 'Save Changes' }}
            </v-btn>
          </div>

          <v-divider />

          <v-card-text class="pa-4">
            <draggable
              v-model="controller.permission.permissions"
              tag="div"
              class="workspace-drop-zone masonry-grid"
              item-key="id"
              :group="{
                name: 'permission-roles',
                pull: false,
                put: (to, from) => from.options.group.name === 'roles',
              }"
              @add="onAddRole"
            >
              <template #item="{ element: role }">
                <div class="masonry-item">
                  <div class="pa-2">
                    <v-card elevation="0" rounded="lg" class="role-workspace-card overflow-hidden d-flex flex-column">
                      <div class="role-card-header d-flex align-center px-4 py-2">
                        <v-icon color="primary" class="mr-2" size="16">mdi-account-group-outline</v-icon>
                        <span class="font-weight-bold text-subtitle-2 section-title">{{ role.title }}</span>
                        <v-spacer />
                        <v-btn
                          icon="mdi-close"
                          variant="text"
                          size="x-small"
                          color="error"
                          @click.stop="removeRole(role.id)"
                        />
                      </div>

                      <v-card-text class="pa-2" style="background: #f9fafb">
                        <draggable
                          v-model="role.modules"
                          item-key="id"
                          class="module-drop-zone rounded-lg"
                          :group="{
                            name: 'role-modules',
                            pull: false,
                            put: (to, from) => from.options.group.name === 'modules',
                          }"
                          @add="onAddModule(role, $event)"
                        >
                          <template #item="{ element: module }">
                            <v-card variant="flat" border rounded="lg" class="module-workspace-card mb-2">
                              <div class="module-card-header d-flex align-center px-3 py-1">
                                <v-icon size="14" color="grey-darken-1" class="mr-2">mdi-package-variant-closed</v-icon>
                                <span class="font-weight-bold text-caption" style="color: #374151">{{ module.title }}</span>
                                <v-spacer />
                                <v-btn
                                  icon="mdi-delete-outline"
                                  variant="text"
                                  size="x-small"
                                  color="grey"
                                  @click.stop="removeModule(role.id, module.id)"
                                />
                              </div>

                              <v-card-text class="pa-2 pt-0">
                                <draggable
                                  v-model="module.actions"
                                  item-key="id"
                                  class="actions-drop-zone d-flex flex-wrap align-center"
                                  :group="{
                                    name: 'module-actions',
                                    pull: false,
                                    put: (to, from) => from.options.group.name === 'actions',
                                  }"
                                  @change="onModuleActionsChange(module)"
                                >
                                  <template #item="{ element: action }">
                                    <v-chip
                                      closable
                                      class="ma-1 text-caption font-weight-bold"
                                      color="primary"
                                      variant="tonal"
                                      label
                                      @click:close="removeAction(role.id, module.id, action.id)"
                                    >
                                      {{ action.title }}
                                    </v-chip>
                                  </template>
                                  <template #header v-if="module.actions.length === 0">
                                    <div class="px-2 py-1 text-caption section-subtitle" style="font-style: italic">
                                      {{ $t('projectDetails.dropActionsHere') || 'Drop actions here...' }}
                                    </div>
                                  </template>
                                </draggable>
                              </v-card-text>
                            </v-card>
                          </template>
                          <template #header v-if="!role.modules || role.modules.length === 0">
                            <div class="d-flex align-center justify-center pa-4 rounded-lg section-subtitle" style="border: 1px dashed #e5e7eb">
                              <v-icon start size="16">mdi-plus</v-icon>
                              <span class="text-caption">{{ $t('projectDetails.dropModulesHere') || 'Drop modules here' }}</span>
                            </div>
                          </template>
                        </draggable>
                      </v-card-text>
                    </v-card>
                  </div>
                </div>
              </template>

              <template
                #header
                v-if="!controller.permission.permissions || controller.permission.permissions.length === 0"
              >
                <div class="masonry-item-full px-2">
                  <div class="d-flex flex-column align-center justify-center py-16 rounded-lg" style="border: 2px dashed #e5e7eb">
                    <v-icon size="48" color="grey-lighten-2" class="mb-3">mdi-drag-variant</v-icon>
                    <p class="text-subtitle-2 font-weight-medium text-grey-darken-1 mb-1">
                      {{ $t('projectDetails.emptyPermissions') || 'Empty workspace' }}
                    </p>
                    <p class="text-body-2 text-grey text-center px-8" style="max-width: 400px">
                      {{ $t('projectDetails.emptyPermissionsDesc') || 'Drag and drop roles from the right panel to define permissions.' }}
                    </p>
                  </div>
                </div>
              </template>
            </draggable>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Source Panel -->
      <v-col cols="12" md="4" lg="3" class="pa-2">
        <div class="sticky-panel">
          <v-card elevation="0" rounded="lg" class="source-card-container">
            <v-tabs
              v-model="activeSourceTab"
              color="primary"
              grow
              density="comfortable"
            >
              <v-tab value="roles" class="text-capitalize font-weight-medium text-caption">
                <v-icon start size="16">mdi-account-group</v-icon>
                {{ $t('projectDetails.roles') || 'Roles' }}
              </v-tab>
              <v-tab value="modules" class="text-capitalize font-weight-medium text-caption">
                <v-icon start size="16">mdi-package-variant</v-icon>
                {{ $t('projectDetails.modules') || 'Modules' }}
              </v-tab>
              <v-tab value="actions" class="text-capitalize font-weight-medium text-caption">
                <v-icon start size="16">mdi-lightning-bolt</v-icon>
                {{ $t('projectDetails.actions') || 'Actions' }}
              </v-tab>
            </v-tabs>

            <v-divider />

            <v-window
              v-model="activeSourceTab"
              class="pa-4 overflow-y-auto"
              style="max-height: calc(100vh - 300px); background: #f9fafb"
            >
              <!-- Roles -->
              <v-window-item value="roles">
                <div class="mb-4">
                  <v-text-field
                    v-model="controller.role.form.title"
                    :label="$t('projectDetails.newRoleName') || 'New Role'"
                    variant="outlined"
                    density="compact"
                    rounded="lg"
                    hide-details="auto"
                    :error-messages="controller.role.errors.title"
                    bg-color="white"
                    @keyup.enter="controller.role.save"
                  >
                    <template v-slot:append-inner>
                      <v-btn color="primary" icon="mdi-plus" size="x-small" variant="flat" rounded="lg" @click="controller.role.save" />
                    </template>
                  </v-text-field>
                </div>
                <v-divider class="mb-4" />
                <draggable
                  :list="controller.role.roles"
                  item-key="id"
                  :group="{ name: 'roles', pull: 'clone', put: false }"
                  :clone="cloneRole"
                  class="source-list"
                >
                  <template #item="{ element }">
                    <v-card elevation="0" rounded="lg" class="source-item mb-2 py-2 px-3 cursor-move">
                      <div class="d-flex align-center">
                        <v-icon size="16" class="mr-2 section-subtitle">mdi-account-group-outline</v-icon>
                        <span class="text-caption font-weight-bold section-title">{{ element.title }}</span>
                        <v-spacer />
                        <v-icon size="14" color="grey-lighten-2">mdi-drag-vertical</v-icon>
                      </div>
                    </v-card>
                  </template>
                </draggable>
              </v-window-item>

              <!-- Modules -->
              <v-window-item value="modules">
                <div class="mb-4">
                  <v-text-field
                    v-model="controller.module.form.title"
                    :label="$t('projectDetails.newModuleName') || 'New Module'"
                    variant="outlined"
                    density="compact"
                    rounded="lg"
                    hide-details="auto"
                    :error-messages="controller.module.errors.title"
                    bg-color="white"
                    @keyup.enter="controller.module.save"
                  >
                    <template v-slot:append-inner>
                      <v-btn color="primary" icon="mdi-plus" size="x-small" variant="flat" rounded="lg" @click="controller.module.save" />
                    </template>
                  </v-text-field>
                </div>
                <v-divider class="mb-4" />
                <draggable
                  :list="controller.module.modules"
                  item-key="id"
                  :group="{ name: 'modules', pull: 'clone', put: false }"
                  :clone="cloneModule"
                  class="source-list"
                >
                  <template #item="{ element }">
                    <v-card elevation="0" rounded="lg" class="source-item mb-2 py-2 px-3 cursor-move">
                      <div class="d-flex align-center">
                        <v-icon size="16" class="mr-2 section-subtitle">mdi-package-variant-closed</v-icon>
                        <span class="text-caption font-weight-bold section-title">{{ element.title }}</span>
                        <v-spacer />
                        <v-icon size="14" color="grey-lighten-2">mdi-drag-vertical</v-icon>
                      </div>
                    </v-card>
                  </template>
                </draggable>
              </v-window-item>

              <!-- Actions -->
              <v-window-item value="actions">
                <div class="mb-4">
                  <v-text-field
                    v-model="controller.action.form.title"
                    :label="$t('projectDetails.newActionName') || 'New Action'"
                    variant="outlined"
                    density="compact"
                    rounded="lg"
                    hide-details="auto"
                    :error-messages="controller.action.errors.title"
                    bg-color="white"
                    @keyup.enter="controller.action.save"
                  >
                    <template v-slot:append-inner>
                      <v-btn color="primary" icon="mdi-plus" size="x-small" variant="flat" rounded="lg" @click="controller.action.save" />
                    </template>
                  </v-text-field>
                </div>
                <v-divider class="mb-4" />
                <draggable
                  :list="controller.action.actions"
                  item-key="id"
                  :group="{ name: 'actions', pull: 'clone', put: false }"
                  :clone="cloneAction"
                  class="source-list"
                >
                  <template #item="{ element }">
                    <v-card elevation="0" rounded="lg" class="source-item mb-2 py-2 px-3 cursor-move">
                      <div class="d-flex align-center">
                        <v-icon size="16" class="mr-2 section-subtitle">mdi-lightning-bolt-outline</v-icon>
                        <span class="text-caption font-weight-bold section-title">{{ element.title }}</span>
                        <v-spacer />
                        <v-icon size="14" color="grey-lighten-2">mdi-drag-vertical</v-icon>
                      </div>
                    </v-card>
                  </template>
                </draggable>
              </v-window-item>
            </v-window>
          </v-card>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import draggable from 'vuedraggable'

const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})

const activeSourceTab = ref('roles')

function cloneRole(role) {
  return { id: role.id, title: role.title, modules: [] }
}
function cloneModule(module) {
  return { id: module.id, title: module.title, actions: [] }
}
function cloneAction(action) {
  return { ...action }
}

function onAddRole(evt) {
  const permissions = props.controller.permission.permissions
  const item = evt.item.__draggable_context?.element || permissions[evt.newIndex]
  if (!item) return
  const exists = permissions.some((r, index) => r.id === item.id && index !== evt.newIndex)
  if (exists) permissions.splice(evt.newIndex, 1)
}

function onAddModule(role, evt) {
  const modules = role.modules
  const item = evt.item.__draggable_context?.element || modules[evt.newIndex]
  if (!item) return
  const exists = modules.some((m, index) => m.id === item.id && index !== evt.newIndex)
  if (exists) modules.splice(evt.newIndex, 1)
}

function onModuleActionsChange(module) {
  if (!module.actions) return
  module.actions = module.actions.filter((a, i, arr) => i === arr.findIndex((x) => x.id === a.id))
}

function removeRole(roleId) {
  const permissions = props.controller.permission.permissions
  const index = permissions.findIndex((r) => r.id === roleId)
  if (index !== -1) permissions.splice(index, 1)
}

function removeModule(roleId, moduleId) {
  const role = props.controller.permission.permissions.find((r) => r.id === roleId)
  if (role) {
    const index = role.modules.findIndex((m) => m.id === moduleId)
    if (index !== -1) role.modules.splice(index, 1)
  }
}

function removeAction(roleId, moduleId, actionId) {
  const role = props.controller.permission.permissions.find((r) => r.id === roleId)
  if (!role) return
  const module = role.modules.find((m) => m.id === moduleId)
  if (!module) return
  const index = module.actions.findIndex((a) => a.id === actionId)
  if (index !== -1) module.actions.splice(index, 1)
}
</script>

<style scoped>
.section-title {
  color: #111827;
}
.section-subtitle {
  color: #6b7280;
}
.workspace-card {
  background: white;
  border: 1px solid #eef0f6;
}
.workspace-header {
  background: white;
}
.workspace-icon {
  width: 30px;
  height: 30px;
  border-radius: 8px;
  background: linear-gradient(135deg, rgba(var(--v-theme-primary), 0.12), rgba(var(--v-theme-primary), 0.06));
  display: flex;
  align-items: center;
  justify-content: center;
}
.source-card-container {
  background: white;
  border: 1px solid #eef0f6;
}
.workspace-drop-zone {
  min-height: 70vh;
  margin: -8px;
  width: calc(100% + 16px);
}
.masonry-grid {
  display: block !important;
  column-width: 450px;
  column-gap: 0;
}
.masonry-item {
  width: 100%;
  display: inline-block;
  break-inside: avoid-column;
}
.masonry-item-full {
  column-span: all;
  width: 100%;
}
.module-drop-zone {
  min-height: 80px;
  border: 1px dashed #e5e7eb;
  padding: 8px;
  border-radius: 8px;
}
.actions-drop-zone {
  min-height: 32px;
}
.role-workspace-card {
  background: white;
  border: 1px solid #eef0f6;
  min-height: 120px;
  transition: border-color 0.2s;
}
.role-workspace-card:hover {
  border-color: rgba(var(--v-theme-primary), 0.25) !important;
}
.role-card-header {
  background: white;
  border-bottom: 1px solid #f3f4f6;
}
.module-card-header {
  background: white;
  border-bottom: 1px solid #f3f4f6;
}
.module-workspace-card {
  background: white;
}
.source-item {
  background: white;
  border: 1px solid #eef0f6;
  user-select: none;
  transition: box-shadow 0.2s ease, border-color 0.2s ease;
}
.source-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.06) !important;
  border-color: rgb(var(--v-theme-primary)) !important;
  transform: translateY(-1px);
}
.source-list {
  min-height: 200px;
}
.sticky-panel {
  position: sticky;
  top: 0;
}
.cursor-move {
  cursor: move;
}
::-webkit-scrollbar {
  width: 4px;
}
::-webkit-scrollbar-track {
  background: transparent;
}
::-webkit-scrollbar-thumb {
  background: #e5e7eb;
  border-radius: 10px;
}
</style>
