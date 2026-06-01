<script setup>
const testimonialsRow1 = [
  { id: 1, name: "Budi Santoso", role: "Orang Tua Murid", comment: "Kualitas seragamnya jempolan, bahannya adem dan tidak mudah kusut. Anak saya suka sekali!", rating: 5 },
  { id: 2, name: "Siti Aminah", role: "Guru SMP", comment: "Sangat membantu untuk pengadaan alat tulis kantor. Prosesnya cepat dan barang sampai dengan aman.", rating: 5 },
  { id: 3, name: "Andi Wijaya", role: "Mahasiswa", comment: "Kalkulator scientific-nya original dan harganya paling bersahabat dibanding toko lain.", rating: 4 },
  { id: 4, name: "Rina Kartika", role: "Orang Tua Murid", comment: "Tas ranselnya awet banget, sudah dipakai satu tahun masih terlihat seperti baru. Mantap FourMart!", rating: 5 },
];

const testimonialsRow2 = [
  { id: 5, name: "Dedi Kurniawan", role: "Wirausaha", comment: "Cari buku tulis lusinan di sini paling praktis. Pengirimannya cepat banget untuk area Jakarta.", rating: 5 },
  { id: 6, name: "Lani Cahya", role: "Siswa SMA", comment: "Kotak pensilnya estetik banget! Teman-teman di sekolah pada nanyain beli di mana.", rating: 5 },
  { id: 7, name: "Hendra Putra", role: "Guru SD", comment: "Pilihan alat peraga edukasinya lengkap. Sangat membantu kegiatan belajar mengajar di kelas.", rating: 4 },
  { id: 8, name: "Maya Sari", role: "Orang Tua Murid", comment: "Belanja kebutuhan sekolah anak jadi nggak ribet lagi. Tinggal klik, bayar, barang sampai.", rating: 5 },
];
</script>

<template>
  <section id="testimoni" class="testimonials-section">
    <div class="container">
      <div class="section-header">
        <h2 class="section-title">Apa Kata Mereka?</h2>
        <p class="section-subtitle">Dengarkan pengalaman tulus dari para pelanggan setia kami tentang produk dan layanan FourMart.</p>
      </div>
    </div>

    <div class="marquee-wrapper">
      <!-- Baris 1: Jalan ke Kiri -->
      <div class="marquee">
        <div class="marquee-content">
          <div v-for="t in [...testimonialsRow1, ...testimonialsRow1]" :key="t.id" class="testimonial-card md-card md-card--elevated">
            <div class="rating">
              <span v-for="star in t.rating" :key="star" class="material-symbols-outlined star">star</span>
            </div>
            <p class="comment">"{{ t.comment }}"</p>
            <div class="user-info">
              <div class="avatar">{{ t.name.charAt(0) }}</div>
              <div class="user-details">
                <span class="user-name">{{ t.name }}</span>
                <span class="user-role">{{ t.role }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Baris 2: Jalan ke Kanan -->
      <div class="marquee marquee--reverse">
        <div class="marquee-content">
          <div v-for="t in [...testimonialsRow2, ...testimonialsRow2]" :key="t.id" class="testimonial-card md-card md-card--elevated">
            <div class="rating">
              <span v-for="star in t.rating" :key="star" class="material-symbols-outlined star">star</span>
            </div>
            <p class="comment">"{{ t.comment }}"</p>
            <div class="user-info">
              <div class="avatar avatar--alt">{{ t.name.charAt(0) }}</div>
              <div class="user-details">
                <span class="user-name">{{ t.name }}</span>
                <span class="user-role">{{ t.role }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.testimonials-section {
  padding: 80px 0;
  overflow: hidden;
  background-color: var(--md-sys-color-surface);
}

.section-header {
  text-align: center;
  margin-bottom: 48px;
}

.section-title {
  font-size: 2.2rem;
  font-weight: 700;
  color: var(--md-sys-color-on-surface);
  margin-bottom: 16px;
}

.section-subtitle {
  font-size: 1.1rem;
  color: var(--md-sys-color-on-surface-variant);
  max-width: 700px;
  margin: 0 auto;
}

.marquee-wrapper {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.marquee {
  display: flex;
  user-select: none;
  gap: 32px;
  mask-image: linear-gradient(to right, transparent, black 10%, black 90%, transparent);
}

.marquee-content {
  display: flex;
  flex-shrink: 0;
  gap: 32px;
  animation: scroll-left 40s linear infinite;
}

.marquee--reverse .marquee-content {
  animation: scroll-right 40s linear infinite;
}

.marquee:hover .marquee-content {
  animation-play-state: paused;
}

@keyframes scroll-left {
  from { transform: translateX(0); }
  to { transform: translateX(calc(-50% - 16px)); }
}

@keyframes scroll-right {
  from { transform: translateX(calc(-50% - 16px)); }
  to { transform: translateX(0); }
}

.testimonial-card {
  width: 350px;
  padding: 24px !important;
  border-radius: var(--md-shape-corner-extra-large);
  background-color: var(--md-sys-color-surface-container-lowest) !important;
  display: flex;
  flex-direction: column;
  gap: 16px;
  flex-shrink: 0;
}

.star {
  font-size: 20px;
  color: #FFB300;
  font-variation-settings: 'FILL' 1;
}

.comment {
  font-size: 0.95rem;
  line-height: 1.6;
  color: var(--md-sys-color-on-surface);
  font-style: italic;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: auto;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: var(--md-sys-color-primary-container);
  color: var(--md-sys-color-on-primary-container);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
}

.avatar--alt { background-color: var(--md-sys-color-tertiary-container); color: var(--md-sys-color-on-tertiary-container); }

.user-name { font-weight: 700; font-size: 0.9rem; color: var(--md-sys-color-on-surface); display: block; }
.user-role { font-size: 0.8rem; color: var(--md-sys-color-on-surface-variant); }

</style>