document.addEventListener('DOMContentLoaded', function() {
    const searchBarSvg = document.getElementById('search-bar-svg');
    const searchBar = document.getElementById('search-bar');

    searchBarSvg.addEventListener('click', function() {
        searchBar.focus();
    });
});