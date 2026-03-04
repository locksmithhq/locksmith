<template>
  <v-row no-gutters>
    <v-col cols="12" lg="4" md="5" class="pa-0 pr-md-6">
      <div class="mb-6">
        <h3 class="text-h6 font-weight-bold mb-1">
          {{ $t('clientDetails.loginSettings') }}
        </h3>
        <p class="text-caption text-grey-darken-1">
          {{ $t('clientDetails.loginSettingsDescription') }}
        </p>
      </div>

      <v-card elevation="0" border rounded="lg" class="pa-4 mb-4">
        <div class="d-flex align-center justify-space-between">
          <div>
            <div class="text-subtitle-2 font-weight-bold">
              {{ $t('clientDetails.loginEnabled') }}
            </div>
            <div class="text-caption text-grey-darken-1">
              {{ $t('clientDetails.loginEnabledDescription') }}
            </div>
          </div>
          <v-switch
            v-model="controller.loginConfig.enabled"
            color="primary"
            density="compact"
            hide-details
            inset
          ></v-switch>
        </div>
      </v-card>

      <v-expansion-panels
        v-model="activePanels"
        multiple
        class="custom-panels mb-4"
      >
        <!-- Visual & Layout -->
        <v-expansion-panel elevation="0" border rounded="lg" class="mb-3">
          <v-expansion-panel-title class="py-3">
            <v-icon color="primary" class="mr-3">mdi-view-quilt-outline</v-icon>
            <span class="font-weight-bold">{{
              $t('clientDetails.layoutAndStyle')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-row dense>
              <v-col cols="12">
                <v-select
                  v-model="controller.loginConfig.layout"
                  :label="$t('clientDetails.pageLayout')"
                  :items="[
                    { title: 'Split (Right)', value: 'split-right' },
                    { title: 'Split (Left)', value: 'split-left' },
                    { title: 'Centered', value: 'centered' },
                  ]"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                  class="mb-4"
                ></v-select>
              </v-col>
              <v-col cols="12">
                <v-select
                  v-model="controller.loginConfig.input_variant"
                  :label="$t('clientDetails.inputStyle')"
                  :items="[
                    { title: 'Outlined', value: 'outlined' },
                    { title: 'Filled', value: 'filled' },
                    { title: 'Underlined', value: 'underlined' },
                  ]"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                ></v-select>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>

        <!-- Branding -->
        <v-expansion-panel elevation="0" border rounded="lg" class="mb-3">
          <v-expansion-panel-title class="py-3">
            <v-icon color="primary" class="mr-3">mdi-palette-outline</v-icon>
            <span class="font-weight-bold">{{
              $t('clientDetails.brandingAndColors')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-row dense>
              <v-col cols="12">
                <v-label class="text-caption font-weight-bold text-grey-darken-1 mb-2"
                  >{{ $t('clientDetails.logoUpload') }}</v-label
                >
                <input
                  ref="logoFileInput"
                  type="file"
                  accept="image/*"
                  style="display: none"
                  @change="handleLogoUpload"
                />
                
                <!-- Sem logo - mostra botão de upload -->
                <div v-if="!controller.loginConfig.logo_url">
                  <v-btn
                    color="primary"
                    variant="outlined"
                    block
                    height="48"
                    @click="triggerLogoUpload"
                  >
                    <v-icon start>mdi-upload</v-icon>
                    {{ $t('clientDetails.logoUpload') }}
                  </v-btn>
                  <div class="text-caption text-grey mt-1">
                    {{ $t('clientDetails.logoSizeHint') }}
                  </div>
                </div>
                
                <!-- Com logo - mostra preview com opções -->
                <div v-else>
                  <v-card variant="outlined" rounded="lg" class="pa-3">
                    <div class="d-flex align-center ga-3">
                      <v-img
                        :src="controller.loginConfig.logo_url"
                        max-height="60"
                        max-width="120"
                        contain
                        class="border rounded"
                      />
                      <div class="flex-grow-1">
                        <div class="text-caption text-grey mb-2">
                          {{ $t('clientDetails.logoCurrent') }}
                        </div>
                        <div class="d-flex ga-2">
                          <v-btn
                            size="small"
                            color="primary"
                            variant="tonal"
                            @click="triggerLogoUpload"
                          >
                            <v-icon start size="16">mdi-upload</v-icon>
                            {{ $t('clientDetails.logoChange') }}
                          </v-btn>
                          <v-btn
                            size="small"
                            color="error"
                            variant="text"
                            @click="controller.loginConfig.logo_url = ''"
                          >
                            <v-icon start size="16">mdi-delete</v-icon>
                            {{ $t('clientDetails.removeLogo') }}
                          </v-btn>
                        </div>
                      </div>
                    </div>
                  </v-card>
                </div>
              </v-col>

              <v-col cols="12">
                <v-menu :close-on-content-click="false" location="bottom start">
                  <template v-slot:activator="{ props }">
                    <v-text-field
                      v-model="controller.loginConfig.primary_color"
                      :label="$t('clientDetails.primaryColor')"
                      variant="outlined"
                      density="comfortable"
                      hide-details
                      class="mb-4"
                      v-bind="props"
                      readonly
                    >
                      <template v-slot:prepend-inner>
                        <div
                          :style="{
                            backgroundColor:
                              controller.loginConfig.primary_color || '#667eea',
                            width: '24px',
                            height: '24px',
                            borderRadius: '4px',
                            border: '1px solid rgba(0,0,0,0.1)',
                          }"
                        ></div>
                      </template>
                    </v-text-field>
                  </template>
                  <v-color-picker
                    v-model="controller.loginConfig.primary_color"
                    mode="hex"
                    hide-inputs
                    elevation="4"
                  ></v-color-picker>
                </v-menu>
              </v-col>

              <v-col cols="12">
                <v-select
                  v-model="controller.loginConfig.background_type"
                  :label="$t('clientDetails.backgroundType')"
                  :items="[
                    { title: 'Color', value: 'color' },
                    { title: 'Image', value: 'image' },
                  ]"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                  class="mb-4"
                ></v-select>
              </v-col>

              <v-col
                cols="12"
                v-if="controller.loginConfig.background_type === 'image'"
                key="bg-image-input"
              >
                <v-text-field
                  v-model="controller.loginConfig.background_image"
                  :label="$t('clientDetails.backgroundImageUrl')"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                  prepend-inner-icon="mdi-image"
                ></v-text-field>
              </v-col>

              <v-col cols="12" v-else key="bg-color-input">
                <v-menu :close-on-content-click="false" location="bottom start">
                  <template v-slot:activator="{ props }">
                    <v-text-field
                      v-model="controller.loginConfig.background_color"
                      :label="$t('clientDetails.backgroundColor')"
                      variant="outlined"
                      density="comfortable"
                      hide-details
                      v-bind="props"
                      readonly
                    >
                      <template v-slot:prepend-inner>
                        <div
                          :style="{
                            backgroundColor:
                              controller.loginConfig.background_color ||
                              '#ffffff',
                            width: '24px',
                            height: '24px',
                            borderRadius: '4px',
                            border: '1px solid rgba(0,0,0,0.1)',
                          }"
                        ></div>
                      </template>
                    </v-text-field>
                  </template>
                  <v-color-picker
                    v-model="controller.loginConfig.background_color"
                    mode="hex"
                    hide-inputs
                    elevation="4"
                  ></v-color-picker>
                </v-menu>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>

        <!-- Components -->
        <v-expansion-panel elevation="0" border rounded="lg" class="mb-3">
          <v-expansion-panel-title class="py-3">
            <v-icon color="primary" class="mr-3">mdi-tune-variant</v-icon>
            <span class="font-weight-bold">{{
              $t('clientDetails.functionalities')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-row dense>
              <v-col cols="6">
                <v-switch
                  v-model="controller.loginConfig.show_social"
                  :label="$t('clientDetails.socialLogin')"
                  color="primary"
                  density="compact"
                  hide-details
                ></v-switch>
              </v-col>
              <v-col cols="6">
                <v-switch
                  v-model="controller.loginConfig.show_remember_me"
                  :label="$t('clientDetails.rememberMe')"
                  color="primary"
                  density="compact"
                  hide-details
                ></v-switch>
              </v-col>
              <v-col cols="6">
                <v-switch
                  v-model="controller.loginConfig.show_forgot_password"
                  :label="$t('clientDetails.forgotPassword')"
                  color="primary"
                  density="compact"
                  hide-details
                ></v-switch>
              </v-col>
              <v-col cols="6">
                <v-switch
                  v-model="controller.loginConfig.show_sign_up"
                  :label="$t('clientDetails.signUp')"
                  color="primary"
                  density="compact"
                  hide-details
                ></v-switch>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>

        <!-- Custom HTML/CSS -->
        <v-expansion-panel elevation="0" border rounded="lg">
          <v-expansion-panel-title class="py-3">
            <v-icon color="primary" class="mr-3">mdi-code-braces</v-icon>
            <span class="font-weight-bold">{{
              $t('clientDetails.advancedHtmlCss')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-switch
              v-model="controller.loginConfig.use_custom_html"
              :label="$t('clientDetails.enableCustomHtml')"
              color="primary"
              density="compact"
              inset
              class="mb-2"
            ></v-switch>

            <template v-if="controller.loginConfig.use_custom_html">
              <v-textarea
                v-model="controller.loginConfig.custom_html"
                :label="$t('clientDetails.htmlCode')"
                variant="outlined"
                density="comfortable"
                rows="6"
                placeholder="<div id='custom-login'>...</div>"
                class="code-textarea mb-3"
              ></v-textarea>

              <v-textarea
                v-model="controller.loginConfig.custom_css"
                :label="$t('clientDetails.cssCode')"
                variant="outlined"
                density="comfortable"
                rows="6"
                placeholder="#custom-login { ... }"
                class="code-textarea"
              ></v-textarea>

              <v-alert
                type="info"
                variant="tonal"
                density="compact"
                class="mt-2 text-caption"
              >
                {{
                  $t('clientDetails.customHtmlHint', {
                    emailId: 'id="email"',
                    passwordId: 'id="password"',
                  })
                }}
              </v-alert>
            </template>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>

      <v-btn
        color="primary"
        block
        size="large"
        class="text-capitalize mb-6"
        rounded="lg"
        elevation="2"
        :loading="controller.saving"
        @click="controller.save"
      >
        <v-icon start>mdi-content-save-outline</v-icon>
        {{ $t('clientDetails.saveChanges') }}
      </v-btn>

      <!-- Test Links -->
      <v-card elevation="0" border rounded="lg" class="pa-4 bg-grey-lighten-5">
        <div class="d-flex align-center mb-4">
          <v-icon color="grey-darken-2" class="mr-2">mdi-link-variant</v-icon>
          <span class="text-subtitle-2 font-weight-bold">{{
            $t('clientDetails.testLinksTitle')
          }}</span>
        </div>
        <div v-for="link in testLinks" :key="link.url" class="mb-2">
          <v-btn
            variant="flat"
            color="white"
            size="small"
            block
            border
            class="text-none justify-start"
            :href="link.url"
            target="_blank"
            rounded="md"
          >
            <v-icon start size="16" color="primary">mdi-open-in-new</v-icon>
            <span class="text-truncate text-grey-darken-3">{{ link.uri }}</span>
          </v-btn>
        </div>
        <div
          v-if="testLinks.length === 0"
          class="text-caption text-grey text-center py-2"
        >
          {{ $t('clientDetails.noRedirectUris') }}
        </div>
      </v-card>
    </v-col>

    <v-col cols="12" lg="8" md="7" class="pa-0">
      <div class="d-flex justify-space-between align-center mb-4">
        <div>
          <h2 class="text-h6 font-weight-bold d-flex align-center">
            {{ $t('clientDetails.loginScreen') }}
            <v-chip size="x-small" color="primary" variant="tonal" class="ml-2"
              >Preview Vivo</v-chip
            >
          </h2>
        </div>
        <div class="d-flex">
          <v-btn-toggle
            v-model="previewMode"
            mandatory
            density="compact"
            color="primary"
            rounded="lg"
          >
            <v-btn value="desktop" size="small">
              <v-icon>mdi-monitor</v-icon>
            </v-btn>
            <v-btn value="mobile" size="small">
              <v-icon>mdi-cellphone</v-icon>
            </v-btn>
          </v-btn-toggle>
        </div>
      </div>

      <!-- Preview Container -->
      <div
        class="preview-wrapper d-flex align-center justify-center bg-grey-lighten-4 rounded-xl border pa-8 overflow-hidden"
        style="height: 800px; position: relative"
      >
        <!-- Browser Frame Decorator for Desktop -->
        <div
          v-if="previewMode === 'desktop'"
          key="browser-frame"
          class="browser-frame-top d-flex align-center px-4"
        >
          <div class="circle red"></div>
          <div class="circle yellow"></div>
          <div class="circle green"></div>
          <div class="browser-address-bar ml-4 text-caption text-grey-darken-1">
            {{ currentOrigin }}/auth...
          </div>
        </div>

        <v-card
          elevation="24"
          :class="['preview-content', previewMode]"
          class="rounded-lg overflow-hidden border-0"
          :style="{
            width: previewMode === 'mobile' ? '375px' : '100%',
            height: previewMode === 'mobile' ? '667px' : '100%',
            transition: 'all 0.4s cubic-bezier(0.4, 0, 0.2, 1)',
          }"
        >
          <!-- Custom HTML Preview -->
          <div
            v-if="controller.loginConfig.use_custom_html"
            key="custom-html-preview"
            class="fill-height pa-0 d-flex flex-column justify-center align-center"
            ref="customPreviewHost"
          ></div>

          <v-row
            v-else
            key="default-preview"
            no-gutters
            class="fill-height"
            :class="{
              'justify-center align-center':
                controller.loginConfig.layout === 'centered',
            }"
            :style="
              controller.loginConfig.layout === 'centered'
                ? backgroundStyle
                : {}
            "
          >
            <!-- Branding Side -->
            <v-col
              v-if="controller.loginConfig.layout !== 'centered'"
              cols="12"
              md="6"
              class="fill-height"
              :class="previewMode === 'mobile' ? 'd-none' : 'd-none d-md-flex'"
              :order="controller.loginConfig.layout === 'split-left' ? 0 : 1"
            >
              <div
                class="bg-gradient-primary d-flex flex-column justify-center align-center pa-12 w-100 fill-height"
                :style="{
                  position: 'relative',
                  overflow: 'hidden',
                  ...backgroundStyle,
                }"
              >
                <div class="decorative-circle circle-1"></div>
                <div class="decorative-circle circle-2"></div>

                <div
                  style="position: relative; z-index: 1"
                  class="text-center text-white"
                >
                  <v-img
                    v-if="controller.loginConfig.logo_url"
                    :src="controller.loginConfig.logo_url"
                    max-height="100"
                    contain
                  />
                  <v-icon
                    v-else
                    key="branding-icon"
                    size="80"
                    color="white"
                    class="mb-4"
                    style="opacity: 0.9"
                  >
                    mdi-shield-lock-outline
                  </v-icon>
                  <h3 class="text-h5 font-weight-bold mb-1">Secure Login</h3>
                  <p class="text-body-2" style="opacity: 0.8">
                    OAuth 2.0 Protection
                  </p>
                </div>
              </div>
            </v-col>

            <!-- Form Side -->
            <v-col
              cols="12"
              :md="
                previewMode === 'mobile'
                  ? 12
                  : controller.loginConfig.layout === 'centered'
                    ? 5
                    : 6
              "
              class="d-flex flex-column justify-center align-center"
              :class="[
                previewMode === 'mobile' ? 'pa-6' : 'pa-8 pa-md-16',
                { 'rounded-xl': controller.loginConfig.layout === 'centered' },
              ]"
              :style="previewMode === 'mobile' ? { minHeight: '100%' } : {}"
            >
              <div class="w-100" style="max-width: 450px">
                <div
                  class="text-center mb-6"
                  v-if="controller.loginConfig.layout === 'centered'"
                >
                  <v-img
                    v-if="controller.loginConfig.logo_url"
                    :src="controller.loginConfig.logo_url"
                    max-height="100"
                  />
                </div>

                <v-form class="login-form">
                  <v-text-field
                    v-model="previewEmail"
                    :label="$t('clientDetails.email')"
                    :variant="controller.loginConfig.input_variant"
                    density="comfortable"
                    rounded="lg"
                    prepend-inner-icon="mdi-email-outline"
                    hide-details
                    class="mb-4"
                    placeholder="user@example.com"
                    :color="controller.loginConfig.primary_color || 'primary'"
                  ></v-text-field>

                  <v-text-field
                    v-model="previewPassword"
                    :label="$t('clientDetails.password')"
                    :variant="controller.loginConfig.input_variant"
                    density="comfortable"
                    rounded="lg"
                    prepend-inner-icon="mdi-lock-outline"
                    type="password"
                    hide-details
                    class="mb-4"
                    placeholder="••••••••"
                    :color="controller.loginConfig.primary_color || 'primary'"
                  ></v-text-field>

                  <div
                    class="d-flex justify-space-between align-center mb-6"
                    v-if="
                      controller.loginConfig.show_remember_me ||
                      controller.loginConfig.show_forgot_password
                    "
                  >
                    <v-checkbox
                      v-model="previewRememberMe"
                      v-if="controller.loginConfig.show_remember_me"
                      :label="$t('clientDetails.rememberMe')"
                      hide-details
                      density="compact"
                      :color="controller.loginConfig.primary_color || 'primary'"
                    ></v-checkbox>
                    <a
                      v-if="controller.loginConfig.show_forgot_password"
                      href="#"
                      class="text-decoration-none text-body-2 font-weight-medium"
                      :style="{
                        color:
                          controller.loginConfig.primary_color ||
                          'rgb(var(--v-theme-primary))',
                      }"
                    >
                      {{ $t('clientDetails.forgotPassword') }}
                    </a>
                  </div>

                  <v-btn
                    :color="controller.loginConfig.primary_color || 'primary'"
                    size="large"
                    block
                    elevation="2"
                    rounded="lg"
                    class="text-capitalize mb-4"
                  >
                    {{ $t('clientDetails.signIn') }}
                  </v-btn>

                  <template v-if="controller.loginConfig.show_social">
                    <v-divider class="my-6">
                      <span class="text-caption text-grey-darken-1 px-3">
                        {{ $t('clientDetails.orContinueWith') }}
                      </span>
                    </v-divider>
                    <v-row dense>
                      <v-col cols="6">
                        <v-btn
                          variant="outlined"
                          block
                          rounded="lg"
                          class="text-capitalize px-0"
                          size="small"
                        >
                          <v-icon start size="small">mdi-google</v-icon> Google
                        </v-btn>
                      </v-col>
                      <v-col cols="6">
                        <v-btn
                          variant="outlined"
                          block
                          rounded="lg"
                          class="text-capitalize px-0"
                          size="small"
                        >
                          <v-icon start size="small">mdi-github</v-icon> GitHub
                        </v-btn>
                      </v-col>
                    </v-row>
                  </template>

                  <div
                    v-if="controller.loginConfig.show_sign_up"
                    class="text-center mt-6"
                  >
                    <span class="text-body-2 text-grey-darken-1">
                      {{ $t('clientDetails.dontHaveAccount') }}
                    </span>
                    <a
                      href="#"
                      class="text-primary text-decoration-none text-body-2 font-weight-bold ml-1"
                    >
                      {{ $t('clientDetails.signUp') }}
                    </a>
                  </div>
                </v-form>
              </div>
            </v-col>
          </v-row>
        </v-card>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'

