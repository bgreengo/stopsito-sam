

// Gallery

var galleryIndexes = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18];
var gallery = document.getElementById('gallery');

if (gallery != null) {
	for (let i = 0; i < galleryIndexes.length; i++) {
		var imgUrl = 'https://stopsito.net/media/gallery/gallery_' + galleryIndexes[i] + '.jpg';
		var img = document.createElement('img');
		img.setAttribute('src', imgUrl);
		gallery.appendChild(img);
	}
}

// Overlay & Menu
function dismissOverlay() {
	var overlay = document.getElementById('overlay');
	overlay.classList.remove('open');
}

function showOverlay(e) {
	e.preventDefault();
	var overlay = document.getElementById('overlay');
	overlay.classList.add('open');
}

var btnOpenMenu = document.getElementById('nav-open-menu');
btnOpenMenu.addEventListener('click', showOverlay);

var btnCloseMenu = document.getElementById('nav-close-menu');
btnCloseMenu.addEventListener('click', dismissOverlay);

var overlay = document.getElementById('overlay');
overlay.addEventListener('click', dismissOverlay);

document.onkeyup = function(key) {
	if (key.keyCode === 27) {
		dismissOverlay();
	}
}
