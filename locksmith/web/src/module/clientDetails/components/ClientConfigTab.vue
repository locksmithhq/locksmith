<template>
  <v-container fluid class="pa-0">
    <!-- Section Header -->
    <div class="d-flex align-center justify-space-between mb-4">
      <div>
        <p class="text-subtitle-2 font-weight-bold section-title mb-0">
          {{ $t('clientDetails.oauthSettings') }}
        </p>
        <p class="text-caption section-subtitle mt-1">
          {{ $t('clientDetails.description') }}
        </p>
      </div>
      <v-btn
        v-if="!controller.isEdit"
        prepend-icon="mdi-pencil-outline"
        variant="tonal"
        color="primary"
        height="32"
        rounded="lg"
        elevation="0"
        class="text-capitalize font-weight-medium"
        @click="controller.editClient"
      >
        {{ $t('common.edit') }}
      </v-btn>
    </div>

    <!-- Form -->
    <v-form :disabled="!controller.isEdit">
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="controller.client.name"
            :label="$t('clientDetails.clientName')"
            variant="outlined"
            density="compact"
            rounded="lg"
            hide-details="auto"
          />
        </v-col>

        <v-col cols="12" md="6">
          <v-text-field
            v-model="controller.client.custom_domain"
            :label="$t('clientDetails.customDomain')"
            placeholder="auth.example.com"
            variant="outlined"
            density="compact"
            rounded="lg"
            hide-details="auto"
          />
        </v-col>

        <v-col cols="12" md="6">
          <v-text-field
            v-model="controller.client.client_id"
            :label="$t('clientDetails.clientId')"
            variant="outlined"
            density="compact"
            rounded="lg"
            hide-details="auto"
            readonly
            bg-color="#f9fafb"
          />
        </v-col>

        <v-col cols="12" md="6">
          <v-text-field
            v-model="controller.client.client_secret"
            :label="$t('clientDetails.clientSecret')"
            variant="outlined"
            density="compact"
            rounded="lg"
            hide-details="auto"
            type="password"
            readonly
            bg-color="#f9fafb"
          />
        </v-col>

        <v-col cols="12">
          <v-divider class="my-1" />
        </v-col>

        <v-col cols="12">
          <v-text-field
            v-model="controller.client.grant_types"
            :label="$t('clientDetails.grantTypes')"
            variant="outlined"
            density="compact"
            rounded="lg"
            hide-details="auto"
            placeholder="authorization_code refresh_token"
          />
        </v-col>

        <v-col cols="12">
          <v-textarea
            v-model="controller.client.redirect_uris"
            :label="$t('clientDetails.redirectUris')"
            variant="outlined"
            density="compact"
            rounded="lg"
            rows="3"
            hide-details="auto"
            placeholder="https://example.com/callback"
          />
        </v-col>
      </v-row>
    </v-form>

    <!-- Action Buttons -->
    <div class="d-flex justify-end mt-4" style="gap: 8px" v-if="controller.isEdit">
      <v-btn
        color="grey"
        variant="tonal"
        rounded="lg"
        class="text-capitalize"
        @click="controller.resetClient"
      >
        {{ $t('clientDetails.cancel') }}
      </v-btn>
      <v-btn
        color="primary"
        variant="flat"
        rounded="lg"
        class="text-capitalize"
        @click="controller.saveClient"
      >
        {{ $t('clientDetails.save') }}
      </v-btn>
    </div>

    <!-- Metadata Footer -->
    <div class="d-flex flex-wrap mt-5 footer-bar pa-3 rounded-lg" style="gap: 16px">
      <div class="d-flex align-center" style="gap: 6px">
        <v-icon size="13" color="grey">mdi-clock-outline</v-icon>
        <span class="text-caption section-subtitle">
          {{ $t('clientDetails.createdAt') }}: <strong style="color: #374151">{{ controller.client.created_at }}</strong>
        </span>
      </div>
      <div class="d-flex align-center" style="gap: 6px">
        <v-icon size="13" color="grey">mdi-update</v-icon>
        <span class="text-caption section-subtitle">
          {{ $t('clientDetails.updatedAt') }}: <strong style="color: #374151">{{ controller.client.updated_at }}</strong>
        </span>
      </div>
    </div>
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
  border: 1px solid #eef0f6;
}
</style>