const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})

const { locale } = useI18n()
const previewMode = ref('desktop')
const activePanels = ref([0, 1])
const currentOrigin = window.location.origin
const customPreviewHost = ref(null)
const logoFileInput = ref(null)

// Dummy state for preview interactions
const previewEmail = ref('')
const previewPassword = ref('')
const previewRememberMe = ref(false)

const triggerLogoUpload = () => {
  logoFileInput.value?.click()
}

const handleLogoUpload = (event) => {
  const file = event.target.files?.[0]
  if (!file) return

  if (file.size > 5 * 1024 * 1024) {
    alert('File size must be less than 5MB')
    return
  }

  const reader = new FileReader()
  reader.onload = (e) => {
    const result = e.target?.result
    if (result && typeof result === 'string') {
      props.controller.loginConfig.logo_url = result
    }
  }
  reader.readAsDataURL(file)
  
  event.target.value = ''
}

// Shadow DOM logic for custom preview isolation
watch(
  () => [
    props.controller.loginConfig.custom_html,
    props.controller.loginConfig.custom_css,
    props.controller.loginConfig.use_custom_html,
  ],
  () => {
    if (props.controller.loginConfig.use_custom_html) {
      setTimeout(() => {
        if (customPreviewHost.value) {
          if (!customPreviewHost.value.shadowRoot) {
            customPreviewHost.value.attachShadow({ mode: 'open' })
          }

          const shadow = customPreviewHost.value.shadowRoot
          const html = props.controller.loginConfig.custom_html || ''
          const css = props.controller.loginConfig.custom_css || ''

          shadow.innerHTML = `
            <style>
              :host { display: block; height: 100%; width: 100%; overflow: auto; background: white; }
              ${css}
            </style>
            ${html}
          `
        }
      }, 50)
    }
  },
  { immediate: true },
)

