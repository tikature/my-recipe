document.addEventListener('DOMContentLoaded', function () {
    // Favorite button functionality
    const favoriteButtons = document.querySelectorAll('.favorite-btn');
    favoriteButtons.forEach(btn => {
        btn.addEventListener('click', function () {
            const heart = this.querySelector('.heart');
            heart.textContent = heart.textContent === '🤍' ? '❤️' : '🤍';
            this.style.transform = 'scale(1.2)';
            setTimeout(() => this.style.transform = 'scale(1)', 200);
        });
    });

    // Share button functionality
    const shareButtons = document.querySelectorAll('.share-btn');
    shareButtons.forEach(btn => {
        btn.addEventListener('click', function () {
            if (navigator.share) {
                navigator.share({
                    title: 'Resep dari ResePin',
                    text: 'Lihat resep lezat ini!',
                    url: window.location.href
                });
            } else {
                navigator.clipboard.writeText(window.location.href);
                this.innerHTML = '<span>✅</span>';
                setTimeout(() => this.innerHTML = '<span>📤</span>', 2000);
            }
        });
    });

    // Load more functionality
    const cards = document.querySelectorAll('.recipe-card');
    const loadMoreBtn = document.getElementById('loadMoreBtn');
    const allLoadedMsg = document.getElementById('allLoadedMsg');
    let visibleCount = 9; // tampilkan 9 dulu

    // sembunyikan yang lewat dari 9
    cards.forEach((card, index) => {
        if (index >= visibleCount) {
            card.style.display = 'none';
        }
    });

    if (loadMoreBtn) {
        loadMoreBtn.addEventListener('click', function () {
            let nextCount = visibleCount + 6; // tiap klik tambah 6 resep
            for (let i = visibleCount; i < nextCount && i < cards.length; i++) {
                cards[i].style.display = 'block';
            }
            visibleCount = nextCount;

            if (visibleCount >= cards.length) {
                loadMoreBtn.style.display = 'none'; // sembunyikan tombol
                if (allLoadedMsg) {
                    allLoadedMsg.style.display = 'block'; // tampilkan teks
                }
            }
        });
    }

    // Newsletter form
    const newsletterForm = document.querySelector('.newsletter-form');
    if (newsletterForm) {
        newsletterForm.addEventListener('submit', function (e) {
            e.preventDefault();
            const btn = this.querySelector('.newsletter-btn');
            const input = this.querySelector('.newsletter-input');

            btn.textContent = 'Berhasil! ✅';
            input.value = '';

            setTimeout(() => {
                btn.textContent = 'Berlangganan';
            }, 3000);
        });
    }
});