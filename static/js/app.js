fetch('/api/skills')
    .then(response => response.json())
    .then(skills => {
        const skillsList = document.getElementById('skills-list');
        skillsList.innerHTML = skills.map(skill => 
            `<span class="skill-tag">${skill}</span>`
        ).join('');
    })
    .catch(error => {
        console.error('Error loading skills:', error);
    });

document.addEventListener("DOMContentLoaded", function() {
    const contactForm = document.getElementById('contact-form');

    if(contactForm) {
        contactForm.addEventListener('submit', function(e) {
            e.preventDefault();

            const formData = new FormData(contactForm);

            fetch('/contact', {
                method: 'POST',
                body: formData
            })
            .then(response => response.json())
            .then(data => {
                if (data.status == 'success') {
                    alert('✅ Сообщение отправлено! Спасибо!');
                    contactForm.reset();
                } else {
                    alert('❌ Ошибка отправки: ' + data.message);
                }
            })
            .catch(error => {
                console.error('Error:', error);
                alert('⚠️ Произошла ошибка при отправке!');
            })
        })
    }
})