<template>
  <NuxtLayout name="blank">
    <div class="login-wrapper md-card md-card--elevated">
      <div class="login-header" @click="navigateTo('/')">
        <span class="material-symbols-outlined text-primary">school</span>
        <h2>FourMart</h2>
      </div>

      <div class="login-intro">
        <h3>Selamat Datang</h3>
        <p class="text-muted">Masuk ke akun Anda untuk mulai berbelanja peralatan sekolah.</p>
      </div>

      <div v-if="error" class="error-alert">
        <span class="material-symbols-outlined">error</span>
        <span>{{ error }}</span>
      </div>

      <form @submit.prevent="handleSubmit" class="login-form">
        <div class="md-field-group">
          <input 
            type="email" 
            id="email" 
            placeholder=" " 
            v-model="email" 
            required 
            class="md-field"
            :class="{ 'md-field-error': error }"
          />
          <label for="email" class="md-field-label">Email</label>
        </div>

        <div class="md-field-group" style="margin-bottom: 24px;">
          <input 
            type="password" 
            id="password" 
            placeholder=" " 
            v-model="password" 
            required 
            class="md-field"
            :class="{ 'md-field-error': error }"
          />
          <label for="password" class="md-field-label">Kata Sandi</label>
        </div>

        <button 
          type="submit" 
          class="md-btn md-btn--filled submit-btn" 
          :disabled="loading"
        >
          <span v-if="loading" class="btn-spinner"></span>
          <span v-else>Masuk</span>
        </button>
      </form>

      <div class="login-footer">
        <p class="text-muted">
          Belum memiliki akun? 
          <NuxtLink to="/register">Daftar Sekarang</NuxtLink>
        </p>
        <p class="text-muted" style="font-size: 0.8rem; margin-top: 16px;">
          Demo Admin: <strong>admin@fourmart.com</strong> / <strong>admin123</strong><br/>
          Demo Customer: <strong>user@fourmart.com</strong> / <strong>user123</strong>
        </p>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import { useToast } from '~/composables/useToast'

definePageMeta({
  layout: false
})

const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

const router = useRouter()
const { login } = useAuth()
const { showToast } = useToast()

const handleSubmit = async () => {
  error.value = ''
  loading.value = true

  try {
    const loggedUser = await login(email.value, password.value)
    showToast(`Selamat datang kembali, ${loggedUser.name}!`)
    // Redirect based on role
    if (loggedUser.role === 'admin') {
      router.push('/admin')
    } else {
      router.push('/')
    }
  } catch (err) {
    error.value = err.message || 'Login gagal. Periksa kembali email dan kata sandi Anda.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-wrapper {
  width: 100%;
  max-width: 440px;
  padding: 32px !important;
  background-color: var(--md-sys-color-surface-container-low);
  border-radius: var(--md-shape-corner-extra-large);
  display: flex;
  flex-direction: column;
}

.login-header {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  cursor: pointer;
  margin-bottom: 24px;
}

.login-header h2 {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--md-sys-color-primary);
}

.login-header span {
  font-size: 32px;
}

.login-intro {
  text-align: center;
  margin-bottom: 28px;
}

.login-intro h3 {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface);
  margin-bottom: 6px;
}

.login-form {
  display: flex;
  flex-direction: column;
}

.submit-btn {
  height: 48px;
}

.login-footer {
  text-align: center;
  margin-top: 24px;
}

.error-alert {
  background-color: var(--md-sys-color-error-container);
  color: var(--md-sys-color-on-error-container);
  padding: 12px 16px;
  border-radius: var(--md-shape-corner-medium);
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.85rem;
  margin-bottom: 20px;
}

.error-alert span {
  font-size: 20px;
}

.btn-spinner {
  width: 20px;
  height: 20px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  border-top-color: #fff;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
</style>
