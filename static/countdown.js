
document.querySelectorAll('.countdown').forEach(element => {
    setInterval(() => {
        let count = parseInt(element.innerText)
        if (count > 0) {
            count--;
            element.textContent = count;
        } else {
            element.parentElement.parentElement.remove()
        }
    }, 950);
});