const testLinks = computed(() => {
  if (!props.controller.client?.redirect_uris) return []
  return props.controller.client.redirect_uris
    .split(/[\s,]+/)
    .filter((uri) => uri.trim())
    .map((uri) => ({
      uri,
      url: `${window.location.origin}/${locale.value}/auth?client_id=${props.controller.client.client_id}&redirect_uri=${uri}`,
    }))
})

const backgroundStyle = computed(() => {
  const { background_type, background_color, background_image } =
    props.controller.loginConfig

  if (background_type === 'image' && background_image) {
    return {
      backgroundImage: `url(${background_image})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
    }
  }

  if (background_color) {
    return {
      background: background_color,
    }
  }

  return {}
})
</script>

<style scoped>
.custom-panels :deep(.v-expansion-panel-text__wrapper) {
  padding: 16px;
}

.custom-panels :deep(.v-expansion-panel-title--active) {
  color: rgb(var(--v-theme-primary));
}

.preview-wrapper {
  background-image: radial-gradient(#d1d1d1 1px, transparent 1px);
  background-size: 20px 20px;
}

.browser-frame-top {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 40px;
  background: #f1f3f4;
  border-bottom: 1px solid #ddd;
  z-index: 10;
}

.browser-frame-top .circle {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  margin-right: 8px;
}

.circle.red {
  background: #ff5f56;
}
.circle.yellow {
  background: #ffbd2e;
}
.circle.green {
  background: #27c93f;
}

.browser-address-bar {
  background: white;
  height: 24px;
  flex-grow: 1;
  border-radius: 12px;
  display: flex;
  align-items: center;
  padding: 0 12px;
  border: 1px solid #e0e0e0;
}

.preview-content.desktop {
  margin-top: 40px;
  height: calc(100% - 40px) !important;
}

.preview-content.mobile {
  box-shadow:
    0 0 0 12px #333,
    0 0 0 15px #555,
    0 30px 60px rgba(0, 0, 0, 0.3) !important;
}

.login-form {
  width: 100%;
}

.code-textarea :deep(textarea) {
  font-family: 'Fira Code', 'Courier New', monospace;
  font-size: 13px;
  line-height: 1.5;
  background-color: #f8f9fa;
  color: #2c3e50;
}

.bg-gradient-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.decorative-circle {
  position: absolute;
  border-radius: 50%;
}

.circle-1 {
  top: -50px;
  right: -50px;
  width: 300px;
  height: 300px;
  background: rgba(255, 255, 255, 0.1);
}

.circle-2 {
  bottom: -100px;
  left: -100px;
  width: 400px;
  height: 400px;
  background: rgba(255, 255, 255, 0.05);
}

@media (max-width: 960px) {
  .preview-wrapper {
    height: 600px !important;
    padding: 16px !important;
  }

  .preview-content.mobile {
    transform: scale(0.8);
  }
}
</style>
