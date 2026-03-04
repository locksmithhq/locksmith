<template>
  <v-container fluid class="pa-0">
    <!-- Section Header -->
    <div class="d-flex align-center justify-space-between mb-3">
      <div>
        <p class="text-subtitle-2 font-weight-bold section-title mb-0">
          {{ $t('projectDetails.users') || 'User Directory' }}
        </p>
        <p class="text-caption section-subtitle mt-1">
          {{ controller.accounts.length }} items registered
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
      @update:modelValue="controller.fetchAccountsByProjectID"
    />

    <!-- Account List -->
    <v-card v-if="controller.accounts.length > 0" elevation="0" rounded="lg" class="account-table-card">
      <div
        v-for="(account, index) in controller.accounts"
        :key="index"
      >
        <div class="account-row px-4 py-2 d-flex align-center justify-space-between">
          <div class="d-flex align-center" style="gap: 10px; min-width: 0">
            <div class="item-avatar flex-shrink-0">
              <v-icon color="primary" size="16">mdi-account-outline</v-icon>
            </div>
            <div style="min-width: 0">
              <p class="text-body-2 font-weight-bold section-title mb-0 text-truncate">{{ account.name }}</p>
              <p class="text-caption section-subtitle mb-0">@{{ account.username }} · {{ account.email }}</p>
            </div>
          </div>

          <div class="d-flex align-center flex-shrink-0" style="gap: 8px">
            <v-chip color="primary" variant="tonal" class="font-weight-bold">
              {{ account.role_name || $t('users.noRole') }}
            </v-chip>
            <v-btn
              icon="mdi-pencil-outline"
              variant="text"
              size="x-small"
              color="primary"
              @click="controller.account.openEditDialog(account)"
            />
            <v-btn icon="mdi-delete-outline" variant="text" size="x-small" color="error" />
          </div>
        </div>
        <v-divider v-if="index < controller.accounts.length - 1" />
      </div>
    </v-card>

    <!-- Pagination -->
    <div class="d-flex justify-center mt-4" v-if="controller.account.filter.totalPages > 1">
      <v-pagination
        v-model="controller.account.filter.page"
        :length="controller.account.filter.totalPages"
        density="comfortable"
        rounded="lg"
        active-color="primary"
        variant="flat"
        @update:model-value="controller.fetchAccountsByProjectID"
      />
    </div>

    <!-- Empty State -->
    <div v-if="controller.accounts.length === 0" class="d-flex flex-column align-center justify-center py-16 text-center">
      <div class="empty-icon-wrap mb-4">
        <v-icon size="28" color="grey-lighten-1">mdi-account-off-outline</v-icon>
      </div>
      <p class="text-subtitle-2 font-weight-semibold text-grey-darken-2">{{ $t('users.noUsersTitle') }}</p>
      <p class="text-body-2 text-grey mt-1">{{ $t('users.noUsersDesc') }}</p>
    </div>

    <!-- Dialog -->
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
const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})
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
  background: white;
  border: 1px solid #eef0f6;
}
.account-row {
  transition: background 0.15s ease;
}
.account-row:hover {
  background: #f9fafb;
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
</style>
