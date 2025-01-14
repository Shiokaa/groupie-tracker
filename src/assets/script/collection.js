const params = new URLSearchParams(window.location.search);
const page = parseInt(params.get("page"), 10) || 1; // Default to page 1 if no page parameter

for (let i = 1; i <= 10; i++) {
  // Assuming there are 10 pages
  const elem = document.getElementById(`page-${i}`);
  if (elem) {
    if (page === i) {
      elem.style.color = "#80c342";
    } else {
      elem.style.color = ""; // Reset color for non-matching pages
    }
  }
}

const filters = document.querySelector(".filters");
const filter = document.querySelector(".filter");
const close1 = document.querySelector(".close");
const test = document.querySelector(".black-background");

filters.addEventListener("click", () => {
  filter.classList.toggle("active");
  test.style.display = "block";
});

close1.addEventListener("click", () => {
  filter.classList.remove("active");
  test.style.display = "none";
});
