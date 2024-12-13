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
