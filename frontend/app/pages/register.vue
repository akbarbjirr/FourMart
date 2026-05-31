<template>
  <NuxtLayout name="blank">
    <div class="register-wrapper md-card md-card--elevated">
      <div class="register-header" @click="navigateTo('/')">
        <span class="material-symbols-outlined text-primary">school</span>
        <h2>FourMart</h2>
      </div>

      <div class="register-intro">
        <h3>Pendaftaran Akun Baru</h3>
        <p class="text-muted">Buat akun untuk memesan peralatan sekolah favorit Anda.</p>
      </div>

      <div v-if="error" class="error-alert">
        <span class="material-symbols-outlined">error</span>
        <span>{{ error }}</span>
      </div>

      <form @submit.prevent="handleSubmit" class="register-form">
        <div class="md-field-group">
          <input 
            type="text" 
            id="name" 
            placeholder=" " 
            v-model="name" 
            required 
            class="md-field"
          />
          <label for="name" class="md-field-label">Nama Lengkap</label>
        </div>

        <div class="md-field-group">
          <input 
            type="email" 
            id="email" 
            placeholder=" " 
            v-model="email" 
            required 
            class="md-field"
          />
          <label for="email" class="md-field-label">Alamat Email</label>
        </div>

        <div class="md-field-group">
          <input 
            type="tel" 
            id="phone" 
            placeholder=" " 
            v-model="phone" 
            required 
            class="md-field"
          />
          <label for="phone" class="md-field-label">Nomor Telepon / WhatsApp</label>
        </div>

        <div class="md-field-group">
          <input 
            type="password" 
            id="password" 
            placeholder=" " 
            v-model="password" 
            required 
            class="md-field"
          />
          <label for="password" class="md-field-label">Kata Sandi (Min 6 karakter)</label>
        </div>

        <div class="md-field-group" style="margin-bottom: 24px;">
          <textarea 
            id="address" 
            placeholder=" " 
            v-model="address" 
            required 
            class="md-field"
            style="height: 100px; padding-top: 16px; resize: none;"
          ></textarea>
          <label for="address" class="md-field-label" style="top: 24px;">Alamat Pengiriman Lengkap</label>
        </div>

        <button 
          type="submit" 
          class="md-btn md-btn--filled submit-btn" 
          :disabled="loading"
        >
          <span v-if="loading" class="btn-spinner"></span>
          <span v-else>Daftar Akun</span>
        </button>
      </form>

      <div class="register-footer">
        <p class="text-muted">
          Sudah memiliki akun? 
          <NuxtLink to="/login">Masuk Disini</NuxtLink>
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

const name = ref('')
const email = ref('')
const phone = ref('')
const password = ref('')
const address = ref('')
const error = ref('')
const loading = ref(false)

const router = useRouter()
const { register } = useAuth()
const { showToast } = useToast()

const handleSubmit = async () => {
  error.value = ''
  
  if (password.value.length < 6) {
    error.value = 'Kata sandi minimal harus 6 karakter.'
    return
  }

  loading.value = true
  try {
    const newUser = await register(
      name.value,
      email.value,
      password.value,
      address.value,
      phone.value
    )
    showToast(`Registrasi berhasil! Selamat datang, ${newUser.name}.`)
    router.push('/')
  } catch (err) {
    error.value = err.message || 'Registrasi gagal. Email mungkin sudah terdaftar.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.register-wrapper {
  width: 100%;
  max-width: 460px;
  padding: 32px !important;
  background-color: var(--md-sys-color-surface-container-low);
  border-radius: var(--md-shape-corner-extra-large);
  display: flex;
  flex-direction: column;
}

.register-header {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  cursor: pointer;
  margin-bottom: 20px;
}

.register-header h2 {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--md-sys-color-primary);
}

.register-header span {
  font-size: 32px;
}

.register-intro {
  text-align: center;
  margin-bottom: 24px;
}

.register-intro h3 {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--md-sys-color-on-surface);
  margin-bottom: 4px;
}

.register-form {
  display: flex;
  flex-direction: column;
}

.submit-btn {
  height: 48px;
}

.register-footer {
  text-align: center;
  margin-top: 20px;
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
